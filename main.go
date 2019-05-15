package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
FOR_LABEL:
	for {
		res := bufio.NewScanner(os.Stdin)
		fmt.Printf("以下のいずれかを入力してください\n1. [ENG] CNN News\n2. [JPN] Yahoo!ニュース\nq. 終了\n")
		res.Scan()
		//fmt.Println("input is", res.Text())
		switch res.Text() {
		case "1":
			fmt.Println("CNN News を読み上げます！")
			RunCNN()
			break FOR_LABEL
		case "2":
			fmt.Println("Yahoo!ニュースを読み上げます！")
			RunYahooJapan()
			break FOR_LABEL
		case "q":
			break FOR_LABEL
		default:
			continue
		}
	}

}
