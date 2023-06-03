package aisvc

type CompileReq struct {
	TypeOfAi string `json:"type_of_ai"`
	IDs      []int  `json:"ids"`
}

type QueryReq struct {
	Query string `json:"query"`
}
