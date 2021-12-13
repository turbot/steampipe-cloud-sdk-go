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

// TypesListWorkspacesResponse struct for TypesListWorkspacesResponse
type TypesListWorkspacesResponse struct {
	Items     *[]TypesWorkspace `json:"items,omitempty"`
	NextToken *string           `json:"next_token,omitempty"`
}

// NewTypesListWorkspacesResponse instantiates a new TypesListWorkspacesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesListWorkspacesResponse() *TypesListWorkspacesResponse {
	this := TypesListWorkspacesResponse{}
	return &this
}

// NewTypesListWorkspacesResponseWithDefaults instantiates a new TypesListWorkspacesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesListWorkspacesResponseWithDefaults() *TypesListWorkspacesResponse {
	this := TypesListWorkspacesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *TypesListWorkspacesResponse) GetItems() []TypesWorkspace {
	if o == nil || o.Items == nil {
		var ret []TypesWorkspace
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesListWorkspacesResponse) GetItemsOk() (*[]TypesWorkspace, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TypesListWorkspacesResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []TypesWorkspace and assigns it to the Items field.
func (o *TypesListWorkspacesResponse) SetItems(v []TypesWorkspace) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *TypesListWorkspacesResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesListWorkspacesResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *TypesListWorkspacesResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *TypesListWorkspacesResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o TypesListWorkspacesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableTypesListWorkspacesResponse struct {
	value *TypesListWorkspacesResponse
	isSet bool
}

func (v NullableTypesListWorkspacesResponse) Get() *TypesListWorkspacesResponse {
	return v.value
}

func (v *NullableTypesListWorkspacesResponse) Set(val *TypesListWorkspacesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesListWorkspacesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesListWorkspacesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesListWorkspacesResponse(val *TypesListWorkspacesResponse) *NullableTypesListWorkspacesResponse {
	return &NullableTypesListWorkspacesResponse{value: val, isSet: true}
}

func (v NullableTypesListWorkspacesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesListWorkspacesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
