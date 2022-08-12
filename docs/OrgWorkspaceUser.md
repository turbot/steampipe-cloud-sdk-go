# OrgWorkspaceUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**Id** | **string** | The unique identifier of the org member. | 
**OrgId** | **string** | The identifier of an org. | 
**Role** | Pointer to **string** | The role of the workspace user. | [optional] 
**Scope** | Pointer to **string** | The scope of the role. Can be either of org / workspace | [optional] 
**Status** | **string** | The status of the org member i.e pending or accepted. | 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**UserHandle** | **string** | The user handle of the member. | 
**UserId** | **string** | The identifier of a user. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceHandle** | **string** | The handle of the workspace with identifier WorkspaceID. | 
**WorkspaceId** | **string** | The identifier of a workspace belonging to the organization. | 

## Methods

### NewOrgWorkspaceUser

`func NewOrgWorkspaceUser(createdAt string, createdById string, id string, orgId string, status string, updatedById string, userHandle string, userId string, versionId int32, workspaceHandle string, workspaceId string, ) *OrgWorkspaceUser`

NewOrgWorkspaceUser instantiates a new OrgWorkspaceUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWorkspaceUserWithDefaults

`func NewOrgWorkspaceUserWithDefaults() *OrgWorkspaceUser`

NewOrgWorkspaceUserWithDefaults instantiates a new OrgWorkspaceUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrgWorkspaceUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgWorkspaceUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgWorkspaceUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *OrgWorkspaceUser) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrgWorkspaceUser) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrgWorkspaceUser) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *OrgWorkspaceUser) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *OrgWorkspaceUser) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *OrgWorkspaceUser) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *OrgWorkspaceUser) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetId

`func (o *OrgWorkspaceUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgWorkspaceUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgWorkspaceUser) SetId(v string)`

SetId sets Id field to given value.


### GetOrgId

`func (o *OrgWorkspaceUser) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrgWorkspaceUser) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrgWorkspaceUser) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetRole

`func (o *OrgWorkspaceUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrgWorkspaceUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrgWorkspaceUser) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrgWorkspaceUser) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetScope

`func (o *OrgWorkspaceUser) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OrgWorkspaceUser) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OrgWorkspaceUser) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OrgWorkspaceUser) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatus

`func (o *OrgWorkspaceUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrgWorkspaceUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrgWorkspaceUser) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *OrgWorkspaceUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrgWorkspaceUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrgWorkspaceUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrgWorkspaceUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *OrgWorkspaceUser) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *OrgWorkspaceUser) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *OrgWorkspaceUser) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *OrgWorkspaceUser) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *OrgWorkspaceUser) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *OrgWorkspaceUser) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *OrgWorkspaceUser) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetUser

`func (o *OrgWorkspaceUser) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrgWorkspaceUser) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrgWorkspaceUser) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *OrgWorkspaceUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserHandle

`func (o *OrgWorkspaceUser) GetUserHandle() string`

GetUserHandle returns the UserHandle field if non-nil, zero value otherwise.

### GetUserHandleOk

`func (o *OrgWorkspaceUser) GetUserHandleOk() (*string, bool)`

GetUserHandleOk returns a tuple with the UserHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHandle

`func (o *OrgWorkspaceUser) SetUserHandle(v string)`

SetUserHandle sets UserHandle field to given value.


### GetUserId

`func (o *OrgWorkspaceUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrgWorkspaceUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrgWorkspaceUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetVersionId

`func (o *OrgWorkspaceUser) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *OrgWorkspaceUser) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *OrgWorkspaceUser) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceHandle

`func (o *OrgWorkspaceUser) GetWorkspaceHandle() string`

GetWorkspaceHandle returns the WorkspaceHandle field if non-nil, zero value otherwise.

### GetWorkspaceHandleOk

`func (o *OrgWorkspaceUser) GetWorkspaceHandleOk() (*string, bool)`

GetWorkspaceHandleOk returns a tuple with the WorkspaceHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceHandle

`func (o *OrgWorkspaceUser) SetWorkspaceHandle(v string)`

SetWorkspaceHandle sets WorkspaceHandle field to given value.


### GetWorkspaceId

`func (o *OrgWorkspaceUser) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *OrgWorkspaceUser) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *OrgWorkspaceUser) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


