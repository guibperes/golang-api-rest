package post

// Service definition to access functions
type Service struct{}

// ServiceBuilder build the service object
func ServiceBuilder() Service {
	return Service{}
}

var database = []Post{}

// Save a new post
func (s Service) Save(post Post) []Post {
	database = append(database, post)
	return database
}

// UpdateByID a post
func (s Service) UpdateByID(id int, post Post) Post {
	database[id] = post
	return database[id]
}

// DeleteByID a post
func (s Service) DeleteByID(id int) {
	database = append(database[:id], database[id+1:]...)
}

// GetAll saved posts
func (s Service) GetAll() []Post {
	return database
}

// GetByID a post
func (s Service) GetByID(id int) Post {
	return database[id]
}

// GetLength returns the length of the list
func (s Service) GetLength() int {
	return len(database)
}
