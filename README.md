# nubes-custos
Multi-cloud audit log analyzer written in Go. Detects suspicious activity
and potential security incidents across AWS, Azure, GCP and Yandex Cloud
from a single CLI tool.

## Supported Providers
| Provider | Log Format | Status |
|---|---|---|
| AWS | CloudTrail JSON | In Progress |
| Azure | Activity Log JSON | Planned |
| GCP | Cloud Audit Logs JSON | Planned |
| Yandex Cloud | Audit Trails JSON | Planned |
| Oracle | Audit Log JSON | Planned |

## Detection Rules
| Rule | Severity | Providers |
|---|---|---|
| Root/Admin account usage | Critical | AWS, Azure, GCP |
| MFA not authenticated | High | AWS, Azure |
| Disable/Delete audit logging | Critical | All |
| Secrets access | High | AWS, GCP |
| IAM privilege escalation | High | All |
| Unexpected region activity | Medium | AWS, Azure, GCP |
| Public storage bucket | High | AWS, GCP |
| Mass deletion events | High | All |

## Usage
```bash
# AWS CloudTrail
nubes-custos aws -f cloudtrail.json

# Azure Activity Log
nubes-custos azure -f activity_log.json

# GCP Cloud Audit
nubes-custos gcp -f audit_log.json

# Yandex Cloud
nubes-custos yandex -f audit_trail.json

# Oracle Cloud
nubes-custos oracle -f audit_log.json

# output to file
nubes-custos aws -f cloudtrail.json -o report.json

# filter by user
nubes-custos aws -f cloudtrail.json -u root

# filter by time range
nubes-custos aws -f cloudtrail.json --from 2026-05-01 --to 2026-05-28

# filter by severity
nubes-custos aws -f cloudtrail.json --severity critical
```

## File Structure
| File | Description |
|---|---|
| main.go | Entrypoint and CLI setup |
| cmd/aws.go | AWS subcommand |
| cmd/azure.go | Azure subcommand |
| cmd/gcp.go | GCP subcommand |
| cmd/yandex.go | Yandex subcommand |
| cmd/oracle.go | Oracle subcommand |
| parser/aws.go | CloudTrail JSON parser |
| parser/azure.go | Azure Activity Log parser |
| parser/gcp.go | GCP Audit Log parser |
| parser/yandex.go | Yandex Audit Trail parser |
| parser/oracle.go | Oracle Audit Log parser |
| detector/rules.go | Detection rules and severity logic |
| detector/common.go | Shared detection interfaces |
| output/report.go | Terminal and JSON report formatting |

## Build
```bash
go build -o nubes-custos .
```

## Requirements
- Go 1.21+
- Sample log files or live cloud access

## Roadmap
- MITRE ATT&CK mapping per detected event
- Live S3/blob storage polling
- Webhook alerting on critical findings
- Multi-file and directory scanning

## Author
Bogdan Ermakov
