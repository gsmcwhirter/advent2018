package day1

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
	"github.com/pkg/errors"
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

	var freq int64
	freqs := map[int64]bool{
		0: true,
	}

	for {
		for _, e := range entries {
			if bytes.Equal(e, []byte("")) {
				continue
			}

			delta, err := strconv.ParseInt(string(e), 10, 64)
			if err != nil {
				return errors.Wrap(err, "could not parse int")
			}

			freq += delta

			if freqs[freq] {
				fmt.Println(freq)
				return nil
			}

			freqs[freq] = true
		}
	}
}
