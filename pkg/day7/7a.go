package day7

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

var parser = regexp.MustCompile(`Step (.+) must be finished before step (.+) can begin`)

func remove(list []string, elem string) []string {
	ret := make([]string, 0, len(list))
	for _, e := range list {
		if e != elem {
			ret = append(ret, e)
		}
	}

	return ret
}

func RunA(args []string) error {
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

	order := make([]string, 0, len(allLetters))

MAIN_LOOP:
	for len(lettersLeft) > 0 {
		for _, e := range lettersLeft {
			if len(deps[e]) > 0 {
				continue
			}

			order = append(order, e)
			lettersLeft = remove(lettersLeft, e)
			for _, ds := range downstreams[e] {
				deps[ds] = remove(deps[ds], e)
			}

			continue MAIN_LOOP
		}
	}

	for _, e := range order {
		fmt.Print(e)
	}
	fmt.Println()

	return nil
}

func RunAToposort(args []string) error {
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

	shells := make([][]string, 0, len(allLetters))
	for len(lettersLeft) > 0 {
		fmt.Println(lettersLeft)
		for _, e := range lettersLeft {
			fmt.Printf("%s: %v\n", e, deps[e])
		}

		shell := make([]string, 0, len(lettersLeft))

		for _, e := range lettersLeft {
			if len(deps[e]) > 0 {
				continue
			}

			shell = append(shell, e)
		}

		fmt.Printf("shell: %v\n\n", shell)

		shells = append(shells, shell)

		for _, e := range shell {
			lettersLeft = remove(lettersLeft, e)
			for _, ds := range downstreams[e] {
				deps[ds] = remove(deps[ds], e)
			}
		}
	}

	for _, shell := range shells {
		for _, e := range shell {
			fmt.Print(e)
		}
	}

	fmt.Println()

	return nil
}
