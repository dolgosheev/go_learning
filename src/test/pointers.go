// Package test Package  теста работы указателей
package test

import "fmt"

// GetMe функция для обработки указателя на строку
func GetMe(value *string) {
	*value = fmt.Sprintf("%s, World!", *value)
}
