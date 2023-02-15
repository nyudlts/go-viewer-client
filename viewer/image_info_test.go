package viewer

import (
	"net/http/httptest"
	"testing"
)

//  wget https://sites.dlib.nyu.edu/viewer/api/image/photos/MSS208_ref5830/1/info.json -O image-set-first-image-info.json
//  wget https://sites.dlib.nyu.edu/viewer/api/image/photos/tam439_ref154/1/info.json -O image-info.json

func TestImageInfoGetByNOIDImageSet(t *testing.T) {

	aPathsfPaths := [][2]string{
		{"/viewer/api/v1/noid/xgxd28gq", "testdata/image-set.json"},
		{"/viewer/api/image/photos/MSS208_ref5830/1/info.json", "testdata/image-set-first-image-info.json"},
	}

	mux := setupMuxMultiPath(aPathsfPaths)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	setupTestServerClient(ts)

	t.Run("Get Image Set First Image Info", func(t *testing.T) {
		resource, err := ResourceGetByNOID("xgxd28gq")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		imageInfo, err := ImageInfoGetByURL(resource.IIIF.Image.Items[0])
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		want_str := "http://iiif.io/api/image/2/context.json"
		got_str := imageInfo.Context
		if want_str != got_str {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want_str, got_str)
		}
		/*
			if want_int != got_int {
				t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want_int, got_int)
			}

			want_string := "https://sites.dlib.nyu.edu/viewer/api/image/photos/MSS208_ref5830/1/info.json"
			got_string := resource.IIIF.Image.Items[0]
			if want_string != got_string {
				t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want_string, got_string)
			}
		*/
	})
}

func TestImageInfoGetByNOIDImage(t *testing.T) {

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
