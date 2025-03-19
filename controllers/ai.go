package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const geminiAPIKey = "GEMINI_API_KEY"
const geminiAPIURL = "GEMINI_API_URL"

type SuggestRequest struct {
	Description string        `json:"description"`
	Tasks       []interface{} `json:"tasks"`
}

func SuggestTasks(c *gin.Context) {
	var input SuggestRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Build prompt with all task info
	prompt := "Suggest subtasks or improvements for: " + input.Description + ".\nExisting tasks:\n"
	for _, task := range input.Tasks {
		taskJSON, _ := json.Marshal(task)
		prompt += string(taskJSON) + "\n"
	}

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": prompt},
				},
			},
		},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to marshal payload"})
		return
	}

	req, err := http.NewRequest("POST", geminiAPIURL+"?key="+geminiAPIKey, bytes.NewBuffer(payloadBytes))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to call Gemini API: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		c.JSON(500, gin.H{"error": "Gemini API error", "status": resp.Status, "raw_response": string(body)})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read response"})
		return
	}

	var result struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		c.JSON(500, gin.H{"error": "Failed to parse Gemini response"})
		return
	}

	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		suggestions := result.Candidates[0].Content.Parts[0].Text
		c.JSON(200, gin.H{"suggestions": suggestions})
	} else {
		c.JSON(500, gin.H{"error": "No suggestions returned from Gemini", "raw_response": string(body)})
	}
}
