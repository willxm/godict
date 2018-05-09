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
	q := ""
	l := len(input)
	for k, v := range input {
		if k > 0 {
			if l-1 == k {
				q += v
			} else {
				q += v + " "
			}
		}
	}
	var yd service.Youdao
	result := yd.Translate(q)
	if len(result.Translation) < 1 {
		fmt.Println("error words")
		return
	}
	for _, v := range result.Translation {
		fmt.Println(v)
	}
	for _, v := range result.Basic.Explains {
		fmt.Println(v)
	}
}
