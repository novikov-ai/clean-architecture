package api

import (
	"clean-architecture/robot/functional_di_v2/internal/domain"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func New(execute Execute) http.HandlerFunc {
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

			cmd := domain.Cmd{Name: domain.Command(words[0])}
			arg := ""
			if len(words) > 1 {
				arg = words[1]
			}

			switch cmd.Name {
			case domain.Move:
				cmd.Steps, _ = strconv.Atoi(arg)
			case domain.Turn:
				cmd.Angle, _ = strconv.Atoi(arg)
			case domain.Set:
				cmd.State = domain.State(arg)
			}

			execute(cmd)
		}
	}
}
