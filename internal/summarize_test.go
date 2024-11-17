package internal_test

import (
	"testing"

	"github.com/ei-sugimoto/datapuppy/internal"
	"github.com/ei-sugimoto/datapuppy/internal/helper"
)

func TestSummarizeByURI(t *testing.T) {

	tc := []struct {
		name string
		log  internal.Log
		want map[string]internal.SummarizeURI
	}{
		{
			name: "normal",
			log:  helper.DummyLog,
			want: map[string]internal.SummarizeURI{
				"/": {
					Count:      2,
					MaxLatency: 200,
				},
				"/search": {
					Count:      1,
					MaxLatency: 300,
				},
			},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := internal.SummarizeByURI(tt.log)
			if len(got) != len(tt.want) {
				t.Errorf("unexpected length: got %d, want %d", len(got), len(tt.want))
			}
			for k, v := range tt.want {
				if got[k].Count != v.Count {
					t.Errorf("unexpected count: got %d, want %d", got[k].Count, v.Count)
				}
				if got[k].MaxLatency != v.MaxLatency {
					t.Errorf("unexpected max latency: got %d, want %d", got[k].MaxLatency, v.MaxLatency)
				}
			}
		})
	}

}
