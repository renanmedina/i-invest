package utils

import (
	"bufio"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	selected_option, _ := reader.ReadString('\n')
	selected_option = strings.TrimSpace(selected_option)
	return selected_option
}

func ReadOption() uint64 {
	option, _ := strconv.ParseUint(ReadLine(), 10, 64)
	return option
}

func ReadFloat() float64 {
	number, _ := strconv.ParseFloat(ReadLine(), 32)
	return number
}

func ClearConsole() {
	command := exec.Command("clear")
	command.Stdout = os.Stdout
	command.Run()
}
