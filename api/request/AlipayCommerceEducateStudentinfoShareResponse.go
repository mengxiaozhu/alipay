package request

import "github.com/cocotyty/alipay/api/response"

type AlipaySocialBaseMcommentStudentQueryRequest struct {

}

func (r *AlipaySocialBaseMcommentStudentQueryRequest) GetApiMethod() string {
	return "alipay.social.base.mcomment.student.query"
}

func (r *AlipaySocialBaseMcommentStudentQueryRequest) GetTextParams() map[string]string {
	params := make(map[string]string)
	return params
}

func (r *AlipaySocialBaseMcommentStudentQueryRequest) GetResponse() response.AlipayResponse {
	resp := new(response.AlipaySocialBaseMcommentStudentQueryResponse)
	resp.Name = "AlipaySocialBaseMcommentStudentQueryResponse"
	return resp
}

func (r *AlipaySocialBaseMcommentStudentQueryRequest) GetApiVersion() string {
	return "1.0"
}
