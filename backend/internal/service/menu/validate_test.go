package menu

import (
	"strings"
	"testing"

	"github.com/kyawzinthant-coding/froky/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func strPtr(s string) *string { return &s }

func validProduct() types.ProductInput {
	return types.ProductInput{Name: "Pad Thai", Description: "Noodles", Price: 6500, CategoryID: 1}
}

func TestNormalizeProductInput_Valid(t *testing.T) {
	t.Parallel()

	in := validProduct()
	in.Name = "  Pad Thai  "
	in.ImageURL = strPtr(" https://cdn.example.com/pad-thai.jpg ")

	got, err := normalizeProductInput(in)
	require.NoError(t, err)
	assert.Equal(t, "Pad Thai", got.Name)
	assert.Equal(t, "https://cdn.example.com/pad-thai.jpg", *got.ImageURL)
}

func TestNormalizeProductInput_EmptyImageURLBecomesNil(t *testing.T) {
	t.Parallel()

	in := validProduct()
	in.ImageURL = strPtr("   ")

	got, err := normalizeProductInput(in)
	require.NoError(t, err)
	assert.Nil(t, got.ImageURL)
}

func TestNormalizeProductInput_Invalid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		modify func(*types.ProductInput)
	}{
		{name: "blank name", modify: func(in *types.ProductInput) { in.Name = "   " }},
		{name: "name too long", modify: func(in *types.ProductInput) { in.Name = strings.Repeat("a", maxNameLength+1) }},
		{name: "description too long", modify: func(in *types.ProductInput) { in.Description = strings.Repeat("a", maxDescriptionLength+1) }},
		{name: "negative price", modify: func(in *types.ProductInput) { in.Price = -1 }},
		{name: "missing price", modify: func(in *types.ProductInput) { in.Price = 0 }},
		{name: "missing category", modify: func(in *types.ProductInput) { in.CategoryID = 0 }},
		{name: "javascript url", modify: func(in *types.ProductInput) { in.ImageURL = strPtr("javascript:alert(1)") }},
		{name: "relative url", modify: func(in *types.ProductInput) { in.ImageURL = strPtr("/images/a.jpg") }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			in := validProduct()
			tc.modify(&in)

			_, err := normalizeProductInput(in)
			assert.ErrorIs(t, err, ErrInvalidInput)
		})
	}
}

func TestNormalizeName_CountsCharactersNotBytes(t *testing.T) {
	t.Parallel()

	// Thai characters are 3 bytes each in UTF-8; 100 of them must still be accepted.
	_, err := normalizeName(strings.Repeat("ก", maxNameLength))
	assert.NoError(t, err)
}
