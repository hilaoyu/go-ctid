package ctideasy

import (
	"encoding/json"
	"fmt"
	"github.com/hilaoyu/go-utils/utilEnc"
	"github.com/hilaoyu/go-utils/utilHttp"
	"github.com/hilaoyu/go-utils/utilRandom"
	"github.com/hilaoyu/go-utils/utils"
	"time"
)

func NewCtidEasyClient(orgCode string) (client *CtidEasyClient) {
	client = &CtidEasyClient{
		httpClient:    utilHttp.NewHttpClient(EASYCTID_API_BASE),
		orgCode:       orgCode,
		signEncryptor: utilEnc.NewRsaEncryptor(),
		dataEncryptor: utilEnc.NewRsaEncryptor(),
	}
	return
}

func (cc *CtidEasyClient) SetSignPrivateKey(privateKey string) (err error) {
	if "" != privateKey {
		_, err = cc.signEncryptor.SetPrivateKey([]byte(privateKey))
	}
	return
}

func (cc *CtidEasyClient) SetSignPublicKey(publicKey string) (err error) {
	if "" != publicKey {
		_, err = cc.signEncryptor.SetPublicKey([]byte(publicKey))
	}
	return
}

func (cc *CtidEasyClient) SetDataEncodePublicKey(publicKey string) (err error) {
	if "" != publicKey {
		_, err = cc.dataEncryptor.SetPublicKey([]byte(publicKey))
	}
	return
}

func (cc *CtidEasyClient) SetTimeout(timeout time.Duration) *CtidEasyClient {
	cc.httpClient.SetTimeout(timeout)
	return cc
}
func (cc *CtidEasyClient) SignBizPackageToRequestData(bizPackage interface{}) (data map[string]string, err error) {
	bizPackageJson, err := json.Marshal(bizPackage)
	if nil != err {
		return
	}
	sign, err := cc.signEncryptor.RsaPrivateKeySign(bizPackageJson)
	if nil != err {
		return
	}

	data = map[string]string{
		"bizPackage": string(bizPackageJson),
		"sign":       utils.Base64EncodeUrlSafe(sign),
	}
	return
}

func (cc *CtidEasyClient) AuthApply(authMode string) (applyBizData *EasyAuthApplyResponseBizPackageBizData, err error) {
	bizPackage := &EasyAuthApplyRequestBizPackage{
		OrgCode: cc.orgCode,
		BizType: 6000,
		BizData: &EasyAuthApplyRequestBizPackageBizData{AuthMode: authMode},
	}

	reqData, err := cc.SignBizPackageToRequestData(bizPackage)
	if nil != err {
		return
	}

	result := &EasyAuthApplyResponse{}
	err = cc.httpClient.WithJsonData(reqData).RequestJson(&result, "post", "/v1/apply", map[string]string{})
	if nil != err {
		return
	}

	if "0" != result.BizPackage.ResultCode {
		err = fmt.Errorf("接口返回错误, code: %s ,msg: %s ", result.BizPackage.ResultCode, result.BizPackage.ResultDesc)
	}
	applyBizData = result.BizPackage.BizData
	return
}
func (cc *CtidEasyClient) Verification(bizData *EasyVerificationRequestBizPackageBizData, retainData *EasyVerificationRequestAuthApplyRetainData) (verifyBizData *EasyVerificationResponseBizPackageBizData, err error) {
	authApplyRespData, err := cc.AuthApply(bizData.AuthMode)
	if nil != err {
		err = fmt.Errorf("申请出错:%v", err)
		return
	}

	retainDataEn, err := cc.EncodeRetainData(retainData)
	if nil != err {
		err = fmt.Errorf("数据加密出错:%v", err)
		return
	}
	bizData.AuthApplyRetainData = retainDataEn
	bizPackage := &EasyVerificationRequestBizPackage{
		OrgCode: cc.orgCode,
		BizType: 6000,
		Bsn:     authApplyRespData.Bsn,
		BizData: bizData,
	}

	reqData, err := cc.SignBizPackageToRequestData(bizPackage)
	if nil != err {
		return
	}

	result := &EasyVerificationResponse{}
	err = cc.httpClient.WithJsonData(reqData).RequestJson(&result, "post", "/ctid/v1/verification", map[string]string{})
	if nil != err {
		return
	}

	if "0" != result.BizPackage.ResultCode {
		err = fmt.Errorf("接口返回错误, code: %s ,msg: %s ", result.BizPackage.ResultCode, result.BizPackage.ResultDesc)
	}
	verifyBizData = result.BizPackage.BizData
	return
}
func (cc *CtidEasyClient) EncodeRetainData(retainData *EasyVerificationRequestAuthApplyRetainData) (str string, err error) {
	jsonByte, err := json.Marshal(retainData)
	if nil != err {
		return
	}

	aesKey := utilRandom.RandString(16)
	enData, iv, err := utilEnc.AesCBCEncrypt(jsonByte, []byte(aesKey))
	if nil != err {
		return
	}
	str2, err := cc.dataEncryptor.RsaPublicKeyEncrypt(append([]byte(aesKey), iv...))
	if nil != err {
		return
	}
	str = utils.Base64EncodeUrlSafe(append(str2, enData...))
	return
}
