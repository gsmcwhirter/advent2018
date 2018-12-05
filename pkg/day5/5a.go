package day5

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

var null byte

func RunA(args []string) error {
	if len(args) < 1 {
		return errors.New("need input file")
	}

	contents, err := lib.GetContents(args[0])
	if err != nil {
		return err
	}

	startStack := make([]int, 0, len(contents))
	for i := 0; i < len(contents); i++ {
		// fmt.Printf("%d %v %v %v\n", i, contents[i], string(contents[i:i+1]), startStack)

		if contents[i] == null {
			continue
		}

		if len(startStack) == 0 {
			startStack = append(startStack, i)
			continue
		}

		if i == 0 {
			continue
		}

		last := startStack[len(startStack)-1]

		if contents[last] == contents[i] {
			startStack = append(startStack, i)
			continue
		}

		if bytes.Equal(bytes.ToLower(contents[last:last+1]), bytes.ToLower(contents[i:i+1])) {
			startStack = startStack[:len(startStack)-1] //pop off the stack
			contents[last] = null
			contents[i] = null
			continue
		}

		startStack = append(startStack, i)
	}

	ct := 0
	for _, v := range contents {
		if v != null {
			ct++
		}
	}

	fmt.Println(ct)
	return nil
}
