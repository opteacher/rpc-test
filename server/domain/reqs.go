package domain

type Reqs struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}
