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

// CreateUserPasswordRequest struct for CreateUserPasswordRequest
type CreateUserPasswordRequest struct {
	ExpiresInMinutes *int32 `json:"expires_in_minutes,omitempty"`
}

// NewCreateUserPasswordRequest instantiates a new CreateUserPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserPasswordRequest() *CreateUserPasswordRequest {
	this := CreateUserPasswordRequest{}
	return &this
}

// NewCreateUserPasswordRequestWithDefaults instantiates a new CreateUserPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserPasswordRequestWithDefaults() *CreateUserPasswordRequest {
	this := CreateUserPasswordRequest{}
	return &this
}

// GetExpiresInMinutes returns the ExpiresInMinutes field value if set, zero value otherwise.
func (o *CreateUserPasswordRequest) GetExpiresInMinutes() int32 {
	if o == nil || o.ExpiresInMinutes == nil {
		var ret int32
		return ret
	}
	return *o.ExpiresInMinutes
}

// GetExpiresInMinutesOk returns a tuple with the ExpiresInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserPasswordRequest) GetExpiresInMinutesOk() (*int32, bool) {
	if o == nil || o.ExpiresInMinutes == nil {
		return nil, false
	}
	return o.ExpiresInMinutes, true
}

// HasExpiresInMinutes returns a boolean if a field has been set.
func (o *CreateUserPasswordRequest) HasExpiresInMinutes() bool {
	if o != nil && o.ExpiresInMinutes != nil {
		return true
	}

	return false
}

// SetExpiresInMinutes gets a reference to the given int32 and assigns it to the ExpiresInMinutes field.
func (o *CreateUserPasswordRequest) SetExpiresInMinutes(v int32) {
	o.ExpiresInMinutes = &v
}

func (o CreateUserPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpiresInMinutes != nil {
		toSerialize["expires_in_minutes"] = o.ExpiresInMinutes
	}
	return json.Marshal(toSerialize)
}

type NullableCreateUserPasswordRequest struct {
	value *CreateUserPasswordRequest
	isSet bool
}

func (v NullableCreateUserPasswordRequest) Get() *CreateUserPasswordRequest {
	return v.value
}

func (v *NullableCreateUserPasswordRequest) Set(val *CreateUserPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserPasswordRequest(val *CreateUserPasswordRequest) *NullableCreateUserPasswordRequest {
	return &NullableCreateUserPasswordRequest{value: val, isSet: true}
}

func (v NullableCreateUserPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
