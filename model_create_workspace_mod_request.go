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

// CreateWorkspaceModRequest struct for CreateWorkspaceModRequest
type CreateWorkspaceModRequest struct {
	Constraint *string `json:"constraint,omitempty"`
	Path       string  `json:"path"`
}

// NewCreateWorkspaceModRequest instantiates a new CreateWorkspaceModRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceModRequest(path string) *CreateWorkspaceModRequest {
	this := CreateWorkspaceModRequest{}
	this.Path = path
	return &this
}

// NewCreateWorkspaceModRequestWithDefaults instantiates a new CreateWorkspaceModRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceModRequestWithDefaults() *CreateWorkspaceModRequest {
	this := CreateWorkspaceModRequest{}
	return &this
}

// GetConstraint returns the Constraint field value if set, zero value otherwise.
func (o *CreateWorkspaceModRequest) GetConstraint() string {
	if o == nil || o.Constraint == nil {
		var ret string
		return ret
	}
	return *o.Constraint
}

// GetConstraintOk returns a tuple with the Constraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceModRequest) GetConstraintOk() (*string, bool) {
	if o == nil || o.Constraint == nil {
		return nil, false
	}
	return o.Constraint, true
}

// HasConstraint returns a boolean if a field has been set.
func (o *CreateWorkspaceModRequest) HasConstraint() bool {
	if o != nil && o.Constraint != nil {
		return true
	}

	return false
}

// SetConstraint gets a reference to the given string and assigns it to the Constraint field.
func (o *CreateWorkspaceModRequest) SetConstraint(v string) {
	o.Constraint = &v
}

// GetPath returns the Path field value
func (o *CreateWorkspaceModRequest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceModRequest) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *CreateWorkspaceModRequest) SetPath(v string) {
	o.Path = v
}

func (o CreateWorkspaceModRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Constraint != nil {
		toSerialize["constraint"] = o.Constraint
	}
	if true {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWorkspaceModRequest struct {
	value *CreateWorkspaceModRequest
	isSet bool
}

func (v NullableCreateWorkspaceModRequest) Get() *CreateWorkspaceModRequest {
	return v.value
}

func (v *NullableCreateWorkspaceModRequest) Set(val *CreateWorkspaceModRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceModRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceModRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceModRequest(val *CreateWorkspaceModRequest) *NullableCreateWorkspaceModRequest {
	return &NullableCreateWorkspaceModRequest{value: val, isSet: true}
}

func (v NullableCreateWorkspaceModRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceModRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
