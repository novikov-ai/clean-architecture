package command_handler

import (
	"clean-architecture/robot/stream_processing/internal/domain"
	"clean-architecture/robot/stream_processing/internal/event_store"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type commandHandler struct {
	eventStore    event_store.EventStore
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

	switch commands[0] {
	case "move":
		value, _ := strconv.Atoi(commands[1])
		ch.eventStore.ProduceMoveRequestedEvent(value)
	case "turn":
		value, _ := strconv.Atoi(commands[1])
		ch.eventStore.ProduceTurnRequestedEvent(value)
	case "set":
		ch.eventStore.ProduceSetRequestedEvent(domain.State(commands[1]))
	case "start":
		ch.eventStore.ProduceStartRequestedEvent()
	case "stop":
		ch.eventStore.ProduceStopRequestedEvent()
	default:
		// log error
		return
	}
}
