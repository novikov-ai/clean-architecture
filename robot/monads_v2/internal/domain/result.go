package domain

type MoveResult string

const (
	MovedOk    MoveResult = "MOVE_OK"
	HitBarrier MoveResult = "HIT_BARRIER"
)

type SetStateResult string

const (
	SetStateOk SetStateResult = "STATE_OK"
	NoWater    SetStateResult = "OUT_OF_WATER"
	NoSoap     SetStateResult = "OUT_OF_SOAP"
)
