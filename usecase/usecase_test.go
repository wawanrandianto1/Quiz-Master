package usecase_test

import (
	"quiz_master/repository"
	"quiz_master/usecase"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Usecase", func() {

	type testCaseInput struct {
		Text       string
		FindLength int
		Message    string
	}

	type testCaseSingleQuestion struct {
		Id      string
		Status  bool
		Message string
	}

	type testCaseCreateQuestion struct {
		Quiz    repository.Quiz
		Status  bool
		Message string
	}

	type testCaseAnswerQuestion struct {
		Id      string
		Answer  string
		Status  string // Correct! or Wrong Answer!
		Message string
	}

	var (
		err  error
		data []repository.Quiz
	)

	Describe("TestDefineInput", func() {
		cases := []testCaseInput{
			{
				Text:       "help",
				FindLength: 1,
				Message:    "try input: 'help' length: 1",
			},
			{
				Text:       "questions",
				FindLength: 1,
				Message:    "try input: 'questions' length: 1",
			},
			{
				Text:       "question 1",
				FindLength: 2,
				Message:    "try input: 'question 1' length: 2",
			},
			{
				Text:       "create_question 1 \"How many letters are in the English alphabet?\" 26",
				FindLength: 4,
				Message:    "try input: 'create_question 1 \"How many letters are in the English alphabet?\" 26' length: 4",
			},
			{
				Text:       "update_question 1 \"How many letters are in the English alphabet?\" 26",
				FindLength: 4,
				Message:    "try input: 'update_question 1 \"How many letters are in the English alphabet?\" 26' length: 4",
			},
			{
				Text:       "answer_question 1 26",
				FindLength: 3,
				Message:    "try input: 'answer_question 1 26' length: 3",
			},
			{
				Text:       "delete_question 1",
				FindLength: 2,
				Message:    "try input: 'delete_question 1' length: 3",
			},
		}

		for _, tc := range cases {
			It(tc.Message, func() {
				result := usecase.DefineInput(tc.Text)
				Expect(len(result)).To(Equal(tc.FindLength))
			})
		}
	})

	BeforeEach(func() {
		data = append(data, repository.Quiz{
			Id:       1,
			Question: "How many letters are there in the English alphabet?",
			Answer:   26,
		}, repository.Quiz{
			Id:       2,
			Question: "How many vowels are there in the English alphabet?",
			Answer:   5,
		})
		repository.Writejson(data)
	})

	AfterEach(func() {
		repository.Clearjson()
	})

	Describe("TestGetQuestionList", func() {
		It("Question list work", func() {
			_, err = usecase.QuestionList()
			Expect(err).To(BeNil())
		})
	})

	Describe("TestGetQuestionSingle", func() {
		cases := []testCaseSingleQuestion{
			{Id: "1", Status: true, Message: "try GetID: 1, data exist, expect: true"},
			{Id: "2", Status: true, Message: "try GetID: 2, data exist, expect: true"},
			{Id: "99", Status: false, Message: "try GetID: 99, new data, expect: false"},
		}

		for _, tc := range cases {
			It(tc.Message, func() {
				_, err = usecase.QuestionSingle([]string{tc.Id})
				result := err == nil
				Expect(result).To(Equal(tc.Status))
			})
		}
	})

	Describe("TestDeleteQuestion", func() {
		cases := []testCaseSingleQuestion{
			{Id: "1", Status: true, Message: "try GetID: 1, data exist, expect: true"},
			{Id: "2", Status: true, Message: "try GetID: 2, data exist, expect: true"},
			{Id: "99", Status: false, Message: "try GetID: 99, new data, expect: false"},
		}

		for _, tc := range cases {
			It(tc.Message, func() {
				_, err = usecase.DeleteQuestion([]string{tc.Id})
				result := err == nil
				Expect(result).To(Equal(tc.Status))
			})
		}
	})

	Describe("TestCreateQuestion", func() {
		cases := []testCaseCreateQuestion{
			{
				Quiz: repository.Quiz{
					Id:       99,
					Question: "1 + 1 = ?",
					Answer:   2,
				},
				Status:  true,
				Message: "create ID: 99, new data, expect: true",
			},
			{
				Quiz: repository.Quiz{
					Id:       1,
					Question: "1 + 1 = ?",
					Answer:   2,
				},
				Status:  false,
				Message: "create ID: 1, data exists (duplicate), expect: false",
			},
		}

		for _, tc := range cases {
			It(tc.Message, func() {
				new_id := strconv.Itoa(tc.Quiz.Id)
				new_answer := strconv.Itoa(tc.Quiz.Answer)
				_, err = usecase.CreateQuestion([]string{new_id, tc.Quiz.Question, new_answer})
				result := err == nil
				Expect(result).To(Equal(tc.Status))
			})
		}
	})

	Describe("TestUpdateQuestion", func() {
		cases := []testCaseCreateQuestion{
			{
				Quiz: repository.Quiz{
					Id:       1,
					Question: "1 + 1 = ?",
					Answer:   2,
				},
				Status:  true,
				Message: "update ID: 1, data exists, expect: true",
			},
			{
				Quiz: repository.Quiz{
					Id:       99,
					Question: "1 + 1 = ?",
					Answer:   2,
				},
				Status:  false,
				Message: "update ID: 1, new data, expect: false",
			},
		}

		for _, tc := range cases {
			It(tc.Message, func() {
				new_id := strconv.Itoa(tc.Quiz.Id)
				new_answer := strconv.Itoa(tc.Quiz.Answer)
				_, err = usecase.UpdateQuestion([]string{new_id, tc.Quiz.Question, new_answer})
				result := err == nil
				Expect(result).To(Equal(tc.Status))
			})
		}
	})

	Describe("TestAnswerQuestion", func() {
		cases := []testCaseAnswerQuestion{
			{
				Id:      "1",
				Answer:  "26",
				Status:  "Correct!",
				Message: "Answer ID: 1, data exists, expect: success",
			},
			{
				Id:      "2",
				Answer:  "5",
				Status:  "Correct!",
				Message: "Answer ID: 2, data exists, expect: success",
			},
			{
				Id:      "2",
				Answer:  "five",
				Status:  "Correct!",
				Message: "Answer ID: 2, data exists, expect: success",
			},
			{
				Id:      "1",
				Answer:  "7",
				Status:  "Wrong Answer!",
				Message: "Answer ID: 1, data exists, expect: fail",
			},
			{
				Id:      "99",
				Answer:  "7",
				Status:  "Wrong Answer!",
				Message: "Answer ID: 99, new data, expect: fail",
			},
		}

		for _, tc := range cases {
			It(tc.Message, func() {
				result, _ := usecase.AnswerQuestion([]string{tc.Id, tc.Answer})
				Expect(result).To(Equal(tc.Status))
			})
		}
	})

})
