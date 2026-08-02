package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hra42/openrouter-go"
)

func OllamaDirect(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	c := openrouter.NewClient(
		openrouter.WithAPIKey("ollama-apikey"),
		openrouter.WithBaseURL("http://localhost:11434/v1"),
		openrouter.WithAppName("ollama"),
		// openrouter.WithAPIKey("todo"),
		// openrouter.WithBaseURL("https://api.deepseek.com"),
		// openrouter.WithAppName("DeepSeek"),
	)

	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Connection", "keep-alive")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	// convert messages
	if len(request.Messages) == 0 {
		http.Error(w, "Invalid Request, messages is empty", http.StatusBadRequest)
		return
	}
	m := []openrouter.Message{}
	for _, message := range request.Messages {
		b, err := json.Marshal(message)
		if err != nil {
			fmt.Printf("marshal message error: %v\n", err)
			continue
		}
		om := openrouter.Message{}
		err = json.Unmarshal(b, &om)
		if err != nil {
			fmt.Printf("unmarshal message error: %v\n", err)
			continue
		}
		m = append(m, om)
	}
	if len(m) == 0 {
		http.Error(w, "Invalid Request, messages is empty", http.StatusBadRequest)
		return
	}

	// convert tools
	tools := []openrouter.Tool{}
	if len(request.Tools) > 0 {
		for _, t := range request.Tools {
			f := t["function"].(map[string]any)
			b, err := json.Marshal(f["parameters"])
			if err != nil {
				fmt.Printf("marshal tool parameters error: %v\n", err)
				continue
			}
			p := make(map[string]any)
			err = json.Unmarshal(b, &p)
			if err != nil {
				fmt.Printf("unmarshal tool parameters error: %v\n", err)
				continue
			}
			tool := openrouter.Tool{
				Type: t["type"].(string),
				Function: openrouter.Function{
					Name:        f["name"].(string),
					Description: f["description"].(string),
					Parameters:  p,
				},
			}
			tools = append(tools, tool)
		}
	}
	if len(request.Tools) > 0 && len(tools) == 0 {
		http.Error(w, "Invalid Request, tools all convert failed", http.StatusBadRequest)
		return
	}

	stream, err := c.ChatCompleteStream(context.Background(), m,
		// openrouter.WithModel("functiongemma:270m"),
		// openrouter.WithModel("deepseek-r1:7b"),
		openrouter.WithModel("granite4.1:3b"),
		// openrouter.WithModel("deepseek-v4-pro"),
		openrouter.WithTools(tools...),
	)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer stream.Close()

	for c := range stream.Events() {
		j, err := json.Marshal(c)
		if err != nil {
			fmt.Printf("json marshal error: %v\n", err)
			continue
		}
		fmt.Fprintf(w, "%s\n", j)
		flusher.Flush()
	}

	err = stream.Err()
	if err != nil {
		fmt.Printf("end meet error: %v\n", err)
	}
}
