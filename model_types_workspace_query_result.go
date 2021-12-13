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

// TypesWorkspaceQueryResult struct for TypesWorkspaceQueryResult
type TypesWorkspaceQueryResult struct {
	Columns []TypesWorkspaceQueryResultColumn `json:"columns"`
	Rows    [][]map[string]interface{}        `json:"rows"`
}

// NewTypesWorkspaceQueryResult instantiates a new TypesWorkspaceQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesWorkspaceQueryResult(columns []TypesWorkspaceQueryResultColumn, rows [][]map[string]interface{}) *TypesWorkspaceQueryResult {
	this := TypesWorkspaceQueryResult{}
	this.Columns = columns
	this.Rows = rows
	return &this
}

// NewTypesWorkspaceQueryResultWithDefaults instantiates a new TypesWorkspaceQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesWorkspaceQueryResultWithDefaults() *TypesWorkspaceQueryResult {
	this := TypesWorkspaceQueryResult{}
	return &this
}

// GetColumns returns the Columns field value
func (o *TypesWorkspaceQueryResult) GetColumns() []TypesWorkspaceQueryResultColumn {
	if o == nil {
		var ret []TypesWorkspaceQueryResultColumn
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *TypesWorkspaceQueryResult) GetColumnsOk() (*[]TypesWorkspaceQueryResultColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value
func (o *TypesWorkspaceQueryResult) SetColumns(v []TypesWorkspaceQueryResultColumn) {
	o.Columns = v
}

// GetRows returns the Rows field value
func (o *TypesWorkspaceQueryResult) GetRows() [][]map[string]interface{} {
	if o == nil {
		var ret [][]map[string]interface{}
		return ret
	}

	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value
// and a boolean to check if the value has been set.
func (o *TypesWorkspaceQueryResult) GetRowsOk() (*[][]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rows, true
}

// SetRows sets field value
func (o *TypesWorkspaceQueryResult) SetRows(v [][]map[string]interface{}) {
	o.Rows = v
}

func (o TypesWorkspaceQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["columns"] = o.Columns
	}
	if true {
		toSerialize["rows"] = o.Rows
	}
	return json.Marshal(toSerialize)
}

type NullableTypesWorkspaceQueryResult struct {
	value *TypesWorkspaceQueryResult
	isSet bool
}

func (v NullableTypesWorkspaceQueryResult) Get() *TypesWorkspaceQueryResult {
	return v.value
}

func (v *NullableTypesWorkspaceQueryResult) Set(val *TypesWorkspaceQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesWorkspaceQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesWorkspaceQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesWorkspaceQueryResult(val *TypesWorkspaceQueryResult) *NullableTypesWorkspaceQueryResult {
	return &NullableTypesWorkspaceQueryResult{value: val, isSet: true}
}

func (v NullableTypesWorkspaceQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesWorkspaceQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
