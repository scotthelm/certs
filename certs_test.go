// Package main provides ...
package main

import (
	"io/ioutil"
	"testing"
)

func TestFlagParsing(t *testing.T) {
	directory := flags()
	if *directory != "" {
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
		t.Error("expected 2 got ", len(file_list))
	}
}

func TestCertificateFromPath(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./ssl_certs/github.crt")
	crt, err := certificate(bytes)
	if err != nil {
		t.Error(err)
	}
	if crt.DNSNames[0] != "github.com" {
		t.Error("got ", crt.DNSNames)
	}
}

func TestStdIn(t *testing.T) {

}
