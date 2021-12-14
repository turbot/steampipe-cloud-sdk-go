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

// OrgUser struct for OrgUser
type OrgUser struct {
	CreatedAt  *string `json:"created_at,omitempty"`
	Email      string  `json:"email"`
	Id         string  `json:"id"`
	OrgId      string  `json:"org_id"`
	Role       *string `json:"role,omitempty"`
	Status     string  `json:"status"`
	UpdatedAt  *string `json:"updated_at,omitempty"`
	User       *User   `json:"user,omitempty"`
	UserHandle string  `json:"user_handle"`
	UserId     string  `json:"user_id"`
	VersionId  *int32  `json:"version_id,omitempty"`
}

// NewOrgUser instantiates a new OrgUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgUser(email string, id string, orgId string, status string, userHandle string, userId string) *OrgUser {
	this := OrgUser{}
	this.Email = email
	this.Id = id
	this.OrgId = orgId
	this.Status = status
	this.UserHandle = userHandle
	this.UserId = userId
	return &this
}

// NewOrgUserWithDefaults instantiates a new OrgUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgUserWithDefaults() *OrgUser {
	this := OrgUser{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrgUser) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgUser) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrgUser) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *OrgUser) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value
func (o *OrgUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *OrgUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *OrgUser) SetEmail(v string) {
	o.Email = v
}

// GetId returns the Id field value
func (o *OrgUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrgUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrgUser) SetId(v string) {
	o.Id = v
}

// GetOrgId returns the OrgId field value
func (o *OrgUser) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *OrgUser) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *OrgUser) SetOrgId(v string) {
	o.OrgId = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *OrgUser) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgUser) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *OrgUser) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *OrgUser) SetRole(v string) {
	o.Role = &v
}

// GetStatus returns the Status field value
func (o *OrgUser) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrgUser) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrgUser) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrgUser) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgUser) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrgUser) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *OrgUser) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OrgUser) GetUser() User {
	if o == nil || o.User == nil {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgUser) GetUserOk() (*User, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OrgUser) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *OrgUser) SetUser(v User) {
	o.User = &v
}

// GetUserHandle returns the UserHandle field value
func (o *OrgUser) GetUserHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserHandle
}

// GetUserHandleOk returns a tuple with the UserHandle field value
// and a boolean to check if the value has been set.
func (o *OrgUser) GetUserHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserHandle, true
}

// SetUserHandle sets field value
func (o *OrgUser) SetUserHandle(v string) {
	o.UserHandle = v
}

// GetUserId returns the UserId field value
func (o *OrgUser) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *OrgUser) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *OrgUser) SetUserId(v string) {
	o.UserId = v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *OrgUser) GetVersionId() int32 {
	if o == nil || o.VersionId == nil {
		var ret int32
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgUser) GetVersionIdOk() (*int32, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *OrgUser) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given int32 and assigns it to the VersionId field.
func (o *OrgUser) SetVersionId(v int32) {
	o.VersionId = &v
}

func (o OrgUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["org_id"] = o.OrgId
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if true {
		toSerialize["user_handle"] = o.UserHandle
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if o.VersionId != nil {
		toSerialize["version_id"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableOrgUser struct {
	value *OrgUser
	isSet bool
}

func (v NullableOrgUser) Get() *OrgUser {
	return v.value
}

func (v *NullableOrgUser) Set(val *OrgUser) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgUser) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgUser(val *OrgUser) *NullableOrgUser {
	return &NullableOrgUser{value: val, isSet: true}
}

func (v NullableOrgUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}