package usecase

import (
	"fmt"
	"quiz_master/repository"
	"testing"
)

func TestQuestionList(t *testing.T) {
	repository.Filldefaultjson()

	err := QuestionList()
	if err != nil {
		t.Errorf("cant get question list")
	}
}

type testCaseSingleQuestion struct {
	arg1   string // id question
	errors bool   // data exist (true) or not (false)
}

func TestQuestionSingle(t *testing.T) {
	repository.Filldefaultjson()

	cases := []testCaseSingleQuestion{
		{"1", true},
		{"2", true},
		{"99", false},
	}

	for _, tc := range cases {
		strMsg := "not found"
		if tc.errors {
			strMsg = "found"
		}
		t.Run(fmt.Sprintf("Get data from Id = %v, expect %v", tc.arg1, strMsg), func(t *testing.T) {
			err := QuestionSingle([]string{tc.arg1})
			if err != nil && tc.errors {
				strErr := "not found, but got 'found'"
				if tc.errors {
					strErr = "found, but got 'not found'"
				}
				t.Errorf("Expected Get Id '%v', '%v'", tc.arg1, strErr)
			}
		})
	}
}

type testCaseDeleteQuestion struct {
	arg1   string // id question
	errors bool   // data exist (true) or not (false)
}

func TestDeleteQuestion(t *testing.T) {
	repository.Filldefaultjson()

	cases := []testCaseDeleteQuestion{
		{"1", true},
		{"2", true},
		{"99", false},
	}

	for _, tc := range cases {
		strMsg := "not found"
		if tc.errors {
			strMsg = "found"
		}
		t.Run(fmt.Sprintf("Delete data from Id = %v, expect %v", tc.arg1, strMsg), func(t *testing.T) {
			err := DeleteQuestion([]string{tc.arg1})
			if err != nil && tc.errors {
				strErr := "not found, but got 'found'"
				if tc.errors {
					strErr = "found, but got 'not found'"
				}
				t.Errorf("Expected Delete Id '%v', '%v'", tc.arg1, strErr)
			}
		})
	}
}

type testCaseCreateQuestion struct {
	arg1   string // id
	arg2   string // question
	arg3   string // answer
	errors bool   // data exist (true) or not (false)
}

func TestCreateQuestion(t *testing.T) {
	repository.Filldefaultjson()

	cases := []testCaseCreateQuestion{
		{"99", "2 % 2 = ?", "0", false}, // new data
		{"98", "3 - 2 = ?", "1", false}, // new data
		{"1", "five", "5", true},        // duplicate, cant insert
	}

	for _, tc := range cases {
		strMsg := "success"
		if tc.errors {
			strMsg = "fail"
		}
		t.Run(fmt.Sprintf("Create data with Id = %v, expect %v", tc.arg1, strMsg), func(t *testing.T) {
			err := CreateQuestion([]string{tc.arg1, tc.arg2, tc.arg3})
			if err != nil && !tc.errors {
				strErr := "got 'fail'"
				if tc.errors {
					strErr = "got 'success'"
				}
				t.Errorf("Expected Create Id '%v', '%v'", tc.arg1, strErr)
			}
		})
	}
}

func TestUpdateQuestion(t *testing.T) {
	repository.Filldefaultjson()

	cases := []testCaseCreateQuestion{
		{"1", "five", "5", true},        // exist, can update
		{"2", "two", "2", true},         // exist, can update
		{"99", "2 % 2 = ?", "0", false}, // new data , no-update
	}

	for _, tc := range cases {
		strMsg := "fail"
		if tc.errors {
			strMsg = "success"
		}
		t.Run(fmt.Sprintf("Update data with Id = %v, expect %v", tc.arg1, strMsg), func(t *testing.T) {
			err := UpdateQuestion([]string{tc.arg1, tc.arg2, tc.arg3})
			if err != nil && tc.errors {
				strErr := "got 'success'"
				if tc.errors {
					strErr = "got 'fail'"
				}
				t.Errorf("Expected Update Id '%v', '%v'", tc.arg1, strErr)
			}
		})
	}
}

type testCaseAnswerQuestion struct {
	arg1   string // id
	arg2   string // answer
	answer string // Correct! or Wrong Answer!
	errors bool   // data exist (true) or not (false)
}

func TestAnswerQuestion(t *testing.T) {
	repository.Filldefaultjson()

	cases := []testCaseAnswerQuestion{
		{"2", "five", "Correct!", true},      // exist, can answer, correct answer (string)
		{"2", "5", "Correct!", true},         // exist, can answer, correct answer (int)
		{"1", "two", "Wrong Answer!", true},  // exist, can answer, wrong answer
		{"99", "0", "data not found", false}, // not-exist , forbidden to answer
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Asnwer data with Id = %v, expect %v ", tc.arg1, tc.answer), func(t *testing.T) {
			msg, err := AnswerQuestion([]string{tc.arg1, tc.arg2})
			if err != nil {
				if tc.errors {
					t.Errorf("Expected Answer Id '%v' data not-exists, got 'success'", tc.arg1)
				}
			} else {
				if msg != tc.answer {
					t.Errorf("Expected Answer Id '%v' '%v', got '%v'", tc.arg1, tc.answer, msg)
				}
			}
		})
	}
}
