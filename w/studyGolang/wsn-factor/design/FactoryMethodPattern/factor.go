package main

import (
	"fmt"
	. "github.com/sdgmf/go-sample-service/w/studyGolang/wsn-factor/models"
	"time"
)

func Create(input int) Translator {

	var translator Translator

	switch input {
	case 1:
		translator = new(GermanTranslator)
		break
	case 2:
		translator = new(EnglishTranslator)
		break
	case 3:
		translator = new(JapaneseTranslator)
		break
	default:
		fmt.Println("输入错误")
		return nil
	}
	return translator
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		time.Sleep(3 * time.Second)
	}()

	var lan int
	fmt.Printf("%s\r\n%s\r\n", "以下是可翻译的语言种类，请输入代表数字", "1：德语、2：英语、3：日语")
	fmt.Scanln(&lan)

	fmt.Println("请输入要翻译成中文的文本：")
	var inputWords string
	fmt.Scanln(&inputWords)

	//客户端只关注如何获取翻译类，而不用关注创建翻译类的细节
	translator:=Create(lan)

	fmt.Println(translator.Translator(inputWords))
}
