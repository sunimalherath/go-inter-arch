package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Base64 basic authentication
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
	// the output of the above (dXNlcjpwYXNz) is the same as what you get from the below curl command.
	// curl -u user:pass -v google.com will output => "Authorization: Basic dXNlcjpwYXNz"
	// -u = basic authentication
	// -v = verbose mode
	// send the user:pass combination to google.com
}
