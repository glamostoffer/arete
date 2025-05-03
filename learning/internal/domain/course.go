package domain

type Course struct {
	ID          int64  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Duration    string `json:"duration" db:"duration"`
	Difficulty  string `json:"difficulty" db:"difficulty"`
	Category    string `json:"category" db:"category"`
	ImageURL    string `json:"imageURL" db:"image_url"`
	IsEnrolled  bool   `json:"isEnrolled" db:"is_enrolled"`
}
