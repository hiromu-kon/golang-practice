package main

import "fmt"

func main() {
	s := make([]int, 3, 8)
	fmt.Println(s)

	type HTTPStatusCode int

	const (
		StatusOK                  HTTPStatusCode = 200 + iota
		StatusBadRequest                         = 400
		StatusUnauthorized                       = 401
		StatusForbidden                          = 403
		StatusNotFound                           = 404
		StatusInternalServerError                = 500
		StatusServiceUnavailable                 = 503
	)

	fmt.Println(StatusOK)
	fmt.Println(StatusBadRequest)

}
