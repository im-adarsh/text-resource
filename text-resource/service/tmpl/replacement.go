package tmpl

import "strings"

func FillTmpl(tmpl string, values map[string]string) string {

	if values == nil {
		return tmpl
	}

	for k, v := range values {
		tmpl = strings.Replace(tmpl, "{{"+k+"}}", v, 1)
	}
	return tmpl
}
