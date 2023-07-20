package model

import "gorm.io/gorm"

type Questions struct {
	gorm.Model
	ID       int    `json:"id"`
	Question string `json:"question" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Options  string `json:"options,omitempty"` // Slice untuk menyimpan opsi jawaban pada tipe pertanyaan pilihan ganda
	Answer   string `json:"answer,omitempty"`  // String untuk menyimpan jawaban pada tipe pertanyaan isian
	IsTrue   int    `json:"is_true"`           // Boolean untuk menyimpan jawaban benar atau salah pada tipe pertanyaan true/false
}
