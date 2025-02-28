package core

type Config struct {
	Target     string
	Header     string
	Extract    bool
	Subdomains bool
	AI         bool
	Filter     bool
	Exploit    bool
	Output     string
	Threads    int
	Timeout    int
	Proxy      string
	Debug      bool
}