package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	_ "log"
	"os"
)

// TODO: Configure any missing variable to pass through to env var
type YamlConfig struct {
	Project         string `yaml:"project"`
	GitCommit       string `yaml:"git_commit"`
	HttpProxy       string `yaml:"http_proxy"`
	HttpsProxy      string `yaml:"https_proxy"`
	ArchiveFilesDir string `yaml:"archive_files_dir"`
	ArchiveDstDir   string `yaml:"archive_dst_dir"`
}

type EnvConfig struct {
	Project         string
	GitCommit       string
	HttpProxy       string
	HttpsProxy      string
	ArchiveFilesDir string
	ArchiveDstDir   string
}

var Y YamlConfig
var E EnvConfig

type Configuration struct {
	y YamlConfig
	e EnvConfig
}

var I interface{}

func (C *Configuration) DoConfig(c string, p bool) (string, string, string, string, string, string) {
	if c != "" {
		fmt.Printf("Reading configuration from yaml file: %v\n", c)
		Y.ParseYaml(c)
		I = Y
	} else {
		fmt.Println("Couldn't read from yaml file, reading configuration from environment")
		E.SetEnv()
		I = E
	}

	if p {
		fmt.Printf("\n%v\n\n", I)
	}
	// TODO: Return values as a Map
	switch I.(type) {
	case YamlConfig:
		result := I.(YamlConfig)
		return result.Project, result.GitCommit, result.HttpProxy, result.HttpsProxy, result.ArchiveFilesDir, result.ArchiveDstDir
	case EnvConfig:
		result := I.(EnvConfig)
		return result.Project, result.GitCommit, result.HttpProxy, result.HttpsProxy, result.ArchiveFilesDir, result.ArchiveDstDir
	}
	return "", "", "", "", "", ""
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

func (env *EnvConfig) SetEnv() {
	env.Project = os.Getenv("PROJECT")
	env.GitCommit = os.Getenv("GIT_COMMIT")
	env.HttpProxy = os.Getenv("http_proxy")
	env.HttpsProxy = os.Getenv("https_proxy")
	env.ArchiveFilesDir = os.Getenv("archive_files_dir")
	env.ArchiveDstDir = os.Getenv("archive_dst_dir")
}
