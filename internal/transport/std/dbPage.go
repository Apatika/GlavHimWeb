package std

import (
	"fmt"
	"glavhim-app/internal/config"
	"net/http"
)

func dbPageGet(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	switch path {
	case config.Cfg.DB.Coll.Users:
		dbPageGetManagers(w, r)
	case config.Cfg.DB.Coll.Cargos:
		dbPageGetCargo(w, r)
	case config.Cfg.DB.Coll.Chems:
		dbPageGetChemistry(w, r)
	default:
		errorResponse(w, fmt.Sprintf("invalid path /db/%v", path), http.StatusInternalServerError)
	}
}

func dbPagePost(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	switch path {
	case config.Cfg.DB.Coll.Users:
		dbPagePostManagers(w, r)
	case config.Cfg.DB.Coll.Cargos:
		dbPagePostCargo(w, r)
	case config.Cfg.DB.Coll.Chems:
		dbPagePostChemistry(w, r)
	default:
		errorResponse(w, fmt.Sprintf("invalid path /db/%v", path), http.StatusInternalServerError)
	}
}

func dbPagePut(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	switch path {
	case config.Cfg.DB.Coll.Users:
		dbPagePutManagers(w, r)
	case config.Cfg.DB.Coll.Cargos:
		dbPagePutCargo(w, r)
	case config.Cfg.DB.Coll.Chems:
		dbPagePutChemistry(w, r)
	default:
		errorResponse(w, fmt.Sprintf("invalid path /db/%v", path), http.StatusInternalServerError)
	}
}

func dbPageDelete(w http.ResponseWriter, r *http.Request) {
	path := r.PathValue(pathVar)
	switch path {
	case config.Cfg.DB.Coll.Users:
		dbPageDeleteManagers(w, r)
	case config.Cfg.DB.Coll.Cargos:
		dbPageDeleteCargo(w, r)
	case config.Cfg.DB.Coll.Chems:
		dbPageDeleteChemistry(w, r)
	default:
		errorResponse(w, fmt.Sprintf("invalid path /db/%v", path), http.StatusInternalServerError)
	}
}
