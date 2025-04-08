package main

import (
	"fmt"
)

var approve string

// var downloadask string
// var downloadask1 string

/* func main() {
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
			// i should have made this a dedicated function. what are you
			// gonna do? bite me!
			fmt.Printf(`
	> `)
			fmt.Scan(&downloadask)
			if downloadask == "ur" {
				fmt.Printf(`
	> `)
				fmt.Scan(&downloadask1)
				exec.Command("mkdir", "rice/"+downloadask1)
				fmt.Printf(`
	> `)
				// should have made it so it moves into the directory before it does all this stuff.
				// its like shooting an arrow from a km away.
				fmt.Scan(&downloadask)
				exec.Command("mkdir", "rice/"+downloadask1+"/"+downloadask)
				exec.Command("cd", "rice/"+downloadask1+"/"+downloadask)
				exec.Command("curl", "-O", "https://rt.rhhen.xyz/ur/"+downloadask1+"/"+downloadask+"/install.sh")
				// fuck. now its time to move all the fucking way back.
				exec.Command("cd", "../../../")
			}
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
} */
// yeah recoding this...
func main() {
	for {
		fmt.Printf(`
	   r-t
	 ,     ,
	< \,_./ >

	_0_____0_
	|> _    |
	|_______|
	 h help!
	 c exit!`)
		fmt.Printf(`

	> `)
		fmt.Scan(&approve)
		if approve == "h" {
			fmt.Printf(`
	r rice (to open the rice menu)
	d download (to download anything from the UR)
	u update (updates everything that needs to be)
	t health-check (runs a health check)
			`)
		}
	}
}
