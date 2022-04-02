package utils

import (
	"fmt"
	"log"
	"runtime"
)

func trace() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return file, line, "?"
	}

	return file, line, fn.Name()
}

func HandleError(err error) {

	if err != nil {
		fmt.Printf("{ERROR}: %s", err.Error())
		log.Fatal(err)
	}
}
