// Package main provides ability to determine cert validity and such
package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// get a directory from a command line argument
	directory := flags()
	if *directory == "" {
		doStdIn(ioutil.ReadAll(os.Stdin))
	} else {
		// get a range of files from the directory
		doDirectory(*directory)
	}
}

func doStdIn(file_bytes []byte, err error) {
	if err == nil {
		showOutput(file_bytes)
	}
}

func doDirectory(directory string) {
	if directory != "" {
		file_list, _ := files(directory)
		for _, f := range file_list {
			// get the certs and create the output struct
			bytes, _ := ioutil.ReadFile(f)
			showOutput(bytes)
		}
	}
}
func flags() *string {
	d := flag.String("d", "", "cert directory to find *.crt")
	flag.Parse()
	return d
}

func showOutput(bytes []byte) {
	cert, err := certificate(bytes)
	if err != nil {
		fmt.Println(err)
	} else {
		o := output(cert)
		fmt.Printf("%s\t'%s'\t%d\n", o.DNSNames, o.Issuer, o.DaysTilExpiration)
	}
}

func files(directory string) ([]string, error) {
	return filepath.Glob(filepath.Join(directory, "*.crt"))
}

func certificate(file_bytes []byte) (*x509.Certificate, error) {
	pemBlock, _ := pem.Decode(file_bytes)
	return x509.ParseCertificate(pemBlock.Bytes)
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
