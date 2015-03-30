package controllers

import "path"

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func viewPath(elem ...string) string {
	components := append([]string{"views"}, elem...)
	return path.Join(components...)
}

func layoutPath(elem ...string) string {
	components := append([]string{"views", "layouts"}, elem...)
	return path.Join(components...)
}
