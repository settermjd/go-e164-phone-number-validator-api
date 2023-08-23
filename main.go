package main

import (
	"fmt"
	"regexp"
	"strings"
)

const E164Regex = `^\+[1-9]\d{1,14}$`

func main() {
	re := regexp.MustCompile(E164Regex)
	phone_number := strings.ReplaceAll("<<YOUR_PHONE_NUMBER>>", " ", "")
	fmt.Printf("%q\n", re.Find([]byte(phone_number)))
}
