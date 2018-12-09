package day9

import (
	"errors"
	"fmt"
)

const players = 446
const lastMarble = 71522

type node struct {
	val   int
	left  *node
	right *node
}

func (n *node) moveLeft(steps int) (*node, error) {
	if steps == 0 {
		return n, nil
	}

	var next *node
	var nextSteps int

	if steps > 0 {
		next = n.left
		nextSteps = steps - 1
	} else {
		next = n.right
		nextSteps = steps + 1
	}

	if next == nil {
		return nil, errors.New("no next")
	}

	return next.moveLeft(nextSteps)
}

func (n *node) moveRight(steps int) (*node, error) {
	return n.moveLeft(-steps)
}

func (n *node) pop() (popped, current *node, err error) {
	if n.left == nil || n.right == nil {
		return nil, nil, errors.New("non-circle")
	}

	n.left.right = n.right
	n.right.left = n.left
	return n, n.right, nil
}

func (n *node) insert(val int) (*node, error) {
	newN := &node{
		val:   val,
		left:  nil,
		right: nil,
	}

	left, err := n.moveRight(1)
	if err != nil {
		return nil, err
	}

	right, err := left.moveRight(1)
	if err != nil {
		return nil, err
	}

	left.right = newN
	newN.left = left
	right.left = newN
	newN.right = right

	return newN, nil
}

func RunA(args []string) error {
	current := &node{
		val:   0,
		left:  nil,
		right: nil,
	}

	current.left = current
	current.right = current

	scores := map[int]int{}

	var err error
	var removed *node

	currPlayer := -1
	for marble := 1; marble <= lastMarble; marble++ {
		currPlayer++
		currPlayer %= players

		if marble%23 != 0 {
			current, err = current.insert(marble)
			if err != nil {
				return err
			}
			continue
		}

		scores[currPlayer] += marble
		removed, err = current.moveLeft(7)
		if err != nil {
			return err
		}

		removed, current, err = removed.pop()
		if err != nil {
			return err
		}

		scores[currPlayer] += removed.val
	}

	maxScore := -1
	for _, sc := range scores {
		if sc > maxScore {
			maxScore = sc
		}
	}

	fmt.Println(maxScore)
	return nil
}
