package nluclient

type QueryRequest struct {
	Input    InputStruct `json:"input,omitempty"`
	Debug    bool        `json:"debug,omitempty"`
	Location interface{} `json:"location,omitempty"`
	User     UserStruct  `json:"user,omitempty"`
	Device   interface{} `json:"device,omitempty"`
	App      interface{} `json:"app,omitempty"`
}

type InputStruct struct {
	Intent            string      `json:"intent,omitempty"`
	Text              string      `json:"text,omitempty"`
	Fields            interface{} `json:"fields,omitempty"`
	DisplayContext    interface{} `json:"display_context,omitempty"`
	BackgroundContext interface{} `json:"background_context,omitempty"`
}

type UserStruct struct {
	Id           string      `json:"id,omitempty"`
	Uuid         string      `json:"uuid,omitempty"`
	AccessTokens interface{} `json:"access_tokens,omitempty"`
}

type QueryResponse struct {
	AppName        string      `json:"app_name"`
	AppVersion     string      `json:"app_version"`
	Awaiting       interface{} `json:"awaiting"`
	ConversationId string      `json:"conversation_id"`
	DebugInfo      interface{} `json:"debug_info"`
	Fields         interface{} `json:"fields"`
	Language       string      `json:"language"`
	ResponseId     string      `json:"response_id"`
	ResponseText   string      `json:"response_text"`
	RuleType       string      `json:"rule_type"`
	SkillData      interface{} `json:"skill_data"`
	Intent         string      `json:"intent"`
	Version        string      `json:"version"`
}
