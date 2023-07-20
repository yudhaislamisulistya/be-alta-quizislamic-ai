package model

type QuizCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}
