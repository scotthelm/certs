Certs
=====

This utility will read an x509 certificate from stdin or all crt files in a given
directory.

## Usage

certs is invoked this way `certs -d="/path/to/cert/dir"` If no path is given,
then it will read from stdin `cat /path/to/certificate.crt | certs`

## Output

The output is tab delimited and in the form of:

    [dnsnames]  issuer  days-til-expiration

