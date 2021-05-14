package discord

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

/*
	SendWebhook
	Send a webhook message to discord.

	@params {string} webhook - The webhook to send the message to.
	@params {string} message - The content of the message to be sent.

	@returns {string} message - The message if a error occurred.
	@returns {bool}   success - If the request was successful or not.
*/
func SendWebhook(webhook string, message string)(string, bool){

	formattedMessage := fmt.Sprintf(`{"content":"%v"}`, message)

	// If it contains @@@ we want to override and split it into embed and content.
	if strings.Contains(message, "@!@!@") {

		// Split the message up into two chunks.
		v := strings.Split(message, "@!@!@")

		// If we have big enough generate the value.
		if len(v) > 1 {
			formattedMessage = fmt.Sprintf(`{"content":"%v", "embeds": %v }`, v[0], v[1])
		}
	}

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
		return "discord_rejection", false
	}

	return "", true
}