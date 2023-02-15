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

		// test parsed strings in ImageInfo type
		// order is [want, got]
		strAssertions := [][2]string{
			{"http://iiif.io/api/image/2/context.json", imageInfo.Context},
			{"https://image1.dlib.nyu.edu:8183/iiif/2/photo%2FMSS208_ref5830%2FMSS208_ref5830_n000001_d.jp2", imageInfo.ID},
			{"http://iiif.io/api/image", imageInfo.Protocol},
		}

		for _, strAssertion := range strAssertions {
			if strAssertion[0] != strAssertion[1] {
				t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", strAssertion[0], strAssertion[1])
			}
		}

		// test image sizes
		// order is [want, got]
		uint32Assertions := [][2]uint32{
			{2811, imageInfo.Width},
			{2780, imageInfo.Height},
			{88, imageInfo.Sizes[0].Width},
			{87, imageInfo.Sizes[0].Height},
			{176, imageInfo.Sizes[1].Width},
			{174, imageInfo.Sizes[1].Height},
			{351, imageInfo.Sizes[2].Width},
			{348, imageInfo.Sizes[2].Height},
			{703, imageInfo.Sizes[3].Width},
			{695, imageInfo.Sizes[3].Height},
			{1406, imageInfo.Sizes[4].Width},
			{1390, imageInfo.Sizes[4].Height},
			{2811, imageInfo.Sizes[5].Width},
			{2780, imageInfo.Sizes[5].Height},
		}
		for _, uint32Assertion := range uint32Assertions {
			if uint32Assertion[0] != uint32Assertion[1] {
				t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", uint32Assertion[0], uint32Assertion[1])
			}
		}
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
