package types

import "errors"

// Flags
const (
	FlagChildCount = "child-count"
	FlagChildUrls  = "child-urls"
)

// child chain params
var (
	ChildCount int
	ChildUrls  []string
)

// SetChildParams ...
func SetChildParams(count int, urls []string) error {
	if count != len(urls) {
		return errors.New("child-count should be equal to length of child-urls")
	}

	ChildCount = count
	ChildUrls = urls
	return nil
}
