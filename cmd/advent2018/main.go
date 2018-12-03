package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gsmcwhirter/advent2018/pkg/day1"
	"github.com/gsmcwhirter/advent2018/pkg/day2"
	"github.com/gsmcwhirter/advent2018/pkg/day3"
)

func main() {
	os.Exit(run())
}

func reportErr(err error) int {
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		return 1
	}
	return 0
}

func run() int {
	if len(os.Args) < 2 {
		fmt.Println("Need a command to run")
		return 1
	}

	cmd, args := os.Args[1], os.Args[2:]
	switch cmd {
	case "1a":
		return reportErr(day1.RunA(args))
	case "1b":
		return reportErr(day1.RunB(args))
	case "2a":
		return reportErr(day2.RunA(args))
	case "2b":
		return reportErr(day2.RunB(args))
	case "3a":
		return reportErr(day3.RunA(args))
	case "3b":
		return reportErr(day3.RunB(args))
	default:
		return reportErr(errors.New("missing day command"))
	}

	return 0
}
