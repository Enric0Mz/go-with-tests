package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "http://anothersiteouthere.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.random.com",
		"http://anothersiteouthere.com",
	}

	want := map[string]bool{
		"http://google.com":             true,
		"http://blog.random.com":        true,
		"http://anothersiteouthere.com": false,
	}

	actual := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, actual) {
		t.Fatalf("wanted %v, got %v", want, actual)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebistes(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	for b.Loop() {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
