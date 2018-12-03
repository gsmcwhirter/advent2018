package day2

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

func oneDiff(a, b []byte) (i int) {
	i = -1

	if len(a) != len(b) {
		return
	}

	for j := 0; j < len(a); j++ {
		if a[j] == b[j] {
			continue
		}

		if i != -1 {
			return -1
		}

		i = j
	}

	return
}

func RunB(args []string) error {
	if len(args) < 1 {
		return errors.New("need input file")
	}

	contents, err := lib.GetContents(args[0])
	if err != nil {
		return err
	}

	entries := bytes.Split(contents, []byte{'\n'})

	for i, ei := range entries {
		for _, ej := range entries[i+1:] {
			if di := oneDiff(ei, ej); di != -1 {
				fmt.Printf("%s%s\n", string(ei[:di]), string(ei[di+1:]))
				return nil
			}
		}
	}

	return errors.New("no match found")
}
