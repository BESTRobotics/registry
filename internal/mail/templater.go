package mail

import (
	"bytes"
	"log"
	"path/filepath"
	"text/template"

	"github.com/spf13/viper"
)

// RenderLetter loads a file containing the text of letters and
// templates it into the provided letter template.  In the event not
// letter is found for that name, the system will return
// ErrNoSuchLetter.
func RenderLetter(name string, tmplVals *LetterContext) (Letter, error) {
	t := templates.Lookup(name)
	if t == nil {
		return Letter{}, ErrNoSuchLetter
	}

	// Render the template into the letter
	l := NewLetter()
	var tmp bytes.Buffer
	err := t.Execute(&tmp, tmplVals)
	if err != nil {
		return Letter{}, ErrNoSuchLetter
	}
	l.Body = tmp.String()

	return l, nil
}

// loadTemplates attempts to load all templates from the viper
// specified location.
func loadTemplates() error {
	path := filepath.Join(viper.GetString("storage.root"), "tmpl", "*")
	t, err := template.ParseGlob(path)
	if err != nil {
		log.Println("Error loading templates:", err)
		return err
	}

	log.Println("The following templates are available for mail:")
	for _, tmpl := range t.Templates() {
		log.Printf("  %s\n", tmpl.Name())
	}

	templates = t
	return nil
}
