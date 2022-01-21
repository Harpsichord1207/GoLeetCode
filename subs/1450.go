package subs

func busyStudent(startTime []int, endTime []int, queryTime int) int {

	r := 0

	for i, s := range startTime {
		e := endTime[i]

		if s <= queryTime && queryTime <= e {
			r ++
		}
	}

	return r
}

