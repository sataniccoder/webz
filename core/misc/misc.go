package misc

import "fmt"

func Print(data string, silent bool) {
	if !silent {
		fmt.Println(data)
	}
}
