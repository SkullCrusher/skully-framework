package sms

import (
	"net/http"
	"net/url"
	"strings"
)

/*
	SendSMS
	Send a new sms message using twilio.

	@params {string} accountSid - Twilio user id.
	@params {string} authToken  - Twilio auth token.
	@params {string} to         - The number we are trying to send the sms to.
	@params {string} from       - The number we are sending from (it's required to specifically say which).
	@params {string} message    - The message we want to send to the user.

	@returns {bool} success
*/
func SendSMS(accountSid string, authToken string, to string, from string, message string) bool {

	// Generate the url.
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	msgData := url.Values{}

	msgData.Set("To",   to)
	msgData.Set("From", from)
	msgData.Set("Body", message)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	// If an error occurred when trying to do the request.
	if err != nil {
		return false
	}

	// If a error status code was provided.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false
	}

	return true
}