package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	io.Copy(os.Stdout, decoder)

	if barr, err := base64.StdEncoding.DecodeString(data); err == nil {
		fmt.Print(string(barr))
	}
}

const data = `YmFzZTY0IHN0cmluZwo=`
