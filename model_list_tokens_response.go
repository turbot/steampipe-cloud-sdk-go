/*
Steampipe Cloud

Interrogate your CloudOps data with the simplicity and power of SQL, then share your discoveries using Steampipe Cloud.

API version: 1.0
Contact: help@steampipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package steampipecloud

import (
	"encoding/json"
)

// ListTokensResponse struct for ListTokensResponse
type ListTokensResponse struct {
	Items     *[]Token `json:"items,omitempty"`
	NextToken *string  `json:"next_token,omitempty"`
}

// NewListTokensResponse instantiates a new ListTokensResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTokensResponse() *ListTokensResponse {
	this := ListTokensResponse{}
	return &this
}

// NewListTokensResponseWithDefaults instantiates a new ListTokensResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTokensResponseWithDefaults() *ListTokensResponse {
	this := ListTokensResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListTokensResponse) GetItems() []Token {
	if o == nil || o.Items == nil {
		var ret []Token
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTokensResponse) GetItemsOk() (*[]Token, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListTokensResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Token and assigns it to the Items field.
func (o *ListTokensResponse) SetItems(v []Token) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListTokensResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTokensResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListTokensResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListTokensResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListTokensResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListTokensResponse struct {
	value *ListTokensResponse
	isSet bool
}

func (v NullableListTokensResponse) Get() *ListTokensResponse {
	return v.value
}

func (v *NullableListTokensResponse) Set(val *ListTokensResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTokensResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTokensResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTokensResponse(val *ListTokensResponse) *NullableListTokensResponse {
	return &NullableListTokensResponse{value: val, isSet: true}
}

func (v NullableListTokensResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTokensResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
