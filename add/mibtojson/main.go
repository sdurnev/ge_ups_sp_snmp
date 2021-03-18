package main

import (
	"fmt"
	"github.com/hallidave/mibtool/smi"
	"log"
)

func main() {
	mib := smi.NewMIB("/usr/share/snmp/mibs", "/Users/svdu/Nextcloud/Develop/Golang/ge_ups_sp_snmp/add/mibtojson/MIB")
	mib.Debug = true
	err := mib.LoadModules("GESINGLEUPS-MIB")
	if err != nil {
		log.Fatal(err)
	}

	// Walk all symbols in MIB
	mib.VisitSymbols(func(sym *smi.Symbol, oid smi.OID) {
		fmt.Printf("%-40s %s\n", sym, oid)
	})

	// Look up OID for an OID string
	oid, err := mib.OID("IF-MIB::ifOperStatus.4")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(oid.String())
	//fmt.Println()
}
