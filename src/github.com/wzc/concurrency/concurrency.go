package concurrency

//Websitechecker a func get string return bool
type Websitechecker func(string) bool

type result struct {
	string
	bool
}

//CheckWebsites check if website is available
func CheckWebsites(wc Websitechecker, urls []string) map[string]bool {
	//map[keyType]valueType
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
