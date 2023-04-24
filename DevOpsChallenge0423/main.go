package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Hosts    []string `yaml:"hosts"`
	Interval int      `yaml:"interval"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong\n")
}

func pingHosts(hosts []string, interval time.Duration) {
	for {
		for _, host := range hosts {
			resp, err := http.Get(fmt.Sprintf("%s/ping", host))
			if err != nil {
				log.Printf("Host: %s - Error: %v", host, err)
				continue
			}
			body, _ := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			log.Printf("Host: %s - Response: %s", host, body)
		}
		time.Sleep(interval)
	}
}

func main() {
	// Load configuration
	configBytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		log.Fatalf("Failed to parse config.yaml: %v", err)
	}

	// Start background goroutine to ping other hosts periodically
	go pingHosts(config.Hosts, time.Duration(config.Interval)*time.Second)

	http.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
