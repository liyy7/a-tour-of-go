package string

import "unicode/utf8"

type String string

func (s String) Length() int {
	return utf8.RuneCountInString(string(s))
}
