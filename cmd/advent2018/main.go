package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gsmcwhirter/advent2018/pkg/day1"
	"github.com/gsmcwhirter/advent2018/pkg/day2"
	"github.com/gsmcwhirter/advent2018/pkg/day3"
	"github.com/gsmcwhirter/advent2018/pkg/day4"
	"github.com/gsmcwhirter/advent2018/pkg/day5"
	"github.com/gsmcwhirter/advent2018/pkg/day6"
	"github.com/gsmcwhirter/advent2018/pkg/day7"
	"github.com/gsmcwhirter/advent2018/pkg/day8"
	"github.com/gsmcwhirter/advent2018/pkg/day9"
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
	case "4a":
		return reportErr(day4.RunA(args))
	case "4b":
		return reportErr(day4.RunB(args))
	case "5a":
		return reportErr(day5.RunA(args))
	case "5b":
		return reportErr(day5.RunB(args))
	case "6a":
		return reportErr(day6.RunA(args))
	case "6b":
		return reportErr(day6.RunB(args))
	case "7a":
		return reportErr(day7.RunA(args))
	case "7b":
		return reportErr(day7.RunB(args))
	case "8a":
		return reportErr(day8.RunA(args))
	case "8b":
		return reportErr(day8.RunB(args))
	case "9a":
		return reportErr(day9.RunA(args))
	case "9b":
		return reportErr(day9.RunB(args))
	default:
		return reportErr(errors.New("missing day command"))
	}
}
