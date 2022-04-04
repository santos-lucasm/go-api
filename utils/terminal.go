package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func InputHandle(msg string) string {
	fmt.Print(msg)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func ClearTerminal() {
	currOS := runtime.GOOS

	switch currOS {

	case "linux":
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Printf("ClearTerminal() call not supported on %s", currOS)
	}

}

func ShowHeader() {
	ClearTerminal()
	fmt.Println("=====================================================")
	fmt.Println("MANGADEX API PROGRAM")
	fmt.Println("Operations: ")
	fmt.Println("LOGIN\tCHECK_LOGIN\tREFRESH_TOKEN")
	fmt.Println("LOGOUT\tSCANLATION_LIST")
	fmt.Println("1. Type the operation and press ENTER to execute it")
	fmt.Println("2. Type ENTER to exit")
	fmt.Println("=====================================================")
}
