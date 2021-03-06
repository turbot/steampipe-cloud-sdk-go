/*
Steampipe Cloud

Steampipe Cloud is a hosted version of Steampipe (https://steampipe.io), an open source tool to instantly query your cloud services (e.g. AWS, Azure, GCP and more) with SQL. No DB required.

API version: 1.0
Contact: help@steampipe.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package steampipecloud

import (
	"encoding/json"
)

// OrgQuota struct for OrgQuota
type OrgQuota struct {
	Association map[string]Quota `json:"association"`
	Conn        Quota            `json:"conn"`
	Mod         map[string]Quota `json:"mod"`
	Workspace   Quota            `json:"workspace"`
}

// NewOrgQuota instantiates a new OrgQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgQuota(association map[string]Quota, conn Quota, mod map[string]Quota, workspace Quota) *OrgQuota {
	this := OrgQuota{}
	this.Association = association
	this.Conn = conn
	this.Mod = mod
	this.Workspace = workspace
	return &this
}

// NewOrgQuotaWithDefaults instantiates a new OrgQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgQuotaWithDefaults() *OrgQuota {
	this := OrgQuota{}
	return &this
}

// GetAssociation returns the Association field value
func (o *OrgQuota) GetAssociation() map[string]Quota {
	if o == nil {
		var ret map[string]Quota
		return ret
	}

	return o.Association
}

// GetAssociationOk returns a tuple with the Association field value
// and a boolean to check if the value has been set.
func (o *OrgQuota) GetAssociationOk() (*map[string]Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Association, true
}

// SetAssociation sets field value
func (o *OrgQuota) SetAssociation(v map[string]Quota) {
	o.Association = v
}

// GetConn returns the Conn field value
func (o *OrgQuota) GetConn() Quota {
	if o == nil {
		var ret Quota
		return ret
	}

	return o.Conn
}

// GetConnOk returns a tuple with the Conn field value
// and a boolean to check if the value has been set.
func (o *OrgQuota) GetConnOk() (*Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conn, true
}

// SetConn sets field value
func (o *OrgQuota) SetConn(v Quota) {
	o.Conn = v
}

// GetMod returns the Mod field value
func (o *OrgQuota) GetMod() map[string]Quota {
	if o == nil {
		var ret map[string]Quota
		return ret
	}

	return o.Mod
}

// GetModOk returns a tuple with the Mod field value
// and a boolean to check if the value has been set.
func (o *OrgQuota) GetModOk() (*map[string]Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mod, true
}

// SetMod sets field value
func (o *OrgQuota) SetMod(v map[string]Quota) {
	o.Mod = v
}

// GetWorkspace returns the Workspace field value
func (o *OrgQuota) GetWorkspace() Quota {
	if o == nil {
		var ret Quota
		return ret
	}

	return o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value
// and a boolean to check if the value has been set.
func (o *OrgQuota) GetWorkspaceOk() (*Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workspace, true
}

// SetWorkspace sets field value
func (o *OrgQuota) SetWorkspace(v Quota) {
	o.Workspace = v
}

func (o OrgQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["association"] = o.Association
	}
	if true {
		toSerialize["conn"] = o.Conn
	}
	if true {
		toSerialize["mod"] = o.Mod
	}
	if true {
		toSerialize["workspace"] = o.Workspace
	}
	return json.Marshal(toSerialize)
}

type NullableOrgQuota struct {
	value *OrgQuota
	isSet bool
}

func (v NullableOrgQuota) Get() *OrgQuota {
	return v.value
}

func (v *NullableOrgQuota) Set(val *OrgQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgQuota(val *OrgQuota) *NullableOrgQuota {
	return &NullableOrgQuota{value: val, isSet: true}
}

func (v NullableOrgQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
