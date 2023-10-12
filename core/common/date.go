package common

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func NewDate(y, m, d int) *Date {

	return &Date{y, m, d}

}
