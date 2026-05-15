package llm

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/flitsinc/go-llms/llms"
	"github.com/flitsinc/go-llms/openai"
	"github.com/flitsinc/go-llms/tools"
)

func Ollama(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	api := openai.New(
		"ollama", "functiongemma:270m",
	).WithEndpoint("http://localhost:11434/v1/chat/completions", "ollama")
	llmapi := llms.New(api)

	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Connection", "keep-alive")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	// llmapi.AddTool(toolx.ReadTool()) // it's almost work
	// llmapi.AddTool(toolx.ReadFileTool()) // it's work

	// fill tools
	if len(request.Tools) > 0 {
		for _, t := range request.Tools {
			vp := tools.ValueSchema{}
			castedT := map[string]any(t)
			err := vp.UnmarshalJSON([]byte(castedT["parameters"].(string)))
			if err != nil {
				fmt.Printf("tool: %s parameters: %s cannot unmarshal\n", castedT["name"], castedT["parameters"])
				continue
			}
			llmapi.AddTool(tools.External(castedT["name"].(string), &tools.FunctionSchema{
				Name:        castedT["name"].(string),
				Description: castedT["description"].(string),
				Parameters:  vp,
			}, func(r tools.Runner, params json.RawMessage) tools.Result {
				p, err := params.MarshalJSON()
				if err != nil {
					fmt.Printf("tool: %s, param marshal json error: %v\n", castedT["name"], err)
				} else {
					fmt.Printf("tool: %s param: %s\n", castedT["name"], p)
				}
				return tools.Successf("tool will executed by client")
			}))
		}
	}

	for update := range llmapi.Chat(request.Messages[0]["content"].(string)) {
		switch casted := update.(type) {
		case llms.TextUpdate:
			fmt.Fprintf(w, "text: %s\n", casted.Text)
			flusher.Flush()
		case llms.ToolStartUpdate:
			fmt.Fprintf(w, "toolCallId: %s\n", casted.ToolCallID)
			flusher.Flush()
		case llms.ToolDeltaUpdate:
			fmt.Fprintf(w, "toolCallId: %s\n", casted.ToolCallID)
			ds, err := casted.Delta.MarshalJSON()
			if err != nil {
				fmt.Printf("delta marshal json error: %v", err)
			} else {
				fmt.Fprintf(w, "%s\n", ds)
			}
			flusher.Flush()
		}
	}

	if err := llmapi.Err(); err != nil {
		fmt.Printf("llm call error: %v", err)
	}
}
