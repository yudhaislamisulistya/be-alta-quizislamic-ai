package controller

import (
	"fmt"
	"net/http"
	"project/config"
	"project/lib/database"
	"project/model"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateQuizController(c echo.Context) error {
	amount := c.FormValue("amount")
	instruction := c.FormValue("instruction")
	typeQuestion := c.FormValue("type")
	quiz := model.Quiz{}
	c.Bind(&quiz)
	quizData := ""
	amountInt, err := strconv.Atoi(amount)
	fmt.Println(amountInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}
	if amountInt == 0 || amountInt > 5 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "amount must be 1 - 5",
		})
	}

	if !strings.Contains(instruction, "islam") || strings.Contains(instruction, "Islam") {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "pertanyaan harus seputar islam",
		})
	}

	if typeQuestion == "multiple_choice" {
		quizData := database.OpenAICreateQuiz(typeQuestion, amountInt, instruction)
		quizDataSplit, ok := quizData.(map[string]interface{})["data"].(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    "500",
				"message": "error",
			})
		}
		quizDataSplitPart := strings.SplitN(quizDataSplit, "\n", 2)
		quiz.Question = quizDataSplitPart[0]
		quiz.Options = quizDataSplitPart[1]
		quiz.Answer = "-"
		quiz.IsTrue = 9

	} else if typeQuestion == "true_false" {
		quizData := database.OpenAICreateQuiz(typeQuestion, amountInt, instruction)
		quizDataSplit, ok := quizData.(map[string]interface{})["data"].(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    "500",
				"message": "error",
			})
		}
		quizDataSplitPart := strings.Split(quizDataSplit, "\n")
		quiz.Question = quizDataSplitPart[0]
		totalLine := len(quizDataSplitPart)
		correctAnswer := quizDataSplitPart[totalLine-1]
		fmt.Println(correctAnswer)
		if strings.Contains(strings.ToLower(correctAnswer), "true") || strings.Contains(strings.ToLower(correctAnswer), "benar") {
			fmt.Println("masuk true")
			quiz.IsTrue = 1
		} else if strings.Contains(strings.ToLower(correctAnswer), "false") || strings.Contains(strings.ToLower(correctAnswer), "salah") {
			fmt.Println("masuk false")
			quiz.IsTrue = 0
		}
		quiz.Options = "-"
		quiz.Answer = "-"
	} else if typeQuestion == "fill_in" {
		quizData := database.OpenAICreateQuiz(typeQuestion, amountInt, instruction)
		quizDataSplit, ok := quizData.(map[string]interface{})["data"].(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    "500",
				"message": "error",
			})
		}
		quizDataSplitPart := strings.SplitN(quizDataSplit, "\n", 2)
		quiz.Question = quizDataSplitPart[0]
		quiz.Answer = quizDataSplitPart[1]
		quiz.Options = "-"
		quiz.IsTrue = 9
	} else {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "type question must be multiple_choice, true_false, or fill_in",
		})
	}

	quiz.Type = typeQuestion
	result := config.DB.Save(&quiz)
	errResult := result.Error

	if errResult != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": errResult.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     "200",
		"message":  "success create quiz",
		"data":     quiz,
		"quizData": quizData,
	})
}
