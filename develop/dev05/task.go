package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Определение флагов
	after := flag.Int("A", 0, "Print N lines after each match")
	before := flag.Int("B", 0, "Print N lines before each match")
	context := flag.Int("C", 0, "Print N lines of context around each match (A+B)")
	count := flag.Bool("c", false, "Print only a count of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case distinctions")
	invert := flag.Bool("v", false, "Invert the sense of matching, to select non-matching lines")
	fixed := flag.Bool("F", false, "Interpret pattern as a literal string")
	lineNum := flag.Bool("n", false, "Print line numbers with output")

	flag.Parse()

	// Определение паттерна для поиска
	var pattern string
	if *fixed {
		pattern = flag.Arg(0)
	} else {
		pattern = regexp.QuoteMeta(flag.Arg(0))
	}

	// Компиляция регулярного выражения
	var regex *regexp.Regexp
	if *ignoreCase {
		regex = regexp.MustCompile("(?i)" + pattern)
	} else {
		regex = regexp.MustCompile(pattern)
	}

	var file *os.File
	var err error
	if flag.NArg() > 1 {
		file, err = os.Open(flag.Arg(1))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
	} else {
		fmt.Fprintf(os.Stderr, "Error! No file path: %v\n", err)
	}

	scanner := bufio.NewScanner(file)
	var contextLines []string
	matches := 0

	printLine := func(line string) {
		if *count {
			return
		}
		if *lineNum {
			fmt.Printf("%d:", matches)
		}
		fmt.Println(line)
	}

	for scanner.Scan() {
		line := scanner.Text()

		match := regex.MatchString(line)

		if *invert {
			match = !match
		}

		if match {
			// Печать строк до совпадения
			for _, contextLine := range contextLines {
				printLine(contextLine)
			}

			printLine(line)

			// Печать строк после совпадения
			for i := 1; i <= *after && scanner.Scan(); i++ {
				printLine(scanner.Text())
			}

			matches++
			contextLines = nil
		} else {
			contextLines = append(contextLines, line)

			// Урезание буфера до N строк перед совпадением
			if len(contextLines) > *before {
				contextLines = contextLines[len(contextLines)-*before:]
			}
		}

		// Урезание буфера строк контекста до N строк вокруг совпадения
		if len(contextLines) > *context {
			contextLines = contextLines[len(contextLines)-*context:]
		}
	}

	if *count {
		fmt.Println(matches)
	}
}
