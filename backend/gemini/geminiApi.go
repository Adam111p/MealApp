package geminiApi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/genai"
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

// Struktura odpowiedzi, którą chcemy dostać od AI
type SearchIntent struct {
	CleanQuery  string   `json:"clean_query"`
	Tags        []string `json:"tags"`
	MaxPrice    float32  `json:"maxPrice"`
	WithoutTags []string `json:"withoutTags"`
}

func AnalyzeIntentWithGemini(ctx context.Context, apiKey string, userInput string, availableTags []string) (*SearchIntent, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}

	tagsList := strings.Join(availableTags, ", ")

	// Budowanie "System Promptu"
	prompt := fmt.Sprintf(`
		Jesteś asystentem bazy danych restauracji. Twoim zadaniem skatalogowanie menu i tagowanie posiłków  na format JSON.
		
		DOSTĘPNE TAGI: [%s]
		
		ZASADY:
		1. Wyodrębnij słowa kluczowe do wyszukiwania wektorowego (clean_query).
		2. Jeśli użytkownik wskazuje preferencje , składniki , typ dania (np. "wege", "ostre", "pomidory" ,"kebab"), dodaj tagi do "tags".
		3. Staraj się zunifikować tagi , makaron i pasta to  to samo
		4. Używaj możliwie tagów z listy DOSTĘPNE TAGI, jeżeli to koniecznie dodaj nowe.
		
		ZAPYTANIE UŻYTKOWNIKA: "%s"
	`, tagsList, userInput)

	// Konfiguracja modelu na tryb JSON
	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
	}

	result, err := client.Models.GenerateContent(ctx, "gemini-3.1-flash-lite-preview", genai.Text(prompt), config)
	if err != nil {
		return nil, err
	}

	// Pobieranie tekstu z odpowiedzi
	var intent SearchIntent
	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		jsonText := result.Candidates[0].Content.Parts[0].Text

		// Parsowanie JSONa na strukturę Go
		err := json.Unmarshal([]byte(jsonText), &intent)
		if err != nil {
			return nil, fmt.Errorf("failed to parse AI response: %v", err)
		}
	}

	return &intent, nil
}

func AnalyzeSearchIntentWithGemini(ctx context.Context, apiKey string, userInput string, availableTags []string) (*SearchIntent, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}

	// Przygotowanie listy tagów jako string
	tagsList := strings.Join(availableTags, ", ")

	// Budowanie "System Promptu"
	prompt := fmt.Sprintf(`
		Jesteś asystentem bazy danych restauracji. Twoim zadaniem jest przetworzenie zapytania użytkownika na format JSON.
		
		DOSTĘPNE TAGI: [%s]
		
		ZASADY:
		1. Wyodrębnij słowa kluczowe do wyszukiwania wektorowego (clean_query).
		2. Jeśli użytkownik wskazuje preferencje , składniki , typ dania (np. "wege", "ostre", "pomidory" ,"kebab"), dodaj tagi do "tags".
		3. Jeśli użytkownik wskazuje że czegoś nie chce (np. "nie chce mięsa"), dodaj tagi do "withoutTags".
		4. Używaj tagów z listy DOSTĘPNE TAGI,
		5. kreatywnie dopasuj tagi np. jeżeli jest ptak w zdaniu,a DOSTĘPNE TAGI to "kurczak" użyj tagu kurczak.
		6. jeżeli wspomni o cenie to takie to do 30 zł a bardzo drogie to powyżej 100 zł dodaj kwotę do "maxPrice"
		ZAPYTANIE UŻYTKOWNIKA: "%s"
	`, tagsList, userInput)

	// Konfiguracja modelu na tryb JSON
	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
	}

	result, err := client.Models.GenerateContent(ctx, "gemini-3.1-flash-lite-preview", genai.Text(prompt), config)
	if err != nil {
		return nil, err
	}

	// Pobieranie tekstu z odpowiedzi
	var intent SearchIntent
	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		jsonText := result.Candidates[0].Content.Parts[0].Text

		// Parsowanie JSONa na strukturę Go
		err := json.Unmarshal([]byte(jsonText), &intent)
		if err != nil {
			return nil, fmt.Errorf("failed to parse AI response: %v", err)
		}
	}

	return &intent, nil
}
