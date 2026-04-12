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
}

func TestCookie(t *testing.T) {
	var cookies []*http.Cookie
	rawCookie := `_UP_A4A_11_=wba291af4c4247a995b4eaabc70db07f; b-user-id=dc877d0a-a43b-b36d-729c-ee0e4159b198; b-user-id=dc877d0a-a43b-b36d-729c-ee0e4159b198; __wpkreporterwid_=35ccc575-547d-4e8c-bf0c-9d5676fef972; isg=BEFBJZs9zBdpGSAjGcpFTkdtUIRbbrVgnF4mxKOVeMinim1c67xmMHHJaP7Mgk2Y; _UP_D_=pc; tfstk=gYFnLLDYUKY5k9xNi3lCB_8akE6TRXGSmud-2bnPQco6pDh8zbRz2zhKUMa8j0q8PDPKJD3zqlnswDFLp8YQsoAJ9M6QzQc-aiIAMsUBRbGPDHW0wmoIPPoPWvkW5fyKaiIA6dLZdMhzJ1M5e54Zf4AEzbry7N0orbREUBJZbcgrabJEa1uZuqtyaulPSPoszbuU4v7g7cgra0l5qQo0a-PN09XGeUgvP5unKmzEYyUL_rtx0yrWaQon-vXg8cAya5wz07LEq1-mV5rogbkDkM4Is5qmkkCwto4zl8kgZhSoYWUUjxehmBq0A4ri9S75dAMztWlgBsTmbW4usvw9kQaEEJE-ukBVcok0H-kgNsOITknK9WhVGHo8BrVor7s5so4zl8kgZHjrwdJVvcRS7aFMFLME5Vmv38_5gFZEAQ_GSK6ILVgjDNbMFLME5VmASNvXfvuslmC..; ctoken=SMh_Ue-CbuGoZXRmkcNgzFZV; grey-id=84e4ebf1-1038-9682-8c39-5420ea625bac; grey-id.sig=FE0ZvD82uoXkGWdIyUU-ZXVaw6JpECXJrrUc2RwgJmY; isQuark=true; isQuark.sig=hUgqObykqFom5Y09bll94T1sS9abT1X-4Df_lzgl8nM; _UP_F7E_8D_=i74i77uRTzvBGir6eD%2BMUFf%2BNNKTcBEqHeH0Qe%2B%2FPM6Tgsl3joSGZxfDU0dSsd6SstZghYFUULAc8umX3Lj2gnAl%2BT0uI3e2TOxuay7ZojksIBt5gp3JIBFHs%2BQ%2FpshPB42jNy%2FSQECUHaJDGHIdd8sH4vzcOfSUgOHfqGdFx%2BVb9AlhXg01OxHO9PRsEKg20jxP1zs9%2BlUts96g52ErQX7gJXHB34DmTJyVbneFy5M6bSold6u4Kz7f1qPddgiUDanFXJZrwBrsZkRmg588m0L3ni%2FcaBtzgicLGXWDULSzVjEvMmznaVZ2giTX2e9Z%2FK0%2FzVwkrF%2B%2FUZ%2FAFOtHRtL8Ea33V8aFkzDJm84ptT%2FxmRuCsa%2BbDGrOK%2BW7NepAeuYdTBrH19WVzMQZ1GracmRKLnouDIOtemzJh0ZgDQXgKP%2Be2vzS%2BRRNxrkTQuotwES3hQ3qMNg%3D; __pus=2f40c4da663bcceb633107749dcb71abAARqiqiwZNspre6DNqdrILkwn+0hlvc9B+mwk3dtusGRSJ2j6bQwuamyqN734FA9r/8cMmKU19r03iXB1F/Nhgoy; __kp=d548b650-364d-11f1-861b-3d1925a0ee36; __kps=AAQmU06jk4Ha6OmSSWctro1p; __ktd=B32ZcPTLr1AjNvYo91HB4g==; __uid=AAQmU06jk4Ha6OmSSWctro1p; web-grey-id=c6e5a194-7ade-d636-833b-607121936f97; web-grey-id.sig=HdCaB-qvsHpMONHYOo_f0lyUz5G7nY0tmj9xplDhhxA; __puus=37ba73b94d1ecd2b7f765fd2bfd4e34cAASFR1GGiWN85iMe8nHrQqIWlQ4Zr2YvVtn5hVF11FAc+PK3pzL4hOI9u2wlDMPDMlkEN4J1+gDUnimj9L+0mludPqHt8/89SsCW69dcHVshfygTyLQW+3wRorzYFUBRIrrz9wSF5xK0VmH0dSzRFKCQHBKmpZFdVecB8YJJp9MGoCddAm1O0bSOiKiZZAzt1qjnS5UaACBVltFSm4xk89O7
`
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
