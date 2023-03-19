package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	db "trackit/db/sqlc"
	"trackit/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userResponse struct {
	UserID    int64  `json:"userid"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	UserType  string `json:"user_type"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		UserID:    user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Phone:     user.Phone,
		UserType:  user.UserType,
	}
}

func (s *Server) CreateUser(ctx *gin.Context) {
	type createUserParams struct {
		Firstname string `json:"firstname" binding:"required"`
		Lastname  string `json:"lastname" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
		Phone     string `json:"phone" binding:"required"`
		UserType  string `json:"user_type"`
	}

	var params createUserParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error getting user input", err))
		return
	}

	hashPassword, err := util.HashPassword(params.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("Error hashing password", err))
		return
	}

	switch params.Email {
	case "blessedmadukoma@gmail.com":
		params.UserType = "admin"
	default:
		params.UserType = "user"
	}

	arg := db.CreateUserParams{
		Firstname: params.Firstname,
		Lastname:  params.Lastname,
		Email:     params.Email,
		Password:  hashPassword,
		Phone:     params.Phone,
		UserType:  params.UserType,
	}

	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error creating user", err))
		return
	}

	ctx.JSON(http.StatusCreated, newUserResponse(user))
	return
}

func (srv *Server) loginUser(ctx *gin.Context) {
	type loginUserRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=5"`
	}

	type LoginUserResponse struct {
		SessionID             uuid.UUID    `json:"session_id"`
		AccessToken           string       `json:"access_token"`
		AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
		RefreshToken          string       `json:"refresh_token"`
		RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
		User                  userResponse `json:"user"`
	}

	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error in login parameters", err))
		return
	}

	fmt.Println("req:", req)

	user, err := srv.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse("user not found", err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse("error communicating with db", err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse("password does not match", err))
		return
	}

	accessToken, accessPayload, err := srv.tokenMaker.CreateToken(user.Email, srv.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("error creating access token", err))
		return
	}

	refreshToken, refreshPayload, err := srv.tokenMaker.CreateToken(
		user.Email,
		srv.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("error creating refresh token", err))
		return
	}

	session, err := srv.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Userid:       user.ID,
		Email:        user.Email,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("error creating session", err))
		return
	}

	response := LoginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, response)
	return
}

// func (srv *Server) LogoutUser(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
// 	return
// }

// getCurrentUserBySession gets the current user by session
func (srv *Server) getCurrentUserBySession(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	payload, err := srv.tokenMaker.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse("error verifying token", err))
		return
	}

	user, err := srv.store.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse("error getting user", err))
		return
	}

	type getCurrentUserResponse struct {
		User userResponse `json:"user"`
	}

	response := getCurrentUserResponse{
		User: newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, response)
	return
}

func (srv *Server) GetUserByEmail(ctx *gin.Context) {

	type getUserByEmailParams struct {
		Email string `json:"email" binding:"required"`
	}

	var params getUserByEmailParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error getting user input", err))
		return
	}

	user, err := srv.store.GetUserByEmail(ctx, params.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("error getting user", err))
		return
	}

	type getUserResponse struct {
		User userResponse `json:"user"`
	}

	response := getUserResponse{
		User: newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, response)
	return
}

func (srv *Server) GetUserByID(ctx *gin.Context) {
	type getUserByIDParams struct {
		UserID int64 `uri:"id" binding:"required,min=1"`
	}

	var params getUserByIDParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error getting user input", err))
		return
	}

	user, err := srv.store.GetUserByID(ctx, params.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("error getting user", err))
		return
	}

	type getUserResponse struct {
		User userResponse `json:"user"`
	}

	response := getUserResponse{
		User: newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, response)
	return
}

func (srv *Server) UpdateUser(ctx *gin.Context) {
	type updateUserParams struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Phone     string `json:"phone"`
		UserType  string `json:"user_type"`
	}

	var params updateUserParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error getting user input", err))
		return
	}

	arg := db.UpdateUserParams{
		Firstname: params.Firstname,
		Lastname:  params.Lastname,
		Phone:     params.Phone,
		UserType:  params.UserType,
	}

	user, err := srv.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error updating user", err))
		return
	}

	ctx.JSON(http.StatusOK, newUserResponse(user))
	return
}

func (srv *Server) DeleteUser(ctx *gin.Context) {
	type getUserByIDParams struct {
		UserID int64 `uri:"id" binding:"required,min=1"`
	}

	var params getUserByIDParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error getting user input", err))
		return
	}

	err := srv.store.DeleteUser(ctx, params.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("error deleting user", err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	return
}

func (srv *Server) GetUsers(ctx *gin.Context) {
	type listAccountsRequest struct {
		PageID   int32 `form:"page_id" binding:"required,min=1"`
		PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
	}
	var req listAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("user input not valid", err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := srv.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("error getting users", err))
		return
	}

	type getUsersResponse struct {
		Users []userResponse `json:"users"`
	}

	response := getUsersResponse{
		Users: make([]userResponse, len(users)),
	}

	for i, user := range users {
		response.Users[i] = newUserResponse(user)
	}

	ctx.JSON(http.StatusOK, response)
	return
}
