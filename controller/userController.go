package controller


import(
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"connection"
	"model"

	"github.com/gorilla/mux"
)

func UserGet(w http.ResponseWriter, r *http.Request)  {

	w.Header().Add("Content-Type", "aplication/json")

	res := connection.Db()

	row, errRow := res.Query("SELECT * FROM client")

	if errRow != nil {

		w.WriteHeader(http.StatusInternalServerError)

		return 

	}

	var users []model.UserModel = make([]model.UserModel, 0)

	for row.Next() {

		var user model.UserModel

		erroScan := row.Scan(&user.Id, &user.Email, &user.Name, &user.Value)

		if erroScan != nil {

			w.WriteHeader(http.StatusInternalServerError)

			continue 

		}

		users = append(users, user)

	}

	erroClose := row.Close()

	if erroClose != nil {

		log.Println("ErroClose: " + erroClose.Error())

	}

	encoder := json.NewEncoder(w)

	encoder.Encode(users)
	
}

func UserGetId(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	id := vars["userId"]

	res := connection.Db()

	row := res.QueryRow("SELECT * FROM client WHERE id = ?", id)

	var user model.UserModel

	erroScan := row.Scan(&user.Id, &user.Email, &user.Name, &user.Value)

	if erroScan != nil {

		w.WriteHeader(http.StatusNotFound)

		return

	}

	encoder := json.NewEncoder(w)

	encoder.Encode(user)

}

func UserPost(w http.ResponseWriter, r *http.Request)  {

	w.Header().Add("Content-Type", "aplication/json")

	body, errBody := ioutil.ReadAll(r.Body)

	if errBody != nil {

		w.WriteHeader(http.StatusBadRequest)

		return

	}

	var newUser model.UserModel

	json.Unmarshal(body, &newUser)

	res := connection.Db()

	row, errRes := res.Exec(
		"INSERT iNTO client (email, name, value) VALUES (?, ?, ?)", 
		newUser.Email, 
		newUser.Name, 
		newUser.Value,
	)

	if errRes != nil {

		w.WriteHeader(http.StatusInternalServerError)

		return

	}

	idGenerate, errId := row.LastInsertId()

	if errId != nil {

		w.WriteHeader(http.StatusInternalServerError)

		return

	}

	newUser.Id = int(idGenerate)

	w.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(w)

	encoder.Encode(newUser)
	
}

func UserUpdata(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	id := vars["userId"]

	body, errBody := ioutil.ReadAll(r.Body)

	if errBody != nil {

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var newUpdate model.UserModel

	errShal := json.Unmarshal(body, &newUpdate)

	if errShal != nil {

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res := connection.Db()

	row := res.QueryRow("SELECT * FROM client WHERE id = ?", id)

	var user model.UserModel

	errSchan := row.Scan(&user.Id, &user.Email, &user.Name, &user.Value)

	if errSchan != nil {

		w.WriteHeader(http.StatusNotFound)

		return

	}

	_, errRes := res.Exec(
		"UPDATE client SET email = ?, name = ?, value = ? WHERE id = ?", 
		newUpdate.Email, 
		newUpdate.Name, 
		newUpdate.Value, 
		id,
	)

	if errRes != nil {

		w.WriteHeader(http.StatusInternalServerError)

		return 

	}

}

func UserDelete(w http.ResponseWriter, r *http.Request){

	w.Header().Add("Content-Type", "aplication/json")

	vars := mux.Vars(r)

	id := vars["userId"]

	res := connection.Db()

	row := res.QueryRow("SELECT id From client WHERE id =?", id)

	var user model.UserModel

	errScan := row.Scan(&user.Id)

	if errScan != nil {

		w.WriteHeader(http.StatusInternalServerError)

		return

	}

	_, errRes := res.Exec("DELETE FROM client WHERE id = ?", id)

	if errRes != nil {

		w.WriteHeader(http.StatusInternalServerError)

		return

	}

	w.WriteHeader(http.StatusNoContent)
}