package internal

// URIごとのリクエスト数を集計する
func SummarizeRequestCountByURI(logs Logs) map[string]int {
	uriCount := make(map[string]int)
	for _, log := range logs {
		uriCount[log.URI]++
	}
	return uriCount
}
