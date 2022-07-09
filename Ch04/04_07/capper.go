package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	wtr io.Writer
}

// Area returns the area of the square
func (c Capper) Write(p []byte) (n int, err error) {
	return  c.wtr.Write([]byte(strings.ToUpper(string(p))))
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
