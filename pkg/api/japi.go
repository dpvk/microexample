package api

type JapiRequest struct {
	Client string `json:"client"`
}

type JapiResponse struct {
	Client  string `json:"client"`
	Adapter string `json:"adapter"`
	Orch    string `json:"orch"`
	Dummy   string `json:"dummy"`
	Db      string `json:"db"`
	Data    string `json:"data"`
}
