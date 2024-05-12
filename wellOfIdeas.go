package main

func Well(x []string) string {
	count := 0
	for _, idea := range x {
		switch idea {
		case "good":
			count++
		}
	}
	switch {
	case count > 2:
		return "I smell a series!"
	case count > 0:
		return "Publish!"
	default:
		return "Fail!"
	}
}
