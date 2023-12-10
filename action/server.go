package action

import (
	"github.com/xuxiaowei-com-cn/go-aliyundrive"
	"log"
	"net/http"
)

var (
	appIdVar     string
	appSecretVar string
	scopeVar     []string
	widthVar     int
	heightVar    int
)

func QrCodeAction(appId string, appSecret string, scope []string, width int, height int, httpListenAddr string) error {

	appIdVar = appId
	appSecretVar = appSecret
	scopeVar = scope
	widthVar = width
	heightVar = height

	http.HandleFunc("/qrcode", qrcodeHandler)
	err := http.ListenAndServe(httpListenAddr, nil)
	if err != nil {
		return err
	}

	return nil
}

func qrcodeHandler(w http.ResponseWriter, r *http.Request) {

	client, err := aliyundrive.NewClient("")
	if err != nil {
		panic(err)
	}

	body := aliyundrive.QrCodeBody{
		ClientId:     appIdVar,
		ClientSecret: appSecretVar,
		Scopes:       scopeVar,
		Height:       heightVar,
		Width:        widthVar,
	}

	qrCode, response, err := client.QrCode.PostQrCode(body)
	if err != nil {
		panic(err)
	}

	log.Println("QrCode Response Status", response.Status)
	log.Println("QrCode Response QrCodeUrl", qrCode.QrCodeUrl)
	log.Println("QrCode Response Sid", qrCode.Sid)

	http.Redirect(w, r, qrCode.QrCodeUrl, http.StatusMovedPermanently)
}
