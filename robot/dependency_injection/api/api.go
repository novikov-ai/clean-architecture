package api

import (
	"io"
	"net/http"
	"strings"
)

type handler struct{
	robot RoboDoer
}

func New(robot RoboDoer) handler{
	return handler{
		robot: robot,
	}
}

func (rd handler) Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	commands := strings.Split(strings.TrimSpace(string(body)), "\n")
	rd.robot.Do(commands...)
}
