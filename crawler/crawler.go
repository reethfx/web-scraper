package crawler

import (
	"fmt"
	"net/http"
	"time"
	"web-scrapper/config"
	"web-scrapper/storage"
	"web-scrapper/utils"
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

	data, err := ParseToJSON(resp.Body)
	if err != nil {
		return err
	}

	err = storage.SaveToJSON(cfg.OutputFile, data)
	if err != nil {
		return err
	}

	fmt.Printf("Resultados guardados en %s\n", cfg.OutputFile)
	return nil
}
