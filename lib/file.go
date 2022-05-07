package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func CreateFile(dirname string, basename string, output OutPut) {
	_, err := os.Create("results/" + dirname + basename)
	if err != nil {
		log.Fatal(err)
	}
	file, _ := json.MarshalIndent(output, "", " ")
	_ = ioutil.WriteFile("results/"+dirname+basename, file, 0644)

}

func CreateDirectory(fps []string) {
	for _, fp := range fps {
		dirname, _ := filepath.Split(fp)
		if err := os.MkdirAll("results/"+dirname, 0777); err != nil {
			log.Fatal(err, "failed to create submit Directory")
		}
	}
}
