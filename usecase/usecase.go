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
	return nil
}

func QuestionSingle(inputs []string) error {
	id, _ := strconv.Atoi(inputs[0])
	data, err := repository.GetByID(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Q: %v \n", data.Question)
	fmt.Printf("A: %d \n", data.Answer)
	return nil
}

func DeleteQuestion(inputs []string) error {
	id, err := strconv.Atoi(inputs[0])
	err = repository.Delete(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Question no %d was deleted!\n", id)
	return nil
}

func CreateQuestion(inputs []string) error {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Println(err)
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
	fmt.Printf("A: %d \n", qData.Answer)
	return nil
}

func UpdateQuestion(inputs []string) error {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Println(err)
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

	err = repository.Update(&qData)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Question no %d updated:\n", qData.Id)
	fmt.Printf("Q: %v \n", qData.Question)
	fmt.Printf("A: %d \n", qData.Answer)
	return nil
}

func AnswerQuestion(inputs []string) (string, error) {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		fmt.Println(err)
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

	fmt.Println(msg)
	return msg, nil
}
