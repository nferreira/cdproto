package media

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// PlayerID players will get an ID that is unique within the agent context.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-PlayerId
type PlayerID string

// String returns the PlayerID as string value.
func (t PlayerID) String() string {
	return string(t)
}

// Timestamp [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-Timestamp
type Timestamp float64

// Float64 returns the Timestamp as float64 value.
func (t Timestamp) Float64() float64 {
	return float64(t)
}

// PlayerMessage have one type per entry in MediaLogRecord::Type Corresponds
// to kMessage.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-PlayerMessage
type PlayerMessage struct {
	Level   PlayerMessageLevel `json:"level"` // Keep in sync with MediaLogMessageLevel We are currently keeping the message level 'error' separate from the PlayerError type because right now they represent different things, this one being a DVLOG(ERROR) style log message that gets printed based on what log level is selected in the UI, and the other is a representation of a media::PipelineStatus object. Soon however we're going to be moving away from using PipelineStatus for errors and introducing a new error type which should hopefully let us integrate the error log level into the PlayerError type.
	Message string             `json:"message"`
}

// PlayerProperty corresponds to kMediaPropertyChange.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-PlayerProperty
type PlayerProperty struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// PlayerEvent corresponds to kMediaEventTriggered.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-PlayerEvent
type PlayerEvent struct {
	Timestamp Timestamp `json:"timestamp"`
	Value     string    `json:"value"`
}

// PlayerError corresponds to kMediaError.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-PlayerError
type PlayerError struct {
	Type      PlayerErrorType `json:"type"`
	ErrorCode string          `json:"errorCode"` // When this switches to using media::Status instead of PipelineStatus we can remove "errorCode" and replace it with the fields from a Status instance. This also seems like a duplicate of the error level enum - there is a todo bug to have that level removed and use this instead. (crbug.com/1068454)
}

// PlayerMessageLevel keep in sync with MediaLogMessageLevel We are currently
// keeping the message level 'error' separate from the PlayerError type because
// right now they represent different things, this one being a DVLOG(ERROR)
// style log message that gets printed based on what log level is selected in
// the UI, and the other is a representation of a media::PipelineStatus object.
// Soon however we're going to be moving away from using PipelineStatus for
// errors and introducing a new error type which should hopefully let us
// integrate the error log level into the PlayerError type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-PlayerMessage
type PlayerMessageLevel string

// String returns the PlayerMessageLevel as string value.
func (t PlayerMessageLevel) String() string {
	return string(t)
}

// PlayerMessageLevel values.
const (
	PlayerMessageLevelError   PlayerMessageLevel = "error"
	PlayerMessageLevelWarning PlayerMessageLevel = "warning"
	PlayerMessageLevelInfo    PlayerMessageLevel = "info"
	PlayerMessageLevelDebug   PlayerMessageLevel = "debug"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PlayerMessageLevel) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PlayerMessageLevel) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PlayerMessageLevel) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch PlayerMessageLevel(in.String()) {
	case PlayerMessageLevelError:
		*t = PlayerMessageLevelError
	case PlayerMessageLevelWarning:
		*t = PlayerMessageLevelWarning
	case PlayerMessageLevelInfo:
		*t = PlayerMessageLevelInfo
	case PlayerMessageLevelDebug:
		*t = PlayerMessageLevelDebug

	default:
		in.AddError(errors.New("unknown PlayerMessageLevel value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PlayerMessageLevel) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// PlayerErrorType [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Media#type-PlayerError
type PlayerErrorType string

// String returns the PlayerErrorType as string value.
func (t PlayerErrorType) String() string {
	return string(t)
}

// PlayerErrorType values.
const (
	PlayerErrorTypePipelineError PlayerErrorType = "pipeline_error"
	PlayerErrorTypeMediaError    PlayerErrorType = "media_error"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PlayerErrorType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PlayerErrorType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PlayerErrorType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch PlayerErrorType(in.String()) {
	case PlayerErrorTypePipelineError:
		*t = PlayerErrorTypePipelineError
	case PlayerErrorTypeMediaError:
		*t = PlayerErrorTypeMediaError

	default:
		in.AddError(errors.New("unknown PlayerErrorType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PlayerErrorType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
