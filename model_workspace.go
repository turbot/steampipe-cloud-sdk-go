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

// Workspace struct for Workspace
type Workspace struct {
	ApiVersion *string `json:"api_version,omitempty"`
	CliVersion *string `json:"cli_version,omitempty"`
	// The time of creation in ISO 8601 UTC.
	CreatedAt string `json:"created_at"`
	CreatedBy *User  `json:"created_by,omitempty"`
	// The ID of the user that created this.
	CreatedById string `json:"created_by_id"`
	// The name of the database.
	DatabaseName *string `json:"database_name,omitempty"`
	// The time of the last update in ISO 8601 UTC.
	DeletedAt *string `json:"deleted_at,omitempty"`
	DeletedBy *User   `json:"deleted_by,omitempty"`
	// The ID of the user that performed the deletion.
	DeletedById  string `json:"deleted_by_id"`
	DesiredState string `json:"desired_state"`
	// The handle name for the workspace.
	Handle string `json:"handle"`
	// The database hive for this workspace.
	Hive *string `json:"hive,omitempty"`
	Host *string `json:"host,omitempty"`
	// The unique identifier for the workspace.
	Id string `json:"id"`
	// The unique identifier for an identity where the workspace is created.
	IdentityId string                  `json:"identity_id"`
	Notices    *map[string]interface{} `json:"notices,omitempty"`
	PublicKey  *string                 `json:"public_key,omitempty"`
	// The current state of the workspace.
	State       *string `json:"state,omitempty"`
	StateReason *string `json:"state_reason,omitempty"`
	// The time of the last update in ISO 8601 UTC.
	UpdatedAt *string `json:"updated_at,omitempty"`
	UpdatedBy *User   `json:"updated_by,omitempty"`
	// The ID of the user that performed the last update.
	UpdatedById string `json:"updated_by_id"`
	// The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item.
	VersionId int32 `json:"version_id"`
}

// NewWorkspace instantiates a new Workspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspace(createdAt string, createdById string, deletedById string, desiredState string, handle string, id string, identityId string, updatedById string, versionId int32) *Workspace {
	this := Workspace{}
	this.CreatedAt = createdAt
	this.CreatedById = createdById
	this.DeletedById = deletedById
	this.DesiredState = desiredState
	this.Handle = handle
	this.Id = id
	this.IdentityId = identityId
	this.UpdatedById = updatedById
	this.VersionId = versionId
	return &this
}

// NewWorkspaceWithDefaults instantiates a new Workspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceWithDefaults() *Workspace {
	this := Workspace{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *Workspace) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *Workspace) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *Workspace) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetCliVersion returns the CliVersion field value if set, zero value otherwise.
func (o *Workspace) GetCliVersion() string {
	if o == nil || o.CliVersion == nil {
		var ret string
		return ret
	}
	return *o.CliVersion
}

// GetCliVersionOk returns a tuple with the CliVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetCliVersionOk() (*string, bool) {
	if o == nil || o.CliVersion == nil {
		return nil, false
	}
	return o.CliVersion, true
}

// HasCliVersion returns a boolean if a field has been set.
func (o *Workspace) HasCliVersion() bool {
	if o != nil && o.CliVersion != nil {
		return true
	}

	return false
}

// SetCliVersion gets a reference to the given string and assigns it to the CliVersion field.
func (o *Workspace) SetCliVersion(v string) {
	o.CliVersion = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Workspace) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Workspace) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Workspace) GetCreatedBy() User {
	if o == nil || o.CreatedBy == nil {
		var ret User
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetCreatedByOk() (*User, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Workspace) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given User and assigns it to the CreatedBy field.
func (o *Workspace) SetCreatedBy(v User) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value
func (o *Workspace) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *Workspace) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *Workspace) GetDatabaseName() string {
	if o == nil || o.DatabaseName == nil {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetDatabaseNameOk() (*string, bool) {
	if o == nil || o.DatabaseName == nil {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *Workspace) HasDatabaseName() bool {
	if o != nil && o.DatabaseName != nil {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *Workspace) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Workspace) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Workspace) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *Workspace) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *Workspace) GetDeletedBy() User {
	if o == nil || o.DeletedBy == nil {
		var ret User
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetDeletedByOk() (*User, bool) {
	if o == nil || o.DeletedBy == nil {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *Workspace) HasDeletedBy() bool {
	if o != nil && o.DeletedBy != nil {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given User and assigns it to the DeletedBy field.
func (o *Workspace) SetDeletedBy(v User) {
	o.DeletedBy = &v
}

// GetDeletedById returns the DeletedById field value
func (o *Workspace) GetDeletedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedById
}

// GetDeletedByIdOk returns a tuple with the DeletedById field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetDeletedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedById, true
}

// SetDeletedById sets field value
func (o *Workspace) SetDeletedById(v string) {
	o.DeletedById = v
}

// GetDesiredState returns the DesiredState field value
func (o *Workspace) GetDesiredState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetDesiredStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredState, true
}

// SetDesiredState sets field value
func (o *Workspace) SetDesiredState(v string) {
	o.DesiredState = v
}

// GetHandle returns the Handle field value
func (o *Workspace) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *Workspace) SetHandle(v string) {
	o.Handle = v
}

// GetHive returns the Hive field value if set, zero value otherwise.
func (o *Workspace) GetHive() string {
	if o == nil || o.Hive == nil {
		var ret string
		return ret
	}
	return *o.Hive
}

// GetHiveOk returns a tuple with the Hive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetHiveOk() (*string, bool) {
	if o == nil || o.Hive == nil {
		return nil, false
	}
	return o.Hive, true
}

// HasHive returns a boolean if a field has been set.
func (o *Workspace) HasHive() bool {
	if o != nil && o.Hive != nil {
		return true
	}

	return false
}

// SetHive gets a reference to the given string and assigns it to the Hive field.
func (o *Workspace) SetHive(v string) {
	o.Hive = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *Workspace) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *Workspace) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *Workspace) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value
func (o *Workspace) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Workspace) SetId(v string) {
	o.Id = v
}

// GetIdentityId returns the IdentityId field value
func (o *Workspace) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *Workspace) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetNotices returns the Notices field value if set, zero value otherwise.
func (o *Workspace) GetNotices() map[string]interface{} {
	if o == nil || o.Notices == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Notices
}

// GetNoticesOk returns a tuple with the Notices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetNoticesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Notices == nil {
		return nil, false
	}
	return o.Notices, true
}

// HasNotices returns a boolean if a field has been set.
func (o *Workspace) HasNotices() bool {
	if o != nil && o.Notices != nil {
		return true
	}

	return false
}

// SetNotices gets a reference to the given map[string]interface{} and assigns it to the Notices field.
func (o *Workspace) SetNotices(v map[string]interface{}) {
	o.Notices = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *Workspace) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *Workspace) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *Workspace) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Workspace) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Workspace) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Workspace) SetState(v string) {
	o.State = &v
}

// GetStateReason returns the StateReason field value if set, zero value otherwise.
func (o *Workspace) GetStateReason() string {
	if o == nil || o.StateReason == nil {
		var ret string
		return ret
	}
	return *o.StateReason
}

// GetStateReasonOk returns a tuple with the StateReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetStateReasonOk() (*string, bool) {
	if o == nil || o.StateReason == nil {
		return nil, false
	}
	return o.StateReason, true
}

// HasStateReason returns a boolean if a field has been set.
func (o *Workspace) HasStateReason() bool {
	if o != nil && o.StateReason != nil {
		return true
	}

	return false
}

// SetStateReason gets a reference to the given string and assigns it to the StateReason field.
func (o *Workspace) SetStateReason(v string) {
	o.StateReason = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Workspace) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Workspace) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Workspace) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Workspace) GetUpdatedBy() User {
	if o == nil || o.UpdatedBy == nil {
		var ret User
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workspace) GetUpdatedByOk() (*User, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Workspace) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given User and assigns it to the UpdatedBy field.
func (o *Workspace) SetUpdatedBy(v User) {
	o.UpdatedBy = &v
}

// GetUpdatedById returns the UpdatedById field value
func (o *Workspace) GetUpdatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetUpdatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedById, true
}

// SetUpdatedById sets field value
func (o *Workspace) SetUpdatedById(v string) {
	o.UpdatedById = v
}

// GetVersionId returns the VersionId field value
func (o *Workspace) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *Workspace) GetVersionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *Workspace) SetVersionId(v int32) {
	o.VersionId = v
}

func (o Workspace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.CliVersion != nil {
		toSerialize["cli_version"] = o.CliVersion
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if true {
		toSerialize["created_by_id"] = o.CreatedById
	}
	if o.DatabaseName != nil {
		toSerialize["database_name"] = o.DatabaseName
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.DeletedBy != nil {
		toSerialize["deleted_by"] = o.DeletedBy
	}
	if true {
		toSerialize["deleted_by_id"] = o.DeletedById
	}
	if true {
		toSerialize["desired_state"] = o.DesiredState
	}
	if true {
		toSerialize["handle"] = o.Handle
	}
	if o.Hive != nil {
		toSerialize["hive"] = o.Hive
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["identity_id"] = o.IdentityId
	}
	if o.Notices != nil {
		toSerialize["notices"] = o.Notices
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StateReason != nil {
		toSerialize["state_reason"] = o.StateReason
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if true {
		toSerialize["updated_by_id"] = o.UpdatedById
	}
	if true {
		toSerialize["version_id"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspace struct {
	value *Workspace
	isSet bool
}

func (v NullableWorkspace) Get() *Workspace {
	return v.value
}

func (v *NullableWorkspace) Set(val *Workspace) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspace) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspace(val *Workspace) *NullableWorkspace {
	return &NullableWorkspace{value: val, isSet: true}
}

func (v NullableWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
