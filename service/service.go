package service

import (
	"fmt"

	"github.com/danmrichards/xkpassgo/pkg/config"
	"github.com/danmrichards/xkpassgo/pkg/generator"
)

type s struct {
	// a database dependency would go here but instead we're going to have a static map
	m map[string]string
}

// New instantiates a new service.
func New( /* a database connection would be injected here */ ) *s {
	return &s{
		m: map[string]string{
			"abc": "https://github.com/cameronbrill/create-go-app",
		},
	}
}

func (s *s) NewLink(orig string) (Link, error) {
	l, err := generator.NewXKPassword(&config.GeneratorConfig{
		NumWords:           3,
		WordLenMin:         3,
		WordLenMax:         6,
		SeparatorCharacter: "-",
		CaseTransform:      "LOWER",
		PaddingType:        "FIXED",
	}).Generate()
	link := Link{
		Original: orig,
		Short:    l,
	}
	if err != nil {
		return link, fmt.Errorf("generating link: %w", err)
	}
	return link, nil
}

func (s *s) ShortURLToLink(short string) (Link, error) {
	return Link{}, nil
}