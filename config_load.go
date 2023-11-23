package hw_server_config

import (
    "gopkg.in/yaml.v3"
    "io"
    "os"
    "strconv"
    "time"
)

type innerConfig struct {
    Host    string `yaml:"host"`
    Port    int    `yaml:"port"`
    Timeout int    `yaml:"timeout"`
}

type ServerConfig struct {
    Addr    string
    Timeout time.Duration
}

// LoadYamlConfig makes server config from .yaml file. Timeout must be specified in milliseconds
func LoadYamlConfig(path string) (*ServerConfig, error) {
    var ic innerConfig
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    b, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }
    err = yaml.Unmarshal(b, &ic)
    if err != nil {
        return nil, err
    }
    return &ServerConfig{
        Addr:    ic.Host + ":" + strconv.Itoa(ic.Port),
        Timeout: time.Millisecond * time.Duration(ic.Timeout),
    }, nil
}
