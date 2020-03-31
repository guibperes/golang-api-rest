package post

// Post entity model
type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  User   `json:"author"`
}

// User entity model
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
