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

// ListLogsResponse struct for ListLogsResponse
type ListLogsResponse struct {
	Items     *[]LogRecord `json:"items,omitempty"`
	NextToken *string      `json:"next_token,omitempty"`
}

// NewListLogsResponse instantiates a new ListLogsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLogsResponse() *ListLogsResponse {
	this := ListLogsResponse{}
	return &this
}

// NewListLogsResponseWithDefaults instantiates a new ListLogsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLogsResponseWithDefaults() *ListLogsResponse {
	this := ListLogsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListLogsResponse) GetItems() []LogRecord {
	if o == nil || o.Items == nil {
		var ret []LogRecord
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogsResponse) GetItemsOk() (*[]LogRecord, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListLogsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []LogRecord and assigns it to the Items field.
func (o *ListLogsResponse) SetItems(v []LogRecord) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListLogsResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLogsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListLogsResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListLogsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListLogsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListLogsResponse struct {
	value *ListLogsResponse
	isSet bool
}

func (v NullableListLogsResponse) Get() *ListLogsResponse {
	return v.value
}

func (v *NullableListLogsResponse) Set(val *ListLogsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListLogsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListLogsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLogsResponse(val *ListLogsResponse) *NullableListLogsResponse {
	return &NullableListLogsResponse{value: val, isSet: true}
}

func (v NullableListLogsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLogsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
