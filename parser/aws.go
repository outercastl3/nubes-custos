package parser

import (
	"os"
	"time"
	"encoding/json"
	"fmt"
	"strings"
	"path/filepath"
)

type CloudTrailLog struct {
  Records []CloudTrailRecord `json:"Records"`
}

type CloudTrailRecord struct {
  EventVersion string `json:"eventVersion"`
  UserIdentity UserIdentity `json:"userIdentity"`
  EventTime time.Time `json:"eventTime"`
  EventSource string `json:"eventSource"`
  EventName string `json:"eventName"`
  AwsRegion string `json:"awsRegion"`
  SourceIP string `json:"sourceIPAddress"`
  UserAgent string `json:"userAgent"`
  EventType string `json:"eventType"`
  EventID string `json:"eventID"`
}

type UserIdentity struct {
  Type string `json:"type"`
  PrincipalID string `json:"principalId"`
  ARN string `json:"arn"`
  AccountID string `json:"accountId"`
  AccessKeyID string `json:"accessKeyId"`
  UserName string `json:"userName"`
  SessionContext SessionContext `json:"sessionContext"`
}

type SessionContext struct {
  Attributes SessionAttributes `json:"attributes"`
}

type SessionAttributes struct {
  MFAAuthenticated string `json:"mfaAuthenticated"`
  CreationDate string `json:"creationDate"`
}

func (r CloudTrailRecord) GetEventTime() time.Time { return r.EventTime }
func (r CloudTrailRecord) GetEventName() string { return r.EventName }
func (r CloudTrailRecord) GetSourceIP() string { return r.SourceIP }
func (r CloudTrailRecord) GetUserName() string { return r.UserIdentity.UserName }
func (r CloudTrailRecord) GetUserType() string         { return r.UserIdentity.Type }
func (r CloudTrailRecord) GetRegion() string           { return r.AwsRegion }
func (r CloudTrailRecord) GetMFAAuthenticated() bool   { return r.UserIdentity.SessionContext.Attributes.MFAAuthenticated == "true" }
func (r CloudTrailRecord) GetProvider() string         { return "aws" }



func ParseFile(path string) ([]CloudTrailRecord, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("failed to read file: %w", err)
    }
    var log CloudTrailLog
    if err := json.Unmarshal(data, &log); err != nil {
        return nil, fmt.Errorf("failed to parse JSON: %w", err)
    }
    return log.Records, nil
}

func ParseDirectory(path string) ([]CloudTrailRecord, error) {
    var allRecords []CloudTrailRecord
    entries, err := os.ReadDir(path)
    if err != nil {
        return nil, fmt.Errorf("failed to read directory: %w", err)
    }
    for _, entry := range entries {
        if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
            continue
        }
        records, err := ParseFile(filepath.Join(path, entry.Name()))
        if err != nil {
            fmt.Printf("warning: skipping %s: %v\n", entry.Name(), err)
            continue
        }
        allRecords = append(allRecords, records...)
    }
    return allRecords, nil
}
