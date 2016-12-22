package request

import "github.com/cocotyty/alipay/api/response"

type AlipayCommerceEducateStudentinfoShareRequest struct {

}

func (r *AlipayCommerceEducateStudentinfoShareRequest) GetApiMethod() string {
	return "alipay.commerce.educate.studentinfo.share"
}

func (r *AlipayCommerceEducateStudentinfoShareRequest) GetTextParams() map[string]string {
	params := make(map[string]string)
	return params
}

func (r *AlipayCommerceEducateStudentinfoShareRequest) GetResponse() response.AlipayResponse {
	resp := new(response.AlipayCommerceEducateStudentinfoShareResponse)
	resp.Name = "AlipayCommerceEducateStudentinfoShareResponse"
	return resp
}

func (r *AlipayCommerceEducateStudentinfoShareRequest) GetApiVersion() string {
	return "1.0"
}
