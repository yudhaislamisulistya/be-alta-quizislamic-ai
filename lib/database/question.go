package database

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"project/config"
	"project/model"
	"strconv"

	"github.com/sashabaranov/go-openai"
)

func OpenAICreateQuestion(typeQuestion string, amount int, instruction string) interface{} {
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

func GetQuestions() (interface{}, error) {
	questions := []model.Questions{}
	result := config.DB.Find(&questions)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return questions, nil
}

func GetByTypeQuestions(typeQuestions string) (interface{}, error) {
	questions := []model.Questions{}
	result := config.DB.Where("type = ?", typeQuestions).Find(&questions)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return questions, nil

}

func CreateQuestionByMultipleChoice(userId string, quizId string, question string, optionsA string, optionsB string, optionsC string, optionsD string, answer string, point string) (interface{}, error) {

	userIdInt, errUserIdInt := strconv.Atoi(userId)
	if errUserIdInt != nil {
		return nil, errors.New("user id should be integer")
	}

	quizIdInt, errQuizIdInt := strconv.Atoi(quizId)
	if errQuizIdInt != nil {
		return nil, errors.New("quiz id should be integer")
	}

	pointInt, errPointInt := strconv.Atoi(point)
	if errPointInt != nil {
		return nil, errors.New("point should be integer")
	}

	options := optionsA + "\n" + optionsB + "\n" + optionsC + "\n" + optionsD + "\nJawaban: " + answer

	questionData := model.Questions{
		UserID:   userIdInt,
		QuizID:   quizIdInt,
		Question: question,
		Type:     "multiple_choice",
		Options:  options,
		Answer:   "-",
		IsTrue:   9,
		Point:    pointInt,
	}

	result := config.DB.Create(&questionData)
	errResult := result.Error

	if errResult != nil {
		return nil, errResult
	}

	return questionData, nil
}

func CreateQuestionByTrueFalse(userId string, quizId string, question string, isTrue string, point string) (interface{}, error) {

	userIdInt, errUserIdInt := strconv.Atoi(userId)
	if errUserIdInt != nil {
		return nil, errors.New("user id should be integer")
	}

	quizIdInt, errQuizIdInt := strconv.Atoi(quizId)
	if errQuizIdInt != nil {
		return nil, errors.New("quiz id should be integer")
	}

	if isTrue == "true" {
		isTrue = "1"
	} else if isTrue == "false" {
		isTrue = "0"
	} else {
		return nil, errors.New("is true should be true or false")
	}

	isTrueInt, errIsTrueInt := strconv.Atoi(isTrue)
	if errIsTrueInt != nil {
		return nil, errors.New("is true should be integer")
	}

	pointInt, errPointInt := strconv.Atoi(point)
	if errPointInt != nil {
		return nil, errors.New("point should be integer")
	}

	questionData := model.Questions{
		UserID:   userIdInt,
		QuizID:   quizIdInt,
		Question: question,
		Type:     "true_false",
		Options:  "-",
		Answer:   "-",
		IsTrue:   isTrueInt,
		Point:    pointInt,
	}

	result := config.DB.Create(&questionData)
	errResult := result.Error

	if errResult != nil {
		return nil, errResult
	}

	return questionData, nil
}

func CreateQuestionByFillIn(userId string, quizId string, question string, answer string, point string) (interface{}, error) {

	userIdInt, errUserIdInt := strconv.Atoi(userId)
	if errUserIdInt != nil {
		return nil, errors.New("user id should be integer")
	}

	quizIdInt, errQuizIdInt := strconv.Atoi(quizId)
	if errQuizIdInt != nil {
		return nil, errors.New("quiz id should be integer")
	}

	pointInt, errPointInt := strconv.Atoi(point)
	if errPointInt != nil {
		return nil, errors.New("point should be integer")
	}

	questionData := model.Questions{
		UserID:   userIdInt,
		QuizID:   quizIdInt,
		Question: question,
		Type:     "fill_in",
		Options:  "-",
		Answer:   answer,
		IsTrue:   9,
		Point:    pointInt,
	}

	result := config.DB.Create(&questionData)
	errResult := result.Error

	if errResult != nil {
		return nil, errResult
	}

	return questionData, nil
}
