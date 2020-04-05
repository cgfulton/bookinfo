package istio

import (
	"github.com/c-bata/go-prompt"
)

func NewCompleter() (*Completer, error) {
	return &Completer{}, nil
}

type Completer struct {
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "bookstore1", Description: "Namespace Bookstore 1"},
		{Text: "bookstore2", Description: "Namespace Bookstore 2"},
		{Text: "bookstore3", Description: "Namespace Bookstore 3"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
