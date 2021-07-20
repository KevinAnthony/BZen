package server

import (
	"fmt"
	"net/http"
)

func AddTwoNumbers(w http.ResponseWriter, r *http.Request) {
	numbers := struct {
		First  int `json:"first"`
		Second int `json:"second"`
	}{}

	if err := UnmarshalFromRequest(r, &numbers); err != nil {
		WriteError(w, err)
	}

	answer := numbers.First + numbers.Second

	WriteSuccess(w, map[string]int{"answer": answer})
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	numbers := struct {
		Alpha int `json:"alpha"`
		Beta  int `json:"beta"`
	}{}
	if err := UnmarshalFromRequest(r, &numbers); err != nil {
		WriteError(w, err)
	}

	answer := numbers.Alpha * numbers.Beta

	type Input struct {
		Alpha int `json:"alpha"`
		Beta  int `json:"beta"`
	}

	WriteSuccess(w, struct {
		Input  Input `json:"values"`
		Answer int   `json:"answer"`
	}{
		Answer: answer,
		Input: Input{
			Alpha: numbers.Alpha,
			Beta:  numbers.Beta,
		},
	})
}

func Sum(w http.ResponseWriter, r *http.Request) {
	input := struct {
		Numbers []int `json:"numbers"`
	}{}
	if err := UnmarshalFromRequest(r, &input); err != nil {
		WriteError(w, err)
	}

	sum := 0
	for _, number := range input.Numbers {
		sum = sum + number
		fmt.Printf("Sum: %d Number: %d\n", sum, number)
	}

	WriteSuccess(w, map[string]int{"answer": sum})
}
