package websocket

type Message struct {
	Type   string `json:"type"`
	Text   string `json:"text"`
	Server string `json:"server"`
}
