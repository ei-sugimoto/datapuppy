package helper

import "github.com/ei-sugimoto/datapuppy/internal"

var DummyLog internal.Log = internal.Log{
	Details: []internal.Detail{
		{
			ID:        "1",
			Time:      "2024-01-01T00:00:00Z",
			RemoteIP:  "127.0.0.1",
			Host:      "example.com",
			Method:    "GET",
			URI:       "/",
			UserAgent: "Mozilla/5.0",
			Status:    200,
			Latency:   100,
		},
		{
			ID:        "2",
			Time:      "2024-01-01T00:00:01Z",
			RemoteIP:  "127.0.0.1",
			Host:      "example.com",
			Method:    "GET",
			URI:       "/",
			UserAgent: "Mozilla/5.0",
			Status:    200,
			Latency:   200,
		},
		{
			ID:        "3",
			Time:      "2024-01-01T00:00:02Z",
			RemoteIP:  "125",
			Host:      "example.com",
			Method:    "GET",
			URI:       "/search",
			UserAgent: "Mozilla/5.0",
			Status:    200,
			Latency:   300,
		},
	},
}
