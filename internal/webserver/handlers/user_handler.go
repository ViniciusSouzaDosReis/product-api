package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ViniciusSouzaDosReis/product-api/internal/dto"
	"github.com/ViniciusSouzaDosReis/product-api/internal/entity"
	"github.com/ViniciusSouzaDosReis/product-api/internal/infra/database/interfaces"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB        interfaces.UserInterface
	Jwt           *jwtauth.JWTAuth
	JwtExperiesIn int
}

type Error struct {
	Message string `json:"message"`
}

func NewUserHandler(db interfaces.UserInterface, jwt *jwtauth.JWTAuth, jwtExperiesIn int) *UserHandler {
	return &UserHandler{
		UserDB:        db,
		Jwt:           jwt,
		JwtExperiesIn: jwtExperiesIn,
	}
}

// GenerateToken godoc
// @Summary      Generate Token
// @Description  Generate Token
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request	body dto.GenerateTokenInput true "user credentials"
// @Success      200 	{object} dto.GenerateTokenOutput
// @Failure      400	{object} Error
// @Failure      401	{object} Error
// @Router       /user/generate_token [post]
func (h *UserHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	var user dto.GenerateTokenInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		errorMessage := Error{
			Message: "password is invalid",
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExperiesIn)).Unix(),
	})
	accessToken := dto.GenerateTokenOutput{
		AccessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        request	body dto.CreateUserInpt true "user request"
// @Success      201
// @Failure      500	{object} Error
// @Failure      400	{object} Error
// @Router       /user [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInpt
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := Error{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(errorMessage)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
