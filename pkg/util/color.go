package util

import "fmt"

const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func Colorize(text string, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}
