package viewer

import (
	client "github.com/nyudlts/go-rsbe-client/rsbe"
)

const Version = "0.2.0-20230213T10_17"

// Map client package types, functions into this package
type Config = client.Config

var ConfigureClient = client.ConfigureClient
var GetBody = client.GetBody
