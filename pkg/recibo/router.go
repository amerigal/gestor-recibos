package recibo

import "github.com/gorilla/mux"

func getRouter() *mux.Router {
	myRouter := mux.NewRouter()
	a := NewApiHandler()

	myRouter.HandleFunc("/status", a.GetStatusApi).Methods("GET")
	myRouter.HandleFunc("/recibos", a.GetRecibosApi).Methods("GET")
	myRouter.HandleFunc("/recibos/{usuario}/{nombreFicheroInput}", a.InsertarReciboApi).Methods("POST")
	myRouter.HandleFunc("/recibos/{idR:[0-9]+}", a.GetReciboApi).Methods("GET")
	myRouter.HandleFunc("/recibos/{idR:[0-9]+}", a.EliminarReciboApi).Methods("DELETE")
	myRouter.HandleFunc("/recibos/{idR:[0-9]+}/articulos", a.GetArticulosReciboApi).Methods("GET")
	myRouter.HandleFunc("/recibos/{idR:[0-9]+}/articulos/{idA:[0-9]+}", a.GetArticuloReciboApi).Methods("GET")
	myRouter.HandleFunc("/recibos/{idR:[0-9]+}/articulos/{idA:[0-9]+}/{tipo}", a.SetTipoArticuloReciboApi).Methods("PUT")
	myRouter.HandleFunc("/recuentos/{periodo:semanal|mensual}/{usuario}", a.GetRecuentoApi).Methods("GET")
	myRouter.HandleFunc("/tendencias/{lugar}", a.GetTendenciaApi).Methods("GET")

	return myRouter
}
