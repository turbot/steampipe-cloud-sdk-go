# TypesWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** |  | 
**DatabaseName** | Pointer to **string** |  | [optional] 
**Handle** | **string** |  | 
**Hive** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Identity** | Pointer to [**TypesIdentity**](TypesIdentity.md) |  | [optional] 
**IdentityId** | **string** |  | 
**PublicKey** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VersionId** | **int32** |  | 
**WorkspaceState** | Pointer to **string** |  | [optional] 

## Methods

### NewTypesWorkspace

`func NewTypesWorkspace(createdAt string, handle string, id string, identityId string, versionId int32, ) *TypesWorkspace`

NewTypesWorkspace instantiates a new TypesWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesWorkspaceWithDefaults

`func NewTypesWorkspaceWithDefaults() *TypesWorkspace`

NewTypesWorkspaceWithDefaults instantiates a new TypesWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TypesWorkspace) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TypesWorkspace) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TypesWorkspace) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDatabaseName

`func (o *TypesWorkspace) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *TypesWorkspace) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *TypesWorkspace) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *TypesWorkspace) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetHandle

`func (o *TypesWorkspace) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *TypesWorkspace) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *TypesWorkspace) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetHive

`func (o *TypesWorkspace) GetHive() string`

GetHive returns the Hive field if non-nil, zero value otherwise.

### GetHiveOk

`func (o *TypesWorkspace) GetHiveOk() (*string, bool)`

GetHiveOk returns a tuple with the Hive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHive

`func (o *TypesWorkspace) SetHive(v string)`

SetHive sets Hive field to given value.

### HasHive

`func (o *TypesWorkspace) HasHive() bool`

HasHive returns a boolean if a field has been set.

### GetHost

`func (o *TypesWorkspace) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TypesWorkspace) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TypesWorkspace) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *TypesWorkspace) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *TypesWorkspace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypesWorkspace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypesWorkspace) SetId(v string)`

SetId sets Id field to given value.


### GetIdentity

`func (o *TypesWorkspace) GetIdentity() TypesIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *TypesWorkspace) GetIdentityOk() (*TypesIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *TypesWorkspace) SetIdentity(v TypesIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *TypesWorkspace) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIdentityId

`func (o *TypesWorkspace) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *TypesWorkspace) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *TypesWorkspace) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetPublicKey

`func (o *TypesWorkspace) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *TypesWorkspace) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *TypesWorkspace) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *TypesWorkspace) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TypesWorkspace) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TypesWorkspace) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TypesWorkspace) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TypesWorkspace) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersionId

`func (o *TypesWorkspace) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TypesWorkspace) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TypesWorkspace) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceState

`func (o *TypesWorkspace) GetWorkspaceState() string`

GetWorkspaceState returns the WorkspaceState field if non-nil, zero value otherwise.

### GetWorkspaceStateOk

`func (o *TypesWorkspace) GetWorkspaceStateOk() (*string, bool)`

GetWorkspaceStateOk returns a tuple with the WorkspaceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceState

`func (o *TypesWorkspace) SetWorkspaceState(v string)`

SetWorkspaceState sets WorkspaceState field to given value.

### HasWorkspaceState

`func (o *TypesWorkspace) HasWorkspaceState() bool`

HasWorkspaceState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


