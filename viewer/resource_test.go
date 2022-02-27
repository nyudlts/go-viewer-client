package viewer

import (
	"net/http/httptest"
        "testing"
)


//  wget https://sites.dlib.nyu.edu/viewer/api/v1/noid/xgxd28gq  (image-set.json)
//  wget https://sites.dlib.nyu.edu/viewer/api/v1/noid/j3tx985c  (image.json)

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

                want := 32
                got := resource.Metadata.PageCount.Value
                if want != got {
                        t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
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

                want := 1
                got := resource.Metadata.PageCount.Value
                if want != got {
                        t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
                }
	})
}

