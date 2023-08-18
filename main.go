package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/TylerBrock/colorjson"
	"github.com/fatih/color"
	"github.com/monaco-io/request"
)

var (
	ipFlag      string
	versionFlag bool
	nocolorFlag bool
	version     string
)

func init() {
	flag.StringVar(&ipFlag, "ip", "", "IP To check")
	flag.BoolVar(&versionFlag, "version", false, "Show version")
	flag.BoolVar(&nocolorFlag, "nc", false, "Boring Output")
	flag.Parse()
}

func main() {
	if versionFlag {
		fmt.Printf("\n     📟  ipchk %v\n\n", version)
		return
	}

	printOutput(checkIP(ipFlag), nocolorFlag)

}

func checkIP(api string) string {
	var url = `http://ip-api.com/json/` + api
	c := request.Client{
		URL:    url,
		Method: "GET",
		Header: (map[string]string{
			"User-Agent": "ipchk - The IP Checker",
		}),
		Timeout:   time.Second * 20,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp := c.Send()
	return resp.String()
}

func printOutput(resp string, boring bool) {
	var obj map[string]interface{}
	json.Unmarshal([]byte(resp), &obj)
	f := colorjson.NewFormatter()
	f.Indent = 4
	f.KeyColor.Add(color.FgBlue, color.Bold)
	f.NumberColor.Add(color.FgWhite)
	f.BoolColor.Add(color.FgCyan)

	if !boring {
		s, _ := f.Marshal(obj)
		fmt.Println(string(s))
	} else {
		fmt.Printf("%v", resp)
	}
}
