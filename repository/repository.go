package repository

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"sort"
)

var (
	_, b, _, _ = runtime.Caller(0)
	dataPath   = filepath.Join(filepath.Dir(b), "../quiz.json")
)

var (
	ErrNotFound      = errors.New("data not found")
	ErrConflict      = errors.New("data already exist")
	ErrBadParamInput = errors.New("param is not valid")
)

type Quiz struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Answer   int    `json:"answer"`
}

func FindIndex(id int, result []Quiz) int {
	var foundIndex int = -1
	for i, v := range result {
		if v.Id == id {
			foundIndex = i
			break
		}
	}
	return foundIndex
}

func Clearjson() error {
	jsonString, err := json.Marshal([]map[string]interface{}{})
	if err != nil {
		return err
	}
	return os.WriteFile(dataPath, jsonString, os.ModePerm)
}

func Writejson(values []Quiz) error {
	jsonString, err := json.Marshal(values)
	if err != nil {
		return err
	}
	return os.WriteFile(dataPath, jsonString, os.ModePerm)
}

func Fetch() ([]Quiz, error) {
	datas := []Quiz{}
	file, _ := os.ReadFile(dataPath)
	errors := json.Unmarshal(file, &datas)
	if errors != nil {
		return []Quiz{}, errors
	}

	sort.Slice(datas, func(i, j int) bool {
		return datas[i].Id < datas[j].Id
	})
	return datas, nil
}

func GetByID(id int) (Quiz, error) {
	// read all
	result, err := Fetch()
	if err != nil {
		return Quiz{}, err
	}

	var qz Quiz
	index := FindIndex(id, result)
	if index < 0 {
		return qz, ErrNotFound
	}

	qz = result[index]
	return qz, nil
}

func Store(values *Quiz) error {
	// read all
	result, err := Fetch()
	if err != nil {
		return err
	}

	index := FindIndex(int(values.Id), result)
	if index > -1 {
		return ErrConflict
	}

	// insert at last
	result = append(result, Quiz{
		Id:       values.Id,
		Question: values.Question,
		Answer:   values.Answer,
	})
	return Writejson(result)
}

func Update(values *Quiz) error {
	// read all
	result, err := Fetch()
	if err != nil {
		return err
	}

	index := FindIndex(int(values.Id), result)
	if index < 0 {
		return ErrNotFound
	}

	result[index].Id = values.Id
	result[index].Question = values.Question
	result[index].Answer = values.Answer
	return Writejson(result)
}

func Delete(id int) error {
	result, err := Fetch()
	if err != nil {
		return err
	}

	index := FindIndex(id, result)
	if index < 0 {
		return ErrNotFound
	}

	newSlice := make([]Quiz, len(result)-1)
	var indexing int = 0
	for i, data := range result {
		if i != index {
			newSlice[indexing] = data
			indexing++
		}
	}
	return Writejson(newSlice)
}
