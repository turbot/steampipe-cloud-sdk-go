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

// TypesListOrgUsersResponse struct for TypesListOrgUsersResponse
type TypesListOrgUsersResponse struct {
	Items     *[]TypesOrgUser `json:"items,omitempty"`
	NextToken *string         `json:"next_token,omitempty"`
}

// NewTypesListOrgUsersResponse instantiates a new TypesListOrgUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesListOrgUsersResponse() *TypesListOrgUsersResponse {
	this := TypesListOrgUsersResponse{}
	return &this
}

// NewTypesListOrgUsersResponseWithDefaults instantiates a new TypesListOrgUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesListOrgUsersResponseWithDefaults() *TypesListOrgUsersResponse {
	this := TypesListOrgUsersResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *TypesListOrgUsersResponse) GetItems() []TypesOrgUser {
	if o == nil || o.Items == nil {
		var ret []TypesOrgUser
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesListOrgUsersResponse) GetItemsOk() (*[]TypesOrgUser, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TypesListOrgUsersResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []TypesOrgUser and assigns it to the Items field.
func (o *TypesListOrgUsersResponse) SetItems(v []TypesOrgUser) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *TypesListOrgUsersResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesListOrgUsersResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *TypesListOrgUsersResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *TypesListOrgUsersResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o TypesListOrgUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableTypesListOrgUsersResponse struct {
	value *TypesListOrgUsersResponse
	isSet bool
}

func (v NullableTypesListOrgUsersResponse) Get() *TypesListOrgUsersResponse {
	return v.value
}

func (v *NullableTypesListOrgUsersResponse) Set(val *TypesListOrgUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesListOrgUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesListOrgUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesListOrgUsersResponse(val *TypesListOrgUsersResponse) *NullableTypesListOrgUsersResponse {
	return &NullableTypesListOrgUsersResponse{value: val, isSet: true}
}

func (v NullableTypesListOrgUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesListOrgUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
