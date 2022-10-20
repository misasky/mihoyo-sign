package mihoyo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mihoyo-sign/common/utils"
	"net/http"
)

type GenShin struct {
	cookie string
}

const (
	AppVersion       = "2.35.2"
	UserAgentVersion = "2.3.0"
	ClientType       = "5"
	Referer          = "https://webstatic.mihoyo.com/bbs/event/signin-ys/index.html?bbs_auth_required=true&act_id=e202009291139501&utm_source=bbs&utm_medium=mys&utm_campaign=icon"
	UserAgent        = "Mozilla/5.0 (Linux; Android 10; MIX 2 Build/QKQ1.190825.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.101 Mobile Safari/537.36 miHoYoBBS/2.35.2"
	ActId            = "e202009291139501"
)

func NewGenShin(cookie string) *GenShin {
	return &GenShin{
		cookie: cookie,
	}
}

// Sign 签到
func (gs *GenShin) Sign() error {
	url := "https://api-takumi.mihoyo.com/event/bbs_sign_reward/sign"
	requestJson := map[string]interface{}{
		"act_id": ActId,
		"region": "cn_gf01",
		"uid":    "105770153",
	}
	b, err := json.Marshal(requestJson)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-rpc-device_id", utils.GetDeviceId(32))
	req.Header.Add("x-rpc-device_name", "Mi 10")
	req.Header.Add("x-rpc-app_version", AppVersion)
	req.Header.Add("x-rpc-client_type", "5")
	req.Header.Add("Cookie", gs.cookie)
	req.Header.Add("Referer", Referer)
	req.Header.Add("DS", utils.GetDSToken(true))
	req.Header.Add("x-rpc-device_model", "Mi 10")

	req.Header.Add("User-Agent", UserAgent)
	req.Header.Set("x-rpc-channel", "miyousheluodi")
	//req.Header.Set("x-rpc-sys_version", "6.0.1")
	req.Header.Set("X_Requested_With", "com.mihoyo.hyperion")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("mihoyo 未正常响应，返回 %s", string(b))
	}
	fmt.Println("[返回详情] ", res.Status, res.StatusCode, string(b))
	var resp signResp
	if err := json.Unmarshal(b, &resp); err != nil {
		return err
	}
	if resp.Retcode != 0 {
		return fmt.Errorf("签到失败，返回 %s", string(b))
	}
	return nil
}

type signResp struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Code      string `json:"code"`
		RiskCode  int    `json:"risk_code"`
		Gt        string `json:"gt"`
		Challenge string `json:"challenge"`
		Success   int    `json:"success"`
	} `json:"data"`
}

func (gs *GenShin) GetSignInfo(uid string) (*SignInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(
		"https://api-takumi.mihoyo.com/event/bbs_sign_reward/info?act_id=e202009291139501&region=cn_gf01&uid=%s", uid),
		nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Cookie", gs.cookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("mihoyo 未正常响应，返回 %s", string(b))
	}
	fmt.Println("[签到信息返回详情] ", resp.Status, resp.StatusCode, string(b))
	var ret SignInfo
	if err := json.Unmarshal(b, &ret); err != nil {
		return nil, err
	}
	if ret.Retcode != 0 {
		return nil, fmt.Errorf("签到信息返回失败，返回 %s", string(b))
	}
	return &ret, nil
}

type SignInfo struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		TotalSignDay  int    `json:"total_sign_day"`
		Today         string `json:"today"`
		IsSign        bool   `json:"is_sign"`
		FirstBind     bool   `json:"first_bind"`
		IsSub         bool   `json:"is_sub"`
		MonthFirst    bool   `json:"month_first"`
		SignCntMissed int    `json:"sign_cnt_missed"`
		MonthLastDay  bool   `json:"month_last_day"`
	} `json:"data"`
}
