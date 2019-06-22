package handlers_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Easprey/whiteboard-server/handlers"
)

var _ = Describe("Handlers", func() {
	var (
		db DataBase
	)

	BeforeEach(func() {
		db.ConnString = fmt.Sprintf("%s?multiStatements=true", os.Getenv("DB_CONN"))
		db.Init()

		executeSQLScript("testdata/delete_all.sql", db)
		executeSQLScript("testdata/mock_data.sql", db)
	})

	Context("infrastructure test", func() {
		It("should get to this point", func() {
			Expect(func() {}).NotTo(Panic())
		})
	})
})
