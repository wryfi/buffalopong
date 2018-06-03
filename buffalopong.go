package buffalopong

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/flosch/pongo2"
)

var templateCache = map[string]*pongo2.Template{}
var moot = &sync.Mutex{}

// Pongo2Renderer implements the render.TemplateEngine interface allowing pongo2 to be used as a template engine
// for Buffalo
func Pongo2Renderer(input string, data map[string]interface{}, helpers map[string]interface{}) (string, error) {
	template, err := Parse(input)
	if err != nil {
		return "", err
	}
	context := pongo2.Context{}
	if data != nil {
		// Add data to context
		for key, value := range data {
			context[key] = value
		}
	}
	if helpers != nil {
		// Add helpers to context
		for key, value := range helpers {
			context[key] = value
		}
	}
	// Execute template
	out, err := template.Execute(context)
	if err != nil {
		return "", err
	}
	return out, nil
}

// Parse an input string and return a *pongo2.Template, and caches the parsed template.
func Parse(input string) (*pongo2.Template, error) {
	moot.Lock()
	defer moot.Unlock()
	if template, ok := templateCache[input]; ok {
		return template, nil
	}
	template, err := pongo2.FromString(input)

	if err == nil {
		templateCache[input] = template
	}

	if err != nil {
		return template, errors.WithStack(err)
	}
	return template, nil
}
