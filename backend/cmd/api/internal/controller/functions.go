package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CMS-Enterprise/ztmf/backend/internal/model"
	"github.com/gorilla/mux"
)

func ListFunctions(w http.ResponseWriter, r *http.Request) {

	var (
		functions []*model.Function
		err       error
	)

	findFunctionsInput := model.FindFunctionsInput{}
	err = decoder.Decode(&findFunctionsInput, r.URL.Query())
	if err == nil {
		functions, err = model.FindFunctions(r.Context(), findFunctionsInput)
	}
	respond(w, r, functions, err)
}

func GetFunctionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, ok := vars["functionid"]
	if !ok {
		respond(w, r, nil, ErrNotFound)
		return
	}
	var functionID int32
	fmt.Sscan(ID, &functionID)

	f, err := model.FindFunctionByID(r.Context(), functionID)

	respond(w, r, f, err)
}

func SaveFunction(w http.ResponseWriter, r *http.Request) {
	user := model.UserFromContext(r.Context())
	if !user.IsAdmin() {
		respond(w, r, nil, ErrForbidden)
		return
	}

	f := &model.Function{}

	err := getJSON(r.Body, f)
	if err != nil {
		log.Println(err)
		respond(w, r, nil, ErrMalformed)
		return
	}

	vars := mux.Vars(r)
	if v, ok := vars["functionid"]; ok {
		fmt.Sscan(v, &f.FunctionID)
	}

	f, err = f.Save(r.Context())

	if err != nil {
		respond(w, r, nil, err)
		return
	}

	respond(w, r, f, nil)

}
