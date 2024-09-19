package main

import (
	"JSONtoXML/src/ex00"
	"JSONtoXML/src/ex01"
	"JSONtoXML/src/ex02"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type userMode struct {
	readDB    bool
	compareDB bool
	compareFS bool
	f         string
	old       string
	new       string
}

func (u *userMode) parseExecutable() {
	executable, err := os.Executable()
	if err != nil {
		log.Fatal("Error getting executable path: ", err)
	}

	if strings.HasSuffix(executable, "readDB") {
		u.readDB = true
	} else if strings.HasSuffix(executable, "compareDB") {
		u.compareDB = true
	} else if strings.HasSuffix(executable, "compareFS") {
		u.compareFS = true
	}
}

func (u *userMode) initFlags() {
	if u.readDB {
		flag.StringVar(
			&u.f,
			"f",
			"",
			"specify path to database in .json or .xml file to read",
		)
	} else if u.compareDB || u.compareFS {
		flag.StringVar(&u.old, "old", "", "specify path to old database/dump")
		flag.StringVar(&u.new, "new", "", "specify path to new database/dump")
	}

	flag.Parse()
}

func main() {
	var um userMode
	um.parseExecutable()
	um.initFlags()

	if um.readDB {
		extension, recipes := ex00.ParseDB(&um.f)
		db := ex00.MarshalDB(recipes, extension)
		fmt.Println(string(db))
	} else if um.compareDB {
		ex01.CompareDB(&um.old, &um.new)
	} else if um.compareFS {
		ex02.CompareFS(&um.old, &um.new)
	}
}
