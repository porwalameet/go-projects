package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the domain to check: ")

	// Currently check for single domain
	// TODO: checking for multiple domains
	for scanner.Scan() {
		checkDomain(scanner.Text())
		fmt.Println("Enter another domain to check or Ctrl+C to exit")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read frmo input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Check MX records for the given domain
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error retrieving MX records: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// Check SPF records for given domain
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error retrieving TXT records: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// Check DMARC recods for given domain
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Panicf("Error retreiving DMARC records: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%-20v %-20v %-20v %-40v %-20v %-40v\n", "domain", "hasMX", "hasSPF", "spfRecord", "hasDMARC", "dmarcRecord")
	fmt.Printf("%-20v %-20v %-20v %-40v %-20v %-40v\n\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	return
}
