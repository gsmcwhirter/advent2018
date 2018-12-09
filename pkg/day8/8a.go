package day8

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

type node struct {
	children []node
	metadata []int
}

func parse(input []string) (node, int, error) {
	if len(input) < 2 {
		return node{}, 0, errors.New("not enough data to parse")
	}

	childCt, err := strconv.Atoi(input[0])
	if err != nil {
		return node{}, 0, err
	}

	metaCt, err := strconv.Atoi(input[1])
	if err != nil {
		return node{}, 0, err
	}

	n := node{
		children: make([]node, 0, childCt),
		metadata: make([]int, 0, metaCt),
	}

	offset := 2

	for i := 0; i < childCt; i++ {
		child, offs, err := parse(input[offset:])
		if err != nil {
			return node{}, 0, err
		}

		n.children = append(n.children, child)
		offset += offs
	}

	for i := 0; i < metaCt; i++ {
		md, err := strconv.Atoi(input[offset+i])
		if err != nil {
			return node{}, 0, err
		}
		n.metadata = append(n.metadata, md)
	}

	return n, offset + metaCt, nil
}

func sumMeta(n node) int {
	sum := 0

	for _, c := range n.children {
		sum += sumMeta(c)
	}

	for _, v := range n.metadata {
		sum += v
	}

	return sum
}

func RunA(args []string) error {
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

	fmt.Println(sumMeta(n))

	return nil
}
