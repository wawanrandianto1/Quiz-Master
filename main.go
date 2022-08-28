package main

import (
	"fmt"
	"os"
	"quiz_master/repository"
	ucase "quiz_master/usecase"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(repository.ErrBadParamInput)
	}

	inputs := os.Args[1:]
	switch inputs[0] {
	case "create_question":
		ucase.CreateQuestion(inputs[1:])
	case "update_question":
		ucase.UpdateQuestion(inputs[1:])
	case "answer_question":
		ucase.AnswerQuestion(inputs[1:])
	case "delete_question":
		ucase.DeleteQuestion(inputs[1:])
	case "question":
		ucase.QuestionSingle(inputs[1:])
	default:
		// question list
		ucase.QuestionList()
	}
}
