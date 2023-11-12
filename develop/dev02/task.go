package main

import (
	"errors"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnzipString(str string) (string, error) {
	var res strings.Builder
	var count int
	var symbol rune
	var escape bool
	if str == "" {
		return "", nil
	}
	for _, c := range str {
		if escape {
			symbol = c
			escape = false
		} else {
			if c == '\\' {
				escape = true
			}
			if unicode.IsDigit(c) {
				if count != 0 {
					return "", errors.New("incorrect string")
				}
				count = int(c - '0')
			} else {
				if count == 0 && symbol != 0 {
					count = 1
				}
				for i := 0; i < count; i++ {
					res.WriteRune(symbol)
				}
				symbol = c
				count = 0
			}
		}
	}
	if count != 0 && symbol == 0 {
		return "", errors.New("incorrect string")
	}

	if count == 0 {
		count = 1
	}
	for i := 0; i < count; i++ {
		res.WriteRune(symbol)
	}
	return res.String(), nil
}
