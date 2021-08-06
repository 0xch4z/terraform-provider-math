package testacc

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"text/template"
)

const templateDirName = "tmpl"

var globalTemplates = make(map[string]*template.Template)

func init() {
	// parse global templates; store them in globalTemplates
}

func LoadTemplate(t *testing.T, name string) string {
	t.Helper()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
		return ""
	}

	templatePath := filepath.Join(wd, templateDirName, name)
	fmt.Fprintf(os.Stderr, "templatePath = %s\n", templatePath)
	// get the template!

	return ""
}
