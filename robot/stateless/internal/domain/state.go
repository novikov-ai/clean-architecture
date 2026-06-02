package domain

type State string

const (
	Water State = "water"
	Soap  State = "soap"
	Brush State = "brush"
)

type RoboState struct {
	Position Position `json:"position"`
	Angle    int      `json:"angle"`
	State    State    `json:"state"`
}
