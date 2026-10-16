// Models of the Tag section.

package tag

// TagListResponse
type TagListResponse struct {
	// "success" string
	// Allowed values: `success`.
	Status string `json:"status,omitempty"`

	// An array of objects containing data for user-defined tags.
	Tags []TagsResponseObject `json:"tags,omitempty"`
}

// TagsResponseObject — An object with the properties of the user tags.
type TagsResponseObject struct {
	// Unique tag id.
	TagID int `json:"tag_id,omitempty"`

	// Tag name.
	Tag string `json:"tag,omitempty"`
}

// TagDeleteRequest
type TagDeleteRequest struct {
	// Unique tag id.
	TagID int `json:"tag_id"`
}
