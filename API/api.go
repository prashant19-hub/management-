package api

import (
	"GitHub/management-/Controller"
	 "GitHub/management-/Store"
)
type APIRouts struct {

	Controller controller.ControllerOptions

}

func(api *APIRouts) StartAPI(controller controller.Controller) {
            api.Controller = &controller
	api.Controller.NewController(store.Postgres{})
}