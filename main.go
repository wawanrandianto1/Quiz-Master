package main

import (
	"bufio"
	"fmt"
	"os"
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
				usecase.CreateQuestion(command[1:])
			case "update_question":
				usecase.UpdateQuestion(command[1:])
			case "answer_question":
				usecase.AnswerQuestion(command[1:])
			case "delete_question":
				usecase.DeleteQuestion(command[1:])
			case "question":
				usecase.QuestionSingle(command[1:])
			case "questions":
				usecase.QuestionList()
			case "help":
				usecase.ShowHelp()
			case "exit":
				i = 2
			default:
				fmt.Printf("Quiz Master: no such command '%v'\nSee 'help' for information on a specific command.\n\n", command[0])
			}
		} else {
			fmt.Printf("Quiz Master: no such command '%v'\nSee 'help' for information on a specific command.\n\n", text)
		}
	}

}
