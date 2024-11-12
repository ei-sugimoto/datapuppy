package internal_test

import (
	"testing"

	"github.com/ei-sugimoto/datapuppy/internal"
)

func TestSummarizeRequestCountByURI(t *testing.T) {
	t.Parallel()
	tc := []struct {
		name string
		internal.Logs
		expected map[string]int
	}{
		{
			"10 logs",
			internal.Logs{
				{URI: "/"},
				{URI: "/"},
				{URI: "/"},
				{URI: "/"},
				{URI: "/"},
				{URI: "/"},
				{URI: "/"},
				{URI: "/"},
				{URI: "/"},
				{URI: "/"},
			},
			map[string]int{
				"/": 10,
			},
		},
		{
			"8 logs",
			internal.Logs{
				{URI: "/"},
				{URI: "/"},
				{URI: "/"},
				{URI: "/"},
				{URI: "/serach"},
				{URI: "/serach"},
				{URI: "/serach"},
				{URI: "/serach"},
			},
			map[string]int{
				"/":       4,
				"/serach": 4,
			},
		},
		{
			"empty logs",
			internal.Logs{},
			map[string]int{},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			actual := internal.SummarizeRequestCountByURI(c.Logs)
			if len(actual) != len(c.expected) {
				t.Errorf("unexpected length. expected: %d, actual: %d", len(c.expected), len(actual))
			}
			for k, v := range c.expected {
				if actual[k] != v {
					t.Errorf("unexpected value. key: %s, expected: %d, actual: %d", k, v, actual[k])
				}
			}
		})
	}
}
