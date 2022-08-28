package repository

import (
	"fmt"
	"testing"
)

func TestWritejson(t *testing.T) {
	datas := []Quiz{
		{
			Id:       1,
			Question: "How many letters are there in the English alphabet?",
			Answer:   26,
		},
		{
			Id:       2,
			Question: "How many vowels are there in the English alphabet?",
			Answer:   5,
		},
	}

	err := Writejson(datas)
	if err != nil {
		t.Errorf("cant write json file")
	}
}

func TestClearjson(t *testing.T) {
	err := Clearjson()
	if err != nil {
		t.Errorf("cant clear json file")
	}
}

func TestFilldefaultjson(t *testing.T) {
	err := Filldefaultjson()
	if err != nil {
		t.Errorf("cant fill json file")
	}
}

func TestFetch(t *testing.T) {
	_, err := Fetch()
	if err != nil {
		t.Errorf("error reading json file")
	}
}

type testCaseGetByID struct {
	arg1   int
	errors bool
}

func TestGetByID(t *testing.T) {
	cases := []testCaseGetByID{
		{1, false},
		{2, false},
		{3, false},
		{99, true}, // if not found dont fail the test, set false to fail the test
	}

	for _, tc := range cases {
		strMsg := "found"
		if tc.errors {
			strMsg = "not found"
		}
		t.Run(fmt.Sprintf("Get data from Id = %d, expect %v", tc.arg1, strMsg), func(t *testing.T) {
			_, err := GetByID(tc.arg1)
			if err != nil && !tc.errors {
				strErr := "found, but got 'not found'"
				if tc.errors {
					strErr = "not found, but got 'found'"
				}
				t.Errorf("Expected Id '%d', '%v'", tc.arg1, strErr)
			}
		})
	}
}

type testCaseStore struct {
	arg1   Quiz
	errors bool
}

func TestStore(t *testing.T) {
	TestWritejson(t)

	// store test
	cases := []testCaseStore{
		{
			Quiz{
				Id:       6,
				Question: "1 + 1 = ?",
				Answer:   2,
			},
			false, // id not exists, can store
		},
		{
			Quiz{
				Id:       3,
				Question: "1 + 1 = ?",
				Answer:   2,
			},
			true, // id exists, cannot store
		},
	}

	for _, tc := range cases {
		strMsg := "new data"
		if tc.errors {
			strMsg = "exists"
		}
		t.Run(fmt.Sprintf("Store id %d, expected %v", tc.arg1.Id, strMsg), func(t *testing.T) {
			err := Store(&tc.arg1)
			if err != nil && !tc.errors {
				strErr := "new data, can insert"
				if tc.errors {
					strErr = "data exists, can't insert"
				}
				t.Errorf("Expected Store Id '%d', '%v'", int(tc.arg1.Id), strErr)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	Filldefaultjson()

	cases := []testCaseStore{
		{
			Quiz{
				Id:       3,
				Question: "3 + 1 = ?",
				Answer:   4,
			},
			false, // id exists, can update
		},
		{
			Quiz{
				Id:       5,
				Question: "1 + 1 = ?",
				Answer:   2,
			},
			true, // id not exists, cannot update
		},
	}

	for _, tc := range cases {
		strMsg := "exists"
		if tc.errors {
			strMsg = "new data"
		}
		t.Run(fmt.Sprintf("Update id %d, expected %v", tc.arg1.Id, strMsg), func(t *testing.T) {
			err := Update(&tc.arg1)
			if err != nil && !tc.errors {
				strErr := "data exists, update-yes"
				if tc.errors {
					strErr = "new data, update-no"
				}
				t.Errorf("Expected Update Id '%d', '%v'", int(tc.arg1.Id), strErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	Filldefaultjson()

	cases := []testCaseGetByID{
		{1, false}, // id exists, delete-yes
		{2, false}, // id exists, delete-yes
		{99, true}, // id not exists, delete-no
	}

	for _, tc := range cases {
		strMsg := "delete-yes"
		if tc.errors {
			strMsg = "delete-no"
		}
		t.Run(fmt.Sprintf("Delete id: %d, expected %v", tc.arg1, strMsg), func(t *testing.T) {
			err := Delete(tc.arg1)
			if err != nil && !tc.errors {
				strErr := "data exists, delete-yes"
				if tc.errors {
					strErr = "new data, delete-no"
				}
				t.Errorf("Expected Delete Id: '%d', '%v'", int(tc.arg1), strErr)
			}
		})
	}
}
