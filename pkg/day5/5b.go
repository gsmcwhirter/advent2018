package day5

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

func RunB(args []string) error {
	if len(args) < 1 {
		return errors.New("need input file")
	}

	contentsReal, err := lib.GetContents(args[0])
	if err != nil {
		return err
	}

	minCt := len(contentsReal) + 1

	lengths := map[byte]int{}

	for j := 'a'; j <= 'z'; j++ {
		contents := make([]byte, len(contentsReal))
		copy(contents, contentsReal)

		remove := byte(j)
		startStack := make([]int, 0, len(contents))
		for i := 0; i < len(contents); i++ {
			// fmt.Printf("%d %v %v %v\n", i, contents[i], string(contents[i:i+1]), startStack)

			if contents[i] == null {
				continue
			}

			if bytes.Equal(bytes.ToLower(contents[i:i+1]), []byte{remove}) {
				contents[i] = null
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

		lengths[remove] = ct
	}

	var minRemove byte

	for k, v := range lengths {
		// fmt.Printf("%v %d\n", string([]byte{k}), v)
		if v < minCt {
			minRemove = k
			minCt = v
		}
	}

	fmt.Printf("%v %d\n", string([]byte{minRemove}), minCt)
	return nil
}
