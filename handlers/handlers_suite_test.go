package handlers_test

import (
	"io/ioutil"
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Easprey/whiteboard-server/handlers"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers Suite")
}

func executeSQLScript(file string, db DataBase) {
	script, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panic("couldn't read SQL file", err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		log.Panic("failed to execute SQL script", err)
	}
}
