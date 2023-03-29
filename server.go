package facade

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	archive *Archive
}

func New(archive *Archive) *Server {
	s := &Server{
		router:  gin.Default(),
		archive: archive,
	}

	s.registerHandlers()
	return s
}

func (s *Server) registerHandlers() {
	s.router.GET("/view", s.handleView)
	s.router.GET("/media", s.handleMedia)
}

func (s *Server) Run(addr string, port int) error {
	return s.router.Run(fmt.Sprintf("%s:%d", addr, port))
}

func (s *Server) parsePath(c *gin.Context) []string {
	path := c.Query("path")
	if path == "" {
		return nil
	}
	return strings.Split(path, ".")
}

func (s *Server) handleView(c *gin.Context) {
	path := s.parsePath(c)
	ar := s.archive.Walk(path...)

	if ar != nil {
		refs := make([]ChildRef, 0)

		curPath := strings.Join(path, ".")
		for _, child := range ar.Children() {
			refs = append(refs, ChildRef{
				Path:  fmt.Sprintf("%s.%s", curPath, child.Name()),
				Title: child.Content().Title(),
			})
		}
		result := Response{
			Content:    ar.Content().Render(),
			Attributes: ar.AdditionalAttributes().Render(),
			Children:   refs,
		}

		if desc, ok := ar.instance.(interface {
			Description() string
		}); ok {
			result.Description = desc.Description()
		}

		if suggests, ok := ar.instance.(interface {
			ViewSuggests() ViewSuggests
		}); ok {
			tmp := suggests.ViewSuggests()
			result.ViewSuggests = &tmp
		}

		c.JSON(http.StatusOK, result)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func (s *Server) handleMedia(c *gin.Context) {
	path := c.Query("path")

	if media, ok := s.archive.instance.(interface {
		GetMedia(path string) (contentType string, data []byte, found bool)
	}); ok {
		contentType, data, found := media.GetMedia(path)
		if found {
			c.Data(http.StatusOK, contentType, data)
			return
		}
	}
	c.Status(http.StatusNotFound)
}
