package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	// Parse all arguments
	flag.Parse()

	// flag.Args contains all non-flag arguments
	sites := flag.Args()

	// Args might be 	a single string with spaces to delimit sites, break them apart
	if len(sites) == 1 {
		sites = strings.Split(sites[0], " ")
	}

	// use a  waitgroup to make sure all go routines finish
	var wg sync.WaitGroup

	// Lets keep a reference to when we started
	start := time.Now()

	// Set the value for the wait group
	wg.Add(len(sites))
	processSites(sites, &wg)

	// Block until all routines finish
	wg.Wait()

	fmt.Printf("Entire process took %s\n", time.Since(start))
}

func processSites(sites []string, wg *sync.WaitGroup) {
	for i, site := range sites {
		// Launch each retrieval in a go routine. This makes each request concurrent
		go func(index int, site string) {
			defer wg.Done()

			// start a timer for this request
			begin := time.Now()

			// Retrieve the site
			resp, err := http.Get(site)
			if err != nil {
				fmt.Println(site, err)
				return
			}

			fmt.Printf("%d Site %q took %s to retrieve with status code of %d.\n", i, site, time.Since(begin), resp.StatusCode)
		}(i, site)
	}
}
