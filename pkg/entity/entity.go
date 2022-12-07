package entity

// Level ...
type Level uint

// Rating ...
type Rating uint

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

// https://www.reddit.com/r/Steam/comments/ivz45n/what_does_the_steam_ratings_like_very_negative_or/
const (
	// Negative means do not care
	Negative Rating = iota
	MostlyNegative
	Mixed
	MostlyPositive
	Positive
	OverwhelminglyPositive
)

// String ...
func (d Level) String() string {
	return [...]string{"Easy", "Medium", "Hard"}[d]
}

// Problem ...
type Problem struct {
	Title         string `json:"question__title"`
	TitleSlug     string `json:"question__title_slug"`
	ID            uint   `json:"frontend_question_id"`
	TotalAccept   uint   `json:"total_acs"`
	TotalSubmit   uint   `json:"total_submitted"`
	Hide          bool   `json:"question__hide"`
	IsNewQuestion bool   `json:"is_new_question"`
}

// Stat ...
type Stat struct {
	// TODO: login
	Status     string     `json:"status,omit_empty"`
	Stat       Problem    `json:"stat"`
	Difficulty Difficulty `json:"difficulty"`
	PaidToWin  bool       `json:"paid_only"`
}

// API ...
type API struct {
	StatPair []Stat `json:"stat_status_pairs"`
}

type ProblemGraphql struct {
	Question ProblemLikes `json:"question"`
}

type ProblemLikes struct {
	TitleSlug  string `json:"titleSlug"`
	Difficulty string `json:"difficulty"`
	Likes      int    `json:"likes"`
	DisLikes   int    `json:"dislikes"`
}

/* 2022/12/07 GMT+8
{
	"stat": {
		"question_id": 1079,
		"question__article__live": true,
		"question__article__slug": "sum-root-to-leaf-binary-numbers",
		"question__article__has_video_solution": false,
		"question__title": "Sum of Root To Leaf Binary Numbers",
		"question__title_slug": "sum-of-root-to-leaf-binary-numbers",
		"question__hide": false,
		"total_acs": 180490,
		"total_submitted": 244816,
		"frontend_question_id": 1022,
		"is_new_question": false
	},
	"status": "ac",
	"difficulty": {
		"level": 1
	},
	"paid_only": false,
	"is_favor": false,
	"frequency": 0,
	"progress": 0
}
*/
