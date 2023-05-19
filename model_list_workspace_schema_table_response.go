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

// ListWorkspaceSchemaTableResponse struct for ListWorkspaceSchemaTableResponse
type ListWorkspaceSchemaTableResponse struct {
	Items     *[]WorkspaceSchemaTable `json:"items,omitempty"`
	NextToken *string                 `json:"next_token,omitempty"`
}

// NewListWorkspaceSchemaTableResponse instantiates a new ListWorkspaceSchemaTableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorkspaceSchemaTableResponse() *ListWorkspaceSchemaTableResponse {
	this := ListWorkspaceSchemaTableResponse{}
	return &this
}

// NewListWorkspaceSchemaTableResponseWithDefaults instantiates a new ListWorkspaceSchemaTableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorkspaceSchemaTableResponseWithDefaults() *ListWorkspaceSchemaTableResponse {
	this := ListWorkspaceSchemaTableResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListWorkspaceSchemaTableResponse) GetItems() []WorkspaceSchemaTable {
	if o == nil || o.Items == nil {
		var ret []WorkspaceSchemaTable
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkspaceSchemaTableResponse) GetItemsOk() (*[]WorkspaceSchemaTable, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListWorkspaceSchemaTableResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []WorkspaceSchemaTable and assigns it to the Items field.
func (o *ListWorkspaceSchemaTableResponse) SetItems(v []WorkspaceSchemaTable) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListWorkspaceSchemaTableResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkspaceSchemaTableResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListWorkspaceSchemaTableResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListWorkspaceSchemaTableResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListWorkspaceSchemaTableResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListWorkspaceSchemaTableResponse struct {
	value *ListWorkspaceSchemaTableResponse
	isSet bool
}

func (v NullableListWorkspaceSchemaTableResponse) Get() *ListWorkspaceSchemaTableResponse {
	return v.value
}

func (v *NullableListWorkspaceSchemaTableResponse) Set(val *ListWorkspaceSchemaTableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorkspaceSchemaTableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorkspaceSchemaTableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorkspaceSchemaTableResponse(val *ListWorkspaceSchemaTableResponse) *NullableListWorkspaceSchemaTableResponse {
	return &NullableListWorkspaceSchemaTableResponse{value: val, isSet: true}
}

func (v NullableListWorkspaceSchemaTableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorkspaceSchemaTableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
