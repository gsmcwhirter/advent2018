package day8

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

func valueOf(n node) int {
	if len(n.children) == 0 {
		return sumMeta(n)
	}

	v := 0
	for _, i := range n.metadata {
		if i == 0 {
			continue
		}

		if i > len(n.children) {
			continue
		}

		v += valueOf(n.children[i-1])
	}

	return v
}

func RunB(args []string) error {
	if len(args) < 1 {
		return errors.New("need input file")
	}

	contents, err := lib.GetContents(args[0])
	if err != nil {
		return err
	}

	entries := strings.Split(string(contents), " ")

	n, _, err := parse(entries)
	if err != nil {
		return err
	}

	fmt.Println(valueOf(n))

	return nil
}
