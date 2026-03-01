package generator

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thumbrise/demo/golang-demo/internal/modules/homepage/infrastucture/generator/templates"
)

// RouteInfo представляет данные для шаблона
type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	MethodClass string // get, post, put, delete, etc.
	IsGET       bool
	IsWildcard  bool
	BasePath    string
}

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (h *Generator) Write(routes gin.RoutesInfo, writer io.Writer) error {
	var routeInfos []RouteInfo

	for _, r := range routes {
		methodClass := strings.ToLower(r.Method)
		if r.Method != http.MethodGet {
			methodClass = "post"
		}

		isGET := r.Method == http.MethodGet
		isWildcard := strings.Contains(r.Path, "*")

		basePath := ""
		if isWildcard {
			basePath = strings.Split(r.Path, "*")[0]
			if basePath == "" {
				basePath = "/"
			}
		}

		routeInfos = append(routeInfos, RouteInfo{
			Method:      r.Method,
			Path:        r.Path,
			Handler:     r.Handler,
			MethodClass: methodClass,
			IsGET:       isGET,
			IsWildcard:  isWildcard,
			BasePath:    basePath,
		})
	}

	// Use the embedded filesystem
	tmpl, err := template.ParseFS(templates.FS, "template.html")
	if err != nil {
		return fmt.Errorf("template parse: %w", err)
	}

	err = tmpl.Execute(writer, map[string]interface{}{
		"Routes": routeInfos,
	})
	if err != nil {
		return fmt.Errorf("template execute: %w", err)
	}

	return nil
}
