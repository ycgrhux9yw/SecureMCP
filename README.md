# SecureMCP

**SecureMCP** is a comprehensive security auditing tool designed to detect vulnerabilities and misconfigurations in applications using the [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction). It proactively identifies threats like OAuth token leakage, prompt injection vulnerabilities, rogue MCP servers, and tool poisoning attacks.

---

## 🛡️ Features

### OAuth Token Scanner
- Token format validation and security checks
- Expiration and scope analysis
- Storage security assessment
- Token endpoint validation
- JWT token analysis

### Prompt Injection Tester
- Multiple injection payload types
- Various injection positions testing
- Response analysis
- System prompt override detection
- Role confusion attack detection

### Authentication & Server Integrity Check
- SSL/TLS configuration validation
- Authentication method testing
- Security header verification
- Server security assessment
- HSTS and CSP validation

### Report Generation
- HTML and JSON report formats
- Vulnerability classification
- Severity assessment
- Remediation suggestions
- Summary statistics

---

## 👨‍💻 Who Should Use SecureMCP?
- AI Developers integrating MCP in applications
- Security teams securing AI model interactions
- DevSecOps engineers embedding MCP in CI/CD pipelines
- Researchers studying AI model vulnerabilities
- Security auditors assessing MCP implementations

---

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Docker (optional, for containerized deployment)
- Node.js (for dashboard UI)

### Installation

#### From Source
```bash
git clone https://github.com/makalin/SecureMCP.git
cd SecureMCP
make build
```

#### Using Docker
```bash
docker pull makalin/SecureMCP
```

### Basic Usage

#### Command Line
```bash
# Basic scan
./securemcp scan --target https://your-mcp-server.com

# Scan with specific options
./securemcp scan --target https://your-mcp-server.com \
    --scan-oauth \
    --scan-prompt-injection \
    --scan-authentication \
    --timeout 30s

# Generate HTML report
./securemcp scan --target https://your-mcp-server.com --report html

# Generate JSON report
./securemcp scan --target https://your-mcp-server.com --report json
```

#### Programmatic Usage
```go
import "github.com/makalin/SecureMCP/internal/scanner"

// Create scanner instance
scanner := scanner.NewScanner()

// Basic scan
results, err := scanner.Scan("https://your-mcp-server.com")

// Scan with options
options := &scanner.ScanOptions{
    ScanOAuth:           true,
    ScanPromptInjection: true,
    ScanAuthentication:  true,
    TestPrompt:          "your test prompt",
    Timeout:             30 * time.Second,
}
results, err := scanner.ScanWithOptions(target, options)
```

### Report Generation
```go
import "github.com/makalin/SecureMCP/internal/report"

// Create report generator
generator := report.NewReportGenerator("reports")

// Generate report
report, err := generator.GenerateReport(target, results)

// Save as HTML
err = generator.Save
```

---

## 📝 Notes

> **Personal fork:** I'm using this primarily to audit internal MCP integrations. The default timeout of `30s` works well for most cases, but I've found bumping it to `60s` helpful when scanning slower staging environments.
