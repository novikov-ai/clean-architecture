package domain

type Command string

const (
	Move  Command = "move"
	Turn  Command = "turn"
	Set   Command = "set"
	Start Command = "start"
	Stop  Command = "stop"
)

