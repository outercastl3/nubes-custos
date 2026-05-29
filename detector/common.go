package detector

import "time"

type Event interface {
  GetEventTime() time.Time
  GetEventName() string
  GetSourceIP() string
  GetUserName() string
  GetUserType() string
  GetRegion() string
  GetMFAAuthenticated() bool
  GetProvider() string
}

type Finding struct {
  Severity string
  Rule string
  Description string
  EventTime time.Time
  UserName string
  SourceIP string
  RawEvent interface{}
}

