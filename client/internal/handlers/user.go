package handlers

import (
	"fmt"
	"github.com/spf13/cobra"
)

func (h *Handlers) RegisterUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reg",
		Short: "Регистрация пользователя",
		Run: func(cmd *cobra.Command, args []string) {
			var username, password string
			fmt.Print("Введите логи: ")
			fmt.Scanln(&username)
			fmt.Print("Введите пароль: ")
			fmt.Scanln(&password)

			if username == "admin" && password == "1234" {
				fmt.Println("Авторизация успешна!")
			} else {
				fmt.Println("Неверные учетные данные.")
			}
		},
	}

	return cmd
}
