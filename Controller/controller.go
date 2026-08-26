package controller

import(
   "fmt"
    "GitHub/management-/Store"
)

type Controller struct {
	PostgresDB store.StoreOptions
}

func(s *Controller) NewController(pgstore store.Postgres) {
   s.PostgresDB = &pgstore
   s.PostgresDB.NewStore()
   	fmt.Printf("Controller = %v\n", s)
}

type ControllerOptions interface {

   NewController(pgstore store.Postgres) 
   
}
