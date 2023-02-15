package viewer

import (
	"net/http/httptest"
	"testing"
)

//  wget https://sites.dlib.nyu.edu/viewer/api/v1/noid/xgxd28gq  -O image-set.json
//  wget https://sites.dlib.nyu.edu/viewer/api/v1/noid/j3tx985c  -O image.json

func TestResource(t *testing.T) {
	var resource = Resource{
		Metadata: Metadata{
			PageCount: PageCount{
				Label: "Waffle Fries",
				Value: 97,
			},
		},
	}

	t.Run("Resource Metadata PageCount Label", func(t *testing.T) {
		want := "Waffle Fries"
		got := resource.Metadata.PageCount.Label
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})

	t.Run("Resource Metadata PageCount Value", func(t *testing.T) {
		want := 97
		got := resource.Metadata.PageCount.Value
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})
}

func TestResourceGetByNOIDImageSet(t *testing.T) {

	mux := setupMux("/viewer/api/v1/noid/xgxd28gq", "testdata/image-set.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("Get Image Set data by NOID", func(t *testing.T) {
		resource, err := ResourceGetByNOID("xgxd28gq")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		want_int := 32
		got_int := resource.Metadata.PageCount.Value
		if want_int != got_int {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want_int, got_int)
		}

		want_string := "https://sites.dlib.nyu.edu/viewer/api/image/photos/MSS208_ref5830/1/info.json"
		got_string := resource.IIIF.Image.Items[0]
		if want_string != got_string {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want_string, got_string)
		}
	})
}

func TestResourceGetByNOIDImage(t *testing.T) {

	mux := setupMux("/viewer/api/v1/noid/j3tx985c", "testdata/image.json")
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("Get Image data by NOID", func(t *testing.T) {
		resource, err := ResourceGetByNOID("j3tx985c")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		want_int := 1
		got_int := resource.Metadata.PageCount.Value
		if want_int != got_int {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want_int, got_int)
		}

		want_string := "https://sites.dlib.nyu.edu/viewer/api/image/photos/tam439_ref154/1/info.json"
		got_string := resource.IIIF.Image.Items[0]
		if want_string != got_string {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want_string, got_string)
		}
	})
}
