package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/medianexapp/plugin_api/plugin"
)

func TestPluginImpl(t *testing.T) {
	cookies := ""
	p := NewPluginImpl()
	auth, _ := p.GetAuth()
	method := auth.AuthMethods[0].Method

	method.(*plugin.AuthMethod_Formdata).Formdata.FormItems[0].Value.(*plugin.Formdata_FormItem_StringValue).StringValue.Value = cookies
	authData, err := p.CheckAuthMethod(&plugin.AuthMethod{
		Method: auth.AuthMethods[0].Method,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = p.CheckAuthData(authData.AuthDataBytes)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := p.GetDirEntry(&plugin.GetDirEntryRequest{
		Path:     "/",
		Page:     1,
		PageSize: 30,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetFileResource %s\n", resp.FileEntries[0].RawData)
	resources, err := p.GetFileResource(&plugin.GetFileResourceRequest{
		FilePath:  "/Unicorn.Academy.S01E02.The.Hidden.Temple.1080p.NF.WEB-DL.DDP5.1.H.264-FLUX.mkv",
		FileEntry: resp.FileEntries[0],
	})
	if err != nil {
		t.Fatal(err)
	}
	f := resources.FileResourceData[0]
	t.Log(f.Url)

	dd, _ := json.Marshal(f.Header)
	t.Log(string(dd))

}

func TestQuarkQrcode(t *testing.T) {
	p := NewPluginImpl()

	token, err := p.getQrcodeToken()
	if err != nil {
		t.Fatal(err)
	}

	// 直接输出到控制台
	fmt.Println(p.qrcodeData(token))
	for {
		res, err := p.checkQrcode(token)
		if err == nil && res != nil {
			t.Log("cookie", res)
			break
		}
		time.Sleep(time.Second * 2)
	}
	p.cookies = p.hb.GetCookies()
}

func TestCookie(t *testing.T) {
	var cookies []*http.Cookie
	rawCookie := ``
	for _, cookie := range strings.Split(rawCookie, ";") {
		sp := strings.Split(strings.TrimSpace(cookie), "=")
		if len(sp) != 2 {
			continue
		}
		cookies = append(cookies, &http.Cookie{Name: sp[0], Value: sp[1]})
	}
	ddd, _ := json.Marshal(cookies)
	fmt.Println(string(ddd))
	var cookies3 []*http.Cookie
	fmt.Println(json.Unmarshal([]byte(rawCookie), &cookies3))
}
