package gologger

import (
	"fmt"
)

const Reset string = "\033[0m"
const Red string = "\033[31m"
const Green string = "\033[32m"
const White string = "\033[97m"
const Yellow string = "\033[33m"

type Logger struct{}

func (logger Logger) Info(str string, args ...any) {
	logger.Print(White, str, args...)
}

func (logger Logger) Success(str string, args ...any) {
	logger.Print(Green, str, args...)
}

func (logger Logger) Warn(str string, args ...any) {
	logger.Print(Yellow, str, args...)
}

func (logger Logger) Error(str string, args ...any) {
	logger.Print(Red, str, args...)
}

func (Logger) Print(color string, str string, args ...any) {
	fmt.Println(color + fmt.Sprintf(str, args...) + Reset)
}
