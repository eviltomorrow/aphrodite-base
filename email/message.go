package email

// Message msg
type Message struct {
	From        Contact
	To          []Contact
	Cc          []Contact
	Bcc         []Contact
	Subject     string
	Body        string
	ContentType string
	Attach      []File
}

// Contact contact
type Contact struct {
	Name    string
	Address string
}

// File file
type File struct {
	Path      string
	AliasName string
}
