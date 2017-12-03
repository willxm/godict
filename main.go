package main

import (
	"fmt"
	"os"

	"github.com/willxm/godict/service"
)

func main() {
	input := os.Args
	if len(input) <= 1 {
		fmt.Println("please input words")
		return
	}
	q := input[1]
	var yd service.Youdao
	result := yd.Translate(q)
	if len(result.Basic.Explains) < 1 {
		fmt.Println("error words")
		return
	}
	for _, v := range result.Basic.Explains {
		fmt.Println(v)
	}
}
