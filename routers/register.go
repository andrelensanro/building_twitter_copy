package routers

import (
	"encoding/json"
	"net/http"

	"github.com/andrelensanro/building_twitter_copy/bd"
	"github.com/andrelensanro/building_twitter_copy/models"
)

/*Register es la funcion para crear en la base de datos el registro de usuario*/
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	/*El objeto Body una vez leido ya no existe, ya que es de tipo stream*/
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "No se encuentra el email en los datos", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Se debe especificar una contraseña de al menos 6 digitos", 400)
		return
	}
	_, found, _ := bd.CheckExists(t.Email)

	if found == true {
		http.Error(w, "This item already exists", 400)
		return
	}
	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "Something went wrong: insert register"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "User registration was not performed", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
