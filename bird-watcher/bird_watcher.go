package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirds := 0

	for i := 0; i < len(birdsPerDay); i++ {
		totalBirds += birdsPerDay[i]
	}

	return totalBirds
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	totalBirds := 0

	for i := (week - 1) * 7; i < week*7; i++ {
		totalBirds += birdsPerDay[i]
	}

	return totalBirds
	panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}

	return birdsPerDay
	panic("Please implement the FixBirdCountLog() function")
}
