package usecase

import (
	"fmt"
	"quiz_master/repository"
	"strconv"

	"github.com/divan/num2words"
)

func QuestionList() error {
	data, err := repository.Fetch()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("No | Question | Answer")
	for _, data := range data {
		fmt.Printf("%d  \"%v\" %d \n", data.Id, data.Question, data.Answer)
	}
	fmt.Println()
	return nil
}

func QuestionSingle(inputs []string) error {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Println(repository.ErrBadParamInput)
		return err
	}

	data, err := repository.GetByID(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Q: %v \n", data.Question)
	fmt.Printf("A: %d \n\n", data.Answer)
	return nil
}

func DeleteQuestion(inputs []string) error {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Println(repository.ErrBadParamInput)
		return err
	}

	err = repository.Delete(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Question no %d was deleted!\n\n", id)
	return nil
}

func CreateQuestion(inputs []string) error {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Println(repository.ErrBadParamInput)
		return err
	}

	ans, err := strconv.Atoi(inputs[2])
	if err != nil {
		fmt.Println(err)
		return err
	}

	qData := repository.Quiz{
		Id:       id,
		Question: inputs[1],
		Answer:   ans,
	}

	err = repository.Store(&qData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Question no %d created:\n", qData.Id)
	fmt.Printf("Q: %v \n", qData.Question)
	fmt.Printf("A: %d \n\n", qData.Answer)
	return nil
}

func UpdateQuestion(inputs []string) error {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Println(repository.ErrBadParamInput)
		return err
	}

	ans, err := strconv.Atoi(inputs[2])
	if err != nil {
		fmt.Println(repository.ErrBadParamInput)
		return err
	}

	qData := repository.Quiz{
		Id:       id,
		Question: inputs[1],
		Answer:   ans,
	}

	err = repository.Update(&qData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Question no %d updated:\n", qData.Id)
	fmt.Printf("Q: %v \n", qData.Question)
	fmt.Printf("A: %d \n\n", qData.Answer)
	return nil
}

func AnswerQuestion(inputs []string) (string, error) {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Println(repository.ErrBadParamInput)
		return "", err
	}

	data, err := repository.GetByID(id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	msg := "Wrong Answer!"
	ansStr := num2words.Convert(data.Answer)
	if ansStr == inputs[1] {
		msg = "Correct!"
	} else {
		ans, err := strconv.Atoi(inputs[1])
		if err == nil {
			if ans == data.Answer {
				msg = "Correct!"
			}
		}
	}

	fmt.Printf("%v\n\n", msg)
	return msg, nil
}

func ShowHelp() {
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
