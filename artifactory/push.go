package main

import (
	"flag"
	"fmt"
	_ "io"
	_ "os"

	"github.com/BobbyShaftoe/go-shell-replace/config"
	"github.com/BobbyShaftoe/go-shell-replace/types"
	"github.com/BobbyShaftoe/go-shell-replace/util"
	_ "github.com/BobbyShaftoe/go-shell-replace/util"
	"log"
	"os"
	"time"
	_ "time"
)

// curl -u artifactory:d41d8cd98f00b204e9800998ecf8427e -X PUT "http://10.30.0.51/artifactory/cloud/env-confs/env-confs-1526139982.zip" -T env-confs-1526139982.zip
const configFile = "../artifactory/config.yml"

var project, gitCommit, httpProxy, httpsProxy string

func main() {

	var yamlFile string
	flag.StringVar(&yamlFile, "config", configFile, "a string")
	flag.Parse()

	fmt.Println("Configuration file:", yamlFile)

	if _, err := os.Stat(yamlFile); err == nil {
		y := &config.YamlConfig{}
		y.ParseYaml(yamlFile)
		project, gitCommit, httpProxy, httpsProxy = y.Project, y.GitCommit, y.HttpProxy, y.HttpsProxy
	} else {
		c := &types.EnvConfig{}
		c.SetEnv()
		project, gitCommit, httpProxy, httpsProxy = c.Project, c.GitCommit, c.HttpProxy, c.HttpsProxy
	}

	timeNow := time.Now()
	timeStamp := timeNow.Format("20060102_150405")
	archiveName := project + "-" + timeStamp + ".tar.gz"

	file, err := os.Create(project + "-version.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content := []byte(archiveName + "\n")
	_, err = file.Write(content)
	if err != nil {
		log.Printf("Error while writing to file: %v", err)
	}

	src := "test"
	dst := "tmp/" + project
	c := &util.CopyDirArgs{nil, src, dst, 0755, true}

	if err = c.CopyDir(); err != nil {
		log.Printf("Error copying directory %v\n", err)
	}

	t := &util.TarGzArgs{dst, "."}
	if err = t.CreateArchive(); err != nil {
		log.Printf("Error creating archive %v\n", err)
	}

	//file, err = os.Create(archiveName)
	//if err != nil {
	//		log.Fatal(err)
	//}

}
