package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Определение флагов
	fields := flag.String("f", "", "Select fields (columns)")
	delimiter := flag.String("d", "\t", "Use a different delimiter")
	separated := flag.Bool("s", false, "Only output lines with delimiters")

	flag.Parse()
	selectedFields := parseFields(*fields)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		hasDelimiter := strings.Contains(line, *delimiter)

		// Если флаг "separated" установлен и в строке нет разделителя - скип
		if *separated && !hasDelimiter {
			continue
		}
		fields := strings.Split(line, *delimiter)

		// Вывод выбранных полей
		for _, fieldIndex := range selectedFields {
			if fieldIndex > 0 && fieldIndex <= len(fields) {
				fmt.Print(fields[fieldIndex-1])
			}
			if fieldIndex < len(fields) {
				fmt.Print(*delimiter)
			}
		}
		fmt.Println()
	}
}

// Функция для парсинга выбранных полей
func parseFields(fields string) []int {
	var result []int

	if fields == "" {
		return result
	}

	fieldList := strings.Split(fields, ",")

	for _, field := range fieldList {
		if index, err := strconv.Atoi(field); err == nil {
			result = append(result, index)
		}
	}

	return result
}
