## SNMP oid reader for GE SitePro Parallel UPS Systems

#### ge_ups_sp_snmp

Read SNMP data and returns a json object.

Programm flags:

-n - number of ups in parallel system. From 1 to 8, or 0 - whole system GenericUPS (defaut value 0)

-ip - SNMP ip adress (defaut value localhost)

-p - SNMP tcp port (defaut value 161)

-c - SNMP community (defaut value public)

-t - timeout seconds (defaut value 3)

=======================================

**Note:**  Reads the OID .1.3.6.1.4.1.818.1.1.10.6.1.0 - "upsAlarmsPresent".

Anything deeper than OID .1.3.6.1.4.1.818.1.1.10.6.1.0 (alarm messages) is not read. Maybe in the future.


###### Example:

`ge_ups_sp_snmp -ip=192.168.10.10 -p=162 -n=2 -id=1 -c=pub -t=5000`