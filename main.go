package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Ошибка при получении домашней директории:", err)
		return
	}

	var folderName string
	fmt.Print("Введите имя папки для удаления: ")
	fmt.Scanln(&folderName)

	folderPath := filepath.Join(homeDir, folderName)

	info, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		fmt.Println("Папка не существует:", folderPath)
		return
	} else if !info.IsDir() {
		fmt.Println("Указанный путь не является папкой:", folderPath)
		return
	}

	entries, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Ошибка при чтении папки:", err)
		return
	}

	if len(entries) == 0 {
		err := os.Remove(folderPath)
		if err != nil {
			fmt.Println("Ошибка при удалении папки:", err)
			return
		}
		fmt.Println("Папка успешно удалена:", folderPath)
	} else {
		fmt.Println("Папка не пустая, не удалена:", folderPath)
	}
}
