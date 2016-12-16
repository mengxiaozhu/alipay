package request

import "github.com/cocotyty/alipay/api/response"

var _ AlipayRequest = &AlipayUserUserinfoShare{}

type AlipayUserUserinfoShare struct {

}

func (r *AlipayUserUserinfoShare) GetApiMethod() string {
	return "alipay.user.userinfo.share"
}

func (r *AlipayUserUserinfoShare) GetTextParams() map[string]string {
	params := make(map[string]string)
	return params
}

func (r *AlipayUserUserinfoShare) GetResponse() response.AlipayResponse {
	resp := new(response.AlipayUserUserinfoShareResponse)
	resp.Name = "AlipayUserUserinfoShareResponse"
	return resp
}

func (r *AlipayUserUserinfoShare) GetApiVersion() string {
	return "1.0"
}
