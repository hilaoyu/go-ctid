go get github.com/hilaoyu/go-ctid


c := ctideasy.NewCtidEasyClient("<orgCode>")  
err = c.SetSignPrivateKey(priKey)  
fmt.Println("SetPrivateKey err", err)  
err = c.SetSignPublicKey(pubKey)  
fmt.Println("SetPublicKey err", err)  
err = c.SetDataEncodePublicKey(dataPubKey)  
fmt.Println("SetDataEncodePublicKey err", err)  

resultData, err := c.Verification(&ctideasy.EasyVerificationRequestBizPackageBizData{  
&nbsp;&nbsp;AuthMode:  ctidtypes.CtidAuthModeTwoAndPhoto,  
&nbsp;&nbsp;PhotoData: <photoData(base64)>,  
}, &ctideasy.EasyVerificationRequestAuthApplyRetainData{  
&nbsp;&nbsp;Name: "<name>",  
&nbsp;&nbsp;IdNo: "<idcard no>",  
})  

