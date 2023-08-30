package rapid

import (
	"encoding/json"
	"net/http"
)

type Ctx struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func (ctx *Ctx) JSON(v interface{}) {
	ctx.Response.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(ctx.Response).Encode(v)
	if err != nil {
		return
	}
}
