package content

import (
	"fmt"
	"net/http"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"

	"github.com/nilslice/appfront"
)

type About struct {
	item.Item

	Title   string `json:"title"`
	Content string `json:"content"`
}

// MarshalEditor writes a buffer of html to edit a About within the CMS
// and implements editor.Editable
func (a *About) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(a,
		// Take note that the first argument to these Input-like functions
		// is the string version of each About field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", a, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Content", a, map[string]string{
				"label":       "Content",
				"placeholder": "Enter the Content here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render About editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["About"] = func() interface{} { return new(About) }

	fmt.Println(fmt.Sprintf("%#v", appfront.Router))

	http.Handle("/", appfront.Router())
}
