package e2e_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Easprey/whiteboard-server/restapi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Suite")
}

var _ = Describe("FingerPathsGet", func() {
	var (
		response *http.Response
		server   *httptest.Server
		body     []byte
		err      error
	)

	BeforeEach(func() {
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
			fmt.Printf("%s", body)
			Expect(body).To(HaveLen(0))
		})
	})
})
