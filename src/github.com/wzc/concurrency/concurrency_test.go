package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "www.baidu.com" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"www.baidu.com",
		"www.google.com",
		"www.bing.com",
	}

	want := map[string]bool{
		"www.baidu.com":  false,
		"www.google.com": true,
		"www.bing.com":   true,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	//make返回指定长度的slice
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
