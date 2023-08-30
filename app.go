package rapid

import "net/http"

type HandleFunc func(*Ctx)

type App struct {
	mux *http.ServeMux
}

func NewApp() *App {
	return &App{
		mux: http.NewServeMux(),
	}
}

func (a *App) Run(addr string) (err error) {
	return http.ListenAndServe(addr, a.mux)
}

func (a *App) GET(pattern string, handler HandleFunc) {
	a.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := &Ctx{
			Response: w,
			Request:  r,
		}
		handler(ctx)
	})
}
