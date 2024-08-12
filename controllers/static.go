package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aminghafoory/multisig2/views/home"
)

func Home(w http.ResponseWriter, r *http.Request) {

	c := home.Index("home", r, true)
	err := c.Render(context.Background(), w)
	if err != nil {
		http.Error(w, fmt.Sprintf("Err : %s", err), http.StatusInternalServerError)
	}
}
