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

// CreateWorkspaceConnRequest struct for CreateWorkspaceConnRequest
type CreateWorkspaceConnRequest struct {
	ConnectionHandle string `json:"connection_handle"`
}

// NewCreateWorkspaceConnRequest instantiates a new CreateWorkspaceConnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceConnRequest(connectionHandle string) *CreateWorkspaceConnRequest {
	this := CreateWorkspaceConnRequest{}
	this.ConnectionHandle = connectionHandle
	return &this
}

// NewCreateWorkspaceConnRequestWithDefaults instantiates a new CreateWorkspaceConnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceConnRequestWithDefaults() *CreateWorkspaceConnRequest {
	this := CreateWorkspaceConnRequest{}
	return &this
}

// GetConnectionHandle returns the ConnectionHandle field value
func (o *CreateWorkspaceConnRequest) GetConnectionHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionHandle
}

// GetConnectionHandleOk returns a tuple with the ConnectionHandle field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceConnRequest) GetConnectionHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionHandle, true
}

// SetConnectionHandle sets field value
func (o *CreateWorkspaceConnRequest) SetConnectionHandle(v string) {
	o.ConnectionHandle = v
}

func (o CreateWorkspaceConnRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connection_handle"] = o.ConnectionHandle
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWorkspaceConnRequest struct {
	value *CreateWorkspaceConnRequest
	isSet bool
}

func (v NullableCreateWorkspaceConnRequest) Get() *CreateWorkspaceConnRequest {
	return v.value
}

func (v *NullableCreateWorkspaceConnRequest) Set(val *CreateWorkspaceConnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceConnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceConnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceConnRequest(val *CreateWorkspaceConnRequest) *NullableCreateWorkspaceConnRequest {
	return &NullableCreateWorkspaceConnRequest{value: val, isSet: true}
}

func (v NullableCreateWorkspaceConnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceConnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
