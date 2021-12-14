# OrgUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**Id** | **string** |  | 
**OrgId** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**UserHandle** | **string** |  | 
**UserId** | **string** |  | 
**VersionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewOrgUser

`func NewOrgUser(email string, id string, orgId string, status string, userHandle string, userId string, ) *OrgUser`

NewOrgUser instantiates a new OrgUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgUserWithDefaults

`func NewOrgUserWithDefaults() *OrgUser`

NewOrgUserWithDefaults instantiates a new OrgUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrgUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrgUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *OrgUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrgUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrgUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetId

`func (o *OrgUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgUser) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *OrgUser) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgUser) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgUser) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetRole

`func (o *OrgUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrgUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrgUser) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrgUser) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *OrgUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrgUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrgUser) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *OrgUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrgUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrgUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrgUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *OrgUser) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrgUser) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrgUser) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *OrgUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserHandle

`func (o *OrgUser) GetUserHandle() string`

GetUserHandle returns the UserHandle field if non-nil, zero value otherwise.

### GetUserHandleOk

`func (o *OrgUser) GetUserHandleOk() (*string, bool)`

GetUserHandleOk returns a tuple with the UserHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHandle

`func (o *OrgUser) SetUserHandle(v string)`

SetUserHandle sets UserHandle field to given value.


### GetUserId

`func (o *OrgUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrgUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrgUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetVersionId

`func (o *OrgUser) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *OrgUser) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *OrgUser) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *OrgUser) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


