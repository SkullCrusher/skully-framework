package slack

import (
	"bytes"
	"fmt"
	"net/http"
)

/*
	SendWebhook
	Send a webhook message to slack. https://api.slack.com/messaging/webhooks

	@params {string} webhook - The webhook to send the message to.
	@params {string} message - The content of the message to be sent.

	@returns {string} message - The message if a error occurred.
	@returns {bool}   success - If the request was successful or not.
*/
func SendWebhook(webhook string, message string)(string, bool){

	formattedMessage := fmt.Sprintf(`{"text":"%v"}`, message)

	// Create a new request.
	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer([]byte(formattedMessage)))

	// If we weren't able to process the record.
	if err != nil {
		return "network_error", false
	}

	// Set the headers so it's json.
	req.Header.Set("Content-Type", "application/json")

	// Create a new client and execute the network request.
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "unable_to_send", false
	}

	defer resp.Body.Close()

	// Some error so return error.
	if resp.StatusCode >= 300 {
		return "slack_rejection", false
	}

	return "", true
}