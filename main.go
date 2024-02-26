/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"app_blade/cmd"
	"app_blade/pkg/thread"
)

func main() {
	defer thread.Recover()
	cmd.Execute()
}
