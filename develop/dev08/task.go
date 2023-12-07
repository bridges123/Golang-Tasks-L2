package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качестве аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

func main() {
	fmt.Println("Simple Shell. Type 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Exiting the shell.")
			break
		}

		processInput(input)
	}
}

func processInput(input string) {
	args := strings.Fields(input)

	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "cd":
		changeDirectory(args[1:])
	case "pwd":
		printCurrentDirectory()
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "kill":
		killProcess(args[1:])
	case "ps":
		showProcesses()
	default:
		runCommand(args)
	}
}

func changeDirectory(args []string) {
	if len(args) > 0 {
		err := os.Chdir(args[0])
		if err != nil {
			fmt.Println("Error changing directory:", err)
		}
	} else {
		fmt.Println("cd: missing argument")
	}
}

func printCurrentDirectory() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	} else {
		fmt.Println(currentDir)
	}
}

func killProcess(args []string) {
	if len(args) > 0 {
		pid := args[0]
		err := exec.Command("kill", pid).Run()
		if err != nil {
			fmt.Println("Error killing process:", err)
		}
	} else {
		fmt.Println("kill: missing argument")
	}
}

func showProcesses() {
	cmd := exec.Command("ps")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting process information:", err)
		return
	}
	fmt.Print(string(output))
}

func runCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running command:", err)
	}
}
