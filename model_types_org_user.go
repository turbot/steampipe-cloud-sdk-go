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

// TypesOrgUser struct for TypesOrgUser
type TypesOrgUser struct {
	CreatedAt  *string    `json:"created_at,omitempty"`
	Email      string     `json:"email"`
	Id         string     `json:"id"`
	OrgId      string     `json:"org_id"`
	Role       *string    `json:"role,omitempty"`
	Status     string     `json:"status"`
	UpdatedAt  *string    `json:"updated_at,omitempty"`
	User       *TypesUser `json:"user,omitempty"`
	UserHandle string     `json:"user_handle"`
	UserId     string     `json:"user_id"`
	VersionId  *int32     `json:"version_id,omitempty"`
}

// NewTypesOrgUser instantiates a new TypesOrgUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypesOrgUser(email string, id string, orgId string, status string, userHandle string, userId string) *TypesOrgUser {
	this := TypesOrgUser{}
	this.Email = email
	this.Id = id
	this.OrgId = orgId
	this.Status = status
	this.UserHandle = userHandle
	this.UserId = userId
	return &this
}

// NewTypesOrgUserWithDefaults instantiates a new TypesOrgUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypesOrgUserWithDefaults() *TypesOrgUser {
	this := TypesOrgUser{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TypesOrgUser) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesOrgUser) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TypesOrgUser) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TypesOrgUser) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetEmail returns the Email field value
func (o *TypesOrgUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *TypesOrgUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *TypesOrgUser) SetEmail(v string) {
	o.Email = v
}

// GetId returns the Id field value
func (o *TypesOrgUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TypesOrgUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TypesOrgUser) SetId(v string) {
	o.Id = v
}

// GetOrgId returns the OrgId field value
func (o *TypesOrgUser) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *TypesOrgUser) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *TypesOrgUser) SetOrgId(v string) {
	o.OrgId = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *TypesOrgUser) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesOrgUser) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *TypesOrgUser) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *TypesOrgUser) SetRole(v string) {
	o.Role = &v
}

// GetStatus returns the Status field value
func (o *TypesOrgUser) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TypesOrgUser) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TypesOrgUser) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TypesOrgUser) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesOrgUser) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TypesOrgUser) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TypesOrgUser) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *TypesOrgUser) GetUser() TypesUser {
	if o == nil || o.User == nil {
		var ret TypesUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesOrgUser) GetUserOk() (*TypesUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *TypesOrgUser) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given TypesUser and assigns it to the User field.
func (o *TypesOrgUser) SetUser(v TypesUser) {
	o.User = &v
}

// GetUserHandle returns the UserHandle field value
func (o *TypesOrgUser) GetUserHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserHandle
}

// GetUserHandleOk returns a tuple with the UserHandle field value
// and a boolean to check if the value has been set.
func (o *TypesOrgUser) GetUserHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserHandle, true
}

// SetUserHandle sets field value
func (o *TypesOrgUser) SetUserHandle(v string) {
	o.UserHandle = v
}

// GetUserId returns the UserId field value
func (o *TypesOrgUser) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *TypesOrgUser) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *TypesOrgUser) SetUserId(v string) {
	o.UserId = v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *TypesOrgUser) GetVersionId() int32 {
	if o == nil || o.VersionId == nil {
		var ret int32
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypesOrgUser) GetVersionIdOk() (*int32, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *TypesOrgUser) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given int32 and assigns it to the VersionId field.
func (o *TypesOrgUser) SetVersionId(v int32) {
	o.VersionId = &v
}

func (o TypesOrgUser) MarshalJSON() ([]byte, error) {
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

type NullableTypesOrgUser struct {
	value *TypesOrgUser
	isSet bool
}

func (v NullableTypesOrgUser) Get() *TypesOrgUser {
	return v.value
}

func (v *NullableTypesOrgUser) Set(val *TypesOrgUser) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesOrgUser) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesOrgUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesOrgUser(val *TypesOrgUser) *NullableTypesOrgUser {
	return &NullableTypesOrgUser{value: val, isSet: true}
}

func (v NullableTypesOrgUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesOrgUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
