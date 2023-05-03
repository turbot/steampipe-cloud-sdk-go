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

// OrgQuota struct for OrgQuota
type OrgQuota struct {
	Aggregator  map[string]Quota `json:"aggregator"`
	Association map[string]Quota `json:"association"`
	Conn        Quota            `json:"conn"`
	Member      Quota            `json:"member"`
	Mod         map[string]Quota `json:"mod"`
	Pipeline    map[string]Quota `json:"pipeline"`
	Snapshot    map[string]Quota `json:"snapshot"`
	Workspace   Quota            `json:"workspace"`
}

// NewOrgQuota instantiates a new OrgQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgQuota(aggregator map[string]Quota, association map[string]Quota, conn Quota, member Quota, mod map[string]Quota, pipeline map[string]Quota, snapshot map[string]Quota, workspace Quota) *OrgQuota {
	this := OrgQuota{}
	this.Aggregator = aggregator
	this.Association = association
	this.Conn = conn
	this.Member = member
	this.Mod = mod
	this.Pipeline = pipeline
	this.Snapshot = snapshot
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

// GetAggregator returns the Aggregator field value
func (o *OrgQuota) GetAggregator() map[string]Quota {
	if o == nil {
		var ret map[string]Quota
		return ret
	}

	return o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value
// and a boolean to check if the value has been set.
func (o *OrgQuota) GetAggregatorOk() (*map[string]Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregator, true
}

// SetAggregator sets field value
func (o *OrgQuota) SetAggregator(v map[string]Quota) {
	o.Aggregator = v
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

// GetMember returns the Member field value
func (o *OrgQuota) GetMember() Quota {
	if o == nil {
		var ret Quota
		return ret
	}

	return o.Member
}

// GetMemberOk returns a tuple with the Member field value
// and a boolean to check if the value has been set.
func (o *OrgQuota) GetMemberOk() (*Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Member, true
}

// SetMember sets field value
func (o *OrgQuota) SetMember(v Quota) {
	o.Member = v
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

// GetPipeline returns the Pipeline field value
func (o *OrgQuota) GetPipeline() map[string]Quota {
	if o == nil {
		var ret map[string]Quota
		return ret
	}

	return o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value
// and a boolean to check if the value has been set.
func (o *OrgQuota) GetPipelineOk() (*map[string]Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pipeline, true
}

// SetPipeline sets field value
func (o *OrgQuota) SetPipeline(v map[string]Quota) {
	o.Pipeline = v
}

// GetSnapshot returns the Snapshot field value
func (o *OrgQuota) GetSnapshot() map[string]Quota {
	if o == nil {
		var ret map[string]Quota
		return ret
	}

	return o.Snapshot
}

// GetSnapshotOk returns a tuple with the Snapshot field value
// and a boolean to check if the value has been set.
func (o *OrgQuota) GetSnapshotOk() (*map[string]Quota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snapshot, true
}

// SetSnapshot sets field value
func (o *OrgQuota) SetSnapshot(v map[string]Quota) {
	o.Snapshot = v
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
		toSerialize["aggregator"] = o.Aggregator
	}
	if true {
		toSerialize["association"] = o.Association
	}
	if true {
		toSerialize["conn"] = o.Conn
	}
	if true {
		toSerialize["member"] = o.Member
	}
	if true {
		toSerialize["mod"] = o.Mod
	}
	if true {
		toSerialize["pipeline"] = o.Pipeline
	}
	if true {
		toSerialize["snapshot"] = o.Snapshot
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
