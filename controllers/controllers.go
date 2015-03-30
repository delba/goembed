package controllers

import (
	"path"
	"runtime"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func viewPath(elem ...string) string {
	_, __FILE__, _, _ := runtime.Caller(1)
	components := append([]string{path.Dir(__FILE__), "..", "views"}, elem...)
	return path.Join(components...)
}

func layoutPath(elem ...string) string {
	_, __FILE__, _, _ := runtime.Caller(1)
	components := append([]string{path.Dir(__FILE__), "..", "views", "layouts"}, elem...)
	return path.Join(components...)
}
