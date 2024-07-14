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
	args := os.Args
	reverse := false
	inputfile := ""
	outputfile := ""
	for index, val := range args {
		switch val {
		case "-r":
			reverse = true
		case "-in":
			if len(args)-1 == index {
				fmt.Println("No input file specified with -o")
				return
			}
			inputfile = args[index+1]
			break
		case "-out":
			if len(args)-1 == index {
				fmt.Println("No output file specified with -o")
				return
			}
			outputfile = args[index+1]
			break
		}
	}

	bytes, err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !reverse {
		bytes = Ceaser(bytes)
	} else {
		bytes = ReverseCeaser(bytes)
	}
	if outputfile == "" {
		err = ioutil.WriteFile("invisible.ghost", bytes, 0777)
	} else {
		err = ioutil.WriteFile(outputfile, bytes, 0777)
	}
	if err != nil {
		fmt.Println(err)
	}
}
