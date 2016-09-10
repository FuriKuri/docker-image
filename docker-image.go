package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("img/image.png")
	if err != nil {
		panic(err)
	}

	sEnc := b64.StdEncoding.EncodeToString(b)
	fmt.Println(sEnc)
}
