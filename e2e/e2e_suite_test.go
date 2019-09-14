package e2e_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Easprey/whiteboard-server/handlers"
	"github.com/Easprey/whiteboard-server/restapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Suite")
}

func executeSQLScript(file string, db handlers.DataBase) {
	script, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panic("couldn't read SQL file", err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		log.Panic("failed to execute SQL script", err)
	}
}

func getDB() handlers.DataBase {
	// read in database connection string
	connString, exists := os.LookupEnv("DB_CONN")
	if !exists {
		Fail("Failed to read connection string")
	}

	// allow multiple statements in connection
	connString = fmt.Sprintf("%s?multiStatements=true", connString)

	db := handlers.DataBase{ConnString: connString}
	db.Init()
	return db
}

var _ = Describe("FingerPathsGet", func() {
	var (
		db       handlers.DataBase
		response *http.Response
		server   *httptest.Server
		body     []byte
		err      error
	)

	// set up database
	db = getDB()
	db.Init()

	// start server

	BeforeEach(func() {
		// reset the database
		executeSQLScript("delete_all.sql", db)
		executeSQLScript("mock_data.sql", db)

		// get handler using hack
		handler, err := restapi.GetAPIHandler()
		if err != nil {
			Fail("Failed to get handler")
		}
		server = httptest.NewServer(handler)
	})

	AfterEach(func() {
		server.Close()
	})

	Context("When paths exist", func() {
		BeforeEach(func() {
			response, err = http.Get(server.URL + "/api/v1/boards/test_board/fingerpaths")
			body, err = ioutil.ReadAll(response.Body)
		})

		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return 200", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})

		It("should return some paths", func() {
			Expect(body).NotTo(HaveLen(0))
		})

		It("should return the right paths", func() {
			// TODO: make this not broken
			Expect(body).To(MatchJSON([]byte(`[{"userId": "hello"}]`)))
		})
	})

	Context("When paths don't exist", func() {
		BeforeEach(func() {
			response, err = http.Get(server.URL + "/api/v1/boards/NOT_A_REAL_BOARD/fingerpaths")
			body, err = ioutil.ReadAll(response.Body)

		})

		It("should not return any paths", func() {
			Expect(body).To(HaveLen(0))
		})
	})
})
