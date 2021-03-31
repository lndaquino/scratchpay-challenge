package entity

// Practice models a practice in the system
type Practice struct {
	ID            string
	Category      string
	Name          string
	StateName     string
	StateCode     string
	OpeningHour   int
	OpeningMinute int
	ClosureHour   int
	ClosureMinute int
}
