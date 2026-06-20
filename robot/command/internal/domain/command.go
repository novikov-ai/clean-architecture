package domain

type Commander interface {
	Execute(Robot) Robot
}
