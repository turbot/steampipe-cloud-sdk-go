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

// JSONTime struct for JSONTime
type JSONTime struct {
	TimeTime *string `json:"time.Time,omitempty"`
}

// NewJSONTime instantiates a new JSONTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONTime() *JSONTime {
	this := JSONTime{}
	return &this
}

// NewJSONTimeWithDefaults instantiates a new JSONTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONTimeWithDefaults() *JSONTime {
	this := JSONTime{}
	return &this
}

// GetTimeTime returns the TimeTime field value if set, zero value otherwise.
func (o *JSONTime) GetTimeTime() string {
	if o == nil || o.TimeTime == nil {
		var ret string
		return ret
	}
	return *o.TimeTime
}

// GetTimeTimeOk returns a tuple with the TimeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONTime) GetTimeTimeOk() (*string, bool) {
	if o == nil || o.TimeTime == nil {
		return nil, false
	}
	return o.TimeTime, true
}

// HasTimeTime returns a boolean if a field has been set.
func (o *JSONTime) HasTimeTime() bool {
	if o != nil && o.TimeTime != nil {
		return true
	}

	return false
}

// SetTimeTime gets a reference to the given string and assigns it to the TimeTime field.
func (o *JSONTime) SetTimeTime(v string) {
	o.TimeTime = &v
}

func (o JSONTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TimeTime != nil {
		toSerialize["time.Time"] = o.TimeTime
	}
	return json.Marshal(toSerialize)
}

type NullableJSONTime struct {
	value *JSONTime
	isSet bool
}

func (v NullableJSONTime) Get() *JSONTime {
	return v.value
}

func (v *NullableJSONTime) Set(val *JSONTime) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONTime) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONTime(val *JSONTime) *NullableJSONTime {
	return &NullableJSONTime{value: val, isSet: true}
}

func (v NullableJSONTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
