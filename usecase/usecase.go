package usecase

import (
	"errors"
	"quiz_master/repository"
	"regexp"
	"strconv"
	"strings"

	"github.com/divan/num2words"
)

func DefineInput(txt string) []string {
	var arr []string
	splitstring := strings.Fields(txt)
	if len(splitstring) <= 2 {
		arr = splitstring
	} else {
		arr = append(arr, splitstring[0], splitstring[1])
		// re := regexp.MustCompile(`(\w{4,})( )?(\d+)?( )?("[^"]+")?( )?(\d+)?`)
		re := regexp.MustCompile(`"[^"]+"`)
		newStrs := re.FindAllString(txt, -1)
		if len(newStrs) > 0 {
			s := newStrs[0]
			arr = append(arr, s[1:len(s)-1])
		}
		lastIndex := len(splitstring) - 1
		arr = append(arr, splitstring[lastIndex])
	}
	return arr
}

func QuestionList() ([]repository.Quiz, error) {
	data, err := repository.Fetch()
	return data, err
}

func QuestionSingle(inputs []string) (repository.Quiz, error) {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		return repository.Quiz{}, repository.ErrBadParamInput
	}

	data, err := repository.GetByID(id)
	if err != nil {
		return repository.Quiz{}, err
	}

	return data, nil
}

func DeleteQuestion(inputs []string) (int, error) {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		return 0, repository.ErrBadParamInput
	}

	err = repository.Delete(id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func CreateQuestion(inputs []string) (repository.Quiz, error) {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		return repository.Quiz{}, repository.ErrBadParamInput
	}

	ans, err := strconv.Atoi(inputs[2])
	if err != nil {
		return repository.Quiz{}, err
	}

	result := repository.Quiz{
		Id:       id,
		Question: inputs[1],
		Answer:   ans,
	}

	err = repository.Store(&result)
	if err != nil {
		convStr := strconv.Itoa(result.Id)
		message := "Question no " + convStr + " already existed!"
		return repository.Quiz{}, errors.New(message)
	}

	return result, nil
}

func UpdateQuestion(inputs []string) (repository.Quiz, error) {
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		return repository.Quiz{}, repository.ErrBadParamInput
	}

	ans, err := strconv.Atoi(inputs[2])
	if err != nil {
		return repository.Quiz{}, repository.ErrBadParamInput
	}

	result := repository.Quiz{
		Id:       id,
		Question: inputs[1],
		Answer:   ans,
	}

	err = repository.Update(&result)
	if err != nil {
		return repository.Quiz{}, err
	}

	return result, nil
}

func AnswerQuestion(inputs []string) (string, error) {
	msg := "Wrong Answer!"
	id, err := strconv.Atoi(inputs[0])
	if err != nil {
		return msg, repository.ErrBadParamInput
	}

	data, err := repository.GetByID(id)
	if err != nil {
		return msg, err
	}

	inputStr := strings.ToLower(inputs[1])
	ansStr := num2words.Convert(data.Answer)
	if ansStr == inputStr {
		msg = "Correct!"
	} else {
		ans, err := strconv.Atoi(inputs[1])
		if err == nil {
			if ans == data.Answer {
				msg = "Correct!"
			}
		}
	}
	return msg, nil
}
