package main

import (
	"fmt"
	"os"
	"time"

	"log"

	"github.com/cloudogu/spinners"
)

func main() {
	name := "dots"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	spinner, err := spinners.NewSpinnerByName(name, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	show(name, spinner)
}

func show(name string, spinner *spinners.Spinner) {
	fmt.Printf("Start spinner %s ...\n", name)
	spinner.Start("wait ...")

	time.Sleep(5 * time.Second)
	spinner.Stop()

	fmt.Printf("... %s done\n", name)
}
