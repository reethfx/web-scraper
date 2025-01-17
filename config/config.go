package config

type Config struct {
    StartURL    string
    MaxDepth    int
    UserAgent   string
    Timeout     int
    OutputFile  string
}

func LoadConfig() Config {
    return Config{
        StartURL:   "https://autenticador.pil.es",
        MaxDepth:   3,
        UserAgent:  "GoSecurityScraper/1.0",
        Timeout:    10,
        OutputFile: "data/results.json",
    }
}
