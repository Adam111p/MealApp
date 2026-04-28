package geminiApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Struktury dla API Google
type EmbeddingRequest struct {
	Content              Content `json:"content"`
	OutputDimensionality int     `json:"outputDimensionality"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type EmbeddingResponse struct {
	Embedding struct {
		Values []float32 `json:"values"`
	} `json:"embedding"`
}

func GetEmbedding(apiKey string, text string) ([]float32, error) {
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-embedding-2:embedContent?key=" + apiKey
	// Przygotowanie danych
	reqBody := EmbeddingRequest{
		Content: Content{
			Parts: []Part{{Text: text}},
		},
		OutputDimensionality: 768,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// Strzał do API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	// Dekodowanie wyniku
	var result EmbeddingResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Embedding.Values, nil
}
