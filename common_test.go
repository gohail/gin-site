package main

import (
	"github.com/andreyberezin/gin-site/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var tmpArticleList [] models.Article

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()

	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func testHTTPResponce(t *testing.T, r *gin.Engine, req *http.Request, f func (w *httptest.ResponseRecorder) bool){
	// Create a response recorder
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This function is used to store the main lists into the temporary one
// for testing
func saveLists() {
	tmpArticleList = models.GetAllArticles()
}
