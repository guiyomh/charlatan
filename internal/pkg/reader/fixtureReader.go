package reader

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/azer/logger"
	"github.com/guiyomh/charlatan/internal/app/model"
	yaml "gopkg.in/yaml.v2"
)

var log = logger.New("reader")

type FixtureReader struct {
}

//NewFixtureReader factory to create an instance of FixtureReader
func NewFixtureReader() *FixtureReader {
	return &FixtureReader{}
}

func (f FixtureReader) Read(fixturePath string) (model.FixtureTables, error) {
	files, err := ioutil.ReadDir(fixturePath)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	content := make([]byte, 0)
	for _, fileinfo := range files {
		fileContent, err := f.readFile(fileinfo, fixturePath)
		if err != nil {
			log.Error(err.Error())
			continue
		}
		content = append(content, fileContent...)
	}
	return f.parse(content)
}

func (f FixtureReader) readFile(file os.FileInfo, path string) ([]byte, error) {
	fullPath, _ := filepath.Abs(path + "/" + file.Name())
	log.Info(fmt.Sprintf("Load file : %s", fullPath))
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Error(fmt.Sprintf("File %s cannot be read", fullPath))
		return nil, err
	}
	return content, nil
}

func (f FixtureReader) parse(content []byte) (model.FixtureTables, error) {
	tbls := model.FixtureTables{}
	err := yaml.Unmarshal(content, tbls)
	return tbls, err
}
