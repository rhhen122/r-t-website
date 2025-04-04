package main

import (
	"fmt"
)

var approve string

func main() {
	fmt.Printf(`
	   r-t
	 ,     ,
	< \,_./ >

	_0_____0_
	|> _    |
	|_______|
	' /   \ '
	.,    ,.
	d download
	ls list
	`)
	fmt.Print("> ")
	fmt.Scan(&approve)
}
