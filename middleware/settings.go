package middleware

var SuccessMessageFormat    = "{\"Success\": true,  \"Result\": \"%v\"}"
var SuccessMessageRawFormat = "{\"Success\": true,  \"Result\": %v}"
var ErrorMessageFormat      = "{\"Success\": false, \"Result\": \"%v\"}"
var GenericMessageFormat    = "{\"Success\": %v, \"Result\": \"%v\"}"
var GenericMessageRawFormat = "{\"Success\": %v, \"Result\": %v}"