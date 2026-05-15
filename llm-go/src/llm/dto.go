package llm

type Request struct {
	Messages []map[string]any `json:"messages"`
	Tools    []map[string]any `json:"tools"`
}
