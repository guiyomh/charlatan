package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/guiyomh/go-faker-fixtures/internal/pkg/generator"

	"github.com/azer/logger"
	"github.com/guiyomh/go-faker-fixtures/internal/app/model"
	"github.com/guiyomh/go-faker-fixtures/internal/pkg/parser"
)

// Loader Load file an create Rows
type Loader struct {
	fixturePath string
	logger      *logger.Logger
	parser      *parser.YamlFixture
	Rows        []*model.Row
}

// NewLoader is a factory to create a Loader instance
func NewLoader(fixturePath string) *Loader {
	return &Loader{
		fixturePath: fixturePath,
		logger:      logger.New("file-loader"),
		parser:      &parser.YamlFixture{},
		Rows:        make([]*model.Row, 0),
	}
}

// Load Load file locate in the fixtures path
func (l *Loader) Load() {
	files := l.locateFixturesFiles()
	l.logger.Info("Fixtures found", files)
	l.loadFixtures(files)
}

func (l *Loader) locateFixturesFiles() []os.FileInfo {
	files, err := ioutil.ReadDir(l.fixturePath)
	if err != nil {
		l.logger.Error(err.Error())
		os.Exit(1)
	}
	return files
}

func (l *Loader) loadFixtures(files []os.FileInfo) {

	wg := &sync.WaitGroup{}
	for _, fileinfo := range files {
		wg.Add(1)
		go l.load(fileinfo, wg)
	}
	wg.Wait()
}

func (l *Loader) load(file os.FileInfo, wg *sync.WaitGroup) {
	fullPath, _ := filepath.Abs(l.fixturePath + "/" + file.Name())
	l.logger.Info(fmt.Sprintf("Load file : %s", fullPath))
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		l.logger.Error(fmt.Sprintf("File %s cannot be read", fullPath))
		panic(err)
	}

	templates, objects, _ := l.parser.Load(content)

	rowGenerator := generator.NewRecord(templates, objects)
	rows := rowGenerator.BuildRecord()
	l.Rows = append(l.Rows, rows...)
	wg.Done()
}
