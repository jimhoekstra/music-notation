package parser

// Tokenize scans the input string and returns a sequence of tokens.
func Tokenize(input string) []Token {
	var tokens []Token

	for len(input) > 0 {
		matched := false

		// Check each pattern in order and apply the first one that matches
		for _, p := range Patterns {

			// Find the longest match
			match := p.Regex.FindString(input)
			if match != "" {
				tokens = append(
					tokens,
					Token{Type: p.Type, Value: match},
				)
				input = input[len(match):]
				matched = true

				// Break if a match was found
				break
			}
		}

		if !matched {
			// Skip unrecognised character. TODO: Handle this more
			// gracefully, by reporting an error with the position
			// of the unrecognised character
			input = input[1:]
		}
	}

	return tokens
}
