package models


type JapaneseTranslator struct{}

func (j *JapaneseTranslator) Translator(words string) string {

	return "日语 == >> "+ words
}

