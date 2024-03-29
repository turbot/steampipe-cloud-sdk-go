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

// InviteOrgUserRequest struct for InviteOrgUserRequest
type InviteOrgUserRequest struct {
	Email  *string `json:"email,omitempty"`
	Handle *string `json:"handle,omitempty"`
	Role   string  `json:"role"`
}

// NewInviteOrgUserRequest instantiates a new InviteOrgUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteOrgUserRequest(role string) *InviteOrgUserRequest {
	this := InviteOrgUserRequest{}
	this.Role = role
	return &this
}

// NewInviteOrgUserRequestWithDefaults instantiates a new InviteOrgUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteOrgUserRequestWithDefaults() *InviteOrgUserRequest {
	this := InviteOrgUserRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InviteOrgUserRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteOrgUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InviteOrgUserRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InviteOrgUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *InviteOrgUserRequest) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteOrgUserRequest) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *InviteOrgUserRequest) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *InviteOrgUserRequest) SetHandle(v string) {
	o.Handle = &v
}

// GetRole returns the Role field value
func (o *InviteOrgUserRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *InviteOrgUserRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *InviteOrgUserRequest) SetRole(v string) {
	o.Role = v
}

func (o InviteOrgUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if true {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableInviteOrgUserRequest struct {
	value *InviteOrgUserRequest
	isSet bool
}

func (v NullableInviteOrgUserRequest) Get() *InviteOrgUserRequest {
	return v.value
}

func (v *NullableInviteOrgUserRequest) Set(val *InviteOrgUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteOrgUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteOrgUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteOrgUserRequest(val *InviteOrgUserRequest) *NullableInviteOrgUserRequest {
	return &NullableInviteOrgUserRequest{value: val, isSet: true}
}

func (v NullableInviteOrgUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteOrgUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
