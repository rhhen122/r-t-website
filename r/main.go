package main

import (
	"fmt"
	"os/exec"
)

var approve string

func main() {
	for {
		fmt.Printf(`
	   r-t
	 ,     ,
	< \,_./ >

	_0_____0_
	|> _    |
	|_______|

	r rice
	d download
	h help
	`)
		fmt.Print("> ")
		fmt.Scan(&approve)
		if approve == "d" {
			cmd := exec.Command("python3", "installer.py")
			cmd.Output()
		}
		if approve == "c" {
			break
		}
		if approve == "ls" {
			ls := exec.Command("ls", "rice")
			ls.Output()
		}
		if approve == "r" {
			ricer := exec.Command("python3", "rice.py")
			ricer.Output()
		}
		if approve == "h" {
			fmt.Printf(`
	k clear
	c exit
	ls list
			`)
		}
		if approve == "k" {
			clear := exec.Command("clear")
			clear.Output()
		}
	}
}
