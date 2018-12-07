package day7

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

type task struct {
	remaining int
	letter    string
}

func RunB(args []string) error {
	if len(args) < 1 {
		return errors.New("need input file")
	}

	contents, err := lib.GetContents(args[0])
	if err != nil {
		return err
	}

	entries := strings.Split(string(contents), "\n")

	allLetters := map[string]bool{}
	deps := map[string][]string{}
	downstreams := map[string][]string{}

	for _, e := range entries {
		match := parser.FindStringSubmatch(e)
		if match == nil || len(match) < 3 {
			return errors.New("malformed line")
		}

		allLetters[match[1]] = true
		allLetters[match[2]] = true
		deps[match[2]] = append(deps[match[2]], match[1])
		downstreams[match[1]] = append(downstreams[match[1]], match[2])
	}

	lettersLeft := make([]string, 0, len(allLetters))
	for k := range allLetters {
		lettersLeft = append(lettersLeft, k)
	}
	sort.Strings(lettersLeft)

	workers := make([]task, 5)

	order := make([]string, 0, len(allLetters))

	totalLetters := len(allLetters)
	time := 0
	for len(order) < totalLetters {
		for i := 0; i < len(workers); i++ {
			if workers[i].letter == "" {
				continue
			}

			workers[i].remaining--
			if workers[i].remaining == 0 {
				order = append(order, workers[i].letter)
				for _, ds := range downstreams[workers[i].letter] {
					deps[ds] = remove(deps[ds], workers[i].letter)
				}
				workers[i].letter = ""
			}
		}

		for _, e := range lettersLeft {
			if len(deps[e]) > 0 {
				continue
			}

			for i := 0; i < len(workers); i++ {
				if workers[i].letter != "" {
					continue
				}

				workers[i].letter = e
				workers[i].remaining = 60 + int([]byte(e)[0]-byte('A')) + 1
				lettersLeft = remove(lettersLeft, e)
				break
			}
		}

		fmt.Printf("%d: %v\n", time, workers)

		time++
	}

	for _, e := range order {
		fmt.Print(e)
	}
	fmt.Println()
	fmt.Println(time - 1)

	return nil
}
