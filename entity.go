package main

// Level ...
type Level uint

// Credit ...
type Credit uint

// Difficulty ...
type Difficulty struct {
	Level Level `json:"level"`
}

const (
	// Easy ...
	Easy Level = iota + 1
	// Medium ...
	Medium
	// Hard ...
	Hard
)
const (
	// Negative ...
	Negative Credit = iota
	// MostlyNegative ...
	MostlyNegative
	// Mixed ...
	Mixed
	// MostlyPositive ...
	MostlyPositive
	// Positive ...
	Positive
	// DonotCare ...
	DonotCare
)

// String ...
func (d Level) String() string {
	return [...]string{"Easy", "Medium", "Hard"}[d]
}

// Problem ...
type Problem struct {
	ID            uint   `json:"frontend_question_id"`
	Title         string `json:"question__title"`
	TitleSlug     string `json:"question__title_slug"`
	TotalAccept   uint   `json:"total_acs"`
	PaidToWin     bool   `json:"paid_only"`
	TotalSubmit   uint   `json:"total_submitted"`
	IsNewQuestion bool   `json:"is_new_question"`
}

// Stat ...
type Stat struct {
	Stat       Problem    `json:"stat"`
	Difficulty Difficulty `json:"difficulty"`
}

// API ...
type API struct {
	StatPair []Stat `json:"stat_status_pairs"`
}
