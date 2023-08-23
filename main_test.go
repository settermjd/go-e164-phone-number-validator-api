package main

import "testing"

type phone_number_data struct {
	phone_number string
	is_valid     bool
}

var tests = []phone_number_data{
	{"+49 109 1234 4321", true},
	{"0049 109 1234 4321", false},
	{"+4910912344321", true},
	{"004910912344321", false},
	{"", false},
	{"   ", false},
}

func TestIsValidPhoneNumber(t *testing.T) {
	for _, data := range tests {
		is_valid := IsValidPhoneNumber(data.phone_number)
		if is_valid != data.is_valid {
			t.Error(
				"For", data.phone_number,
				"expected", data.is_valid,
				"got", is_valid,
			)
		}
	}
}
