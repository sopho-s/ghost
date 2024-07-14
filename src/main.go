package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("              ('-. .-.                .-')     .-') _    ")
	fmt.Println("             ( OO )  /               ( OO ).  (  OO) )   ")
	fmt.Println("  ,----.     ,--. ,--.  .-'),-----. (_)---\\_) /     '._  ")
	fmt.Println(" '  .-./-')  |  | |  | ( OO'  .-.  '/    _ |  |'--...__) ")
	fmt.Println(" |  |_( O- ) |   .|  | /   |  | |  |\\  :` `.  '--.  .--' ")
	fmt.Println(" |  | .--, \\ |       | \\_) |  |\\|  | '..`''.)    |  |    ")
	fmt.Println("(|  | '. (_/ |  .-.  |   \\ |  | |  |.-._)   \\    |  |    ")
	fmt.Println(" |  '--'  |  |  | |  |    `'  '-'  '\\       /    |  |    ")
	fmt.Println("  `------'   `--' `--'      `-----'  `-----'     `--'    ")
	_ = os.Args

	bytes, err := ioutil.ReadFile("test.txt") //read the content of file
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bytes)
	bytes = Ceaser(bytes)
	fmt.Println(bytes)
	err = ioutil.WriteFile("invisible.ghost", bytes, 0777)
	if err != nil {
		fmt.Println(err)
	}
}
