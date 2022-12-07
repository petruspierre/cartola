package entity

import "errors"

type ActionTableInterface interface {
	Init()
	GetScore(action string) (int, error)
}

type ActionTable struct {
	Action map[string]int
}

func (a *ActionTable) Init() {
	a.Action = map[string]int{
		"yellow_card": -1,
		"red_card":    -3,
		"goal":        6,
	}
}

func (a *ActionTable) GetScore(action string) (int, error) {
	score, ok := a.Action[action]

	if !ok {
		return 0, errors.New("action not found")
	}

	return score, nil
}
