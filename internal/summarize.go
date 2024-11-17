package internal

type SummarizeURI struct {
	Count      int
	MaxLatency int
}

const INF = 1 << 31

func SummarizeByURI(log Log) map[string]SummarizeURI {
	uriCount := make(map[string]SummarizeURI)
	for _, d := range log.Details {
		if d.URI == "" {
			d.URI = "Unknown"
		}
		if _, ok := uriCount[d.URI]; !ok {
			uriCount[d.URI] = SummarizeURI{}
		}

		uriCount[d.URI] = SummarizeURI{
			Count:      uriCount[d.URI].Count + 1,
			MaxLatency: max(uriCount[d.URI].MaxLatency, d.Latency),
		}
	}
	return uriCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
