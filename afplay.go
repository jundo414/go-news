package main

import "os/exec"

func SayText(file string) {
	exec.Command("afplay", "-r", "1.6", file).Run()
}
