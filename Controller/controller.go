package controller

import (
	model "GitHub/management-/Model"
	store "GitHub/management-/Store"
	util "GitHub/management-/Util"
	"fmt"
)

type Controller struct {
	PostgresDB store.StoreOptions
}

func (s *Controller) NewController(pgstore store.Postgres) {
	util.SetLogger()
	util.Logger.Info("Controller initialized")
	s.PostgresDB = &pgstore
	err := s.PostgresDB.NewStore()
	if err != nil {
		util.Logger.Errorf("Failed to initialize Postgres store, error:=%v", err)
		util.Log(model.LogLevelError, model.Controller, "NewController", "Failed to initialize Postgres store", nil)

	} else {
		util.Logger.Info("Postgres store initialized successfully")
		util.Log(model.LogLevelInfo, model.Controller, model.NewController, "Controller initialized successfully", nil)
	}

	fmt.Printf("Controller = %v\n", s)
}

type ControllerOptions interface {
	NewController(pgstore store.Postgres)
}
