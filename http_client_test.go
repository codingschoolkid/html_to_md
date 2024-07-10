package htmltomd

import "testing"

func TestGet(t *testing.T) {
	resp, err := Get("https://www.baidu.com")
	if err != nil {
		t.Errorf("Get error: %v", err)
	}
	if len(resp) == 0 {
		t.Errorf("Get error: empty response")
	}
}
