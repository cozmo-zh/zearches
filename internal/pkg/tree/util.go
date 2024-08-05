// Package tree .
package tree

import (
	"path"
	"path/filepath"
	"runtime"
)

func GetTemplate() string {
	const (
		templatePath = "./api"
		file         = "node.tmpl"
	)
	_, b, _, _ := runtime.Caller(0)
	pwd := filepath.Dir(b)
	p := path.Join(pwd, templatePath, file)
	return p
}
