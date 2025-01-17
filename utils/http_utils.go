package utils

import (
    "net/http"
    "os"
)

func FetchURL(client *http.Client, url string) (*http.Response, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("User-Agent", "GoSecurityScraper/1.0")

    return client.Do(req)
}

func SaveToFile(filename string, data interface{}) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    // Implementar serialización del data en formato JSON
    // Para simplificar, esto queda pendiente
    return nil
}
