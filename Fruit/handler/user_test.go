package handler

// func TestHandler_SignUpUser(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	var appCtx *appcontext.AppContext
// 	var u utils.AppUtil
// 	var v validation.AppValidation
// 	var valP validation.SignUpRequest
// 	t.Run("1: On Siccessfully Singup", func(t *testing.T) {
// 		mockDomain := new(mock.Domain)
// 		mockValidation := new(mock.Validation)
// 		mockValidation.On("SignUpValidation").Return(valP, nil)
// 		mockDomain.On("GetUser").Return(nil, apperrors.ErrUserNotFound)
// 		mockDomain.On("UserSignup").Return(true, nil)
//
// 		testHandler := NewHandler(appCtx, mockDomain, v, u)
//
// 		// appRouter := routes.NewRouter()
// 		// appRouter.SetupRoutes(testHandler)
// 		router := gin.Default()
// 		router.GET("/signup", testHandler.SignUpUser)
//
// 		w := executeRequest(router, "GET", "/signup")
// 		// testHandler.HealthHandler()
// 		mockDomain.AssertExpectations(t)
//
// 		assert.Equal(t, http.StatusOK, w.Code)
//
// 		var response map[string]bool
// 		err := json.Unmarshal([]byte(w.Body.String()), &response)
// 		assert.Nil(t, err)
// 		body := gin.H{
// 			"message": "Successfully SignedUp",
// 		}
// 		value1, _ := response["message"]
// 		assert.Equal(t, body["message"], value1)
//
// 	})
//
// }
