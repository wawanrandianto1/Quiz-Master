package repository_test

import (
	"quiz_master/repository"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {

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
		Status   string
		Message  string // message
	}

	var (
		err  error
		data []repository.Quiz
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
		It("have list data", func() {
			data, err = repository.Fetch()
			Expect(len(data)).NotTo(Equal(0))
		})
	})

	Describe("TestGetByID", func() {
		cases := []testCaseGetByID{
			{Id: 1, Status: true, Message: "try GetID: 1, expect: true"},
			{Id: 2, Status: true, Message: "try GetID: 2, expect: true"},
			{Id: 99, Status: false, Message: "try GetID: 99, expect: false"},
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
			{
				Quizdata: repository.Quiz{
					Id:       6,
					Question: "1 + 1 = ?",
					Answer:   2,
				},
				Status:  "success",
				Message: "try Store ID: 6, new data, expect: success",
			},
			{
				Quizdata: repository.Quiz{
					Id:       3,
					Question: "1 + 1 = ?",
					Answer:   2,
				},
				Status:  "fail",
				Message: "try Store ID: 3, duplicate, expect: fail",
			},
		}

		for _, tc := range cases {
			It(tc.Message, func() {
				result := "success"
				err = repository.Store(&tc.Quizdata)
				if err != nil {
					result = "fail"
				}
				Expect(result).To(Equal(tc.Status))
			})
		}
	})

	Describe("TestUpdateData", func() {
		cases := []testCaseStore{
			{
				Quizdata: repository.Quiz{
					Id:       3,
					Question: "3 + 1 = ?",
					Answer:   4,
				},
				Status:  "success",
				Message: "try Update ID: 3, data exists, expect: success",
			},
			{
				Quizdata: repository.Quiz{
					Id:       5,
					Question: "1 + 1 = ?",
					Answer:   2,
				},
				Status:  "fail",
				Message: "try Store ID: 5, data not exists, expect: fail",
			},
		}

		for _, tc := range cases {
			It(tc.Message, func() {
				result := "success"
				err = repository.Update(&tc.Quizdata)
				if err != nil {
					result = "fail"
				}
				Expect(result).To(Equal(tc.Status))
			})
		}
	})

	Describe("TestDeleteData", func() {
		cases := []testCaseGetByID{
			{Id: 1, Status: true, Message: "try Delete ID: 1, data exists, expect: true"},
			{Id: 99, Status: false, Message: "try Delete ID: 99, data not exists, expect: false"},
		}

		for _, tc := range cases {
			It(tc.Message, func() {
				err = repository.Delete(tc.Id)
				result := err == nil
				Expect(result).To(Equal(tc.Status))
			})
		}
	})

})
