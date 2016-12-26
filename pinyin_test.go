package pinyin

import "testing"

func TestIsPinyin1(t *testing.T) {
	s := "nihaoshijie"
	if is, result := IsPinyin(s); is != true || len(result) != 4 {
		t.Log(is)
		t.Log(result)
		t.Error(s + " test fail")
	} else {
		t.Log(s + " test success")
	}
}

func TestIsPinyin2(t *testing.T) {
	s := "helloworld"
	if is, _ := IsPinyin(s); is != false {
		t.Error(s + " test fail")
	} else {
		t.Log(s + " test success")
	}
}

func TestIsPinyin3(t *testing.T) {
	s := "你好世界"
	if is, _ := IsPinyin(s); is != false {
		t.Error(s + " test fail")
	} else {
		t.Log(s + " test success")
	}
}
