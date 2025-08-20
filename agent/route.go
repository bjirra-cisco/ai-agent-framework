package agent

type Route struct {
	SelectedAgent   string          `json:"selected_agent"`
	Confidence      float64         `json:"confidence"`
	Reasoning       string          `json:"reasoning"`
	ExtractedParams ExtractedParams `json:"extracted_params"`
	Endpoint        string          `json:"endpoint"`
}

type ExtractedParams struct {
	Intent  string                 `json:"intent"`
	Filters map[string]interface{} `json:"filters"`
	Options map[string]interface{} `json:"options"`
}
