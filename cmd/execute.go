package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/Harry7U/ReconSuite-AI/core"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func Execute(target, header string, extract, subdomains, ai, filter, exploit bool, output string, threads, timeout int, proxy string, debug bool) {
	// Initialize configuration
	config := core.Config{
		Target:     target,
		Header:     header,
		Extract:    extract,
		Subdomains: subdomains,
		AI:         ai,
		Filter:     filter,
		Exploit:    exploit,
		Output:     output,
		Threads:    threads,
		Timeout:    timeout,
		Proxy:      proxy,
		Debug:      debug,
	}

	// Create spinner for visual feedback
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()

	// Run reconnaissance
	if config.Subdomains || config.Extract {
		color.Cyan("🔍 Running reconnaissance...")
		err := core.RunRecon(config)
		if err != nil {
			s.Stop()
			color.Red("Reconnaissance failed: %v", err)
			log.Fatalf("Reconnaissance failed: %v", err)
		}
		color.Green("Reconnaissance completed successfully!")
	}

	// Run filtering
	if config.Filter {
		color.Cyan("🔍 Running filtering...")
		err := core.RunFilter(config)
		if err != nil {
			s.Stop()
			color.Red("Filtering failed: %v", err)
			log.Fatalf("Filtering failed: %v", err)
		}
		color.Green("Filtering completed successfully!")
	}

	// Run exploitation
	if config.Exploit {
		color.Cyan("💥 Running exploitation...")
		err := core.RunExploit(config)
		if err != nil {
			s.Stop()
			color.Red("Exploitation failed: %v", err)
			log.Fatalf("Exploitation failed: %v", err)
		}
		color.Green("Exploitation completed successfully!")
	}

	// Generate AI payloads
	if config.AI {
		color.Cyan("🤖 Generating AI payloads...")
		payload := core.GeneratePayload("vulnerability type", config.Target)
		color.Green("Generated AI payload: %s", payload)
	}

	// Save results
	color.Cyan("💾 Saving results...")
	err := core.SaveResults(config)
	if err != nil {
		s.Stop()
		color.Red("Saving results failed: %v", err)
		log.Fatalf("Saving results failed: %v", err)
	}
	color.Green("Results saved successfully!")

	s.Stop()
	color.Magenta("🎉 ReconSuite-AI finished successfully! 🎉")
}
