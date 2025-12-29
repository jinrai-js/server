package seo_table

import (
	"fmt"
	"strings"
)

type Header struct {
	Coll     string `json:"coll"`
	Title    string `json:"title"`
	Subtitle any    `json:"subtitle"`
}

type Row = map[string]any

type TableProps struct {
	Headers []Header `json:"headers"`
	Rows    []Row    `json:"rows"`
	Else    int      `json:"else"`
}

func Component(props TableProps) string {

	return `<table border="1" style="border-collapse: collapse; width: 100%; text-align: left;" class="ssrtable border">
  <thead>
    <tr style="background-color: #f4f4f4;">
      ` + renderHead(props.Headers) + `
    </tr>
  </thead>
  <tbody>
    ` + renderBody(props.Headers, props.Rows) + `
  </tbody>
</table>`
}

func renderHead(head []Header) string {
	var result strings.Builder

	for _, value := range head {
		result.WriteString(fmt.Sprintf(`<th style="padding: 6px;">%s</th>`, value.Title))
	}

	return result.String()
}

func renderBody(head []Header, rows []Row) string {
	var result strings.Builder

	for _, row := range rows {
		result.WriteString("<tr>")

		for _, coll := range head {
			value, _ := row[coll.Coll]
			result.WriteString(fmt.Sprintf(`<td>%s</td>`, toString(value)))
		}

		result.WriteString("</tr>")

	}

	return result.String()
}

type PagetitleColl struct {
	Pretty string `json:"pretty"`
	Code   string `json:"code"`
}

func toString(row any) string {
	switch v := row.(type) {
	case string:
		return v

	case []any:
		return strings.Join(toStringArray(v), ", ")

	case map[string]any:
		code, _ := v["code"].(string)
		if pretty, ok := v["pretty"].(string); ok {
			if code == "" {
				code = pretty
			}
			return fmt.Sprintf(`<a class="text-blue-500" href="/%s">%s</a>`, code, pretty)
		}
		return fmt.Sprintf("%v", v)

	default:
		return fmt.Sprintf("unknown type (%v)", v)
	}
}

func toStringArray(from []any) []string {
	var result []string

	for _, val := range from {
		result = append(result, toString(val))
	}

	return result
}
