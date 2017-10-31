package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var ViewsDir string = "views/templates"
var ViewsExt string = ".gohtml"
var LayoutDir string = ViewsDir + "/layouts"

func NewView(layout string, files ...string) *View {
	addViewsDirPrefix(files)
	addViewsExtSuffix(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, nil)
}

// ViewsExt if we need to change our template extension
// we simply change the ViewsExt varibles.
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*" + ViewsExt)
	if err != nil {
		panic(err)
	}
	return files
}

// Takes in a slice of strings representing file paths
// and prepends the ViewsDir directory to ecach string in
// the slice.
func addViewsDirPrefix(files []string) {
	for i, f := range files {
		files[i] = ViewsDir + "/" + f
	}
}

// Takes in a slice of strings representing file paths
// and appends the ViewsExt extension to each string in the slice.
func addViewsExtSuffix(files []string) {
	for i, f := range files {
		files[i] = f + ViewsExt
	}
}
