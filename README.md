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
}
```
