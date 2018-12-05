package day4

import "fmt"

func RunB(args []string) error {
	guardMinutes, err := parseGuardLines(args)
	if err != nil {
		return err
	}

	maxGuard := ""
	maxGuardMaxMinute := -1
	maxGuardMaxMinuteCt := -1
	for guard, minutes := range guardMinutes {
		for min, ct := range minutes {
			if ct > maxGuardMaxMinuteCt {
				maxGuard = guard
				maxGuardMaxMinute = min
				maxGuardMaxMinuteCt = ct
			}
		}
	}

	fmt.Printf("%s %d %d\n", maxGuard, maxGuardMaxMinute, maxGuardMaxMinuteCt)

	return nil
}
