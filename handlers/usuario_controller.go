package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"www.github.com/miprimerproyectoengo/entities"
	"www.github.com/miprimerproyectoengo/services"
)

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)

type UsuarioController interface {
	HandleGeneral(writer http.ResponseWriter, request *http.Request)
	HandleOne(writer http.ResponseWriter, request *http.Request)
}

type usuarioController struct {
	service services.UsuarioService
}

func NewUsuarioController() UsuarioController {
	return &usuarioController{
		service: services.NewUsuarioService(),
	}
}

func (h *usuarioController) HandleGeneral(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Println(request.Method)
	if request.Method == GET {
		usuarios := h.service.FindAll()
		json.NewEncoder(writer).Encode(usuarios)
	}
	if request.Method == POST {
		usuario := entities.Usuario{}
		json.NewDecoder(request.Body).Decode(&usuario)

		h.service.Save(usuario)
		json.NewEncoder(writer).Encode(usuario)
	}

	// if request.Method == DELETE {
	// 	json.NewEncoder(writer).Encode(DELETE)
	// }
}
func (h *usuarioController) HandleOne(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if request.Method == GET {
		id, encontrado := mux.Vars(request)["id"]

		if encontrado {
			usuario, err := h.service.FindOne(id)

			if err != nil {
				writer.WriteHeader(http.StatusNotFound)
				json.NewEncoder(writer).Encode(err.Error())
				return
			}
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(usuario)
			return
		}

		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(nil)
	}
	if request.Method == PUT {
		id, encontrado := mux.Vars(request)["id"]
		if encontrado {
			usuario := entities.Usuario{}
			json.NewDecoder(request.Body).Decode(&usuario)

			usuarioNew := h.service.Update(id, usuario)

			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(usuarioNew)
			return
		}
		json.NewEncoder(writer).Encode("Error encontrado..")
	}
}
