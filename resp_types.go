package facade

import "github.com/gin-gonic/gin"

type ChildRef struct {
	Path  string `json:"path"`
	Title string `json:"title"`
}

type ViewSuggests struct {
	// TODO
}

type Response struct {
	Description  string        `json:"description"`
	Content      gin.H         `json:"content"`
	Attributes   gin.H         `json:"attributes"`
	ViewSuggests *ViewSuggests `json:"suggests"`
	Children     []ChildRef    `json:"children"`
}
