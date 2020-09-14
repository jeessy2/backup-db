package notice

import (
	"testing"
)

func TestDingDingSendMessage(test *testing.T) {
	ding := DingDing{
		WebHook: "https://oapi.dingtalk.com/robot/send?access_token=26b9fe0562e9a012cafacaa7adb12e4dd0b3c6a32dbff76d0625363ca52ece1d",
		Secret:  "SEC1297d8e2e56ae0bd7f975ab36beb7cfd53a4037bfac71fd9a7a66ebf34ba8f1b",
	}

	err := ding.SendMessage("", "hello")
	if err != nil {
		test.Fatal(err)
	}

}
