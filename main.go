package main

import (
	"bufio"
	"fmt"
	"os"
	"quiz_master/helper"
	"quiz_master/usecase"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Welcome to Quiz Master\n")

	count := 1
	for i := 0; i < count; {
		fmt.Print("$ ")
		text, _ := reader.ReadString('\n')
		command := usecase.DefineInput(text)
		if len(command) > 0 {
			switch command[0] {
			case "create_question":
				data, err := usecase.CreateQuestion(command[1:])
				if err != nil {
					helper.PrintError(err)
				} else {
					helper.PrintCreateQuestion(data)
				}
			case "update_question":
				data, err := usecase.UpdateQuestion(command[1:])
				if err != nil {
					helper.PrintError(err)
				} else {
					helper.PrintUpdateQuestion(data)
				}
			case "answer_question":
				message, err := usecase.AnswerQuestion(command[1:])
				if err != nil {
					helper.PrintError(err)
				} else {
					fmt.Printf("%v\n\n", message)
				}
			case "delete_question":
				id, err := usecase.DeleteQuestion(command[1:])
				if err != nil {
					helper.PrintError(err)
				} else {
					fmt.Printf("Question no %d was deleted!\n\n", id)
				}
			case "question":
				data, err := usecase.QuestionSingle(command[1:])
				if err != nil {
					helper.PrintError(err)
				} else {
					helper.PrintQuestionSingle(data)
				}
			case "questions":
				data, err := usecase.QuestionList()
				if err != nil {
					helper.PrintError(err)
				} else {
					helper.PrintQuestionList(data)
				}
			case "help":
				helper.PrintHelp()
			case "exit":
				i = 2
			default:
				helper.PrintUnknownCommand(command[0])
			}
		} else {
			helper.PrintUnknownCommand(command[0])
		}
	}

}
