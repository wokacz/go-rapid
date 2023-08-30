package rapid

import "net/http"

type (
	Handler func(*Ctx)
	App     struct {
		mux    *http.ServeMux
		prefix string
	}
)

func NewApp() *App {
	return &App{
		mux: http.NewServeMux(),
	}
}

func (a *App) Run(addr string) (err error) {
	return http.ListenAndServe(addr, a.mux)
}

// HTTP

func createPath() string {
	return ""
}

func (a *App) Get(pattern string, handler Handler) {
	a.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		ctx := &Ctx{
			Response: w,
			Request:  r,
		}
		handler(ctx)
	})
}

func (a App) Post(pattern string, handler Handler) {
	a.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		ctx := &Ctx{
			Response: w,
			Request:  r,
		}
		handler(ctx)
	})
}

func (a *App) Put(pattern string, handler Handler) {
	a.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		ctx := &Ctx{
			Response: w,
			Request:  r,
		}
		handler(ctx)
	})
}

func (a *App) Delete(pattern string, handler Handler) {

	a.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		ctx := &Ctx{
			Response: w,
			Request:  r,
		}
		handler(ctx)
	})
}

// PREFIX

func (a *App) Prefix(pattern string, handler Handler) {

}
