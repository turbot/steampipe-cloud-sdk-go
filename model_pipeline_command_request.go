/*
Steampipe Cloud

Steampipe Cloud is a hosted version of Steampipe (https://steampipe.io), an open source tool to instantly query your cloud services (e.g. AWS, Azure, GCP and more) with SQL. No DB required.

API version: {{OPEN_API_VERSION}}
Contact: help@steampipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package steampipecloud

import (
	"encoding/json"
)

// PipelineCommandRequest struct for PipelineCommandRequest
type PipelineCommandRequest struct {
	Command string `json:"command"`
}

// NewPipelineCommandRequest instantiates a new PipelineCommandRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineCommandRequest(command string) *PipelineCommandRequest {
	this := PipelineCommandRequest{}
	this.Command = command
	return &this
}

// NewPipelineCommandRequestWithDefaults instantiates a new PipelineCommandRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineCommandRequestWithDefaults() *PipelineCommandRequest {
	this := PipelineCommandRequest{}
	return &this
}

// GetCommand returns the Command field value
func (o *PipelineCommandRequest) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *PipelineCommandRequest) GetCommandOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *PipelineCommandRequest) SetCommand(v string) {
	o.Command = v
}

func (o PipelineCommandRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["command"] = o.Command
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineCommandRequest struct {
	value *PipelineCommandRequest
	isSet bool
}

func (v NullablePipelineCommandRequest) Get() *PipelineCommandRequest {
	return v.value
}

func (v *NullablePipelineCommandRequest) Set(val *PipelineCommandRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineCommandRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineCommandRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineCommandRequest(val *PipelineCommandRequest) *NullablePipelineCommandRequest {
	return &NullablePipelineCommandRequest{value: val, isSet: true}
}

func (v NullablePipelineCommandRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineCommandRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}