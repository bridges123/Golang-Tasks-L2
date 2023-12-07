package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// sortString возвращает посимвольно отсортированную строку
func sortString(s string) string {
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}

// FindAnagramsByDict распределяет анаграммы из массива по множествам,
// где ключом является первое слово подходящее под анаграмму
func FindAnagramsByDict(words []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, word := range words {
		lower := strings.ToLower(word)

		// Сортировка букв в слове
		sortedWord := sortString(lower)

		// Поиск множества анаграмм для текущего слова
		if set, found := anagrams[sortedWord]; found {
			// Если множество уже существует, добавляем текущее слово
			set = append(set, lower)
			anagrams[sortedWord] = set
		} else {
			anagrams[sortedWord] = []string{lower}
		}
	}
	// Удаляем множества из одного элемента
	for key, set := range anagrams {
		if len(set) == 1 {
			delete(anagrams, key)
		} else {
			// Сортируем множества по возрастанию
			sort.Strings(set)
			if set[0] != key {
				// Используем первое слово из множества в качестве ключа
				anagrams[set[0]] = set
				delete(anagrams, key)
			}
		}
	}

	return anagrams
}
