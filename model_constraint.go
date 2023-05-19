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

// Constraint struct for Constraint
type Constraint struct {
	// The status of the setting requested for the constraint at this level.
	Approval *string `json:"approval,omitempty"`
	// The datatank id for which the constraint applies to.
	DatatankId *string `json:"datatank_id,omitempty"`
	// A short description of the constraint.
	Description *string     `json:"description,omitempty"`
	Limit       interface{} `json:"limit,omitempty"`
	// The name of the constraint.
	Name string `json:"name"`
	// The org id for which the constraint applies to.
	OrgId *string `json:"org_id,omitempty"`
	// The pipeline id for which the constraint applies to.
	PipelineId *string `json:"pipeline_id,omitempty"`
	// The plan id for which the constraint applies to.
	PlanId     *string     `json:"plan_id,omitempty"`
	Setting    interface{} `json:"setting,omitempty"`
	Statistics interface{} `json:"statistics,omitempty"`
	// The resource which the constraint targets.
	Target string `json:"target"`
	// The type of constraint. Can be one of rate, quota, range, feature or constant.
	Type string `json:"type"`
	// The user id for which the constraint applies to.
	UserId *string     `json:"user_id,omitempty"`
	Value  interface{} `json:"value"`
	// The id of the item from which the value is inherited.
	ValueFrom string `json:"value_from"`
	// The workspace id for which the constraint applies to.
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

// NewConstraint instantiates a new Constraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstraint(name string, target string, type_ string, value interface{}, valueFrom string) *Constraint {
	this := Constraint{}
	this.Name = name
	this.Target = target
	this.Type = type_
	this.Value = value
	this.ValueFrom = valueFrom
	return &this
}

// NewConstraintWithDefaults instantiates a new Constraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstraintWithDefaults() *Constraint {
	this := Constraint{}
	return &this
}

// GetApproval returns the Approval field value if set, zero value otherwise.
func (o *Constraint) GetApproval() string {
	if o == nil || o.Approval == nil {
		var ret string
		return ret
	}
	return *o.Approval
}

// GetApprovalOk returns a tuple with the Approval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetApprovalOk() (*string, bool) {
	if o == nil || o.Approval == nil {
		return nil, false
	}
	return o.Approval, true
}

// HasApproval returns a boolean if a field has been set.
func (o *Constraint) HasApproval() bool {
	if o != nil && o.Approval != nil {
		return true
	}

	return false
}

// SetApproval gets a reference to the given string and assigns it to the Approval field.
func (o *Constraint) SetApproval(v string) {
	o.Approval = &v
}

// GetDatatankId returns the DatatankId field value if set, zero value otherwise.
func (o *Constraint) GetDatatankId() string {
	if o == nil || o.DatatankId == nil {
		var ret string
		return ret
	}
	return *o.DatatankId
}

// GetDatatankIdOk returns a tuple with the DatatankId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetDatatankIdOk() (*string, bool) {
	if o == nil || o.DatatankId == nil {
		return nil, false
	}
	return o.DatatankId, true
}

// HasDatatankId returns a boolean if a field has been set.
func (o *Constraint) HasDatatankId() bool {
	if o != nil && o.DatatankId != nil {
		return true
	}

	return false
}

// SetDatatankId gets a reference to the given string and assigns it to the DatatankId field.
func (o *Constraint) SetDatatankId(v string) {
	o.DatatankId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Constraint) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Constraint) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Constraint) SetDescription(v string) {
	o.Description = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Constraint) GetLimit() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Constraint) GetLimitOk() (*interface{}, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return &o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *Constraint) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given interface{} and assigns it to the Limit field.
func (o *Constraint) SetLimit(v interface{}) {
	o.Limit = v
}

// GetName returns the Name field value
func (o *Constraint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Constraint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Constraint) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Constraint) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Constraint) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Constraint) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPipelineId returns the PipelineId field value if set, zero value otherwise.
func (o *Constraint) GetPipelineId() string {
	if o == nil || o.PipelineId == nil {
		var ret string
		return ret
	}
	return *o.PipelineId
}

// GetPipelineIdOk returns a tuple with the PipelineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetPipelineIdOk() (*string, bool) {
	if o == nil || o.PipelineId == nil {
		return nil, false
	}
	return o.PipelineId, true
}

// HasPipelineId returns a boolean if a field has been set.
func (o *Constraint) HasPipelineId() bool {
	if o != nil && o.PipelineId != nil {
		return true
	}

	return false
}

// SetPipelineId gets a reference to the given string and assigns it to the PipelineId field.
func (o *Constraint) SetPipelineId(v string) {
	o.PipelineId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *Constraint) GetPlanId() string {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetPlanIdOk() (*string, bool) {
	if o == nil || o.PlanId == nil {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *Constraint) HasPlanId() bool {
	if o != nil && o.PlanId != nil {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *Constraint) SetPlanId(v string) {
	o.PlanId = &v
}

// GetSetting returns the Setting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Constraint) GetSetting() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Setting
}

// GetSettingOk returns a tuple with the Setting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Constraint) GetSettingOk() (*interface{}, bool) {
	if o == nil || o.Setting == nil {
		return nil, false
	}
	return &o.Setting, true
}

// HasSetting returns a boolean if a field has been set.
func (o *Constraint) HasSetting() bool {
	if o != nil && o.Setting != nil {
		return true
	}

	return false
}

// SetSetting gets a reference to the given interface{} and assigns it to the Setting field.
func (o *Constraint) SetSetting(v interface{}) {
	o.Setting = v
}

// GetStatistics returns the Statistics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Constraint) GetStatistics() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Constraint) GetStatisticsOk() (*interface{}, bool) {
	if o == nil || o.Statistics == nil {
		return nil, false
	}
	return &o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *Constraint) HasStatistics() bool {
	if o != nil && o.Statistics != nil {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given interface{} and assigns it to the Statistics field.
func (o *Constraint) SetStatistics(v interface{}) {
	o.Statistics = v
}

// GetTarget returns the Target field value
func (o *Constraint) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *Constraint) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *Constraint) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value
func (o *Constraint) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Constraint) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Constraint) SetType(v string) {
	o.Type = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Constraint) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Constraint) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Constraint) SetUserId(v string) {
	o.UserId = &v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Constraint) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Constraint) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Constraint) SetValue(v interface{}) {
	o.Value = v
}

// GetValueFrom returns the ValueFrom field value
func (o *Constraint) GetValueFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValueFrom
}

// GetValueFromOk returns a tuple with the ValueFrom field value
// and a boolean to check if the value has been set.
func (o *Constraint) GetValueFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueFrom, true
}

// SetValueFrom sets field value
func (o *Constraint) SetValueFrom(v string) {
	o.ValueFrom = v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *Constraint) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *Constraint) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *Constraint) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o Constraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Approval != nil {
		toSerialize["approval"] = o.Approval
	}
	if o.DatatankId != nil {
		toSerialize["datatank_id"] = o.DatatankId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.PipelineId != nil {
		toSerialize["pipeline_id"] = o.PipelineId
	}
	if o.PlanId != nil {
		toSerialize["plan_id"] = o.PlanId
	}
	if o.Setting != nil {
		toSerialize["setting"] = o.Setting
	}
	if o.Statistics != nil {
		toSerialize["statistics"] = o.Statistics
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["value_from"] = o.ValueFrom
	}
	if o.WorkspaceId != nil {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableConstraint struct {
	value *Constraint
	isSet bool
}

func (v NullableConstraint) Get() *Constraint {
	return v.value
}

func (v *NullableConstraint) Set(val *Constraint) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraint(val *Constraint) *NullableConstraint {
	return &NullableConstraint{value: val, isSet: true}
}

func (v NullableConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
