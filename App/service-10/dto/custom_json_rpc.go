package dto

type CustomJSONRPCContainer struct {
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Model  string  `json:"model"`
		Method string  `json:"method"`
		Kwargs struct {
			Context struct{
				Lang  string  `json:"lang"`
			} `json:"context"`
			Domain interface{}`json:"domain"`
			Field interface{}`json:"domain"`
		} `json:"kwargs"`
	} `json:"params"`
}
