# Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The workspace created time. | 
**DatabaseName** | Pointer to **string** | The name of the database. | [optional] 
**Handle** | **string** | The handle name for the workspace. | 
**Hive** | Pointer to **string** | The database hive for this workspace. | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | **string** | The unique identifier for the workspace. | 
**Identity** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**IdentityId** | **string** | The unique identifier for an identity where the workspace is created. | 
**PublicKey** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The workspace updated time. | [optional] 
**VersionId** | **int32** | The current version ID for the workspace. | 
**WorkspaceState** | Pointer to **string** | The current state of the workspace. | [optional] 

## Methods

### NewWorkspace

`func NewWorkspace(createdAt string, handle string, id string, identityId string, versionId int32, ) *Workspace`

NewWorkspace instantiates a new Workspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWithDefaults

`func NewWorkspaceWithDefaults() *Workspace`

NewWorkspaceWithDefaults instantiates a new Workspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Workspace) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Workspace) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Workspace) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDatabaseName

`func (o *Workspace) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *Workspace) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *Workspace) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *Workspace) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetHandle

`func (o *Workspace) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Workspace) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Workspace) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetHive

`func (o *Workspace) GetHive() string`

GetHive returns the Hive field if non-nil, zero value otherwise.

### GetHiveOk

`func (o *Workspace) GetHiveOk() (*string, bool)`

GetHiveOk returns a tuple with the Hive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHive

`func (o *Workspace) SetHive(v string)`

SetHive sets Hive field to given value.

### HasHive

`func (o *Workspace) HasHive() bool`

HasHive returns a boolean if a field has been set.

### GetHost

`func (o *Workspace) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Workspace) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Workspace) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Workspace) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *Workspace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workspace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workspace) SetId(v string)`

SetId sets Id field to given value.


### GetIdentity

`func (o *Workspace) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *Workspace) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *Workspace) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *Workspace) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIdentityId

`func (o *Workspace) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *Workspace) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *Workspace) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetPublicKey

`func (o *Workspace) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Workspace) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Workspace) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *Workspace) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Workspace) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Workspace) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Workspace) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Workspace) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersionId

`func (o *Workspace) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Workspace) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Workspace) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceState

`func (o *Workspace) GetWorkspaceState() string`

GetWorkspaceState returns the WorkspaceState field if non-nil, zero value otherwise.

### GetWorkspaceStateOk

`func (o *Workspace) GetWorkspaceStateOk() (*string, bool)`

GetWorkspaceStateOk returns a tuple with the WorkspaceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceState

`func (o *Workspace) SetWorkspaceState(v string)`

SetWorkspaceState sets WorkspaceState field to given value.

### HasWorkspaceState

`func (o *Workspace) HasWorkspaceState() bool`

HasWorkspaceState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


