// Package main provides ...
package main

import "testing"

func TestFlagParsing(t *testing.T) {
	directory := flags()
	if *directory != "./ssl_certs/" {
		t.Error("expected ./ssl_certs/, got ", *directory)
	}
}

func TestDirectoryParsing(t *testing.T) {
	directory := "./ssl_certs/"
	file_list, err := files(directory)
	if err != nil {
		t.Error(err)
	}
	if len(file_list) != 2 {
		t.Error("expected 2 got ", file_list)
	}
}

func TestCertificateFromPath(t *testing.T) {
	crt, err := certificate("./ssl_certs/campaignio.com.crt")
	if err != nil {
		t.Error(err)
	}
	if crt.DNSNames[0] != "*.campaignio.com" {
		t.Error("got ", crt.DNSNames)
	}
}
