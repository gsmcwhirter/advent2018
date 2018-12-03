package day3

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

func RunB(args []string) error {
	if len(args) < 1 {
		return errors.New("need input file")
	}

	contents, err := lib.GetContents(args[0])
	if err != nil {
		return err
	}

	entries := bytes.Split(contents, []byte{'\n'})

	seen := map[pair][]string{}
	dupes := map[string]bool{}

	for _, e := range entries {
		match := parser.FindSubmatch(e)
		if match == nil {
			return errors.New("bad input line")
		}

		id := string(match[1])
		dupes[id] = false

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
				seen[q] = append(seen[q], string(match[1]))

				if len(seen[q]) > 1 {
					for _, dup := range seen[q] {
						dupes[dup] = true
					}
				}
			}
		}
	}

	for id, dup := range dupes {
		if !dup {
			fmt.Println(id)
		}
	}

	return nil
}
