package dto

type ChatUpdateType int

const (
	TypeMessage ChatUpdateType = iota
)

type ChatUpdate struct {
	Type       ChatUpdateType
	Content    string
	Connection uint
}

type MessageRequest struct {
	To uint `json:"to"`
	// PubKey []byte `json:"pub_key"`
}

type MessageRequestResponse struct {
	IsAccepted bool `json:"is_accepted"`
	ReqID      uint `json:"req_id"`
	// nil если IsAccepted == false
	// PubKey []byte `json:"pub_key"`
}
