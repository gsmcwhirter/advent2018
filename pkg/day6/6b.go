package day6

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

	const threshold = 10000
	numInRegion := 0
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			t := point{x: x, y: y}

			totalDist := 0

			for _, p := range points {
				totalDist += p.dist(t)
			}

			if totalDist < threshold {
				numInRegion++
			}
		}
	}

	fmt.Println(numInRegion)

	return nil
}
