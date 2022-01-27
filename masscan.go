package masscan

import (
	"encoding/xml"
)

type MasscanRun nmapRun

type nmapRun struct {
	Scanner          string   `xml:"scanner,attr"`
	Start            string   `xml:"start,attr"`
	Version          string   `xml:"version,attr"`
	XmlOutputVersion string   `xml:"xmloutputversion,attr"`
	ScanInfo         ScanInfo `xml:"scaninfo"`
	Hosts            []Host   `xml:"host"`
}

type ScanInfo struct {
	Type     string `xml:"type,attr"`
	Protocol string `xml:"protocol,attr"`
}

type Host struct {
	Endtime string  `xml:"endtime,attr"`
	Address Address `xml:"address"`
	Ports   []Port  `xml:"ports>port"`
}

type Address struct {
	Addr string `xml:"addr,attr"`
	Type string `xml:"addrtype,attr"`
}

type Port struct {
	Protocol string `xml:"protocol,attr"`
	ID       string `xml:"portid,attr"`
	State    State  `xml:"state"`
}

type State struct {
	State  string `xml:"state,attr"`
	Reason string `xml:"reason,attr"`
	TTL    string `xml:"reason_ttl,attr"`
}

func Parse(content []byte) (*MasscanRun, error) {
	r := &MasscanRun{}
	err := xml.Unmarshal(content, r)
	return r, err
}
