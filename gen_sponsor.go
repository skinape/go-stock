//go:build ignore

package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/duke-git/lancet/v2/cryptor"
)

// 与 main.go 中 checkDir 回退值一致
const aesKeyHex = "cc1e0d684e32f176c56ff1fcf384dcd9"

type SponsorInfo struct {
	VipLevel     string `json:"vipLevel"`
	VipStartTime string `json:"vipStartTime"`
	VipEndTime   string `json:"vipEndTime"`
	VipAuthTime  string `json:"vipAuthTime"`
}

func main() {
	now := time.Now()

	info := SponsorInfo{
		VipLevel:     "2",                                                                      // vip等级: 1 或 2
		VipStartTime: now.Format("2006-01-02 15:04:05"),                                        // VIP生效时间
		VipEndTime:   now.AddDate(100, 0, 0).Format("2006-01-02 15:04:05"),                    // VIP过期时间（这里设了100年）
		VipAuthTime:  now.Add(-24 * time.Hour).Format("2006-01-02 15:04:05"),                   // 授权时间（设为昨天，确保当前时间已过）
	}

	jsonBytes, _ := json.Marshal(info)
	fmt.Println("原始JSON:", string(jsonBytes))

	key, _ := hex.DecodeString(aesKeyHex)
	encrypted := cryptor.AesEcbEncrypt(jsonBytes, key)
	sponsorCode := hex.EncodeToString(encrypted)

	fmt.Println("\n赞助码:", sponsorCode)
}
