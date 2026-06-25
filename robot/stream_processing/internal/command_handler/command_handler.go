package command_handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"clean-architecture/robot/stream_processing/internal/domain"
	"clean-architecture/robot/stream_processing/internal/event_store"
)

type CommandHandler struct {
	robotID string
	store   *event_store.EventStore
}

func New(robotID string, store *event_store.EventStore) *CommandHandler {
	return &CommandHandler{robotID: robotID, store: store}
}

func (ch *CommandHandler) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	parts := strings.Fields(strings.TrimSpace(string(body)))
	if len(parts) == 0 {
		http.Error(w, "empty command", http.StatusBadRequest)
		return
	}

	cmd, err := parseCommand(parts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ch.store.Append(cmd.ToEvents(ch.robotID))
	w.WriteHeader(http.StatusOK)
}

func parseCommand(parts []string) (domain.Command, error) {
	switch parts[0] {
	case "move":
		if len(parts) < 2 {
			return nil, fmt.Errorf("move requires distance")
		}
		val, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid distance: %s", parts[1])
		}
		return domain.MoveCommand{Distance: val}, nil

	case "turn":
		if len(parts) < 2 {
			return nil, fmt.Errorf("turn requires angle")
		}
		val, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid angle: %s", parts[1])
		}
		return domain.TurnCommand{Angle: val}, nil

	case "set":
		if len(parts) < 2 {
			return nil, fmt.Errorf("set requires mode")
		}
		mode, err := parseMode(parts[1])
		if err != nil {
			return nil, err
		}
		return domain.SetStateCommand{Mode: mode}, nil

	case "start":
		return domain.StartCommand{}, nil

	case "stop":
		return domain.StopCommand{}, nil

	default:
		return nil, fmt.Errorf("unknown command: %s", parts[0])
	}
}

func parseMode(s string) (domain.CleaningMode, error) {
	switch strings.ToLower(s) {
	case "water", "1":
		return domain.Water, nil
	case "soap", "2":
		return domain.Soap, nil
	case "brush", "3":
		return domain.Brush, nil
	default:
		return 0, fmt.Errorf("unknown mode: %s", s)
	}
}
