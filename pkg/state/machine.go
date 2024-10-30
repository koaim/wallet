package state

import (
	"fmt"

	"github.com/NicoNex/echotron/v3"
)

type Storage[T comparable] interface {
	Current(id int64) (T, error)
	Set(id int64, name T) error
	Clear(id int64) error
}

type Machine[T comparable] struct {
	storage Storage[T]
	states  map[T]State[T]
}

func NewMachine[T comparable](s Storage[T]) *Machine[T] {
	return &Machine[T]{storage: s, states: map[T]State[T]{}}
}

func (m *Machine[T]) Register(states ...State[T]) {
	for _, v := range states {
		m.states[v.ID] = v
	}
}

func (m *Machine[T]) Handle(msg echotron.Message) error {
	stateID, err := m.storage.Current(msg.From.ID)
	if err != nil {
		return fmt.Errorf("get current state: %w", err)
	}

	state := m.states[stateID]

	if msg.Text != "" {
		if h, ok := state.textHandlers[msg.Text]; ok {
			return h(msg)
		}

		if state.defaultTextHandler != nil {
			return state.defaultTextHandler(msg)
		}
	}

	return nil
}
