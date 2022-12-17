package main

import (
	"fmt"
	"os"
)

func (args *Args) assignFlags() (ArgGroups) {
	return ArgGroups {
		args.assignStringFlags(),
		args.assignTargetFlags(),
		args.assignIssuerFlags(),
		args.assignBooleanFlags(),
	}
}

func printHelp(flags ArgGroups) {
	fmt.Fprintf(os.Stderr, "USAGE\n\n%s [--domain DOMAIN] [--ip-address IP] [OPTIONS…]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, `
SYNOPSIS

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

AUTHORS

Copyright (c) 2016 Jacob Hoffman-Andrews
Copyright (c) 2022 Fredrick R. Brennan <copypaste@kittens.ph>
`)
	fmt.Fprintf(os.Stderr, "version %v (built on %v)\n\n", version, date)

	fmt.Fprintf(os.Stderr, "OPTIONS\n\nSTRINGS:\n")
	flags.sflag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "TARGETS:\n")
	flags.tflag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "FLAGS:\n")
	flags.bflag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "ISSUER INFORMATION (all optional):\n")
	flags.iflag.PrintDefaults()
}
