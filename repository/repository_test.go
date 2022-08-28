package repository_test

import (
	"quiz_master/repository"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type testCaseFindIndex struct {
	Index    int    // index
	Expected int    // expected
	Message  string // message
}

type testCaseGetByID struct {
	Id      int
	Status  bool
	Message string // message
}

type testCaseStore struct {
	Quizdata repository.Quiz
	Status   bool
	Message  string // message
}

var _ = Describe("Repository", func() {

	var (
		err  error
		data []repository.Quiz
		// tcGetId []testCaseGetByID
		// tcStore []testCaseStore
	)

	BeforeEach(func() {
		data = append(data, repository.Quiz{
			Id:       1,
			Question: "How many letters are there in the English alphabet?",
			Answer:   26,
		}, repository.Quiz{
			Id:       2,
			Question: "How many vowels are there in the English alphabet?",
			Answer:   5,
		}, repository.Quiz{
			Id:       4,
			Question: "1 + 1 = ?",
			Answer:   2,
		}, repository.Quiz{
			Id:       3,
			Question: "2 + 2 = ?",
			Answer:   4,
		})

		repository.Writejson(data)
	})

	AfterEach(func() {
		repository.Clearjson()
	})

	Describe("TestFindIndex", func() {
		indexing := []testCaseFindIndex{
			{Index: 1, Expected: 0, Message: "try find index ID: 1, expect: 0"},
			{Index: 2, Expected: 1, Message: "try find index ID: 2, expect: 1"},
			{Index: 99, Expected: -1, Message: "try find index ID: 99, expect: -1"},
		}

		for _, tc := range indexing {
			It(tc.Message, func() {
				result := repository.FindIndex(tc.Index, data)
				Expect(result).To(Equal(tc.Expected))
			})
		}
	})

	Describe("TestFetchingData", func() {
		BeforeEach(func() {
			data, err = repository.Fetch()
		})
		It("have list data", func() {
			Expect(len(data)).NotTo(Equal(0))
		})
	})

	Describe("TestGetByID", func() {
		cases := []testCaseGetByID{
			{Id: 1, Status: true, Message: "try GetID: 1, expect: true"},
			{Id: 2, Status: true, Message: "try GetID: 2, expect: true"},
			{Id: 99, Status: false, Message: "try GetID: 3, expect: false"},
		}

		for _, tc := range cases {
			It(tc.Message, func() {
				_, err = repository.GetByID(tc.Id)
				result := err == nil
				Expect(result).To(Equal(tc.Status))
			})
		}
	})

	Describe("TestStoreData", func() {
		cases := []testCaseStore{
			{Quizdata: 1, Status: true, Message: "try GetID: 1, expect: true"},
			{Quizdata: 2, Status: true, Message: "try GetID: 2, expect: true"},
			{Quizdata: 99, Status: false, Message: "try GetID: 3, expect: false"},
		}

		for _, tc := range cases {
			It(tc.Message, func() {
				_, err = repository.GetByID(tc.Id)
				result := err == nil
				Expect(result).To(Equal(tc.Status))
			})
		}
	})

})
