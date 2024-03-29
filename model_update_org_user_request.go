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

// UpdateOrgUserRequest struct for UpdateOrgUserRequest
type UpdateOrgUserRequest struct {
	Role string `json:"role"`
}

// NewUpdateOrgUserRequest instantiates a new UpdateOrgUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrgUserRequest(role string) *UpdateOrgUserRequest {
	this := UpdateOrgUserRequest{}
	this.Role = role
	return &this
}

// NewUpdateOrgUserRequestWithDefaults instantiates a new UpdateOrgUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrgUserRequestWithDefaults() *UpdateOrgUserRequest {
	this := UpdateOrgUserRequest{}
	return &this
}

// GetRole returns the Role field value
func (o *UpdateOrgUserRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UpdateOrgUserRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UpdateOrgUserRequest) SetRole(v string) {
	o.Role = v
}

func (o UpdateOrgUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrgUserRequest struct {
	value *UpdateOrgUserRequest
	isSet bool
}

func (v NullableUpdateOrgUserRequest) Get() *UpdateOrgUserRequest {
	return v.value
}

func (v *NullableUpdateOrgUserRequest) Set(val *UpdateOrgUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrgUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrgUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrgUserRequest(val *UpdateOrgUserRequest) *NullableUpdateOrgUserRequest {
	return &NullableUpdateOrgUserRequest{value: val, isSet: true}
}

func (v NullableUpdateOrgUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrgUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
