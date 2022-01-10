package recibo

import "net/http"

// LogHttp es un wrapper de http.ResponseWriter para registrar el código de estado de la respuesta HTTP

type logHttp struct {
	http.ResponseWriter
	codigoEstado int
}

// NewLogHttp construye un logHttp a partir de un http.ResponseWriter
func NewLogHttp(w http.ResponseWriter) *logHttp {
	return &logHttp{w, http.StatusOK}
}

// WriteHeader establece el código de estado en el http.ResponseWriter y lo guarda en codigoEstado
func (l *logHttp) WriteHeader(codigo int) {
	l.codigoEstado = codigo
	l.ResponseWriter.WriteHeader(codigo)
}

// LogHttpMiddleware es una función de Middleware para loggear la información asociada a la petición HTTP
func LogHttpMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := NewLogHttp(w)
		h.ServeHTTP(l, r)

		uri := r.URL.String()
		metodo := r.Method
		codigo := l.codigoEstado

		GetLogger().Info("HTTP: %d %s | %s %s", codigo, http.StatusText(codigo), metodo, uri)
	})
}
