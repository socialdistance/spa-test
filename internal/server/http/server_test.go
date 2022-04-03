package internalhttp

import (
	"bytes"
	"context"
	"fmt"
	"github.com/socialdistance/spa-test/internal/app"
	internalconfig "github.com/socialdistance/spa-test/internal/config"
	"github.com/socialdistance/spa-test/internal/logger"
	sqlstorage "github.com/socialdistance/spa-test/internal/storage/sql"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"
)

var configFile = "config.yaml"

func Test_HttpServerHelloWorld(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	httpHandlers := NewRouter(createApp(t))
	httpHandlers.ServeHTTP(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	require.Nil(t, err)
	require.Equal(t, "Hello, world!\n", string(body))
}

func Test_HttpSearch(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "post",
		"description": ""
	}`)

	req := httptest.NewRequest("POST", "/posts/search", body)
	w := httptest.NewRecorder()

	httpHandlers := NewRouter(createApp(t))
	httpHandlers.ServeHTTP(w, req)

	resp := w.Result()
	res, err := io.ReadAll(resp.Body)
	fmt.Println(string(res))
	require.Nil(t, err)
}

func Test_HttpPagination(t *testing.T) {
	req := httptest.NewRequest("GET", "/posts/1", nil)
	w := httptest.NewRecorder()

	httpHandlers := NewRouter(createApp(t))
	httpHandlers.ServeHTTP(w, req)

	resp := w.Result()
	res, err := io.ReadAll(resp.Body)
	fmt.Println(string(res))
	require.Nil(t, err)
}

func Test_HttpLogin(t *testing.T) {
	body := bytes.NewBufferString(`{
		"username": "test1",
		"password": "test1"
	}`)

	req := httptest.NewRequest("POST", "/login", body)
	w := httptest.NewRecorder()

	httpHandlers := NewRouter(createApp(t))
	httpHandlers.ServeHTTP(w, req)

	resp := w.Result()
	res, err := io.ReadAll(resp.Body)
	fmt.Println("RES:", string(res))
	require.Nil(t, err)
}

func Test_HttpSelectedPost(t *testing.T) {
	body := bytes.NewBufferString(`{
		"id": "5753e882-91e0-4e1a-a827-eef8d8271e50",
		"title": "Test title",
		"created": "2022-04-01T18:43:25.391Z",
		"description": "test description",
		"userId": "1528371b-229c-4370-839a-0571d969902a"
	}`)

	req := httptest.NewRequest("POST", "/post", body)
	w := httptest.NewRecorder()

	httpHandlers := NewRouter(createApp(t))
	httpHandlers.ServeHTTP(w, req)

	resp := w.Result()
	_, err := io.ReadAll(resp.Body)
	require.Nil(t, err)
}

func Test_HttpListAll(t *testing.T) {
	req := httptest.NewRequest("GET", "/posts", nil)
	w := httptest.NewRecorder()

	httpHandlers := NewRouter(createApp(t))
	httpHandlers.ServeHTTP(w, req)

	resp := w.Result()
	res, err := io.ReadAll(resp.Body)
	fmt.Println(string(res))
	require.Nil(t, err)
}

func Test_HttpCreateComment(t *testing.T) {
	body := bytes.NewBufferString(`{
		"id": "a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b8",
		"username": "test username",
		"content": "test content",
		"userId": "1528371b-229c-4370-839a-0572d969902a",
		"postId": "a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b0"
	}`)

	req := httptest.NewRequest("POST", "/comments/create", body)
	w := httptest.NewRecorder()

	httpHandlers := NewRouter(createApp(t))
	httpHandlers.ServeHTTP(w, req)

	resp := w.Result()
	_, err := io.ReadAll(resp.Body)
	require.Nil(t, err)

	body = bytes.NewBufferString(`{
		"username": "test username",
		"content": "test content 2",
		"userId": "1528371b-229c-4370-839a-0572d969902a",
		"postId": "a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b0"
	}`)

	request := httptest.NewRequest("PUT", "/comments/update/a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b8", body)
	w = httptest.NewRecorder()

	httpHandlers.ServeHTTP(w, request)

	response := w.Result()
	responseBody, _ := ioutil.ReadAll(response.Body)
	responseExcepted := `{"id":"a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b8","username":"test username","content":"test content 2","userId":"1528371b-229c-4370-839a-0572d969902a","postId":"a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b0"}` //nolint:lll
	require.Equal(t, responseExcepted, string(responseBody))

	request = httptest.NewRequest("DELETE", "/comments/delete/a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b8", body)
	w = httptest.NewRecorder()

	httpHandlers.ServeHTTP(w, request)

	response = w.Result()
	responseBody, _ = ioutil.ReadAll(response.Body)
	fmt.Println(string(responseBody))
	//responseExcepted = ""
	//require.Equal(t, responseExcepted, string(responseBody))
}

func Test_HttpCrudHandler(t *testing.T) {
	body := bytes.NewBufferString(`{
		"id": "a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b0",
		"title": "Test title",
		"created": "2022-04-01T18:43:25.391Z",
		"description": "test description",
		"userId": "1528371b-229c-4370-839a-0571d969902a"
	}`)

	request := httptest.NewRequest("POST", "/posts/create", body)
	w := httptest.NewRecorder()

	testHandlers := NewRouter(createApp(t))
	testHandlers.ServeHTTP(w, request)

	response := w.Result()
	responseBody, _ := ioutil.ReadAll(response.Body)
	responseExcepted := `{"id":"a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b0","title":"Test title","created":"2020-10-20 12:30:00","description":"test description","userId":"1528371b-229c-4370-839a-0571d969902a"}` //nolint:lll
	require.Equal(t, responseExcepted, string(responseBody))

	//body = bytes.NewBufferString(`{
	//	"title": "Test title 2",
	//	"created": "2022-04-01T18:43:25.391Z",
	//	"description": "test description 2",
	//	"userId": "1528371b-229c-4370-839a-0571d969902a"
	//}`)
	//
	//request = httptest.NewRequest("PUT", "/posts/update/a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b0", body)
	//w = httptest.NewRecorder()
	//
	//testHandlers.ServeHTTP(w, request)
	//
	//response = w.Result()
	//responseBody, _ = ioutil.ReadAll(response.Body)
	//responseExcepted = `{"id":"a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b0","title":"Test title 2","created":"2020-10-20 12:30:00","description":"test description 2","userId":"1528371b-229c-4370-839a-0571d969902a"}` //nolint:lll
	//require.Equal(t, responseExcepted, string(responseBody))
	//
	//requestListAll := httptest.NewRequest("GET", "/posts", body)
	//w = httptest.NewRecorder()
	//
	//testHandlers.ServeHTTP(w, requestListAll)
	//
	//response = w.Result()
	//responseBody, _ = ioutil.ReadAll(response.Body)
	//fmt.Println(string(responseBody))
	//responseExcepted = `[{"id":"a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b0","title":"Test title 2","created":"2020-10-20T12:30:00Z","description":"test description 2","userId":"1528371b-229c-4370-839a-0571d969902a"}]` // nolint:lll
	//require.Equal(t, responseExcepted, string(responseBody))
	//
	//request = httptest.NewRequest("DELETE", "/posts/delete/a17b3f01-fbd7-40e5-8d8e-9b4cf1ef21b0", body)
	//w = httptest.NewRecorder()
	//
	//testHandlers.ServeHTTP(w, request)
	//
	//response = w.Result()
	//responseBody, _ = ioutil.ReadAll(response.Body)
	//responseExcepted = ""
	//require.Equal(t, responseExcepted, string(responseBody))
}

func createApp(t *testing.T) *app.App {
	configContent, _ := os.ReadFile(configFile)

	var config struct {
		Storage struct {
			URL string
		}
	}

	err := yaml.Unmarshal(configContent, &config)
	if err != nil {
		t.Fatal("Failed to unmarshal config", err)
	}

	t.Helper()
	logFile, err := os.CreateTemp("", "log")
	if err != nil {
		t.Errorf("failed to open test log file: %s", err)
	}

	logg, _ := logger.New(internalconfig.LoggerConf{
		Level:    internalconfig.Debug,
		Filename: logFile.Name(),
	})

	if err != nil {
		t.Errorf("failed to open test log file: %s", err)
	}

	ctx := context.Background()
	inSQLStorage := sqlstorage.New(ctx, config.Storage.URL)
	if err := inSQLStorage.Connect(ctx); err != nil {
		t.Fatal("Failed to connect to DB server", err)
	}

	return app.New(logg, inSQLStorage)
}
