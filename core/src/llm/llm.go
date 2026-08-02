package llm

import (
	"net/http"
)

func Llm(w http.ResponseWriter, r *http.Request) {
	// defer r.Body.Close()
	// req := make(map[string]any)
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, "Invalid Request", http.StatusBadRequest)
	// 	return
	// }

	// api := openai.New(
	// 	"todo", "deepseek-v4-pro",
	// ).WithEndpoint("https://api.deepseek.com/chat/completions", "DeepSeek")
	// llmapi := llms.New(api)

	// w.Header().Add("Content-Type", "text/event-stream")
	// w.Header().Add("Cache-Control", "no-cache")
	// w.Header().Add("Connection", "keep-alive")
	// flusher, ok := w.(http.Flusher)
	// if !ok {
	// 	http.Error(w, "streaming unsupported", http.StatusInternalServerError)
	// 	return
	// }

	// for update := range llmsUpdateChan {
	// 	switch casted := update.(type) {
	// 	case llms.TextUpdate:
	// 		fmt.Fprintf(w, "data: %s\n\n", casted.Text)
	// 		flusher.Flush()
	// 	}
	// }

	// if err := llmapi.Err(); err != nil {
	// 	fmt.Fprintf(w, "data: [ERROR] %v\n\n", err)
	// 	flusher.Flush()
	// }
}
