package crawler

import (
    "fmt"
    "net/http"
    "time"
    "github.com/reethfx/web-scrapper/config"
    "github.com/reethfx/web-scrapper/utils"
)

func Start(cfg config.Config) error {
    fmt.Println("Comenzando a rastrear:", cfg.StartURL)

    client := &http.Client{Timeout: time.Duration(cfg.Timeout) * time.Second}
    visited := make(map[string]bool)

    return crawl(cfg.StartURL, 0, cfg, client, visited)
}

func crawl(url string, depth int, cfg config.Config, client *http.Client, visited map[string]bool) error {
    if depth > cfg.MaxDepth || visited[url] {
        return nil
    }

    visited[url] = true
    fmt.Println("Visitando:", url)

    resp, err := utils.FetchURL(client, url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Aquí puedes extraer datos sensibles
    data := parse(resp.Body)
    if len(data) > 0 {
        fmt.Println("Datos sensibles encontrados:", data)
        utils.SaveToFile(cfg.OutputFile, data)
    }

    // Continuar con los enlaces encontrados (pendiente implementar lógica para extraer enlaces)
    return nil
}
