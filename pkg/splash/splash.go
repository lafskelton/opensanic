package splash

import (
	"fmt"
	"time"
)

//** This is just for fun, I plan to leave it on if I ever open-source the software.
// 	 I would turn it off for production releases

//Print splash
func Print() {
	//Lol

	x := []string{
		"    ▄████████    ▄████████ ███▄▄▄▄    ▄█   ▄████████ ████████▄  ▀█████████▄  ",
		"   ███    ███   ███    ███ ███▀▀▀██▄ ███  ███    ███ ███   ▀███   ███    ███ ",
		"   ███    █▀    ███    ███ ███   ███ ███▌ ███    █▀  ███    ███   ███    ███ ",
		"   ███          ███    ███ ███   ███ ███▌ ███        ███    ███  ▄███▄▄▄██▀  ",
		" ▀███████████ ▀███████████ ███   ███ ███▌ ███        ███    ███ ▀▀███▀▀▀██▄  ",
		"     	  ███   ███    ███ ███   ███ ███  ███    █▄  ███    ███   ███    ██▄ ",
		"    ▄█    ███   ███    ███ ███   ███ ███  ███    ███ ███   ▄███   ███    ███ ",
		"   ▄████████▀    ███    █▀   ▀█   █▀  █▀   ████████▀  ████████▀  ▄█████████▀ ",
	}

	fmt.Println()
	fmt.Println("==============================================================================")
	fmt.Println()
	for _, v := range x {
		time.Sleep(time.Millisecond * 25)
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println("SanicDB - A database project by: http://github.com/lafskelton")
	fmt.Println()
	fmt.Println("==============================================================================")
	fmt.Println()
	time.Sleep(time.Millisecond * 250)
	return
}
