package responses

import "server/metadata"

type ErrorResponse struct {
	Error   string  `json:"error"`
	Version float32 `json:"version"`
}

var InvalidPayload = ErrorResponse{Error: "Invalid or unrecognized body or payload format.", Version: metadata.Version}
var InvalidAuthentication = ErrorResponse{Error: "Incorrect username or password, please try again.", Version: metadata.Version}
var Unauthorized = ErrorResponse{Error: "You cannot perform this action, please authenticate first.", Version: metadata.Version}
