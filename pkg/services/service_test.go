package services

import "testing"

func TestHash(t *testing.T) {
	testData := []struct {
		input    string
		expected string
	}{
		{
			input:    "test",
			expected: "7365637265749f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		},
		{
			input:    "Qqwerty1!",
			expected: "73656372657414406d3848369e58541696a81a0dbe945227272a4ed57858385b585b8206ed8d",
		},
	}

	for _, td := range testData {
		hash := Hash(td.input)
		if hash != td.expected {
			t.Errorf("Hash is not correct " + hash)
		}
	}
}
