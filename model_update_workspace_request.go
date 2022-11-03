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

// UpdateWorkspaceRequest struct for UpdateWorkspaceRequest
type UpdateWorkspaceRequest struct {
	Handle *string `json:"handle,omitempty"`
}

// NewUpdateWorkspaceRequest instantiates a new UpdateWorkspaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWorkspaceRequest() *UpdateWorkspaceRequest {
	this := UpdateWorkspaceRequest{}
	return &this
}

// NewUpdateWorkspaceRequestWithDefaults instantiates a new UpdateWorkspaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWorkspaceRequestWithDefaults() *UpdateWorkspaceRequest {
	this := UpdateWorkspaceRequest{}
	return &this
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *UpdateWorkspaceRequest) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorkspaceRequest) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *UpdateWorkspaceRequest) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *UpdateWorkspaceRequest) SetHandle(v string) {
	o.Handle = &v
}

func (o UpdateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWorkspaceRequest struct {
	value *UpdateWorkspaceRequest
	isSet bool
}

func (v NullableUpdateWorkspaceRequest) Get() *UpdateWorkspaceRequest {
	return v.value
}

func (v *NullableUpdateWorkspaceRequest) Set(val *UpdateWorkspaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWorkspaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWorkspaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWorkspaceRequest(val *UpdateWorkspaceRequest) *NullableUpdateWorkspaceRequest {
	return &NullableUpdateWorkspaceRequest{value: val, isSet: true}
}

func (v NullableUpdateWorkspaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWorkspaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
