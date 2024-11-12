package internal

// URIごとのリクエスト数を集計する
func SummarizeRequestCountByURI(log Log) map[string]int {
	uriCount := make(map[string]int)
	for _, log := range log.Details {
		if log.URI == "" {
			log.URI = "unknown"
		}

		uriCount[log.URI]++
	}
	return uriCount
}
