package concurrency

// WebsiteChecker is a type based on a function that takes a string and returns a bool.
type WebsiteChecker func(string) bool

// CheckWebsites function takes a Website Checker and a slice of url as arguments.
// It returns a map of urls and whether they can be reached or not.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
