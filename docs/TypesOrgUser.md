# TypesOrgUser

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
**User** | Pointer to [**TypesUser**](TypesUser.md) |  | [optional] 
**UserHandle** | **string** |  | 
**UserId** | **string** |  | 
**VersionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewTypesOrgUser

`func NewTypesOrgUser(email string, id string, orgId string, status string, userHandle string, userId string, ) *TypesOrgUser`

NewTypesOrgUser instantiates a new TypesOrgUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesOrgUserWithDefaults

`func NewTypesOrgUserWithDefaults() *TypesOrgUser`

NewTypesOrgUserWithDefaults instantiates a new TypesOrgUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TypesOrgUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TypesOrgUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TypesOrgUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TypesOrgUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *TypesOrgUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TypesOrgUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TypesOrgUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetId

`func (o *TypesOrgUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypesOrgUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypesOrgUser) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *TypesOrgUser) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *TypesOrgUser) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *TypesOrgUser) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetRole

`func (o *TypesOrgUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TypesOrgUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TypesOrgUser) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *TypesOrgUser) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *TypesOrgUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesOrgUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesOrgUser) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *TypesOrgUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TypesOrgUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TypesOrgUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TypesOrgUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *TypesOrgUser) GetUser() TypesUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TypesOrgUser) GetUserOk() (*TypesUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TypesOrgUser) SetUser(v TypesUser)`

SetUser sets User field to given value.

### HasUser

`func (o *TypesOrgUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserHandle

`func (o *TypesOrgUser) GetUserHandle() string`

GetUserHandle returns the UserHandle field if non-nil, zero value otherwise.

### GetUserHandleOk

`func (o *TypesOrgUser) GetUserHandleOk() (*string, bool)`

GetUserHandleOk returns a tuple with the UserHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHandle

`func (o *TypesOrgUser) SetUserHandle(v string)`

SetUserHandle sets UserHandle field to given value.


### GetUserId

`func (o *TypesOrgUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TypesOrgUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TypesOrgUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetVersionId

`func (o *TypesOrgUser) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TypesOrgUser) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TypesOrgUser) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *TypesOrgUser) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


