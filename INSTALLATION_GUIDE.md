# ReconSuite-AI Installation and Update Guide

ReconSuite-AI is a collection of tools and scripts for comprehensive reconnaissance, combining traditional recon utilities with AI-powered analysis. This guide will walk you through installing all required dependencies and tools – including Go, Git, the OpenAI API, and key recon tools (gf, subfinder, httpx, etc.) – and explain how to keep each of them up to date. We also include official documentation references and troubleshooting tips for common issues.

## 1. Install Go (Golang) – Latest Version

ReconSuite-AI relies on several Go-based tools (e.g. subfinder, httpx, gf), so having the latest Go compiler is essential. Many of these tools require a recent Go version (Go 1.19+ or newer) for installation ([Installing Subfinder - ProjectDiscovery Documentation](https://docs.projectdiscovery.io/tools/subfinder/install#:~:text=go%20install%20)). 

**Steps to Install Go on Linux/macOS:**

1. **Download Go:** Visit the official Go download page and get the latest Go tarball for your platform (e.g. Linux 64-bit). For example, replace `X.Y.Z` with the current version in the commands below:  
   ```bash
   wget https://go.dev/dl/goX.Y.Z.linux-amd64.tar.gz
   sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf goX.Y.Z.linux-amd64.tar.gz
   ```  
   This removes any old Go version and extracts Go into `/usr/local/go` ([Download and install - The Go Programming Language](https://go.dev/doc/install#:~:text=Linux%20Image%20%20%20,58%20%20%20Windows%20Image)). (On macOS, you can use the official .pkg installer or Homebrew: `brew install go`.)

2. **Configure PATH:** Add Go’s binary directory to your PATH. For Linux/macOS, append this line to your shell profile (e.g. `~/.bashrc` or `~/.zshrc`):  
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```  
   Then reload the profile with `source ~/.bashrc`. This makes the `go` command accessible system-wide ([Download and install - The Go Programming Language](https://go.dev/doc/install#:~:text=2,environment%20variable)). (Homebrew and the macOS installer usually handle PATH automatically. On Windows, the MSI installer will set PATH for `go.exe`.)

3. **Verify Installation:** Open a new terminal and check the version:  
   ```bash
   go version
   ```  
   You should see the Go version printed, confirming it’s installed ([Download and install - The Go Programming Language](https://go.dev/doc/install#:~:text=3,and%20typing%20the%20following%20command)).

**Updating Go:** To update Go, repeat the above steps with the newer version (remove the old `/usr/local/go` then install the new version) ([Download and install - The Go Programming Language](https://go.dev/doc/install#:~:text=Linux%20Image%20%20%20,58%20%20%20Windows%20Image)). If you installed via a package manager (e.g. `apt` or Homebrew), you can use those to update (e.g. `sudo apt update && sudo apt install --only-upgrade golang` or `brew upgrade go`). Always verify with `go version` after updating.

**Troubleshooting:** If `go version` still shows an old version or command not found: ensure no older Go exists in your PATH and that you added the new Go path correctly. Open a new shell or run `source ~/.bashrc` to apply PATH changes ([Download and install - The Go Programming Language](https://go.dev/doc/install#:~:text=Note%3A%20Changes%20made%20to%20a,source%20%24HOME%2F.profile)). On Linux, check that `/usr/local/go/bin` comes before any older Go binary in `$PATH`. On Windows, if `go` isn’t recognized, reboot or log out/in so the PATH change takes effect ([Download and install - The Go Programming Language](https://go.dev/doc/install#:~:text=1,the%20prompts%20to%20install%20Go)).

## 2. Install Git

Git is required to clone repositories (for example, fetching the gf patterns and other tool sources). Most Linux systems have Git or allow easy installation via package managers, and macOS can install it via Command Line Tools or Homebrew ([Git - Installing Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git#:~:text=Installing%20on%20macOS)).

**Installation on Linux:** Update your package index and install Git:  
```bash
sudo apt update && sudo apt install -y git
```  
On Fedora/CentOS, use `dnf install git-all`. On Debian/Ubuntu, `apt install git-all` installs Git and common add-ons ([Git - Installing Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git#:~:text=%24%20sudo%20dnf%20install%20git)). This will install the latest version available in your distro’s repositories (often Git 2.x).

**Installation on macOS:** Easiest method is to run `git --version` in Terminal; if Git is not already installed, macOS will prompt to install Command Line Developer Tools which include Git ([Git - Installing Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git#:~:text=Installing%20on%20macOS)). Alternatively, install Homebrew and run `brew install git`, or download the installer from the official site ([Git - Installing Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git#:~:text=%24%20git%20)).

**Installation on Windows:** Download the **Git for Windows** installer from the official site (Git SCM) and run it. During installation, you can opt to add Git to your PATH. After installing, open “Git Bash” or a new Command Prompt and run `git --version` to verify it’s accessible.

**Updating Git:** On Linux, update via your package manager (e.g. `sudo apt upgrade git` for newer releases when they become available, or add the official PPA for the latest Git). On macOS with Homebrew, run `brew update && brew upgrade git`. On Windows, download and run the latest installer, or use Git for Windows’ built-in update feature if available. Always check `git --version` after updating to ensure the new version is active ([Git - Installing Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git#:~:text=Before%20you%20start%20using%20Git%2C,code%20and%20compile%20it%20yourself)).

**Troubleshooting:** If `git` isn’t found after installation, make sure the installation completed and your PATH is set. For Linux, confirm `/usr/bin/git` or `/usr/local/bin/git` exists. On Windows, ensure you opened a new terminal (or use the Git Bash provided by the installer). If you encounter SSL or proxy errors when using Git (cloning), ensure your internet connection or proxy settings are configured (you can set `https_proxy` environment or use `git config --global http.proxy` as needed).

## 3. OpenAI API Setup

ReconSuite-AI integrates with OpenAI’s API for AI-driven analysis. To use the OpenAI API, you’ll need to obtain an API key and set up the OpenAI client library.

**Steps:**

1. **Sign Up and Get API Key:** Create an account at the OpenAI platform (if you haven’t) and generate an API key from the **OpenAI Dashboard** (under your profile > “View API keys”). Copy the secret key – it starts with `sk-...` – and store it securely. *(If you already have a key, you can reuse it. Ensure you have an active billing plan if required, since new accounts may need to set up billing to use the API beyond the free trial limits.)*

2. **Set the API Key as Environment Variable:** It’s best practice not to hard-code the key. Instead, export it as an environment variable so ReconSuite-AI or its scripts can read it. For example, on Linux/macOS use:  
   ```bash
   export OPENAI_API_KEY="your_api_key_here"
   ```  
   add that to your `~/.bashrc`/`.zshrc` for persistence. On Windows, you can set a User Environment Variable via System Settings, or in Powershell:  
   ```powershell
   [Environment]::SetEnvironmentVariable("OPENAI_API_KEY","your_api_key_here","User")
   ```  
   *(You might need to restart your terminal or log out/in for changes to take effect.)*

3. **Install OpenAI SDK (Python):** If ReconSuite-AI uses Python to call the OpenAI API, install the OpenAI Python client library. First ensure **Python 3** and **pip** are installed (`sudo apt install -y python3 python3-pip` on Linux). Then install the library:  
   ```bash
   pip install --upgrade openai
   ```  
   This provides the `openai` package for interacting with the API. (**Note:** If the ReconSuite-AI tool uses a different method or language to call the API, adjust accordingly. For example, OpenAI also provides official SDKs for Node.js, etc., but Python is common.)

4. **Test the API (Optional):** To verify your API key and library are set up, you can run a quick test. For example, in Python:  
   ```python
   import openai
   openai.api_key = "your_api_key_here"  # or rely on OPENAI_API_KEY env var
   resp = openai.Engine.list()
   print(resp)
   ```  
   This should list available engines without error. Or simply use the command-line:  
   ```bash
   curl https://api.openai.com/v1/models \
        -H "Authorization: Bearer $OPENAI_API_KEY"
   ```  
   which should return a JSON (if the key is valid).

**Updating OpenAI Tools:** The OpenAI API itself is cloud-based, so there’s no “software” to update for the service – but you should keep the client library up to date. Use `pip install -U openai` periodically to get the latest features and fixes. Check OpenAI’s API documentation for any changes in endpoints or parameters. If using environment variables, you only need to update the value if you rotate/regenerate your API key.

**Troubleshooting:** Common issues include authentication errors (`401 Unauthorized`) – double-check that your `OPENAI_API_KEY` is correctly set and has no extra quotes or spaces. If you see module import errors in Python, ensure the `openai` library is installed in the correct Python environment (use `pip show openai`). For networking issues or timeouts, ensure you have internet connectivity and that your system’s date/time is correct (SSL can fail if clock is skewed). If behind a proxy or firewall, set the `HTTP_PROXY`/`HTTPS_PROXY` environment variables so the OpenAI library or curl can reach the API.

## 4. Install **gf** (Gf Patterns)

**gf** (by Tomnomnom) is a tool that wraps `grep` with pre-defined patterns to quickly search for potential vulnerabilities or interesting strings (like API keys, SQL injection points, etc.) in text output. We will install the gf binary and then add a collection of pattern files (**Gf-Patterns**) that ReconSuite-AI uses.

**Installation Steps:**

1. **Install the gf tool:** Ensure Go is installed and `$HOME/go/bin` is in your PATH. Then run:  
   ```bash
   go install github.com/tomnomnom/gf@latest
   ```  
   This fetches and builds gf from source and places the `gf` binary in your Go bin directory. (For older Go versions <1.17, the command would be `go get -u github.com/tomnomnom/gf`, but using `go install ...@latest` is preferred for Go 1.17+.)

2. **Enable shell autocompletion (optional):** gf comes with tab-completion scripts. If desired, source the completion script in your shell config. For bash:  
   ```bash
   echo 'source $HOME/go/pkg/mod/github.com/tomnomnom/gf@latest/gf-completion.bash' >> ~/.bashrc
   ```  
   For Zsh:  
   ```bash
   echo 'source $HOME/go/pkg/mod/github.com/tomnomnom/gf@latest/gf-completion.zsh' >> ~/.zshrc
   ```  
   *(The exact path may include the gf version; adjust accordingly. Alternatively, locate `gf-completion.bash` under `$GOPATH/pkg/mod/github.com/tomnomnom/gf*`.)* Then reload your shell.

3. **Install gf pattern files:** Clone the community patterns repository and move the JSON pattern files to your gf directory:  
   ```bash
   git clone https://github.com/1ndianl33t/Gf-Patterns.git ~/Gf-Patterns
   mkdir -p ~/.gf
   mv ~/Gf-Patterns/*.json ~/.gf/
   ```  
   This gives gf a wide range of patterns (SQLi, XSS, LFI, SSRF, etc.) beyond the default examples. Once done, you can remove the cloned `Gf-Patterns` folder if desired.

4. **Verify gf:** Run `gf -h` to see the help text. You can test a pattern, for example:  
   ```bash
   echo "test.php?file=etc/passwd" | gf lfi
   ```  
   If the pattern file is set up, gf will output the matching line (indicating it detected a potential LFI parameter).

**Updating gf:** To update the gf binary to the latest version, rerun the install command:  
```bash
go install github.com/tomnomnom/gf@latest
```  
This will fetch any new updates. To update your pattern files, navigate to the `Gf-Patterns` cloned directory and do `git pull` to get the latest patterns, then copy any new `.json` files into `~/.gf`. (If you removed the local repo, you can clone again or use `wget` to fetch specific new patterns.) Consider periodically checking the Gf-Patterns repo for updates or new patterns.

**Troubleshooting:** If `gf` command is not found, double-check that `~/go/bin` is in your PATH and that your current shell has loaded the updated PATH (open a new terminal or `source ~/.bashrc`). If gf runs but returns no matches where you expect some, ensure the pattern files are in `~/.gf` and named correctly (the pattern name you use with gf corresponds to the filename, e.g. `xss` for `xss.json`). You can list available patterns with `ls ~/.gf` to confirm. 

*Note:* If you use **Oh-My-Zsh**, be aware it sets an alias `gf` = `git fetch` by default. This conflicts with the gf tool. To resolve this, remove or override that alias (e.g., add `unalias gf