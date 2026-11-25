package menu

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"

	"github.com/kyawzinthant-coding/froky/internal/types"
)

const (
	maxNameLength        = 100
	maxDescriptionLength = 500
)

func normalizeName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" || utf8.RuneCountInString(name) > maxNameLength {
		return "", fmt.Errorf("%w: name must be 1-%d characters", ErrInvalidInput, maxNameLength)
	}
	return name, nil
}

func normalizeProductInput(in types.ProductInput) (types.ProductInput, error) {
	name, err := normalizeName(in.Name)
	if err != nil {
		return in, err
	}
	in.Name = name

	in.Description = strings.TrimSpace(in.Description)
	if utf8.RuneCountInString(in.Description) > maxDescriptionLength {
		return in, fmt.Errorf("%w: description must be at most %d characters", ErrInvalidInput, maxDescriptionLength)
	}

	// A missing price decodes as 0, so 0 is rejected too.
	if in.Price <= 0 {
		return in, fmt.Errorf("%w: price must be greater than 0", ErrInvalidInput)
	}

	if in.CategoryID <= 0 {
		return in, fmt.Errorf("%w: categoryId is required", ErrInvalidInput)
	}

	if in.ImageURL != nil {
		trimmed := strings.TrimSpace(*in.ImageURL)
		if trimmed == "" {
			in.ImageURL = nil
			return in, nil
		}
		u, err := url.Parse(trimmed)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
			return in, fmt.Errorf("%w: imageUrl must be an http or https URL", ErrInvalidInput)
		}
		in.ImageURL = &trimmed
	}

	return in, nil
}
