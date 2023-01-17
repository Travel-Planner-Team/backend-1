package util

import (
    "io/ioutil"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

type TokenInfo struct {
    Secret string `yaml:"secret"`
}

type ApplicationConfig struct {
    TokenConfig         *TokenInfo         `yaml:"token"`
}

func LoadApplicationConfig(configDir, configFile string) (*ApplicationConfig, error) {
    content, err := ioutil.ReadFile(filepath.Join(configDir, configFile))
    if err != nil {
        return nil, err
    }

    var config ApplicationConfig
    err = yaml.Unmarshal(content, &config)
    if err != nil {
        return nil, err
    }
    return &config, nil
}