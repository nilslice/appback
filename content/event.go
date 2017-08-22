package content

import (
	"fmt"
	"net/http"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Event struct {
	item.Item

	Title      string   `json:"title"`
	Details    []string `json:"details"`
	TicketLink string   `json:"ticket_link"`
}

// MarshalEditor writes a buffer of html to edit a Event within the CMS
// and implements editor.Editable
func (e *Event) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(e,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Event field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", e, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.InputRepeater("Details", e, map[string]string{
				"label":       "Details",
				"type":        "text",
				"placeholder": "Enter the Details here",
			}),
		},
		editor.Field{
			View: editor.Input("TicketLink", e, map[string]string{
				"label":       "TicketLink",
				"type":        "text",
				"placeholder": "Enter the TicketLink here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Event editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Event"] = func() interface{} { return new(Event) }
}

func (e *Event) Create(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (e *Event) AutoApprove(res http.ResponseWriter, req *http.Request) error {
	return nil
}
