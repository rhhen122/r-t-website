package main

import (
	"fmt"
	"os/exec"
)

var approve string
var dead bool = false

func main() {
	fmt.Printf(`
	   r-t [setup]
	 ,     ,
	< \,_./ >

	_0_____0_
	|> _    |
	|_______|
	`)
	fmt.Print("> ")
	fmt.Scan(&approve)
	cmd := exec.Command("python3", "r-t-build.py")
	output, err := cmd.Output()
	fmt.Println(string(output))
	if dead {
		fmt.Println("Error:", err)
	}
}
