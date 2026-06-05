package api

type RoboDoer interface{
	Do(...string)
}