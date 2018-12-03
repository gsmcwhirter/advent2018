package day3

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

var parser = regexp.MustCompile(`#(.+) @ (.+),(.+): (.+)x(.+)`)

type pair struct {
	x, y int
}

func (p pair) move(x, y int) pair {
	p.x += x
	p.y += y
	return p
}

func RunA(args []string) error {
	if len(args) < 1 {
		return errors.New("need input file")
	}

	contents, err := lib.GetContents(args[0])
	if err != nil {
		return err
	}

	entries := bytes.Split(contents, []byte{'\n'})

	seen := map[pair]int{}

	for _, e := range entries {
		match := parser.FindSubmatch(e)
		if match == nil {
			return errors.New("bad input line")
		}

		var err error
		p := pair{}
		var w, h int

		if p.x, err = strconv.Atoi(string(match[2])); err != nil {
			return err
		}

		if p.y, err = strconv.Atoi(string(match[3])); err != nil {
			return err
		}

		if w, err = strconv.Atoi(string(match[4])); err != nil {
			return err
		}

		if h, err = strconv.Atoi(string(match[5])); err != nil {
			return err
		}

		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				q := p.move(i, j)
				seen[q]++
			}
		}
	}

	duplicates := 0
	for _, v := range seen {
		if v > 1 {
			duplicates++
		}
	}

	fmt.Println(duplicates)
	return nil
}
