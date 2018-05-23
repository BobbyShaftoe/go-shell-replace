package main

import (
	"flag"
	_ "io"
	"log"
	"os"
	"time"

	"github.com/BobbyShaftoe/go-shell-replace/config"
	"github.com/BobbyShaftoe/go-shell-replace/types"
	_ "github.com/BobbyShaftoe/go-shell-replace/types"
	"github.com/BobbyShaftoe/go-shell-replace/util"
)

// curl -u artifactory:d41d8cd98f00b204e9800998ecf8427e -X PUT "http://10.30.0.51/artifactory/cloud/env-confs/env-confs-1526139982.zip" -T env-confs-1526139982.zip
const configFile = "../artifactory/config.yml"

var project, gitCommit, httpProxy, httpsProxy, archiveFilesDir, archiveDstDir string

func main() {

	var yamlFile string
	flag.StringVar(&yamlFile, "config", configFile, "a string")
	flag.Parse()

	if _, err := os.Stat(yamlFile); err != nil {
		yamlFile = ""
	}
	conf := &config.Configuration{}
	project, gitCommit, httpProxy, httpsProxy, archiveFilesDir, archiveDstDir = conf.DoConfig(yamlFile, true)

	versionFile, archiveName := makeFileNames(project)

	// Create version file
	file, err := os.Create(versionFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Write archive name to version file
	content := []byte(archiveName + "\n")
	_, err = file.Write(content)
	if err != nil {
		log.Printf("Error while writing to file: %v", err)
	}

	// Copy files for archiving into destination directory
	src, dst := "test", archiveFilesDir+"/"+project
	//src, dst := "test", "tmp/"+project
	c := &util.CopyDirArgs{Src: src, Dst: dst, Mode: 0755, IgnoreDot: true}
	if err = c.CopyDir(); err != nil {
		log.Printf("Error copying directory %v\n", err)
	}

	// Create project files archive
	src = archiveFilesDir + "/" + project
	t := &util.TarGzArgs{Name: archiveName, Src: src, DstDir: archiveDstDir}
	if err = t.CreateArchive(); err != nil {
		log.Printf("Error creating archive %v\n", err)
	}

}

// Function to create file names
func makeFileNames(p string) (string, string) {
	timeNow := time.Now()
	timeStamp := timeNow.Format(types.TimestampFormat)
	archiveName := p + "-" + timeStamp + ".tar.gz"
	versionFile := p + "-version.txt"
	return versionFile, archiveName
}
