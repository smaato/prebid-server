package smaato

import (
	"testing"
)

func TestRenderAdMarkup(t *testing.T) {
	type args struct {
		adType             adMarkupType
		adapterResponseAdm string
	}
	tests := []struct {
		testName string
		args     args
		result   string
	}{
		{"imageTest", args{"Img", "{\"image\":{\"img\":{\"url\":\"//prebid-test.smaatolabs.net/img/320x50.jpg\"," +
			"\"w\":350,\"h\":50,\"ctaurl\":\"//prebid-test.smaatolabs.net/track/ctaurl/1\"},\"impressiontrackers\":[\"//prebid-test.smaatolabs.net/track/imp/1\",\"//prebid-test.smaatolabs.net/track/imp/2\"],\"clicktrackers\":[\"//prebid-test.smaatolabs.net/track/click/1\",\"//prebid-test.smaatolabs.net/track/click/2\"]}}"},
			`<div onclick="fetch(decodeURIComponent('%2F%2Fprebid-test.smaatolabs.net%2Ftrack%2Fclick%2F1'), {cache: 'no-cache'});fetch(decodeURIComponent('%2F%2Fprebid-test.smaatolabs.net%2Ftrack%2Fclick%2F2'), {cache: 'no-cache'});"><a href="//prebid-test.smaatolabs.net/track/ctaurl/1"><img src="//prebid-test.smaatolabs.net/img/320x50.jpg" width="350" height="50"/></a><img src="//prebid-test.smaatolabs.net/track/imp/1" alt="" width="0" height="0"/><img src="//prebid-test.smaatolabs.net/track/imp/2" alt="" width="0" height="0"/></div>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := renderAdMarkup(tt.args.adType, tt.args.adapterResponseAdm)
			if err != nil {
				t.Errorf("error rendering ad markup: %v", err)
			}
			if got != tt.result {
				t.Errorf("renderAdMarkup() got = %v, result %v", got, tt.result)
			}
		})
	}
}