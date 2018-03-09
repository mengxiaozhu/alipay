package response

import (
	"fmt"
	"log"
	"strings"
)

// AlipayResponse response接口
type AlipayResponse interface {
	// 判断是否成功
	IsSuccess() bool
	// 接口名称
	ToStr() string
	// 保留body
	SetBody(body string)
	// code
	GetCode() string
	// subCode
	GetSubCode() string
	// msg
	GetMsg() string
}

type BaseResponse struct {
	Code    interface{} `json:"code"`
	Msg     string      `json:"msg"`
	SubCode string      `json:"sub_code"`
	SubMsg  string      `json:"sub_msg"`
	Name    string      `json:"name"`
	Body    string      `json:"body"`
}

func (r *BaseResponse) IsSuccess() bool {
	// sub_code如果为空，表明执行成功
	return strings.EqualFold("", r.SubCode)
}

// ToStr 输出类名，用于动态获取支付宝返回值key
func (r *BaseResponse) ToStr() string {
	return r.Name
}

// SetBody 保存请求结果
func (r *BaseResponse) SetBody(body string) {
	r.Body = body
}

// GetCode
func (r *BaseResponse) GetCode() string {
	// code may be string
	str, ok := r.Code.(string)
	if ok {
		return str
	}
	// code may be float64
	integer, ok := r.Code.(float64)
	if ok {
		return fmt.Sprintf("%0.f", integer)
	}
	// both not
	log.Printf("alipay response code type:%s", r.Code)
	return ""
}

// GetSubCode
func (r *BaseResponse) GetSubCode() string {
	return r.SubCode
}

// GetMsg
func (r *BaseResponse) GetMsg() string {
	return r.Msg
}

// AlipayMobilePublicMessageCustomSendResponse
// 与AlipayMobilePublicMessageCustomSendRequest关联
type AlipayMobilePublicMessageCustomSendResponse struct {
	BaseResponse
}

// AlipaySystemOauthTokenResponse
// refer AlipaySystemOauthTokenRequest
type AlipaySystemOauthTokenResponse struct {
	BaseResponse
	AccessToken  string `json:"access_token"`
	AlipayUserId string `json:"user_id"`
	ExpiresIn    int64  `json:"expires_in"`
	ReExpiresIn  int64  `json:"re_expires_in"`
	RefreshToken string `json:"refresh_token"`
	Sign         string `json:"sign"`
}

// AlipayPassTplContentAddResponse
// refer AlipayPassTplContentAddRequest
type AlipayPassTplContentAddResponse struct {
	BaseResponse
	BizResult string `json:"biz_result"`
	ErrorCode string `json:"error_code"`
	Success   string `json:"success"` //T-成功；F-失败
}

// AlipayPassSyncUpdateResponse
// refer AlipayPassSyncUpdateRequest
type AlipayPassSyncUpdateResponse struct {
	BaseResponse
	BizResult string `json:"biz_result"`
	ErrorCode string `json:"error_code"`
	Success   bool   `json:"success"` //T-成功；F-失败
}

// AlipayMobilePublicGisGetResponse
// refer AlipayMobilePublicGisGetRequest
type AlipayMobilePublicGisGetResponse struct {
	BaseResponse
	Accuracy  string `json:"accuracy"`
	City      string `json:"city"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Province  string `json:"province"`
}

// AlipayPassTplContentUpdateResponse
// refer AlipayPassTplContentUpdateRequest
type AlipayPassTplContentUpdateResponse struct {
	BaseResponse
	Result    string `json:"result"`
	ErrorCode string `json:"error_code"`
	Success   bool   `json:"success"`
}
type AlipayUserUserinfoShareResponse struct {
	BaseResponse
	Address               string `db:"address" json:"address"`
	AddressCode           string `db:"address_code" json:"address_code"`
	AlipayUserID          string `db:"alipay_user_id" json:"alipay_user_id"`
	Area                  string `db:"area" json:"area"`
	Avatar                string `db:"avatar" json:"avatar"`
	BalanceFreezeType     string `db:"balance_freeze_type" json:"balance_freeze_type"`
	Birthday              string `db:"birthday" json:"birthday"`
	CertNo                string `db:"cert_no" json:"cert_no"`
	CertTypeValue         string `db:"cert_type_value" json:"cert_type_value"`
	City                  string `db:"city" json:"city"`
	DefaultDeliverAddress string `db:"default_deliver_address" json:"default_deliver_address"`
	DeliverAddressList    []struct {
		Address               string `db:"address" json:"address"`
		AddressCode           string `db:"address_code" json:"address_code"`
		DefaultDeliverAddress string `db:"default_deliver_address" json:"default_deliver_address"`
		DeliverArea           string `db:"deliver_area" json:"deliver_area"`
		DeliverCity           string `db:"deliver_city" json:"deliver_city"`
		DeliverFullname       string `db:"deliver_fullname" json:"deliver_fullname"`
		DeliverMobile         string `db:"deliver_mobile" json:"deliver_mobile"`
		DeliverPhone          string `db:"deliver_phone" json:"deliver_phone"`
		DeliverProvince       string `db:"deliver_province" json:"deliver_province"`
		Zip                   string `db:"zip" json:"zip"`
	} `db:"deliver_address_list" json:"deliver_address_list"`
	DeliverArea        string `db:"deliver_area" json:"deliver_area"`
	DeliverCity        string `db:"deliver_city" json:"deliver_city"`
	DeliverFullname    string `db:"deliver_fullname" json:"deliver_fullname"`
	DeliverMobile      string `db:"deliver_mobile" json:"deliver_mobile"`
	DeliverPhone       string `db:"deliver_phone" json:"deliver_phone"`
	DeliverProvince    string `db:"deliver_province" json:"deliver_province"`
	Email              string `db:"email" json:"email"`
	FamilyName         string `db:"family_name" json:"family_name"`
	FirmName           string `db:"firm_name" json:"firm_name"`
	Gender             string `db:"gender" json:"gender"`
	IsBalanceFrozen    string `db:"is_balance_frozen" json:"is_balance_frozen"`
	IsBankAuth         string `db:"is_bank_auth" json:"is_bank_auth"`
	IsCertified        string `db:"is_certified" json:"is_certified"`
	IsCertifyGradeA    string `db:"is_certify_grade_a" json:"is_certify_grade_a"`
	IsIDAuth           string `db:"is_id_auth" json:"is_id_auth"`
	IsLicenceAuth      string `db:"is_licence_auth" json:"is_licence_auth"`
	IsMobileAuth       string `db:"is_mobile_auth" json:"is_mobile_auth"`
	IsStudentCertified string `db:"is_student_certified" json:"is_student_certified"`
	Mobile             string `db:"mobile" json:"mobile"`
	NickName           string `db:"nick_name" json:"nick_name"`
	Phone              string `db:"phone" json:"phone"`
	Province           string `db:"province" json:"province"`
	RealName           string `db:"real_name" json:"real_name"`
	ReducedBirthday    string `db:"reduced_birthday" json:"reduced_birthday"`
	UserID             string `db:"user_id" json:"user_id"`
	UserStatus         string `db:"user_status" json:"user_status"`
	UserTypeValue      string `db:"user_type_value" json:"user_type_value"`
	Zip                string `db:"zip" json:"zip"`
}
type AlipayCommerceEducateStudentinfoShareResponse struct {
	BaseResponse
	StudentInfoShareResult struct {
		UserID       string `db:"user_id" json:"user_id"`
		BizType      string `db:"biz_type" json:"biz_type"`
		StudentInfos []struct {
			CityNo        string `db:"city_no" json:"city_no"`
			CollegeNo     string `db:"college_no" json:"college_no"`
			CollegeName   string `db:"college_name" json:"college_name"`
			Degree        string `db:"degree" json:"degree"`
			Departments   string `db:"departments" json:"departments"`
			Major         string `db:"major" json:"major"`
			ClassName     string `db:"class_name" json:"class_name"`
			StudentID     string `db:"student_id" json:"student_id"`
			GmtGraduation string `db:"gmt_graduation" json:"gmt_graduation"`
			GmtEnrollment string `db:"gmt_enrollment" json:"gmt_enrollment"`
		} `db:"student_infos" json:"student_infos"`
	} `db:"student_info_share_result" json:"student_info_share_result"`
}

type AlipaySocialBaseMcommentStudentQueryResponse struct {
	BaseResponse
	UserID         string `db:"user_id" json:"user_id"`
	CampusName     string `db:"campus_name" json:"campus_name"`
	CampusCode     string `db:"campus_code" json:"campus_code"`
	Degree         string `db:"degree" json:"degree"`
	EnrollmentTime string `db:"enrollment_time" json:"enrollment_time"`
	GraduationTime string `db:"graduation_time" json:"graduation_time"`
	StatusEnum     int    `db:"status_enum" json:"status_enum"`
	ProvinceCode   string `db:"province_code" json:"province_code"`
	ProvinceName   string `db:"province_name" json:"province_name"`
}
