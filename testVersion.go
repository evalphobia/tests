package tests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-xorm/xorm"
)

type VersionS struct {
	Id   int64
	Name string
	Ver  int `xorm:"version"`
}

func testVersion(engine *xorm.Engine, t *testing.T) {
	err := engine.DropTables(new(VersionS))
	if err != nil {
		t.Error(err)
		panic(err)
	}

	err = engine.CreateTables(new(VersionS))
	if err != nil {
		t.Error(err)
		panic(err)
	}

	ver := &VersionS{Name: "sfsfdsfds"}
	_, err = engine.Insert(ver)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	fmt.Println(ver)
	if ver.Ver != 1 {
		err = errors.New("insert error")
		t.Error(err)
		panic(err)
	}

	newVer := new(VersionS)
	has, err := engine.Id(ver.Id).Get(newVer)
	if err != nil {
		t.Error(err)
		panic(err)
	}

	if !has {
		t.Error(errors.New(fmt.Sprintf("no version id is %v", ver.Id)))
		panic(err)
	}
	fmt.Println(newVer)
	if newVer.Ver != 1 {
		err = errors.New("insert error")
		t.Error(err)
		panic(err)
	}

	newVer.Name = "-------"
	_, err = engine.Id(ver.Id).Update(newVer)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if newVer.Ver != 2 {
		err = errors.New("update should set version back to struct")
		t.Error(err)
	}

	newVer = new(VersionS)
	has, err = engine.Id(ver.Id).Get(newVer)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	fmt.Println(newVer)
	if newVer.Ver != 2 {
		err = errors.New("insert error")
		t.Error(err)
		panic(err)
	}
}
