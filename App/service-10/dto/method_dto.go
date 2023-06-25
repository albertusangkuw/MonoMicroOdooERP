package dto

import "time"

type resultDTO struct{
	ID      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  interface{} `json:"result"`
}

func GetResultDTO(id int,result  interface{}) resultDTO{
	return resultDTO{
		ID: id,
		Jsonrpc: "2.0",
		Result: result,
	}
}


type WriteDTO struct{
	ID      int    `json:"id"`
	Params  struct {
		Args   []interface{} `json:"args"`
		Kwargs struct {
			Context interface{}`json:"context"`
			Domain interface{}`json:"domain"`
			Field interface{}`json:"domain"`
		} `json:"kwargs"`
	} `json:"params"`
}

type ReadDTO struct{
	ID      int    `json:"id"`
	Params  struct {
		Args   []interface{}
		Kwargs struct {
			Context interface{}`json:"context"`
			Domain interface{}`json:"domain"`
			Field interface{}`json:"domain"`
		} `json:"kwargs"`
	} `json:"params"`
}

func TimeToString(t time.Time) string{
	return t.Format("2006-01-02 15:04:05")
}