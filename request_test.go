package HttpRequest

import (
	"fmt"
	"bytes"
	"testing"
)

const testUrl = "https://cvm.tencentcloudapi.com/"

var data = map[string]interface{}{
	"var1": "val1",
	"var2": 2022,
}

func TestGetRequest(t *testing.T) {
	req := NewRequest()
	req.SetDebug(false)

	var resp *Response
	var err error

	resp, err = req.Get(testUrl, nil)
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode() != 200 {
		t.Error("GET "+testUrl, "expected 200", fmt.Sprintf("return %d", resp.StatusCode()))
	}

	resp, err = req.Get(testUrl, data)
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode() != 200 {
		t.Error("GET "+testUrl, "expected 200", fmt.Sprintf("return %d", resp.StatusCode()))
	}

	resp, err = req.Get(testUrl, "var3=val3")
	if err != nil {
		t.Error(err)
		return
	}

	if resp.StatusCode() != 200 {
		t.Error("GET "+testUrl, "expected 200", fmt.Sprintf("return  %d", resp.StatusCode()))
	}
}

func TestPostRequest(t *testing.T) {
	req := NewRequest()
	req.SetDebug(false)

	postData := []interface{}{
		data,
		bytes.NewReader([]byte{97}),
		[]byte{97},
		nil,
		"helloworld",
		`{"var1":"val1","var2":2022}`,
		100,
		int8(100),
		int16(100),
		int32(100),
		int64(100),
		"var1=val1&var2=2022",
	}

	var resp *Response
	var err error

	for _, v := range postData {
		resp, err = req.Post(testUrl, v)
		if err != nil {
			t.Error(err)
			return
		}
	}

	if resp.StatusCode() != 200 {
		t.Error("GET "+testUrl, "expected 200", fmt.Sprintf("return %d", resp.StatusCode()))
	}

}
