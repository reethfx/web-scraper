package main

import (
    "fmt"
    "web-scrapper/config"
    "web-scrapper/crawler"
)

func main() {
    cfg := config.LoadConfig()
    fmt.Println("Iniciando el scraper...")

    err := crawler.Start(cfg)
    if err != nil {
        fmt.Println("Error ejecutando el scraper:", err)
    } else {
        fmt.Println("Scraper finalizado con éxito.")
    }
}
