go get github.com/hilaoyu/go-ctid

c := ctideasy.NewCtidEasyClient("<orgCode>")
err = c.SetSignPrivateKey(priKey)
fmt.Println("SetPrivateKey err", err)
err = c.SetSignPublicKey(pubKey)
fmt.Println("SetPublicKey err", err)
err = c.SetDataEncodePublicKey(dataPubKey)
fmt.Println("SetDataEncodePublicKey err", err)

returnData, err := c.Verification(&ctideasy.EasyVerificationRequestBizPackageBizData{
  AuthMode:  ctidtypes.CtidAuthModeTwoAndPhoto,
  PhotoData: <photoData(base64)>,
}, &ctideasy.EasyVerificationRequestAuthApplyRetainData{
  Name: "<name>",
  IdNo: "<idcard no>",
})
