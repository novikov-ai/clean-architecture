package api

import (
	"clean-architecture/robot/functional_di/internal/domain"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func New(
	move domain.ActionCommand,
	turn domain.ActionCommand,
	set domain.StateCommand,
	start domain.ModeCommand,
	stop domain.ModeCommand,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		lines := strings.Split(strings.TrimSpace(string(body)), "\n")

		for _, line := range lines {
			words := strings.Fields(line)
			if len(words) == 0 {
				continue
			}

			command := domain.Command(words[0])
			arg := ""
			if len(words) > 1 {
				arg = words[1]
			}

			switch command {
			case domain.Start:
				start()
			case domain.Stop:
				stop()
			case domain.Move:
				steps, _ := strconv.Atoi(arg)
				move(steps)
			case domain.Turn:
				angle, _ := strconv.Atoi(arg)
				turn(angle)
			case domain.Set:
				set(domain.State(arg))
			}
		}
	}
}
