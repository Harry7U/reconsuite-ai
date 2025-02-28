package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Harry7U/ReconSuite-AI/cmd"
	"github.com/fatih/color"
)

func main() {
	// Define flags
	target := flag.String("target", "", "Set target domain")
	header := flag.String("header", "", "Set custom HTTP headers")
	extract := flag.Bool("extract", false, "Extract URLs, APIs, JS endpoints")
	subdomains := flag.Bool("subdomains", false, "Run subdomain enumeration")
	ai := flag.Bool("ai", false, "Use AI for payload generation")
	filter := flag.Bool("filter", false, "Apply GF filtering for vulnerabilities")
	exploit := flag.Bool("exploit", false, "Run exploitation (SQLi, XSS, SSRF, RCE)")
	output := flag.String("output", "json", "Save results as JSON/HTML")
	threads := flag.Int("threads", 10, "Set number of threads")
	timeout := flag.Int("timeout", 10, "Set request timeout")
	proxy := flag.String("proxy", "", "Use a proxy for scanning")
	debug := flag.Bool("debug", false, "Enable debug mode")

	// Parse flags
	flag.Parse()

	// Validate required flags
	if *target == "" {
		color.Red("Error: Target domain is required")
		flag.Usage()
		os.Exit(1)
	}

	// Display banner
	color.Cyan("🚀 ReconSuite-AI: The Ultimate AI-Powered Bug Bounty & Exploitation Framework 🚀")
	color.Green("Starting ReconSuite-AI with the following settings:")
	color.Yellow("Target: %s", *target)
	color.Yellow("Custom Headers: %s", *header)
	color.Yellow("Extract: %t", *extract)
	color.Yellow("Subdomains: %t", *subdomains)
	color.Yellow("AI-Powered: %t", *ai)
	color.Yellow("Filter: %t", *filter)
	color.Yellow("Exploit: %t", *exploit)
	color.Yellow("Output Format: %s", *output)
	color.Yellow("Threads: %d", *threads)
	color.Yellow("Timeout: %d seconds", *timeout)
	color.Yellow("Proxy: %s", *proxy)
	color.Yellow("Debug Mode: %t", *debug)

	// Execute command
	cmd.Execute(*target, *header, *extract, *subdomains, *ai, *filter, *exploit, *output, *threads, *timeout, *proxy, *debug)
}