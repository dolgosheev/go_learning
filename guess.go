// Package Игра угадай число
package main

import (
	redisClient "TestProject/src/redis"
	"TestProject/src/test"
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// main
func main() {
	str := "Hello"
	fmt.Println(str)
	test.GetMe(&str)
	fmt.Println(str)

	redisClient.SetValue("1", 89831143406, time.Second*15)

	value := redisClient.GetValue("1")

	fmt.Println("Redis value :", value)

	/* количество попыток */
	attempts := 5

	/* номер попытки */
	attempt := 1

	/* Максимальное число от 0 до ? */
	maxValue := 5

	answerNumber := getRandomNumber(maxValue)

	for attempts >= 0 {

		fmt.Printf("Отгадайте число от 1 до %d : ", maxValue)

		/* чтение из буфера */
		reader := bufio.NewReader(os.Stdin)
		stdIn, err := reader.ReadString('\n')

		/* ошибка и выход */
		logError(err, "Не удалось получить значение из буфера")

		/* обрезание символа окнца строки */
		stdIn = strings.TrimSpace(stdIn)

		/* конвертация stdin в число */
		inputValue, err := strconv.Atoi(stdIn)

		/* ошибка и выход */
		logError(err, "Не удалось привести полученное значение к числу")

		/* отгадали */
		if inputValue == answerNumber {
			break
		}

		fmt.Println("Ответ не верный")

		/* уменьшаем количество попыток */
		attempts--

		/* увеличиваем показатель попыток */
		attempt++
	}

	fmt.Printf("\tВы отгадали число [%d] c %d попытки", answerNumber, attempt)
}

// logError функция для вывода ошибок
func logError(err error, description string) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%s | Error [%v]", description, err))
	}
}

// getRandomNumber функция получения случайного числа
func getRandomNumber(maxValue int) (random int) {
	/* рандомизация по времени */
	rand.Seed(time.Now().Unix())

	return rand.Intn(maxValue) + 1
}
