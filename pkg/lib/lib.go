package lib

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

func GetContents(fname string) ([]byte, error) {
	input, err := os.Open(fname)
	if err != nil {
		return nil, errors.Wrap(err, "could not open input file")
	}
	defer input.Close()

	contents, err := ioutil.ReadAll(input)
	return contents, errors.Wrap(err, "could not read input contents")
}
