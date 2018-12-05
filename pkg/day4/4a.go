package day4

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gsmcwhirter/advent2018/pkg/lib"
)

var guardLine = regexp.MustCompile(`Guard (.+) begins shift`)

func parseGuardLines(args []string) (map[string][]int, error) {
	if len(args) < 1 {
		return nil, errors.New("need input file")
	}

	contents, err := lib.GetContents(args[0])
	if err != nil {
		return nil, err
	}

	entries := strings.Split(string(contents), "\n")
	sort.Strings(entries)

	guardMinutes := map[string][]int{}

	currentGuard := ""
	startMinute := -1

	for _, line := range entries {
		parts := strings.SplitN(line, " ", 3)
		if len(parts) != 3 {
			return nil, errors.New("bad line")
		}

		var err error
		match := guardLine.FindStringSubmatch(parts[2])
		minute, err := strconv.Atoi(parts[1][3:5])
		if err != nil {
			return nil, err
		}

		switch {
		case match != nil:
			if startMinute >= 0 {

			}
			currentGuard = match[1]

			if guardMinutes[currentGuard] == nil {
				guardMinutes[currentGuard] = make([]int, 60)
			}

		case parts[2] == "falls asleep":
			startMinute = minute
		case parts[2] == "wakes up":
			if startMinute < 0 {
				return nil, errors.New("wake without asleep")
			}

			for i := startMinute; i < minute; i++ {
				guardMinutes[currentGuard][i]++
			}

			startMinute = -1
		}
	}

	return guardMinutes, nil
}

func RunA(args []string) error {
	guardMinutes, err := parseGuardLines(args)
	if err != nil {
		return err
	}

	maxGuard := ""
	maxGuardMinutes := -1
	maxGuardMaxMinute := -1
	for guard, minutes := range guardMinutes {
		total := 0
		maxMinute := -1
		maxMinuteCt := -1

		for min, ct := range minutes {
			total += ct

			if ct > maxMinuteCt {
				maxMinute = min
				maxMinuteCt = ct
			}
		}

		if total > maxGuardMinutes {
			maxGuard = guard
			maxGuardMinutes = total
			maxGuardMaxMinute = maxMinute
		}
	}

	fmt.Printf("%s %d %d\n", maxGuard, maxGuardMinutes, maxGuardMaxMinute)

	return nil
}
