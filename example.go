package main

import "fmt"

func main() {
	var1, vars2, vars3 := colors()
	fmt.Println(var1, vars2, vars3,)
}


func colors() (string, string, string) {

	return "red", "blue", "green"

}

