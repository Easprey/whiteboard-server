/*
This hack, taken from https://goswagger.io/faq/faq_testing.html, allows us to
get a configured handler for testing.
*/

package restapi

import (
	"net/http"
	"os"

	"github.com/Easprey/whiteboard-server/handlers"

	"github.com/Easprey/whiteboard-server/restapi/operations"

	loads "github.com/go-openapi/loads"
)

func getAPI() (*operations.WhiteboardAPI, error) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	api := operations.NewWhiteboardAPI(swaggerSpec)
	return api, nil
}

func GetAPIHandler() (http.Handler, error) {
	api, err := getAPI()
	if err != nil {
		return nil, err
	}
	h := configureAPI(api)
	err = api.Validate()
	if err != nil {
		return nil, err
	}

	db = handlers.DataBase{ConnString: os.Getenv("DB_CONN")}
	db.Init() // initialize database for use by tests
	return h, nil
}
