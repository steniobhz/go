package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	strnum := query.Get("number")
	if strnum == "" {
		strnum = "1"
	}
	number, err := strconv.ParseInt(strnum, 0, 64)
	steps := calcCollatz(int(number))
	log.Println(steps, err, reflect.TypeOf(steps))
	w.Write([]byte(fmt.Sprintf("steps: %d\n", steps)))
}

func calcCollatz(n int) int {
	steps := 1
	for n > 1 {
		steps++
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return steps
}

func main() {

	http.HandleFunc("/", handler)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)

}
