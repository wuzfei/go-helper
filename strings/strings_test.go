package strings

import (
	"errors"
	"strings"
	"testing"
)

var str1 = "a=我的1;飞机=司机"

func TestKv2MapFunc(t *testing.T) {
	res := Kv2MapFunc(str1, ";", func(s string) (string, string, error) {
		arr := strings.SplitN(s, "=", 2)
		if len(arr) != 2 {
			return "", "", errors.New("err")
		}
		return arr[0], arr[1], nil
	})
	if res["飞机"] != "司机" {
		t.Errorf("Kv2MapFunc(%s) = %+v", str1, res)
	}
}
