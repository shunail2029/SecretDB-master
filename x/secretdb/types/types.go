package types

import "errors"

// flags
const (
	FlagChildCount   = "child-count"
	FlagChildURL     = "child-url"
	FlagChaldChainID = "child-chainid"
)

// child chain params
var (
	ChildCount    int
	ChildURLs     []string
	ChildChainIDs []string
)

// SetChildParams ...
func SetChildParams(count int, urls, chainIDs []string) error {
	if count != len(urls) {
		return errors.New("child-count should be equal to length of child-urls")
	}

	ChildCount = count
	ChildURLs = urls
	ChildChainIDs = chainIDs
	return nil
}
