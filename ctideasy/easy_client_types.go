package ctideasy

import (
	"encoding/json"
	"github.com/hilaoyu/go-utils/utilEnc"
	"github.com/hilaoyu/go-utils/utilHttp"
)

const (
	EASYCTID_API_BASE = "http://api.easyctid.cn"
)

type CtidEasyClient struct {
	httpClient    *utilHttp.HttpClient
	orgCode       string
	signEncryptor *utilEnc.RsaEncryptor
	dataEncryptor *utilEnc.RsaEncryptor
}

type EasyAuthApplyRequestBizPackageBizData struct {
	AuthMode string `json:"authMode"`
}
type EasyAuthApplyRequestBizPackage struct {
	OrgCode string                                 `json:"orgCode,omitempty"`
	BizType int                                    `json:"bizType,omitempty"`
	BizData *EasyAuthApplyRequestBizPackageBizData `json:"bizData,omitempty"`
}

type EasyAuthApplyResponseBizPackageBizData struct {
	Bsn          string `json:"bsn"`
	RandomNumber string `json:"randomNumber"`
}
type EasyAuthApplyResponseBizPackage struct {
	ResultCode string                                  `json:"resultCode,omitempty"`
	ResultDesc string                                  `json:"resultDesc,omitempty"`
	BizData    *EasyAuthApplyResponseBizPackageBizData `json:"bizData,omitempty"`
}
type EasyAuthApplyResponse struct {
	BizPackage *EasyAuthApplyResponseBizPackage `json:"bizPackage,omitempty"`
	Sign       string                           `json:"sign,omitempty"`
}

func (r *EasyAuthApplyResponse) UnmarshalJSON(b []byte) error {
	tmp := map[string]string{}
	err := json.Unmarshal(b, &tmp)
	bizPackage, _ := tmp["bizPackage"]
	if "" != bizPackage {
		bizTemp := &EasyAuthApplyResponseBizPackage{}
		err = json.Unmarshal([]byte(bizPackage), bizTemp)
		r.BizPackage = bizTemp
	}

	return err
}

type EasyVerificationRequestAuthApplyRetainData struct {
	Name         string `json:"name,omitempty"`
	IdNo         string `json:"idNo,omitempty"`
	IdIssueDate  string `json:"idIssueDate,omitempty"`
	IdExpireDate string `json:"idExpireDate,omitempty"`
	Location     string `json:"location,omitempty"`
	PackageName  string `json:"packageName,omitempty"`
}

type EasyVerificationRequestBizPackageBizData struct {
	AuthMode            string `json:"authMode"`
	PhotoData           string `json:"photoData,omitempty"`
	AuthApplyRetainData string `json:"authApplyRetainData"`
	PhotoData2          string `json:"photoData2,omitempty"`
}
type EasyVerificationRequestBizPackage struct {
	OrgCode string                                    `json:"orgCode,omitempty"`
	Bsn     string                                    `json:"bsn"`
	BizType int                                       `json:"bizType,omitempty"`
	BizData *EasyVerificationRequestBizPackageBizData `json:"bizData,omitempty"`
}

/*func (rd *EasyVerificationRequestAuthApplyRetainData) MarshalJSON() (jsonStr []byte, err error) {

	jsonStr = []byte("\"\"")
	if nil == rd {
		return jsonStr, err
	}
	jsonStr, err = json.Marshal()
	if nil != err {
		return jsonStr, err
	}
	return []byte(fmt.Sprintf("\"%s\"", jsonStr)), nil
}*/

type EasyVerificationResponseBizPackageBizData struct {
	Bid               string `json:"bid,omitempty"`
	PhotoCompareScore string `json:"photoCompareScore,omitempty"`
	CertificationData string `json:"certificationData,omitempty"`
}

type EasyVerificationResponseBizPackage struct {
	ResultCode string                                     `json:"resultCode,omitempty"`
	ResultDesc string                                     `json:"resultDesc,omitempty"`
	BizData    *EasyVerificationResponseBizPackageBizData `json:"bizData,omitempty"`
}
type EasyVerificationResponse struct {
	BizPackage *EasyVerificationResponseBizPackage `json:"bizPackage,omitempty"`
	Sign       string                              `json:"sign,omitempty"`
}

func (r *EasyVerificationResponse) UnmarshalJSON(b []byte) error {
	tmp := map[string]string{}
	err := json.Unmarshal(b, &tmp)
	bizPackage, _ := tmp["bizPackage"]
	if "" != bizPackage {
		bizTemp := &EasyVerificationResponseBizPackage{}
		err = json.Unmarshal([]byte(bizPackage), bizTemp)
		r.BizPackage = bizTemp
	}

	return err
}
