package bitfinex

// Package GoBitfinex provides structs and functions for accessing
// bitfinex.com api version 1.0
//
// Usage:
//   import "github.com/wanth1997/GoBitfinex"
//
// Create new client:
//   api := bitfinex.NewClient()
//
// For access methods that requires authentication use the next code:
//   api := bitfinex.NewClient().Auth(key, secret)
//
// Get all pairs
//   api.Pairs.V1All()
//
// Get account info
//   api.Account.V1Info()
//
// See examples dir for more info.
