// Package main demonstrates type checking on custom errors. This example
// lifted from
// https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go
package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

type RequestError struct {
	StatusCode int
	Err error
}

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
// type error interface {
// 	Error() string
// }
func (r *RequestError) Error() string {
	return r.Err.Error()
}

func (r *RequestError) Temporary() bool {
	return r.StatusCode == http.StatusServiceUnavailable // 503
}

func doRequest() error {
	return &RequestError{
		StatusCode: 503,
		Err:        errors.New("unavailable"),
	}
}

func main() {
	err := doRequest()
	if err != nil {
		fmt.Println(err)
		re, ok := err.(*RequestError)
		if ok {
			if re.Temporary() {
				fmt.Println("This request can be tried again")
			} else {
				fmt.Println("This request cannot be tried again")
			}
		}
		os.Exit(1)
	}

	fmt.Println("success!")
}
