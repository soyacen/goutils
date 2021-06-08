package sliceutils

import "testing"

func TestCloneString(t *testing.T) {
	var tests = []struct {
		slice    []string
		expected []string
	}{
		{
			slice:    nil,
			expected: nil,
		},
		{
			slice:    []string{},
			expected: []string{},
		},
		{
			slice:    []string{"1"},
			expected: []string{"1"},
		},
		{
			slice:    []string{"1", "2"},
			expected: []string{"1", "2"},
		},
	}

	for _, test := range tests {
		actual := CloneString(test.slice)
		if len(actual) != len(test.expected) {
			t.Errorf("expected is %v, but actual is %v", test.expected, actual)
			continue
		}
		for i := range actual {
			if actual[i] != test.expected[i] {
				t.Errorf("expected is %v, but actual is %v", test.expected, actual)
				break
			}
		}
	}

}
