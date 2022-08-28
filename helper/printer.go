package helper

import (
	"fmt"
	"quiz_master/repository"
)

func PrintHelp() {
	fmt.Printf("Command                                   | Description\n")
	fmt.Printf("---------------------------------------------------------------\n")
	fmt.Printf("create_question <no> <question> <answer>  | Creates a question\n")
	fmt.Printf("update_question <no> <question> <answer>  | Updates a question\n")
	fmt.Printf("answer_question <no> <answer>             | Answer  a question\n")
	fmt.Printf("delete_question <no>                      | Deletes a question\n")
	fmt.Printf("question <no>                             | Shows   a question\n")
	fmt.Printf("questions                                 | Shows question list\n")
	fmt.Println()
}

func PrintUnknownCommand(str string) error {
	fmt.Printf("Quiz Master: no such command '%v'\nSee 'help' for information on a specific command.\n\n", str)
	return nil
}

func PrintError(err error) error {
	fmt.Println(err)
	fmt.Println()
	return nil
}

func PrintQuestionList(data []repository.Quiz) error {
	fmt.Println("No | Question | Answer")
	for _, data := range data {
		fmt.Printf("%d  \"%v\" %d \n", data.Id, data.Question, data.Answer)
	}
	fmt.Println()
	return nil
}

func PrintQuestionSingle(data repository.Quiz) error {
	fmt.Printf("Q: %v \n", data.Question)
	fmt.Printf("A: %d \n\n", data.Answer)
	return nil
}

func PrintCreateQuestion(data repository.Quiz) error {
	fmt.Printf("Question no %d created:\n", data.Id)
	fmt.Printf("Q: %v \n", data.Question)
	fmt.Printf("A: %d \n\n", data.Answer)
	return nil
}

func PrintUpdateQuestion(data repository.Quiz) error {
	fmt.Printf("Question no %d updated:\n", data.Id)
	fmt.Printf("Q: %v \n", data.Question)
	fmt.Printf("A: %d \n\n", data.Answer)
	return nil
}
