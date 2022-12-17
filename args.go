package main

import (
	"crypto/x509/pkix"
	"flag"
	"fmt"
	"os"
	"net"
	"regexp"
	"strings"
)

type Args struct {
	caKey, caCert *string
	macValidity *bool
	domains, ipAddresses ArgsArr
	org, iorg, iunit, icountry, ilocale, iaddress, ipostal ArgsArr
}

type ArgsArr struct {
	a []string
}

func (arr *ArgsArr) sumFlagFunc(arg string) error {
	arr.a = append(arr.a, arg)
	return nil
}

func (arr *ArgsArr) sumFlagFuncIP(arg string) error {
	if arg == "" {
		return nil
	}
	if net.ParseIP(arg) == nil {
		return fmt.Errorf("Invalid IP address %q\n", arg)
	}
	arr.a = append(arr.a, arg)
	return nil
}

func (arr *ArgsArr) sumFlagFuncDomain(arg string) error {
	if arg == "" {
		return nil
	}
	domainRe := regexp.MustCompile("^[A-Za-z0-9.*-]+$")
	if !domainRe.MatchString(arg) {
		return fmt.Errorf("Invalid domain name %q\n", arg)
	}
	arr.a = append(arr.a, arg)
	return nil
}

func (args *Args) assignIssuerFlags() (flag.FlagSet) {
	var iflag = flag.NewFlagSet("issuer", flag.ExitOnError)
	iflag.Func("issuer", "Issuing organization common name", args.org.sumFlagFunc)
	iflag.Func("organization", "Issuing organization", args.iorg.sumFlagFunc)
	iflag.Func("unit", "Issuing unit of organization (e.g., IT)", args.iunit.sumFlagFunc)
	iflag.Func("country", "Issuer's country", args.icountry.sumFlagFunc)
	iflag.Func("locality", "Issuer's locality (i.e., city)", args.ilocale.sumFlagFunc)
	iflag.Func("address", "Issuer's address", args.iaddress.sumFlagFunc)
	iflag.Func("postal-code", "Issuer's postal code (in 🇺🇸 called a “ZIP” code)", args.ipostal.sumFlagFunc)
	return *iflag
}

func (args *Args) assignStringFlags() (flag.FlagSet) {
	var sflag = flag.NewFlagSet("string", flag.ExitOnError)
	args.caKey = sflag.String("ca-key", "minica-key.pem", "Root private key filename, PEM encoded.")
	args.caCert = sflag.String("ca-cert", "minica.pem", "Root certificate filename, PEM encoded.")
	return *sflag
}

func (args *Args) assignTargetFlags() (flag.FlagSet) {
	var tflag = flag.NewFlagSet("target", flag.ExitOnError)
	tflag.Func("domain", "Comma separated domain names to include as Server Alternative Names.", args.domains.sumFlagFuncDomain)
	tflag.Func("ip-address", "Comma separated IP addresses to include as Server Alternative Names.", args.ipAddresses.sumFlagFuncIP)
	return *tflag
}

func (args *Args) assignBooleanFlags() (flag.FlagSet) {
	var bflag = flag.NewFlagSet("boolean", flag.ExitOnError)
	args.macValidity = bflag.Bool("mac-validity", false, "Make a valid certificate for macOS / iOS (2 yrs + 30 days validity)")
	return *bflag
}

type ArgGroups struct {
	sflag, tflag, iflag, bflag flag.FlagSet
}

func (args *Args) parseIssuer() (pkix.Name) {
	return pkix.Name{
		CommonName: strings.Join(args.org.a, ","),
		Organization: args.iorg.a,
		OrganizationalUnit: args.iunit.a,
		Country: args.icountry.a,
		Locality: args.ilocale.a,
		StreetAddress: args.iaddress.a,
		PostalCode: args.ipostal.a,
	}
}

func (args *Args) parse() {
	var flags = args.assignFlags()
	flag.Usage = func() { printHelp(flags) }
	flag.Parse()
	var domains = args.domains.a
	var ipAddresses = args.ipAddresses.a
	if len(domains) == 0 && len(ipAddresses) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if len(flag.Args()) > 0 {
		fmt.Printf("Extra arguments: %s (maybe there are spaces in your domain list?)\n", flag.Args())
		os.Exit(1)
	}
}
