package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString("sort: No arguments\n")
		os.Exit(1)
	}
	flags := parseFlags(os.Args[1:])
	fileName := os.Args[len(os.Args)-1]
	sortFile(flags, fileName)
}

func parseFlags(args []string) map[byte]int {
	availableFalgs := []byte{'k', 'n', 'r', 'u'}
	flags := make(map[byte]int, 8)
	for i := 0; i < len(args)-1; i++ {
		arg := args[i]
		// Если аргумент - недопустимый флаг, выкидываем Stderr
		if arg[0] != '-' || bytes.IndexByte(availableFalgs, arg[1]) == -1 {
			os.Stderr.WriteString("sort: Invalid command-line arguments\n")
			os.Exit(1)
		}
		if arg[1] == 'k' {
			i++
			val, err := strconv.Atoi(args[i])
			if err != nil || val <= 0 {
				os.Stderr.WriteString("sort: -k: Invalid command-line arguments\n")
				os.Exit(1)
			}
			flags['k'] = val - 1
		} else {
			flags[arg[1]] = 1
		}
	}
	return flags
}

func sortFile(flags map[byte]int, fileName string) {
	lines := readLines(fileName)
	sortLines(lines, flags)
	if flags['u'] == 1 {
		cutRepeatableLines(lines)
	}
	printLines(lines)
}

func cutRepeatableLines(lines []string) {
	length := len(lines)
	for i := 1; i < length; i++ {
		if lines[i] == lines[i-1] {
			copy(lines[i-1:], lines[i:])
			length--
			lines = lines[:length]
		}
	}
}

func printLines(lines []string) {
	for _, l := range lines {
		fmt.Println(l)
	}
}

func sortLines(lines []string, flags map[byte]int) {
	for i := 0; i < len(lines)-1; i++ {
		for j := i; j < len(lines); j++ {
			if compareLines(lines[i], lines[j], flags) {
				lines[i], lines[j] = lines[j], lines[i]
			}
		}
	}
}

func compareLines(l1 string, l2 string, flags map[byte]int) bool {
	result, col := true, flags['k']
	split1, split2 := strings.Split(l1, " "), strings.Split(l2, " ")
	len1, len2 := len(split1), len(split2)
	maxLen := max(len1, len2)
	if col >= maxLen {
		col = 0
	}
	for i := col; i < maxLen; i++ {
		l1, l2 := "", ""
		if i < len1 {
			l1 = split1[i]
		}
		if i < len2 {
			l2 = split2[i]
		}
		result = l1 >= l2
		if flags['n'] == 1 {
			l1Num, _ := strconv.Atoi(l1)
			l2Num, _ := strconv.Atoi(l2)
			result = l1Num >= l2Num
		}
		if l1 != l2 {
			break
		}
	}
	if flags['r'] == 1 {
		return !result
	}
	return result
}

func readLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		os.Stderr.WriteString("No such file or directory\n")
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
