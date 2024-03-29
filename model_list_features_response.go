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

// ListFeaturesResponse struct for ListFeaturesResponse
type ListFeaturesResponse struct {
	Items *[]Feature `json:"items,omitempty"`
}

// NewListFeaturesResponse instantiates a new ListFeaturesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFeaturesResponse() *ListFeaturesResponse {
	this := ListFeaturesResponse{}
	return &this
}

// NewListFeaturesResponseWithDefaults instantiates a new ListFeaturesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFeaturesResponseWithDefaults() *ListFeaturesResponse {
	this := ListFeaturesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListFeaturesResponse) GetItems() []Feature {
	if o == nil || o.Items == nil {
		var ret []Feature
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFeaturesResponse) GetItemsOk() (*[]Feature, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListFeaturesResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Feature and assigns it to the Items field.
func (o *ListFeaturesResponse) SetItems(v []Feature) {
	o.Items = &v
}

func (o ListFeaturesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableListFeaturesResponse struct {
	value *ListFeaturesResponse
	isSet bool
}

func (v NullableListFeaturesResponse) Get() *ListFeaturesResponse {
	return v.value
}

func (v *NullableListFeaturesResponse) Set(val *ListFeaturesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFeaturesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFeaturesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFeaturesResponse(val *ListFeaturesResponse) *NullableListFeaturesResponse {
	return &NullableListFeaturesResponse{value: val, isSet: true}
}

func (v NullableListFeaturesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFeaturesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
