// Package main provides ability to determine cert validity and such
package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

func main() {
	// get a directory from a command line argument
	directory := flags()
	// get a range of files from the directory
	file_list, _ := files(*directory)
	for _, f := range file_list {
		// get the certs and create the output struct
		cert, err := certificate(f)
		if err != nil {
			fmt.Println(err)
		} else {
			o := output(cert)
			fmt.Printf("%s\t'%s'\t%d\n", o.DNSNames, o.Issuer, o.DaysTilExpiration)
		}
	}
}

func flags() *string {
	d := flag.String("-d", "./ssl_certs/", "cert directory to find *.crt")
	flag.Parse()
	return d
}

func files(directory string) ([]string, error) {
	return filepath.Glob(filepath.Join(directory, "*.crt"))
}

func certificate(filepath string) (*x509.Certificate, error) {
	file_bytes, file_err := ioutil.ReadFile(filepath)
	pemBlock, _ := pem.Decode(file_bytes)
	var cert *x509.Certificate
	var err error
	if file_err == nil {
		cert, err = x509.ParseCertificate(pemBlock.Bytes)
	}
	return cert, err
}

func output(cert *x509.Certificate) Output {
	return Output{
		cert.DNSNames,
		cert.Issuer.CommonName,
		cert.NotBefore,
		cert.NotAfter,
		int(cert.NotAfter.Sub(time.Now()).Hours() / 24),
	}
}

type Output struct {
	DNSNames          []string
	Issuer            string
	NotBefore         time.Time
	NotAfter          time.Time
	DaysTilExpiration int
}
