package starter

import (
	"fmt"
	"net/http"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func OddOrEven(n int) string {
	fmtStr := "%d is an odd number"
	if n%2 == 0 {
		fmtStr = "%d is an even number"
	}
	resp := fmt.Sprintf(fmtStr, n)
	return resp
}

func Checkhealth(writer http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(writer, "health check passed")
}
