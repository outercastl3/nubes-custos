package parser

import "time"

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

