package e2e_test

import (
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
		err      error
		// boardName string
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
			// boardName = "test"
			response, err = http.Get(server.URL + "/api/v1/boards/test/fingerpaths")
		})

		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return 200", func() {
			Expect(response.StatusCode).To(Equal(http.StatusOK))
		})
	})
})
