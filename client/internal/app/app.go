package app

import (
	"bufio"
	"client/internal/config"
	"client/internal/handlers"
	"client/internal/service"
	"fmt"
	"os"
	"strings"
)

func Run(cnf *config.Config) {
	gophKeeper := service.NewGophKeeperClient()
	handlers := handlers.NewHandlers(gophKeeper, *cnf)
	err := handlers.Run()
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}

	// Цикл для работы приложения
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ") // Интерактивная строка ввода
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		// Выход из приложения
		if input == "exit" {
			fmt.Println("Выход из программы...")
			break
		}

		// Разделяем ввод на команду и аргументы
		args := strings.Split(input, " ")

		handlers.SetArgs(args)

		if err := handlers.Execute(); err != nil {
			fmt.Println("Ошибка выполнения команды:", err)
		}
	}
}
