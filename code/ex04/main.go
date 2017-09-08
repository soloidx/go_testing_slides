package main

import "fmt"
import "net/http"

import "../ex04/ex0401"
import "../ex04/ex0402"

func main() {
	resp, _ := ex0401.GetGoogleStatusBad()
	fmt.Printf("%s\n", resp)

	client := &http.Client{}
	resp2, _ := ex0402.GetGoogleStatus(client)
	fmt.Printf("%s\n", resp2)
}
