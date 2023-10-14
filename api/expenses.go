package api

import (
	"fmt"
	"net/http"
	"time"
	db "trackit/db/sqlc"

	"github.com/gin-gonic/gin"
)

type expenseResponse struct {
	ID          int64     `json:"id"`
	Userid      int64     `json:"userid"`
	Email       string    `json:"email"`
	Amount      string    `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	// budgetid int64 `json:"budgetid" binding:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func newExpenseResponse(expense db.Expense) expenseResponse {
	return expenseResponse{
		ID:          expense.ID,
		Userid:      expense.Userid,
		Email:       expense.Email,
		Amount:      expense.Amount,
		Description: expense.Description,
		Date:        expense.Date,
		Created_at:  expense.CreatedAt,
	}
}

type expenseParams struct {
	Userid      int64     `json:"userid"`
	Email       string    `json:"email"`
	Amount      string    `json:"amount" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Tag         string    `json:"tag" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	// budgetid int64 `json:"budgetid" binding:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

// createExpense   godoc
//
//	@Summary		Create Expense Transaction
//	@Description	Responds with a newly created expense record as JSON.
//	@Tags			expense
//	@Produce		json
//	@Param			params	body		expenseParams	true	"Expense JSON"
//	@Success		201		{object}	expenseResponse
//	@Security		BearerAuth
//	@Router			/expense [post]
func (srv *Server) CreateExpense(ctx *gin.Context) {
	fmt.Println("context:", ctx)

	token := ctx.Request.Header.Get("Authorization")

	fmt.Println("token:", token)

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

	var params expenseParams

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	params.Email = user.Email
	params.Userid = user.ID

	fmt.Println("params:", params)

	arg := db.CreateExpenseParams{
		Userid:      params.Userid,
		Email:       params.Email,
		Amount:      params.Amount,
		Description: params.Description,
		Date:        params.Date,
	}

	expense, err := srv.store.CreateExpense(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("error creating expense", err))
		return
	}

	type createExpenseResponse struct {
		Expense expenseResponse `json:"expense"`
	}

	response := createExpenseResponse{
		Expense: newExpenseResponse(expense),
	}

	ctx.JSON(http.StatusCreated, response)
	return
}

type listExpensesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// createExpense   godoc
//
//	@Summary		Get Expenses Transactions
//	@Description	Responds with a list of expense records as JSON.
//	@Tags			expense
//	@Produce		json
//	@Param			params	body	listExpensesRequest	true	"Expense JSON"
//	@Success		200		{array}	expenseResponse
//	@Security		BearerAuth
//	@Router			/expense [get]
func (srv *Server) GetExpenses(ctx *gin.Context) {

	var req listExpensesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse("user input not valid", err))
		return
	}

	arg := db.ListExpensesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	expense, err := srv.store.ListExpenses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse("error getting expense", err))
		return
	}

	type getExpenseResponse struct {
		Expenses []expenseResponse `json:"expenses"`
	}

	response := getExpenseResponse{
		Expenses: make([]expenseResponse, len(expense)),
	}

	for i, expense := range expense {
		response.Expenses[i] = newExpenseResponse(expense)
	}

	ctx.JSON(http.StatusOK, response)
	return
}
