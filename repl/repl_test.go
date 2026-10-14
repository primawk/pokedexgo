package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf(
				"Expected length %d, but got %d",
				len(c.expected),
				len(actual),
			)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf(
					"Expected word %q, but got %q",
					expectedWord,
					word,
				)
			}
		}
	}
}
