package photocontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Dedemahendra1/go_tutorial/helper"
	"github.com/Dedemahendra1/go_tutorial/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ResponseJson = helper.ResponseJSON
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var photos []models.Photo

	if err := models.DB.Find(&photos).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, photos)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var photos models.Photo
	if err := models.DB.First(&photos, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Photo tidak ditemukan")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseJson(w, http.StatusOK, photos)
}

func Create(w http.ResponseWriter, r *http.Request) {

	var photos models.Photo

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&photos); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&photos).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusCreated, photos)
}

func Update(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var photos models.Photo

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&photos); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	if models.DB.Where("id = ?", id).Updates(&photos).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak dapat mengupdate photos")
		return
	}

	photos.Id = id

	ResponseJson(w, http.StatusOK, photos)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	input := map[string]string{"id": ""}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	var photos models.Photo
	if models.DB.Delete(&photos, input["id"]).RowsAffected == 0 {
		ResponseError(w, http.StatusBadRequest, "Tidak dapat menghapus photo")
		return
	}

	response := map[string]string{"message": "Photo berhasil dihapus"}
	ResponseJson(w, http.StatusOK, response)
}
