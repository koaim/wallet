package state

import "github.com/NicoNex/echotron/v3"

type handler = func(msg echotron.Message) error

type State[T comparable] struct {
	ID T

	textHandlers       map[string]handler
	callbackHandlers   map[string]handler
	defaultTextHandler handler
}

func New[T comparable](id T) State[T] {
	return State[T]{ID: id, textHandlers: map[string]handler{}, callbackHandlers: map[string]handler{}}
}

func (s *State[T]) On(msg string, h handler) {
	s.textHandlers[msg] = h
}

func (s *State[T]) OnText(h handler) {
	s.defaultTextHandler = h
}
