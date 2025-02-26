package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/xoxoharsh/go-student-api/internal/storage"
	"github.com/xoxoharsh/go-student-api/internal/types"
	"github.com/xoxoharsh/go-student-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		var student types.Student
		
		err := json.NewDecoder(r.Body).Decode(&student)
		
		if errors.Is(err, io.EOF) {
			slog.Error("Empty request body")
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// validate request body
		if err:=validator.New().Struct(student); err!=nil {

			validateErrs := err.(validator.ValidationErrors)

			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		slog.Info("Creating new student")
		
		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)

		if err != nil {
			response.WriteJSON(w, http.StatusInternalServerError, err)
			return
		}	

		slog.Info("Student created", slog.Int64("id", lastId))
		

		response.WriteJSON(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}