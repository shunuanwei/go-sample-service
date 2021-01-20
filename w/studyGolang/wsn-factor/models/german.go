package models

type GermanTranslator struct {
}

func (g *GermanTranslator) Translator(words string) string {

	return "德语 == >> " + words
}
