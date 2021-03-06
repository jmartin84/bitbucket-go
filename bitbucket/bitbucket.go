package bitbucket

import (
	"golang.org/x/oauth2"
)

type auth struct {
	appID, secret  string
	user, password string
	token          oauth2.Token
	bearerToken    string
}

// Link represents a single link object from Bitbucket object links.
type Link struct {
	Name *string `json:"name,omitempty"`
	HRef *string `json:"href,omitempty"`
}

// Content represents content found in a Bitbucket resource.
type Content struct {
	Raw    *string `json:"raw,omitempty"`
	Markup *string `json:"markup,omitempty"`
	HTML   *string `json:"html,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// PaginationInfo represents the pagination data returned on most LIST functions.
//
// Bitbucket API Docs: https://developer.atlassian.com/bitbucket/api/2/reference/meta/pagination
type PaginationInfo struct {
	// Page number of the current results. This is an optional element that is not provided in all responses.
	Page *int64 `json:"page,omitempty"`

	//  Link to the next page if it exists. The last page of a collection does not have this value.
	//  Use this link to navigate the result set and refrain from constructing your own URLs.
	Next *string `json:"next,omitempty"`

	// Current number of objects on the existing page.
	Pagelen *int64 `json:"pagelen,omitempty"`

	// Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Size *int64 `json:"size,omitempty"`

	//Link to previous page if it exists. A collections first page does not have this value.
	// This is an optional element that is not provided in all responses.
	// Some result sets strictly support forward navigation and never provide previous links.
	// Clients must anticipate that backwards navigation is not always available.
	// Use this link to navigate the result set and refrain from constructing your own URLs.
	Previous *string `json:"previous,omitempty"`
}
