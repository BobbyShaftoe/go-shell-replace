package config

import (
		_ "log"
		"io/ioutil"
		"gopkg.in/yaml.v2"
		"log"
)

type YamlConfig struct {
		Project    string `yaml:"project"`
		GitCommit  string `yaml:"git_commit"`
		HttpProxy  string `yaml:"http_proxy"`
		HttpsProxy string `yaml:"https_proxy"`
}

func (y *YamlConfig) ParseYaml(c string) {
		configFile := c
		yamlData, err := ioutil.ReadFile(configFile)
		if err != nil {
				log.Fatalf("cannot open config file: %v", err)
		}

		err = yaml.Unmarshal(yamlData, y)

		if err != nil {
				log.Fatalf("cannot unmarshal data: %v", err)
		}

}
