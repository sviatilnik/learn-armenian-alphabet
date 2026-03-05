package main

import (
	"bufio"
	"fmt"
	"os"
)

type GameMode int

const (
	Letter GameMode = 1
	Sound  GameMode = 2
)

func main() {
	fmt.Println(`Добро пожаловать в тренажёр армянского алфавита!
Выберите режим:
1. Назвать букву (по символу)
2. Назвать звук/транслитерацию (по символу)
3. Выход
Введите номер режима:`)

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
