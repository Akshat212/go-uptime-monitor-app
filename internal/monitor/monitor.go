package monitor

import (
	"log"
	"net/http"
	"time"
)

type MonitoredURL struct {
	URL string
	Up bool
	LastChecked time.Time
}

var monitored = make(map[string]*MonitoredURL)

// Add URL for monitoring 
func AddURL(url string) {
	if _, exists := monitored[url]; !exists {
		monitored[url] = &MonitoredURL{URL: url}
		go startMonitoring(url)
		log.Printf("Started monitoring %s\n", url)
	}
}

// startMonitoring goroutine to start monitoring the url
func startMonitoring(url string) {
	for {
		resp, err := http.Get(url)
		status := err == nil && resp.StatusCode == 200

		monitored[url].Up = status
		monitored[url].LastChecked = time.Now()

		log.Printf("[%s] Status: %v\n", url, status)

		if resp != nil {
			resp.Body.Close()
		}
		time.Sleep(30 * time.Second)
	}
}

// GetAll returns all monitoring URLs
func GetAll() []*MonitoredURL {
	var urls []*MonitoredURL
	for _, u := range monitored {
		urls = append(urls, u)
	}
	return urls
}


