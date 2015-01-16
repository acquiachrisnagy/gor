package main

import (
	"testing"
)

func TestUrlRewriteMap_1(t *testing.T) {
	var url string

	rewrites := UrlRewriteMap{}

	err := rewrites.Set("/abc:/123")
	if err != nil {
		t.Error("Should not error on /abc:/123")
	}

	url = "/abc"
	if rewrites.Rewrite(url) == url {
		t.Error("Request url should have been rewritten, wasn't")
	}

	url = "/wibble"
	if rewrites.Rewrite(url) != url {
		t.Error("Request url should not have been rewritten, was")
	}
}

func TestUrlRewriteMap_2(t *testing.T) {
	var url string

	rewrites := UrlRewriteMap{}

	err := rewrites.Set("/v1/user/([^\\/]+)/ping:/v2/user/$1/ping")
	if err != nil {
		t.Error("Should not error on /v1/user/([^\\/]+)/ping:/v2/user/$1/ping")
	}

	url = "/v1/user/joe/ping"
	if rewrites.Rewrite(url) == url {
		t.Error("Request url should have been rewritten, wasn't")
	}

	url = "/v1/user/joe/ping"
	if rewrites.Rewrite(url) != "/v2/user/joe/ping" {
		t.Error("Request url should have been rewritten, wasn't")
	}

	url = "/v1/user/ping"
	if rewrites.Rewrite(url) != url {
		t.Error("Request url should not have been rewritten, was")
	}
}

func TestUrlRewriteMap_3(t *testing.T) {
	var rawQuery string

	rewrites := UrlRewriteMap{}

	err := rewrites.Set("a=1:a=2")
	if err != nil {
		t.Error("Should not error on a=1:a=2")
	}

	rawQuery = "a=1&b=2"
	if rewrites.Rewrite(rawQuery) == rawQuery {
		t.Error("Request url should have been rewritten, wasn't")
	}

	rawQuery = "a=1&b=2"
	if rewrites.Rewrite(rawQuery) != "a=2&b=2" {
		t.Error("Request url should have been rewritten, wasn't")
	}

	rawQuery = "a=2&b=2"
	if rewrites.Rewrite(rawQuery) != rawQuery {
		t.Error("Request url should not have been rewritten, was")
	}
}
