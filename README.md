Minica2 is a simple CA intended for use in situations where the CA operator
also operates each host where a certificate will be used. It automatically
generates both a key and a certificate when asked to produce a certificate.
It does not offer OCSP or CRL services. Minica2 is appropriate, for instance,
for generating certificates for RPC systems or microservices.

On first run, minica2 will generate a keypair and a root certificate in the
current directory, and will reuse that same keypair and root certificate
unless they are deleted.

On each run, minica2 will generate a new keypair and sign an end-entity (leaf)
certificate for that keypair. The certificate will contain a list of DNS names
and/or IP addresses from the command line flags. The key and certificate are
placed in a new directory whose name is chosen as the first domain name from
the certificate, or the first IP address if no domain names are present. It
will not overwrite existing keys or certificates.

The certificate will have a validity of 2 years and 30 days.

# Installation

First, install the [Go tools](https://golang.org/dl/) and set up your `$GOPATH`.
Then, run:

`go install github.com/ctrlcctrlv/minica2.go@latest`

When using Go 1.11 or newer you don't need a $GOPATH and can instead do the
following:

```
cd /ANY/PATH
git clone https://github.com/ctrlcctrlv/minica2.go.git
go build
## or
# go install
```

Mac OS users could alternatively use Homebrew: `brew install minica2`

# Example usage

```
# Generate a root key and cert in minica2-key.pem, and minica2.pem, then
# generate and sign an end-entity key and cert, storing them in ./foo.com/
$ minica2 --domains foo.com

# Wildcard
$ minica2 --domains '*.foo.com'
```
