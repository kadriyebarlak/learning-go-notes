package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	apiKey := ""

	file, err := os.Open("thoughts.txt")
	if err != nil {
		fmt.Println("Failed to open file: ", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sentences []string

	for scanner.Scan() {
		sentences = append(sentences, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading file: ", err)
	}

	for _, sentence := range sentences {
		sentiment, err := analyzeSentimentOpenAI(sentence, apiKey)
		if err != nil {
			fmt.Println("Error analyzing sentiment:", err)
			continue
		}
		fmt.Printf("Sentence: \"%s\" -> Sentiment: %s\n", sentence, sentiment)
	}
}

func analyzeSentimentOpenAI(sentence, apiKey string) (string, error) {
	// Hugging Face API URL
	url := "https://api-inference.huggingface.co/models/nlptown/bert-base-multilingual-uncased-sentiment"

	requestBody, _ := json.Marshal(map[string]string{
		"inputs": sentence,
	})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Raw Response:", string(body)) // Print the full response body
	/*
		Raw Response: [[{"label":"5 stars","score":0.8546807765960693},
		{"label":"4 stars","score":0.11446134746074677},
		{"label":"3 stars","score":0.020417453721165657},
		{"label":"1 star","score":0.0054611158557236195},
		{"label":"2 stars","score":0.004979335702955723}]]
	*/

	var result [][]struct {
		Label string  `json:"label"`
		Score float64 `json:"score"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if len(result) > 0 && len(result[0]) > 0 {
		bestSentiment := result[0][0].Label
		return bestSentiment, nil
	}

	return "Unknown", nil

}
