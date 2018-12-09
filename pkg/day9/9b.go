package day9

import "fmt"

const lastMarbleB = 7152200

func RunB(args []string) error {
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
	for marble := 1; marble <= lastMarbleB; marble++ {
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
