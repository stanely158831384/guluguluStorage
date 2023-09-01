package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/stanely158831384/guluguluStorage/db/mock"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
	"github.com/stanely158831384/guluguluStorage/token"
	"github.com/stanely158831384/guluguluStorage/util"
	"github.com/stretchr/testify/require"
)
func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	return
}

func TestCreateCategoryAPI(t *testing.T) {
	user, _ := randomUser(t)
	category :=randomCategory(t)

	testCases := []struct{
		name string
		body gin.H
		setupAuth func(t *testing.T,request *http.Request, tokenMaker token.Maker)
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name": category.Name,
			},
			setupAuth: func(t *testing.T,request *http.Request, tokenMaker token.Maker){
				addAuthorization(t, request, tokenMaker,authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().CreateCategory(gomock.Any(),gomock.Eq(category.Name)).Times(1).Return(category,nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder){
				require.Equal(t,http.StatusOK, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"name": category.Name,
			},
			setupAuth: func(t *testing.T,request *http.Request, tokenMaker token.Maker){
				addAuthorization(t, request, tokenMaker,authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().CreateCategory(gomock.Any(),gomock.Any()).Times(1).Return(db.Category{},sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder){
				require.Equal(t,http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidName",
			body: gin.H{
				"name": "",
			},
			setupAuth: func(t *testing.T,request *http.Request, tokenMaker token.Maker){
				addAuthorization(t, request, tokenMaker,authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().CreateCategory(gomock.Any(),gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder){
				require.Equal(t,http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "NoAuthorization",
			body: gin.H{
				"name": category.Name,
			},
			setupAuth: func(t *testing.T,request *http.Request, tokenMaker token.Maker){
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().CreateCategory(gomock.Any(),gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder){
				require.Equal(t,http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases{
		tc := testCases[i]
		t.Run(tc.name,func(t *testing.T){
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t,store)
			// server, _ := NewServer2(store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t,err)

			url := "/categories"
			request, err := http.NewRequest(http.MethodPost,url,bytes.NewReader(data))
			require.NoError(t,err)

			tc.setupAuth(t,request,server.tokenMaker)
			server.router.ServeHTTP(recorder,request)
			tc.checkResponse(recorder)
		})
	}
}