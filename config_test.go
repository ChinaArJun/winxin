package weixin

import (
	"time"

	"github.com/arstd/log"
)

// 微信公众平台测试号
// http://mp.weixin.qq.com/debug/cgi-bin/sandboxinfo?action=showinfo&t=sandbox/index
// 测试号 毛毛雨
const (
	//originId       = "gh_398ea213a0d3"
	//appId          = "wxad7f6407fcbb9cae"
	//appSecret      = "0884d88aaa8ef70dfcd484fcc52bd484"
	//token          = "0t37dWsIYg6NsVLgEY1fNuB1rSLyyeQE"
	//encodingAESKey = "dkVSmV3CtG3IBCY96A8RSNoCOzaPcx36uGJo8fl9wWn" // 为了通过验证
	originId       = "gh_398ea213a0d3"
	appId          = "wxdfd0fdb2e3de09ce"
	appSecret      = "b3203e44d39a0b0d5b7d8def439997de"
	token          = "0t37dWsIYg6NsVLgEY1fNuB1rSLyyeQE"
	encodingAESKey = "dkVSmV3CtG3IBCY96A8RSNoCOzaPcx36uGJo8fl9wWn"
)

// 小程序
const (
	miniAppId     = "wxd18eafc798ccfe25"
	miniAppSecret = "7a968f61971a3dba1635c95b188fc67c"
)

// 单元测试 testURL
const testURL = "http://127.0.0.1:3080/weixin?signature=292ea5c2b515b3615eecca128f5f90b05d1786dc&timestamp=1449649715&nonce=2109185587"

func init() {
	Initialize(originId, appId, appSecret, token, encodingAESKey)
	log.Info("Initializing, please wait 5 seconds...")
	time.Sleep(5 * time.Second) // waiting to refresh token
}
