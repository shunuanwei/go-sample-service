package main

import (
	"fmt"
	."github.com/sdgmf/go-sample-service/w/studyGolang/wsn-factor/models"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		time.Sleep(3 * time.Second)
	}()

	var lan int
	fmt.Printf("%s\r\n%s\r\n", "以下是可翻译的语言种类，请输入代表数字", "1：德语、2：英语、3：日语")
	_, _ = fmt.Scanln(&lan)

	fmt.Println("请输入要翻译成中文的文本：")
	var inputWords string
	_, _ = fmt.Scanln(&inputWords)



	var translator Translator

	//根据不同的语言种类，实例化不同的翻译类
	switch lan {
	case 1:
		translator = new(GermanTranslator)
	case 2:
		translator = new(EnglishTranslator)
	case 3:
		translator = new(JapaneseTranslator)
	default:
		panic("no such translator")
	}

	fmt.Println(translator.Translator(inputWords))
}
