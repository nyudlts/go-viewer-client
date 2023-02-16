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

func handleFuncBuilderReturnFileContents(path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		data, _ := os.ReadFile(path)
		w.Write(data)
	}
}

// This function takes an array of 2-string arrays
// where:
//
//	the first element of the 2-string array is the API path
//	the second element of the 2-string array is the path of the file to serve at the API path
func setupMuxMultiPath(aPathsfPaths [][2]string) (mux *http.ServeMux) {
	mux = http.NewServeMux()

	for _, apfp := range aPathsfPaths {
		mux.HandleFunc(apfp[0], handleFuncBuilderReturnFileContents(apfp[1]))
	}
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
