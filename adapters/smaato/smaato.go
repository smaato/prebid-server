package smaato

import (
	"github.com/mxmCherry/openrtb"
	"github.com/prebid/prebid-server/adapters"
)

// Adapter provides an adapter for prebid-server
type Adapter struct {
	http         *adapters.HTTPAdapter
	URI          string
	XAPIUsername string
	XAPIPassword string
}

func NewSmaatoBidder() *Adapter {
	return &Adapter{}
}

// MakeRequests makes the HTTP requests which should be made to fetch bids.
//
// Bidder implementations can assume that the incoming BidRequest has:
//
//   1. Only {Imp.Type, Platform} combinations which are valid, as defined by the static/bidder-info.{bidder}.yaml file.
//   2. Imp.Ext of the form {"bidder": params}, where "params" has been validated against the static/bidder-params/{bidder}.json JSON Schema.
//
// nil return values are acceptable, but nil elements *inside* those slices are not.
//
// The errors should contain a list of errors which explain why this bidder's bids will be
// "subpar" in some way. For example: the request contained ad types which this bidder doesn't support.
//
// If the error is caused by bad user input, return an errortypes.BadInput.
func (a *Adapter) MakeRequests(request *openrtb.BidRequest, reqInfo *adapters.ExtraRequestInfo) ([]*adapters.RequestData, []error) {
	return nil, nil
}

// MakeBids unpacks the server's response into Bids.
//
// The bids can be nil (for no bids), but should not contain nil elements.
//
// The errors should contain a list of errors which explain why this bidder's bids will be
// "subpar" in some way. For example: the server response didn't have the expected format.
//
// If the error was caused by bad user input, return a errortypes.BadInput.
// If the error was caused by a bad server response, return a errortypes.BadServerResponse
func (a *Adapter) MakeBids(internalRequest *openrtb.BidRequest, externalRequest *adapters.RequestData, response *adapters.ResponseData) (*adapters.BidderResponse, []error) {
	return nil, nil
}
