package day6

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

type point struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p point) dist(other point) int {
	return abs(p.x-other.x) + abs(p.y-other.y)
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

	points := make([]point, 0, len(entries))
	var minX, maxX, minY, maxY int

	for _, e := range entries {
		parts := bytes.SplitN(e, []byte{',', ' '}, 2)
		var err error
		p := point{}
		p.x, err = strconv.Atoi(string(parts[0]))
		if err != nil {
			return err
		}

		p.y, err = strconv.Atoi(string(parts[1]))
		if err != nil {
			return err
		}

		if len(points) == 0 {
			minX, maxX = p.x, p.x
			minY, maxY = p.y, p.y
		}

		points = append(points, p)
		if p.x < minX {
			minX = p.x
		}

		if p.x > maxX {
			maxX = p.x
		}

		if p.y < minY {
			minY = p.y
		}

		if p.y > maxY {
			maxY = p.y
		}
	}

	areas := map[point]int{}

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			t := point{x: x, y: y}

			nearest := points[0]
			nearestDist := points[0].dist(t)
			tie := false

			for i, p := range points {
				if i == 0 {
					continue
				}

				dist := p.dist(t)
				if dist < nearestDist {
					nearest = p
					nearestDist = dist
					tie = false
					continue
				}

				if dist == nearestDist {
					tie = true
					continue
				}
			}

			// fmt.Printf("%+v nearest to %+v (tie=%v)\n", t, nearest, tie)

			if !tie {
				areas[nearest]++
			}

			if nearestDist == 0 {
				fmt.Print("X")
			} else if tie {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	maxPt := point{}
	maxArea := -1
	for p, area := range areas {
		// on hull?
		if p.x == minX || p.x == maxX || p.y == minY || p.y == maxY {
			continue
		}

		if maxArea == -1 {
			maxPt = p
			maxArea = area
			continue
		}

		if area > maxArea {
			maxArea = area
			maxPt = p
		}
	}

	fmt.Printf("%+v: %d\n", maxPt, maxArea)

	return nil
}
