package bosster

// Post is universe representation of any content
type Post struct {
	ID string

	Body string

	Images    []byte
	ImageURLs []string
}
