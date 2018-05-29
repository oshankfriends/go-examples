package speech_to_text

type SpeechRecognizeResponse struct {
	Results     []Result `json:"results,omitempty"`
	ResultIndex int      `json:"result_index,omitempty"`
	Warnings    []string `json:"warnings,omitempty"`
	Error       string   `json:"error,omitempty"`
}

type Result struct {
	Final            bool                     `json:"final"`
	Alternatives     []Alternative            `json:"alternatives,omitempty"`
	KeywordsResult   KeywordResults           `json:"keyword_results,omitempty"`
	WordAlternatives []WordAlternativeResults `json:"word_alternatives,omitempty"`
}

type Alternative struct {
	Transcript     string          `json:"transcript,omitempty"`
	Confidence     float64         `json:"confidence,omitempty"`
	Timestamps     [][]interface{} `json:"timestamps,omitempty"`
	WordConfidence []string        `json:"word_confidence,omitempty"`
}

type KeywordResults struct {
	Keyword []KeywordResult `json:"keyword"`
}

type KeywordResult struct {
	NormalizedText string  `json:"normalized_text"`
	StartTime      int64   `json:"start_time"`
	EndTime        int64   `json:"end_time"`
	Confidence     float64 `json:"confidence"`
}

type WordAlternativeResults struct {
	StartTime    int `json:"start_time"`
	EndTime      int `json:"end_time"`
	Alternatives []WordAlternativeResult
}

type WordAlternativeResult struct {
	Confidence float64 `json:"confidence"`
	Word       string  `json:"word"`
}

type StateReply struct {
	State string `json:"state"`
}
