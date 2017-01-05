package userService

import (
	"casino_common/common/sys"
	"testing"
)

func init() {
	sys.InitRedis("127.0.0.1:6379", "test")
}

func TestINCRUserCOIN(t *testing.T) {
	b, e := DECRUserDiamond(1, 2)
	if e != nil {
		t.Logf("出错%v", e)
	} else {
		t.Logf("增加之后的金额%v", b)
	}
}
