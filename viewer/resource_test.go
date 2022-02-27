package viewer

import (
        // "net/http/httptest"
        // "sort"
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
		want := uint(97)
		got := resource.Metadata.PageCount.Value
		if want != got {
			t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
		}
	})
}

	
	


// func TestResourceGetByNOID(t *testing.T) {

//         mux := setupMux("/viewer/api/v1/noid/xgxd28gq", "testdata/image-set.json")
//         ts := httptest.NewServer(mux)
//         defer ts.Close()

//         setupTestServerClient(ts)

//         t.Run("ResourceGET", func(t *testing.T) {
//                 resource, err := ResourceGetByNOID("xgxd28gq")
//                 if err != nil {
//                         t.Errorf("Unexpected error: %s", err)
//                 }

//                 want := "2020-11-27T01:05:44Z"
//                 got := resource.TimeStamp
//                 if want != got {
//                         t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
//                 }

//                 want = "c44e95e9-5cca-4c26-8e52-12773334dc95"
//                 got = resource.Info.ID
//                 if want != got {
//                         t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
//                 }

//                 want = "publication"
//                 got = resource.Info.Type
//                 if want != got {
//                         t.Errorf("Mismatch: want: \"%v\", got: \"%v\"", want, got)
//                 }

// 	})
// }

