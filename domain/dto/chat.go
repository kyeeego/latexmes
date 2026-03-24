package dto

type ChatUpdateType int

const (
	TypeMessage ChatUpdateType = iota
	TypeMsgRead
)

type ChatUpdate struct {
	Type      ChatUpdateType
	Content   string
	Recipient uint
	Sender    uint
}
