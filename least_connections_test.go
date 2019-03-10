package leastconnections

import (
	"net/url"
	"reflect"
	"testing"
)

func TestLeastConnections(t *testing.T) {
	lc, err := New([]*url.URL{
		{Host: "192.168.33.10"},
		{Host: "192.168.33.11"},
		{Host: "192.168.33.12"},
	})
	if err != nil {
		t.Error("New returns nil")
	}

	url1, done1 := lc.Next()
	if got, want := url1, (&url.URL{Host: "192.168.33.10"}); !reflect.DeepEqual(got, want) {
		t.Errorf("Next is wrong. want: %v, but got: %v", want, got)
	}

	url2, done2 := lc.Next()
	if got, want := url2, (&url.URL{Host: "192.168.33.11"}); !reflect.DeepEqual(got, want) {
		t.Errorf("Next is wrong. want: %v, but got: %v", want, got)
	}
	done1()

	url3, done3 := lc.Next()
	if got, want := url3, (&url.URL{Host: "192.168.33.10"}); !reflect.DeepEqual(got, want) {
		t.Errorf("Next is wrong. want: %v, but got: %v", want, got)
	}

	url4, done4 := lc.Next()
	if got, want := url4, (&url.URL{Host: "192.168.33.12"}); !reflect.DeepEqual(got, want) {
		t.Errorf("Next is wrong. want: %v, but got: %v", want, got)
	}

	url5, done5 := lc.Next()
	if got, want := url5, (&url.URL{Host: "192.168.33.10"}); !reflect.DeepEqual(got, want) {
		t.Errorf("Next is wrong. want: %v, but got: %v", want, got)
	}

	done2()
	done3()
	done4()
	done5()
}
