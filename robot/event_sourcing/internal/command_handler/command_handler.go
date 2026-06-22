package command_handler

import (
	"clean-architecture/event_sourcing/internal/domain"
	"clean-architecture/event_sourcing/internal/event_store"
	"clean-architecture/event_sourcing/internal/robo"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type commandHandler struct {
	eventStore event_store.EventStore
}

func New(storage event_store.EventStore) commandHandler {
	return commandHandler{
		eventStore: storage,
	}
}

func (ch commandHandler) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	commands := strings.Split(strings.TrimSpace(string(body)), "\n")

	// general command validation (simplified)
	if len(commands) > 2 {
		return
	}

	var command domain.Commander

	switch commands[0] {
	case "move":
		value, _ := strconv.Atoi(commands[1])
		command = robo.NewMoveCommand(value)
	case "turn":
		value, _ := strconv.Atoi(commands[1])
		command = robo.NewTurnCommand(value)
	case "set":
		command = robo.NewSetCommand(domain.State(commands[1]))
	case "start":
		command = robo.StartCommand{}
	case "stop":
		command = robo.StopCommand{}
	default:
		// log error
		return
	}

	// rebuild state every time
	robotState := ch.eventStore.Status()

	command.Execute(robotState)

	ch.eventStore.Produce(event_store.NewEvent(command))
}
