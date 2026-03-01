package generator

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"sort"
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

func (g *Generator) Write(routes gin.RoutesInfo, writer io.Writer) error {
	g.sort(routes)

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
func (g *Generator) sort(routes gin.RoutesInfo) {
	sort.Slice(routes, func(i, j int) bool {
		getGroup := func(path string) string {
			if path == "" || path == "/" {
				return "/"
			}
			trimmed := strings.TrimPrefix(path, "/")
			parts := strings.SplitN(trimmed, "/", 2)
			return "/" + parts[0]
		}
		gi := getGroup(routes[i].Path)
		gj := getGroup(routes[j].Path)
		if gi != gj {
			return gi < gj
		}

		return routes[i].Path < routes[j].Path
	})
}
