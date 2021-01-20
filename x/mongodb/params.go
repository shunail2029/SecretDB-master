package mongodb

// MongoDB start-up flag
const (
	FlagDBURI = "mongodb-uri"
)

var (
	dbURI string
)

// SetURI sets dbURI
func SetURI(uri string) {
	dbURI = uri
}
