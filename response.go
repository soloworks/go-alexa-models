package alexa

// Response is the response back to the response speech service
type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Body              struct {
		OutputSpeech     *Payload     `json:"outputSpeech,omitempty"`
		Card             *Payload     `json:"card,omitempty"`
		Reprompt         *Reprompt    `json:"reprompt,omitempty"`
		Directives       []Directives `json:"directives,omitempty"`
		ShouldEndSession bool         `json:"shouldEndSession"`
	} `json:"response"`
}

// NewEmptyResponse builds an empty response
func NewEmptyResponse() Response {
	return Response{Version: "1.0"}
}

// NewSimpleTerminateResponse builds an empty response
func NewSimpleTerminateResponse() Response {
	r := NewEmptyResponse()
	r.Body.ShouldEndSession = true
	return r
}

// NewSpeechResponse builds a simple speech response
func NewSpeechResponse(speech string) Response {
	r := NewEmptyResponse()
	r.Body.ShouldEndSession = true
	r.Body.OutputSpeech = &Payload{
		Type: "PlainText",
		Text: speech,
	}

	return r
}

// NewDialogDelegateResponse builds a simple response response to advance to the next step
func NewDialogDelegateResponse() Response {
	r := NewEmptyResponse()
	r.Body.ShouldEndSession = false
	r.Body.Directives = append(r.Body.Directives, Directives{Type: DirectiveTypeDialogDelegate})

	return r
}
