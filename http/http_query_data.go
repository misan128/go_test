// find_links_in_page.go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    start := time.Now()
    // Make HTTP request
    response, err := http.Get("https://weather.com/es-ES/tiempo/hoy/l/ae60f7c43022aa56e26e3498ec7ff0293103aa8e67c59886ae7c33b3ce0355dc")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Create a goquery document from the HTTP response
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }

    // Find all links and process them with the function
    // defined earlier
    /*document.Find("div.today_nowcard-temp").First(func(i int, s *goquery.Selection) {
        // For each item found, get the band and title
        //band := s.Find("a").Text()
        //itle := s.Find("i").Text()
        fmt.Printf("%d - %s\n",i, s.Text())
    })*/

    sel := document.Find("div.today_nowcard-temp").First()
    fmt.Printf("Temperatura Donostia: %s\n", sel.Text())

    elapsed := time.Since(start)
    log.Printf("Scrapping took %s", elapsed)
}

//<div class="today_nowcard-temp"><span class="">9<sup>°</sup></span></div>