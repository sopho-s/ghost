package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args
	reverse := false
	inputfile := ""
	outputfile := ""
	crypttype := ""
	foil := false
	for index, val := range args {
		switch val {
		case "-r":
			reverse = true
			break
		case "-f":
			foil = true
			break
		case "-ct":
			crypttype = args[index+1]
			break
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
		case "--help":
			fmt.Println("\n\n              ('-. .-.                .-')     .-') _    ")
			fmt.Println("             ( OO )  /               ( OO ).  (  OO) )   ")
			fmt.Println("  ,----.     ,--. ,--.  .-'),-----. (_)---\\_) /     '._  ")
			fmt.Println(" '  .-./-')  |  | |  | ( OO'  .-.  '/    _ |  |'--...__) ")
			fmt.Println(" |  |_( O- ) |   .|  | /   |  | |  |\\  :` `.  '--.  .--' ")
			fmt.Println(" |  | .--, \\ |       | \\_) |  |\\|  | '..`''.)    |  |    ")
			fmt.Println("(|  | '. (_/ |  .-.  |   \\ |  | |  |.-._)   \\    |  |    ")
			fmt.Println(" |  '--'  |  |  | |  |    `'  '-'  '\\       /    |  |    ")
			fmt.Println("  `------'   `--' `--'      `-----'  `-----'     `--'    ")
			fmt.Println("\n")
			fmt.Println("Options:\n")
			fmt.Println("\t-r : revive, turns a ghost file back to its original form\n")
			fmt.Println("\t-f : foil, wraps the ghost file so that when executed the original file will be restored, no need to use ghost again\n")
			fmt.Println("\t-ct <type> : choose the method that the file will be turned to a ghost\n")
			fmt.Println("IO:\n")
			fmt.Println("\t-out : output file\n")
			fmt.Println("\t-in : input file\n")
			return
		}
	}

	bytes, err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !reverse {
		if crypttype == "ceaser" {
			bytes = Ceaser(bytes)
		} else {
			fmt.Println("No crypt type was chosen")
			return
		}
	} else {
		if crypttype == "ceaser" {
			bytes = ReverseCeaser(bytes)
		} else {
			fmt.Println("No crypt type was chosen")
			return
		}
	}
	if !foil {
		if outputfile == "" {
			err = ioutil.WriteFile("invisible.ghost", bytes, 0777)
		} else {
			err = ioutil.WriteFile(outputfile, bytes, 0777)
		}
	} else {
		if crypttype == "ceaser" {
			arr := fmt.Sprint(bytes)
			arr = arr[1:]
			arr = arr[:len(arr)-1]
			arr = strings.Replace(arr, " ", ",", -1)
			text := ""
			if outputfile == "" {
				text = "package main\nimport (\"os/exec\"\n\"io/ioutil\")\nfunc main() {\nexe := []byte{" + arr + "} \nfor index, currbyte := range exe {\nexe[index] = byte((int(currbyte) + (256 - index)) % 256)\n}\n_ = ioutil.WriteFile(\"ghost.exe\", exe, 0777)\ncommand := exec.Command(\"./revived\")\ncommand.Run()\n}"
			} else {
				text = "package main\nimport (\"os/exec\"\n\"io/ioutil\")\nfunc main() {\nexe := []byte{" + arr + "} \nfor index, currbyte := range exe {\nexe[index] = byte((int(currbyte) + (256 - index)) % 256)\n}\n_ = ioutil.WriteFile(\"ghost.exe\", exe, 0777)\ncommand := exec.Command(\"./" + outputfile + "\")\ncommand.Run()\n}"
			}
			if outputfile == "" {
				err = ioutil.WriteFile("invisible.go", []byte(text), 0777)
				command := exec.Command("go", "mod", "init", "main")
				command.Run()
				command = exec.Command("go", "build", "-o", "invisible")
				command.Run()
			} else {
				err = ioutil.WriteFile("invisible.go", []byte(text), 0777)
				command := exec.Command("go", "mod", "init", "main")
				command.Run()
				command = exec.Command("go", "build", "-o", outputfile)
				command.Run()
			}
			os.Remove("invisible.go")
			os.Remove("go.mod")
		}
	}
	if err != nil {
		fmt.Println(err)
	}
}
