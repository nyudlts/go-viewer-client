package viewer

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func setupMux(apiPath string, filePath string) (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc(apiPath, func(w http.ResponseWriter, _ *http.Request) {
		data, _ := ioutil.ReadFile(filePath)
		w.Write(data)
	})

	return mux
}

func setupTestServerClient(ts *httptest.Server) {
	c := new(Config)
	c.BaseURL = ts.URL
	c.User = "foo"
	c.Password = "bar"

	ConfigureClient(c)
}

func setupLocalhostClient() {
	c := new(Config)
	c.BaseURL = "http://localhost:3000"
	c.User = "foo"
	c.Password = "bar"

	ConfigureClient(c)
}

