package text_to_speech

type TTSQueryParams struct {
	Voice       string `url:"voice"`
	WatsonToken string `url:"watson-token"`
}

type TTSRequest struct {
	Text   string `json:"text"`
	Accept string `json:"accept"`
}

