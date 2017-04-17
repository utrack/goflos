package db

import "github.com/utrack/goflos/domain/static/market"

type GoodDb struct {
	byArchID map[uint64]market.Good
}

func NewUniverseDB(flPath string) error {
	//ini, err := flini.ParseFile(path.Join(flPath, "EXE/freelancer.ini"))
	return nil
}
