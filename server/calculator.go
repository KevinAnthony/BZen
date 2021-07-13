package server

import "net/http"

func AddTwoNumbers(w http.ResponseWriter, r *http.Request) {
	numbers := struct {
		First int `json:"first"`
		Second int `json:"second"`
	}{}

	if err := UnmarshalFromRequest(r, &numbers); err != nil {
		WriteError(w, err)
	}

	answer := numbers.First + numbers.Second

	WriteSuccess(w, map[string]int{"answer": answer})
}


