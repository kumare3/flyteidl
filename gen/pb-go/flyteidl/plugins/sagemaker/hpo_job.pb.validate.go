// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: flyteidl/plugins/sagemaker/hpo_job.proto

package plugins

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _hpo_job_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on HyperparameterTuningObjective with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HyperparameterTuningObjective) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ObjectiveType

	// no validation rules for MetricName

	return nil
}

// HyperparameterTuningObjectiveValidationError is the validation error
// returned by HyperparameterTuningObjective.Validate if the designated
// constraints aren't met.
type HyperparameterTuningObjectiveValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HyperparameterTuningObjectiveValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HyperparameterTuningObjectiveValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HyperparameterTuningObjectiveValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HyperparameterTuningObjectiveValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HyperparameterTuningObjectiveValidationError) ErrorName() string {
	return "HyperparameterTuningObjectiveValidationError"
}

// Error satisfies the builtin error interface
func (e HyperparameterTuningObjectiveValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHyperparameterTuningObjective.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HyperparameterTuningObjectiveValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HyperparameterTuningObjectiveValidationError{}

// Validate checks the field values on HPOJob with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HPOJob) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTrainingJob()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HPOJobValidationError{
				field:  "TrainingJob",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MaxNumberOfTrainingJobs

	// no validation rules for MaxParallelTrainingJobs

	return nil
}

// HPOJobValidationError is the validation error returned by HPOJob.Validate if
// the designated constraints aren't met.
type HPOJobValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HPOJobValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HPOJobValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HPOJobValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HPOJobValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HPOJobValidationError) ErrorName() string { return "HPOJobValidationError" }

// Error satisfies the builtin error interface
func (e HPOJobValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHPOJob.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HPOJobValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HPOJobValidationError{}

// Validate checks the field values on HPOJobConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HPOJobConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHyperparameterRanges()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HPOJobConfigValidationError{
				field:  "HyperparameterRanges",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TuningStrategy

	if v, ok := interface{}(m.GetTuningObjective()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HPOJobConfigValidationError{
				field:  "TuningObjective",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TrainingJobEarlyStoppingType

	return nil
}

// HPOJobConfigValidationError is the validation error returned by
// HPOJobConfig.Validate if the designated constraints aren't met.
type HPOJobConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HPOJobConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HPOJobConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HPOJobConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HPOJobConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HPOJobConfigValidationError) ErrorName() string { return "HPOJobConfigValidationError" }

// Error satisfies the builtin error interface
func (e HPOJobConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHPOJobConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HPOJobConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HPOJobConfigValidationError{}

// Validate checks the field values on HPOJobCustom with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HPOJobCustom) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHpoJobConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HPOJobCustomValidationError{
				field:  "HpoJobConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTrainingJobTaskTemplate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HPOJobCustomValidationError{
				field:  "TrainingJobTaskTemplate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HPOJobCustomValidationError is the validation error returned by
// HPOJobCustom.Validate if the designated constraints aren't met.
type HPOJobCustomValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HPOJobCustomValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HPOJobCustomValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HPOJobCustomValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HPOJobCustomValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HPOJobCustomValidationError) ErrorName() string { return "HPOJobCustomValidationError" }

// Error satisfies the builtin error interface
func (e HPOJobCustomValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHPOJobCustom.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HPOJobCustomValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HPOJobCustomValidationError{}
