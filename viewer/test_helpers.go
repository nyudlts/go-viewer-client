package viewer

import (
	"net/http"
	"net/http/httptest"
	"os"
)

func setupMux(apiPath string, filePath string) (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc(apiPath, func(w http.ResponseWriter, _ *http.Request) {
		data, _ := os.ReadFile(filePath)
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

/* keeping this code in case it's needed in the future
func setupLocalhostClient() {
	c := new(Config)
	c.BaseURL = "http://localhost:3000"
	c.User = "foo"
	c.Password = "bar"

	ConfigureClient(c)
}
*/
