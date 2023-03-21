package main

import "os/exec"

func main() { 
	cmd := exec.Command("xdg-open", "../index.html")
	cmd.Start()
}