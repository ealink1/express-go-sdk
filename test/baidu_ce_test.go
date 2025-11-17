package test

import (
	"fmt"
	"github.com/linbe-ff/express-go-sdk"
	"testing"
)

func TestBaiduAddr(t *testing.T) {

	// 更改为您的Key和Secret等
	kdClient := express.NewBaiCe("pXPUxxxxx", "vZvpfXvpxxxxx")

	resolution, err := kdClient.AnalyzeAddr("张三广东省深圳市南山区粤海街道科技南十二路金蝶软件园13088888888")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resolution)
}
