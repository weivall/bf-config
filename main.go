package main

import "strings"

// Betfair config
type Config struct {
	User     string
	Password string
	AppName  string
	CertUrl  string
	Cert     []byte
	Key      []byte
}

// NormalizeRSA replaces '#' by '\n' in input string
func NormalizeRSA(s string) string {
	const (
		hashCode    = 35 // #
		newLineCode = 10 // \n
	)
	return strings.Replace(s, string(byte(hashCode)), string(byte(newLineCode)), -1)
}
