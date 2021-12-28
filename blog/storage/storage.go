package storage

type Post struct {
	ID           int64  `db:"id"`
	Description  string `db:"description"`
	Title        string `db:"title"`
	CategoryId   int64  `db:"category_id"`
	Image        string `db:"image"`
	IsCompleted  bool   `db:"is_completed"`
	CategoryName string `db:"name"`
}

type Category struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
}
