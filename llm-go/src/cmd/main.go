package main

import (
	"fmt"
	"llm-go/src/llm"
	"net/http"
)

func main() {
	// http.HandleFunc("/llm", llm.Llm)
	// http.HandleFunc("/ollama", llm.Ollama)
	http.HandleFunc("/ollama_direct", llm.OllamaDirect)

	fmt.Println("http start at 6060")
	http.ListenAndServe(":6060", nil)
}
