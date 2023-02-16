## go-viewer-client

This repo contains a Golang client for interacting with the DLTS  
Viewer API.  


**Please note:**   
This is an absolute minimal implementation to support  
work on the finding aids redesign project.  

#### Example:

```
package main

import (
	"github.com/nyudlts/go-viewer-client/viewer"
	"fmt"
)

func main() {
	var c viewer.Config
	c.BaseURL = "https://sites.dlib.nyu.edu"
	viewer.ConfigureClient(&c)

	noid := "tb2rbsmk"
	
	resource, err := viewer.ResourceGetByNOID(noid)
	if err != nil {
		fmt.Errorf("%s\n", err)
	}

	fmt.Printf("Count: %d\n", resource.Metadata.PageCount.Value)

	// assert that we have an image-info URL to access
	if len(resource.IIIF.Image.Items) == 0 {
		return fmt.Errorf("no item URLs found for: %s", noid)
	}

	imageInfo, err := viewer.ImageInfoGetByURL(resource.IIIF.Image.Items[0])
	if err != nil {
		return err
	}

	...
	// find the closest pre-calculated image to the targetWidth
	result := imageInfo.Sizes[0]
	for _, size := range imageInfo.sizes {
		if abs(int64(size.Width)-int64(targetWidth)) < abs(int64(result.Width)-int64(targetWidth)) {
			result = size
		}
	}
	fmt.Printf("Target Width: %d Closest Image: (Width: %d Height: %d)\n", targetWidth, result.Width, result.Height)
}
```
