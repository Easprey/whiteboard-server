package handlers_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Easprey/whiteboard-server/handlers"
)

var _ = Describe("Database", func() {
	var (
		db DataBase
	)

	Context("with valid connection string", func() {
		BeforeEach(func() {
			db.ConnString = os.Getenv("DB_CONN")
		})

		It("should not panic", func() {
			Expect(db.Init).NotTo(Panic())
		})
	})

	Context("with invalid connection string,", func() {
		BeforeEach(func() {
			db.ConnString = "ThisIsNotValid"
		})

		It("should panic", func() {
			Expect(db.Init).To(Panic())
		})
	})
})
