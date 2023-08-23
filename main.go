package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func IsValidPhoneNumber(phone_number string) bool {
	e164Regex := `^\+[1-9]\d{1,14}$`
	re := regexp.MustCompile(e164Regex)
	phone_number = strings.ReplaceAll(phone_number, " ", "")

	return re.Find([]byte(phone_number)) != nil
}

type Response struct {
	Data string
}

func validate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	phone_number := r.FormValue("phone_number")
	if phone_number == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Data: "No phone number was provided\n",
		})
		return
	}

	if IsValidPhoneNumber(phone_number) {
		w.Write([]byte(fmt.Sprintf("[%s] is a valid phone number", phone_number)))
		json.NewEncoder(w).Encode(Response{
			Data: fmt.Sprintf("[%s] is a valid phone number", phone_number),
		})
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Response{
		Data: fmt.Sprintf("[%s] is NOT a valid phone number", phone_number),
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", validate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
