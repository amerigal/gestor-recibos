package recibo

import "github.com/gorilla/mux"

func getRouter() *mux.Router {
	myRouter := mux.NewRouter()
	a := NewApiHandler()

	myRouter.HandleFunc("/status", a.GetStatusApi).Methods("GET")
	myRouter.HandleFunc("/recibos", a.GetRecibosApi).Methods("GET")
	myRouter.HandleFunc("/recibos", a.InsertarReciboApi).Methods("POST")
	myRouter.HandleFunc("/recibos/{idR:[0-9]+}", a.GetReciboApi).Methods("GET")
	myRouter.HandleFunc("/recibos/{idR:[0-9]+}", a.EliminarReciboApi).Methods("DELETE")
	myRouter.HandleFunc("/recibos/{idR:[0-9]+}/articulos", a.GetArticulosReciboApi).Methods("GET")
	myRouter.HandleFunc("/recibos/{idR:[0-9]+}/articulos/{idA:[0-9]+}", a.GetArticuloReciboApi).Methods("GET")
	myRouter.HandleFunc("/recibos/{idR:[0-9]+}/articulos/{idA:[0-9]+}", a.SetTipoArticuloReciboApi).Methods("PATCH")
	myRouter.HandleFunc("/recuentos", a.GetRecuentoApi).Queries("periodo", "{periodo:semanal|mensual}", "usuario", "{usuario}").Methods("GET")
	myRouter.HandleFunc("/tendencias", a.GetTendenciaApi).Queries("lugar", "{lugar}").Methods("GET")

	return myRouter
}
