package database

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/sashabaranov/go-openai"
)

func OpenAICreateQuiz(typeQuestion string, amount int, instruction string) interface{} {
	token := os.Getenv("OPENAI_API_KEY")
	fmt.Println(token)
	cc := openai.NewClient(token)
	ctx := context.Background()

	content := ""
	if typeQuestion == "multiple_choice" {
		content = `
		Jumlah Soal : ` + fmt.Sprintf("%d", amount) + `
		Intruksi : ` + instruction + `
		Format Pilihan : Pilihan Ganda 4 Pilihan
		Contoh Format Pertanyaan, Jawaban dan Jawaban Benar Seperti Dibawah Ini
		1. BLALA BLALA BLALA BLAA BLAA BLAA...;
		a. AAA;
		b. BBB;
		c. CCC;
		d. DDD;
		Jawaban: b. BBB
		Output To The Point, Tidak Perlu Menggunakan Berikut Ini, Dibawah Ini, dan Sebagainya serta Buatkan soal, Pertanyaan, Soal: , Fokus Ke Pertanyaan dan Jawaban dan Kalimat Awal Merupakan Pertanyaan
		`
	} else if typeQuestion == "true_false" {
		content = `
		Jumlah Soal : ` + fmt.Sprintf("%d", amount) + `
		Intruksi : ` + instruction + `
		Format Pilihan : True or False (Benar atau Salah)
		Contoh Format Pertanyaan, Jawaban dan Jawaban Benar Seperti Dibawah Ini
		1. BLALA BLALA BLALA BLAA BLAA BLAA...;
		a. True
		b. False
		Jawaban: true
		Output To The Point, Tidak Perlu Menggunakan Berikut Ini, Dibawah Ini, dan Sebagainya serta Buatkan soal, Pertanyaan, Soal: , Fokus Ke Pertanyaan dan Jawaban dan Kalimat Awal Merupakan Pertanyaan
		`
	} else if typeQuestion == "fill_in" {
		content = `
		Jumlah Soal : ` + fmt.Sprintf("%d", amount) + `
		Intruksi : ` + instruction + `
		Format Pilihan : Essay
		Contoh Format Pertanyaan, Jawaban dan Jawaban Benar Seperti Dibawah Ini
		1. BLALA BLALA BLALA BLAA BLAA BLAA...;
		Jawaban: Lala Lala Lala Lala Lala
		Output To The Point, Tidak Perlu Menggunakan Berikut Ini, Dibawah Ini, dan Sebagainya serta Buatkan soal, Pertanyaan, Soal: , Fokus Ke Pertanyaan dan Jawaban dan Kalimat Awal Merupakan Pertanyaan
		`
	}

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
		Stream: true,
	}
	stream, err := cc.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return map[string]string{
			"code":    "500",
			"message": err.Error(),
		}
	}
	defer stream.Close()

	resultResponse := ""

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return map[string]interface{}{
				"code":    "200",
				"message": "success get user",
				"data":    resultResponse,
			}

		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return map[string]string{
				"code":    "500",
				"message": err.Error(),
			}
		}
		resultResponse += response.Choices[0].Delta.Content
		fmt.Printf(response.Choices[0].Delta.Content)
	}
}
