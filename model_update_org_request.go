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

// UpdateOrgRequest struct for UpdateOrgRequest
type UpdateOrgRequest struct {
	DisplayName      *string   `json:"display_name,omitempty"`
	Handle           *string   `json:"handle,omitempty"`
	TokenMinIssuedAt *JSONTime `json:"token_min_issued_at,omitempty"`
	Url              *string   `json:"url,omitempty"`
}

// NewUpdateOrgRequest instantiates a new UpdateOrgRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrgRequest() *UpdateOrgRequest {
	this := UpdateOrgRequest{}
	return &this
}

// NewUpdateOrgRequestWithDefaults instantiates a new UpdateOrgRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrgRequestWithDefaults() *UpdateOrgRequest {
	this := UpdateOrgRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UpdateOrgRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *UpdateOrgRequest) SetHandle(v string) {
	o.Handle = &v
}

// GetTokenMinIssuedAt returns the TokenMinIssuedAt field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetTokenMinIssuedAt() JSONTime {
	if o == nil || o.TokenMinIssuedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.TokenMinIssuedAt
}

// GetTokenMinIssuedAtOk returns a tuple with the TokenMinIssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetTokenMinIssuedAtOk() (*JSONTime, bool) {
	if o == nil || o.TokenMinIssuedAt == nil {
		return nil, false
	}
	return o.TokenMinIssuedAt, true
}

// HasTokenMinIssuedAt returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasTokenMinIssuedAt() bool {
	if o != nil && o.TokenMinIssuedAt != nil {
		return true
	}

	return false
}

// SetTokenMinIssuedAt gets a reference to the given JSONTime and assigns it to the TokenMinIssuedAt field.
func (o *UpdateOrgRequest) SetTokenMinIssuedAt(v JSONTime) {
	o.TokenMinIssuedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateOrgRequest) SetUrl(v string) {
	o.Url = &v
}

func (o UpdateOrgRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.TokenMinIssuedAt != nil {
		toSerialize["token_min_issued_at"] = o.TokenMinIssuedAt
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrgRequest struct {
	value *UpdateOrgRequest
	isSet bool
}

func (v NullableUpdateOrgRequest) Get() *UpdateOrgRequest {
	return v.value
}

func (v *NullableUpdateOrgRequest) Set(val *UpdateOrgRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrgRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrgRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrgRequest(val *UpdateOrgRequest) *NullableUpdateOrgRequest {
	return &NullableUpdateOrgRequest{value: val, isSet: true}
}

func (v NullableUpdateOrgRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrgRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
