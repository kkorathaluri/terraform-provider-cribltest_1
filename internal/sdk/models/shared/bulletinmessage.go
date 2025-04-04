// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Severity string

const (
	SeverityInfo  Severity = "info"
	SeverityWarn  Severity = "warn"
	SeverityError Severity = "error"
	SeverityFatal Severity = "fatal"
)

func (e Severity) ToPointer() *Severity {
	return &e
}
func (e *Severity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "info":
		fallthrough
	case "warn":
		fallthrough
	case "error":
		fallthrough
	case "fatal":
		*e = Severity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Severity: %v", v)
	}
}

type BulletinMessageMetadata struct {
}

type BulletinMessage struct {
	ID       string                    `json:"id"`
	Severity *Severity                 `json:"severity,omitempty"`
	Title    *string                   `json:"title,omitempty"`
	Text     string                    `json:"text"`
	Time     *float64                  `json:"time,omitempty"`
	Group    *string                   `json:"group,omitempty"`
	Metadata []BulletinMessageMetadata `json:"metadata,omitempty"`
}

func (o *BulletinMessage) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *BulletinMessage) GetSeverity() *Severity {
	if o == nil {
		return nil
	}
	return o.Severity
}

func (o *BulletinMessage) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *BulletinMessage) GetText() string {
	if o == nil {
		return ""
	}
	return o.Text
}

func (o *BulletinMessage) GetTime() *float64 {
	if o == nil {
		return nil
	}
	return o.Time
}

func (o *BulletinMessage) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *BulletinMessage) GetMetadata() []BulletinMessageMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}
