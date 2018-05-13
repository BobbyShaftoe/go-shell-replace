package types

import (
		_ "log"
		"os"
)

type EnvConfig struct {
		Project    string
		GitCommit  string
		HttpProxy  string
		HttpsProxy string
}

func (env *EnvConfig) SetEnv() {

		env.Project = os.Getenv("PROJECT")
		env.GitCommit = os.Getenv("COMMIT_ID")
		env.HttpProxy = os.Getenv("http_proxy")
		env.HttpsProxy = os.Getenv("https_proxy")

}

