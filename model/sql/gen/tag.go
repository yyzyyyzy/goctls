package gen

import (
	"github.com/yyzyyyzy/goctls/model/sql/template"
	"github.com/yyzyyyzy/goctls/util"
	"github.com/yyzyyyzy/goctls/util/pathx"
)

func genTag(table Table, in string) (string, error) {
	if in == "" {
		return in, nil
	}

	text, err := pathx.LoadTemplate(category, tagTemplateFile, template.Tag)
	if err != nil {
		return "", err
	}

	output, err := util.With("tag").Parse(text).Execute(map[string]any{
		"field": in,
		"data":  table,
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
