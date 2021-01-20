package models

type EnglishTranslator struct {}


//实现接口

func (e *EnglishTranslator)Translator(str string) string {

	return "english ==>>"+str

}
