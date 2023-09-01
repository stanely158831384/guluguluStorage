package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/stanely158831384/guluguluStorage/db/mock"
	db "github.com/stanely158831384/guluguluStorage/db/sqlc"
	"github.com/stanely158831384/guluguluStorage/util"
	"github.com/stretchr/testify/require"
)

func randomCategory(t *testing.T) (category db.Category){
		id := util.RandomInt(1,1000)
		name := string("produces")
		category = db.Category{
			ID: id,
			Name: name,
		}
		return category
}

func TestListCategoriesNoAuth(t *testing.T){
	n := 10
	catergories := make([]db.Category,n)
	for i := 0; i < n; i++ {
		catergories[i] = randomCategory(t)
	}
	type Query struct {
		Limit  int32
		Offset *int32
	}
	Limit := int32(5)
	Offset := int32(0)
	
	testCases := []struct{
		name string
		query Query
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			query: Query{
				Limit: Limit,
				Offset: &Offset,
			},
			buildStubs: func(store *mockdb.MockStore){
				arg := db.ListCategoriesParams{
					Limit: Limit,
					Offset: Offset,
				}
				fmt.Printf("arg is %v\n",arg)
				store.EXPECT().ListCategories(gomock.Any(),gomock.Eq(arg)).Times(1).Return(catergories,nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder){
				require.Equal(t,http.StatusOK, recorder.Code)
				requireBodyMatchCategory(t,recorder.Body,catergories)
			},
		},
		{
			name: "InternalError",
			query: Query{
				Limit: Limit,
				Offset: &Offset,
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().ListCategories(gomock.Any(),gomock.Any()).Times(1).Return([]db.Category{},sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder){
				require.Equal(t,http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidLimit",
			query: Query{
				Limit: 0,
				Offset: &Offset,
			},
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().ListCategories(gomock.Any(),gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder){
				require.Equal(t,http.StatusBadRequest, recorder.Code)
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

			url := "/listCategories/noAuth"
			request, err := http.NewRequest(http.MethodGet,url,nil)
			require.NoError(t,err)

			q := request.URL.Query()
			q.Add("limit",fmt.Sprintf("%d", tc.query.Limit))
			q.Add("offset",fmt.Sprintf("%d", *tc.query.Offset))
			request.URL.RawQuery = q.Encode()
			fmt.Printf("request.URL.RawQuery is %v\n",request.URL.RawQuery)
			server.router.ServeHTTP(recorder,request)
			fmt.Printf("listCategoriesNoAuth is called6\n")
			tc.checkResponse(recorder)
			// requireBodyMatchCategory(t, recorder.Body, catergories)

		})
	}
}




func TestGetGategoryAPI(test *testing.T) {
	category := randomCategory(test)

	ctrl := gomock.NewController(test)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().GetCategory(gomock.Any(), gomock.Eq(category.ID)).Times(1).Return(category, nil)

	server, _ := NewServer2(store)

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/getCategory/noAuth/%d", category.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(test, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(test, http.StatusOK, recorder.Code)
}

func TestListGategoryAPI(t *testing.T) {
	n := 10
	catergories := make([]db.Category,n)
	for i := 0; i < n; i++ {
		catergories[i] = randomCategory(t)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	arg := db.ListCategoriesParams{
		Limit: int32(5),
		Offset: int32(0),
	}
	store.EXPECT().ListCategories(gomock.Any(), gomock.Eq(arg)).Times(1).Return(catergories, nil)
	server, _ := NewServer2(store)
	recorder := httptest.NewRecorder()
	url := "/listCategories/noAuth"
	request, err := http.NewRequest(http.MethodGet,url,nil)
	require.NoError(t,err)

	q := request.URL.Query()
	q.Add("limit",fmt.Sprintf("%d", arg.Limit))
	q.Add("offset",fmt.Sprintf("%d", arg.Offset))
	request.URL.RawQuery = q.Encode()

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchCategory(t, recorder.Body, catergories)

}

func requireBodyMatchCategory(t *testing.T, body *bytes.Buffer, category []db.Category){
	data, err := ioutil.ReadAll(body)
	require.NoError(t,err)

	var gotCategories []db.Category
	err = json.Unmarshal(data,&gotCategories)
	require.NoError(t,err)
	require.Equal(t,category,gotCategories)
}