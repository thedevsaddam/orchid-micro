package console

import (
	"fmt"
	"strconv"
)

func ConsoleSuccess(i interface{}) {
	fmt.Println("\x1b[32;1m" + ToString(i) + "\x1b[0m")
}

func ConsoleWarning(i interface{}) {
	fmt.Println("\x1b[33;1m" + ToString(i) + "\x1b[0m")
}

func ConsoleError(i interface{}) {
	fmt.Println("\x1b[31;1m" + ToString(i) + "\x1b[0m")
}

func ToString(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	default:
		return ""
	}
}
