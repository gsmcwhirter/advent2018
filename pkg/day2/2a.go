package day2

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

func RunA(args []string) error {
	if len(args) < 1 {
		return errors.New("need input file")
	}

	contents, err := lib.GetContents(args[0])
	if err != nil {
		return err
	}

	entries := bytes.Split(contents, []byte{'\n'})

	has2 := 0
	has3 := 0

	for _, e := range entries {
		cts := map[byte]int{}

		for _, c := range e {
			cts[c]++
		}

		saw2 := false
		saw3 := false

		for _, v := range cts {
			if v == 2 && !saw2 {
				has2++
				saw2 = true
				continue
			}

			if v == 3 && !saw3 {
				has3++
				saw3 = true
				continue
			}
		}
	}

	fmt.Println(has2 * has3)
	return nil
}
