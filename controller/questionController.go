package controller

import (
	"fmt"
	"net/http"
	"project/config"
	"project/lib/database"
	v "project/lib/validator"
	"project/model"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func CreateQuestionController(c echo.Context) error {
	userId := c.FormValue("user_id")
	quizId := c.FormValue("quiz_id")
	amount := c.FormValue("amount")
	instruction := c.FormValue("instruction")
	typeQuestion := c.FormValue("type")
	question := model.Questions{}
	c.Bind(&question)

	_, errValidator := v.QuestionValidator(question)

	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": errValidator.Error(),
		})
	}

	questionData := ""
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

	userIdInt, errResultAtoi := strconv.Atoi(userId)

	if errResultAtoi != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "error convert user id",
		})
	}

	quizIdInt, errResultAtoiQuizID := strconv.Atoi(quizId)
	fmt.Println("Quiz ID : ", quizIdInt)

	if errResultAtoiQuizID != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "error convert quiz id",
		})
	}

	wallet, err := database.GetWallet(uint(userIdInt))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	balance := wallet.(model.Wallet).Balance

	if balance == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "balance is not enough",
		})
	}

	if balance < int64(amountInt) {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "balance is not enough",
		})
	}

	if !strings.Contains(instruction, "islam") || strings.Contains(instruction, "Islam") {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "pertanyaan harus seputar islam",
		})
	}

	if typeQuestion == "multiple_choice" {
		questionData := database.OpenAICreateQuestion(typeQuestion, amountInt, instruction)
		questionDataSplit, ok := questionData.(map[string]interface{})["data"].(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    "500",
				"message": "error",
			})
		}
		questionDataSplitPart := strings.SplitN(questionDataSplit, "\n", 2)
		question.Question = questionDataSplitPart[0]
		question.Options = questionDataSplitPart[1]
		question.Answer = "-"
		question.IsTrue = 9

	} else if typeQuestion == "true_false" {
		questionData := database.OpenAICreateQuestion(typeQuestion, amountInt, instruction)
		questionDataSplit, ok := questionData.(map[string]interface{})["data"].(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    "500",
				"message": "error",
			})
		}
		questionDataSplitPart := strings.Split(questionDataSplit, "\n")
		question.Question = questionDataSplitPart[0]
		totalLine := len(questionDataSplitPart)
		correctAnswer := questionDataSplitPart[totalLine-1]
		fmt.Println(correctAnswer)
		if strings.Contains(strings.ToLower(correctAnswer), "true") || strings.Contains(strings.ToLower(correctAnswer), "benar") {
			fmt.Println("masuk true")
			question.IsTrue = 1
		} else if strings.Contains(strings.ToLower(correctAnswer), "false") || strings.Contains(strings.ToLower(correctAnswer), "salah") {
			fmt.Println("masuk false")
			question.IsTrue = 0
		}
		question.Options = "-"
		question.Answer = "-"
	} else if typeQuestion == "fill_in" {
		questionData := database.OpenAICreateQuestion(typeQuestion, amountInt, instruction)
		questionDataSplit, ok := questionData.(map[string]interface{})["data"].(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    "500",
				"message": "error",
			})
		}
		questionDataSplitPart := strings.SplitN(questionDataSplit, "\n", 2)
		question.Question = questionDataSplitPart[0]
		question.Answer = questionDataSplitPart[1]
		question.Options = "-"
		question.IsTrue = 9
	} else {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "type question must be multiple_choice, true_false, or fill_in",
		})
	}

	question.Type = typeQuestion
	question.UserID = userIdInt
	question.QuizID = quizIdInt
	result := config.DB.Save(&question)
	errResult := result.Error

	if errResult != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": errResult.Error(),
		})
	}

	balance = balance - int64(amountInt)
	resultUpdateWallet, errUpdateWallet := database.UpdateWallet(uint(userIdInt), balance)

	if errUpdateWallet != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "error update wallet",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":         "200",
		"message":      "success create question",
		"data":         question,
		"questionData": questionData,
		"wallet":       resultUpdateWallet,
	})
}

func GetByUserIDQuizIDQuestionController(c echo.Context) error {
	userID := c.Param("user_id")
	userQuizID := c.Param("quiz_id")
	question := []model.Questions{}
	result := config.DB.Where("user_id = ? AND user_quiz = ?", userID, userQuizID).Find(&question)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "error get question",
		})
	}

	if len == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "question not found, please create question first",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get question",
		"data":    question,
	})
}

func GetQuestionsController(c echo.Context) error {
	resultQuestions, err := database.GetQuestions()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "error get questions",
		})
	}

	if resultQuestions == int64(0) {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "questions not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get questions",
		"data":    resultQuestions,
	})
}

func GetByTypeQuestionsController(c echo.Context) error {
	typeQuestion := c.QueryParam("type")
	resultTypeQuestions, err := database.GetByTypeQuestions(typeQuestion)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "error get questions",
		})
	}

	if resultTypeQuestions == int64(0) {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "questions not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get questions",
		"data":    resultTypeQuestions,
	})
}

func CreateQuestionByMultipleChoiceController(c echo.Context) error {
	userId := c.FormValue("user_id")
	quizId := c.FormValue("quiz_id")
	question := c.FormValue("question")
	optionsA := c.FormValue("options_a")
	optionsB := c.FormValue("options_b")
	optionsC := c.FormValue("options_c")
	optionsD := c.FormValue("options_d")
	correctAnswer := c.FormValue("correct_answer")
	point := c.FormValue("point")

	fmt.Println("User ID : ", userId)
	fmt.Println("Quiz ID : ", quizId)
	fmt.Println("Question : ", question)
	fmt.Println("Options A : ", optionsA)
	fmt.Println("Options B : ", optionsB)
	fmt.Println("Options C : ", optionsC)
	fmt.Println("Options D : ", optionsD)
	fmt.Println("Correct Answer : ", correctAnswer)
	fmt.Println("Point : ", point)

	// check all field is not empty
	if userId == "" || quizId == "" || question == "" || optionsA == "" || optionsB == "" || optionsC == "" || optionsD == "" || correctAnswer == "" || point == "" {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "all field is required",
		})
	}

	result, err := database.CreateQuestionByMultipleChoice(userId, quizId, question, optionsA, optionsB, optionsC, optionsD, correctAnswer, point)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    "201",
		"message": "success create question",
		"data":    result,
	})
}

func CreateQuestionByTrueFalseController(c echo.Context) error {
	userId := c.FormValue("user_id")
	quizId := c.FormValue("quiz_id")
	question := c.FormValue("question")
	isTrue := c.FormValue("is_true")
	point := c.FormValue("point")

	if userId == "" || quizId == "" || question == "" || isTrue == "" || point == "" {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "all field is required",
		})
	}

	result, err := database.CreateQuestionByTrueFalse(userId, quizId, question, isTrue, point)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    "201",
		"message": "success create question",
		"data":    result,
	})
}

func CreateQuestionByFillInController(c echo.Context) error {
	userId := c.FormValue("user_id")
	quizId := c.FormValue("quiz_id")
	question := c.FormValue("question")
	answer := c.FormValue("answer")
	point := c.FormValue("point")

	if userId == "" || quizId == "" || question == "" || answer == "" || point == "" {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "all field is required",
		})
	}

	result, err := database.CreateQuestionByFillIn(userId, quizId, question, answer, point)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    "201",
		"message": "success create question",
		"data":    result,
	})
}

func GetSearchQuestionsController(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	questions := []model.Questions{}

	result, err := database.GetSearchQuestions(&questions, keyword)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data search",
		"data":    result,
	})
}

func GetSortQuestionsController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	questions := []model.Questions{}

	result, err := database.GetSortQuestions(&questions, sortBy, order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data sort",
		"data":    result,
	})
}

func GetPaginationQuestionsController(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	questions := []model.Questions{}

	result, err := database.GetPaginationQuestions(&questions, page, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data pagination",
		"data":    result,
	})
}
