package main

type solver interface {
	solve(string) string
}

type Date struct {
	year int
	day  int
	part int
}

// func main() {
// 	tasks := map[Date]solver{
// 		Date{year: 2025, day: 1, part: 1}: Day1Part1{},
// 		Date{year: 2025, day: 1, part: 2}: Day1Part2{},
// 	}
// }
