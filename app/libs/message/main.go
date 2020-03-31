package message

// Message Data Transfer Object(DTO)
type Message struct {
	Message string `json:"message"`
}

// Builder build the service object
func Builder(message string) Message {
	return Message{Message: message}
}
