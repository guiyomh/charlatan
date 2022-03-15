// Package loader provides tooling to load fixtures from path
package loader

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"github.com/guiyomh/charlatan/internal/db"
	"github.com/guiyomh/charlatan/internal/dto"
	"github.com/guiyomh/charlatan/internal/tree"
)

type Loader interface {
	Load(paths ...string) error
}

type fixtureLoader struct {
	logger    *zap.Logger
	files     []string
	persistor db.Persistor
}

func New(logger *zap.Logger, persistor db.Persistor) Loader {
	return fixtureLoader{
		logger:    logger,
		persistor: persistor,
		files:     make([]string, 0),
	}
}

func (l fixtureLoader) Load(paths ...string) error {
	var err error
	var fixtures dto.FixtureSet
	var tables dto.Tables
	for _, path := range paths {
		var files []string
		l.logger.With(zap.Any("path", path)).Debug("Analyzing path")
		files, err = locateFixtureFile(path)
		if err != nil {
			return err
		}
		l.files = append(l.files, files...)
	}

	if len(l.files) == 0 {
		return ErrNoFixtureFile
	}
	content, err := l.read()
	if err != nil {
		return err
	}

	fixtures, err = l.load(content)
	if err != nil {
		return err
	}

	tables, err = dto.ConvertFixtureSetToTable(fixtures)
	if err != nil {
		return err
	}
	tables = dto.FakeData(tables)
	tree := tree.ConvertTablesToTree(tables)

	err = l.persistor.Persist(tree)
	if err != nil {
		return err
	}

	return nil
}

func (l fixtureLoader) read() ([]byte, error) {
	contents := make([]byte, 0)
	for _, file := range l.files {
		l.logger.Sugar().Debugf("Read fixtures %s", file)
		content, err := ioutil.ReadFile(file)
		if err != nil {
			l.logger.With(zap.Any("file", file)).Error("Cannot read file")

			continue
		}
		contents = append(contents, content...)
	}

	if len(contents) == 0 {
		return []byte{}, ErrNoFixtureFile
	}

	return contents, nil
}

func (l fixtureLoader) load(content []byte) (dto.FixtureSet, error) {
	fixtures := dto.FixtureSet{}
	err := yaml.Unmarshal(content, fixtures)
	if err != nil {
		return dto.FixtureSet{}, errors.Wrap(err, "Unavaible to Unmarshal yaml fixtures")
	}

	return fixtures, nil
}
