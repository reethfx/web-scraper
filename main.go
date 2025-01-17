package main

import (
    "fmt"
    "web-scraper-security/config"
    "web-scraper-security/crawler"
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
