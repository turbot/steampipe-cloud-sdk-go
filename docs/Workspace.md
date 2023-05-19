# Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**CliVersion** | Pointer to **string** |  | [optional] 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DatabaseName** | Pointer to **string** | The name of the database. | [optional] 
**DeletedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) |  | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**DesiredState** | **string** |  | 
**Handle** | **string** | The handle name for the workspace. | 
**Hive** | Pointer to **string** | The database hive for this workspace. | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | **string** | The unique identifier for the workspace. | 
**IdentityId** | **string** | The unique identifier for an identity where the workspace is created. | 
**Notices** | Pointer to **map[string]interface{}** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | The current state of the workspace. | [optional] 
**StateReason** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 

## Methods

### NewWorkspace

`func NewWorkspace(createdAt string, createdById string, deletedById string, desiredState string, handle string, id string, identityId string, updatedById string, versionId int32, ) *Workspace`

NewWorkspace instantiates a new Workspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWithDefaults

`func NewWorkspaceWithDefaults() *Workspace`

NewWorkspaceWithDefaults instantiates a new Workspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Workspace) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Workspace) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Workspace) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *Workspace) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCliVersion

`func (o *Workspace) GetCliVersion() string`

GetCliVersion returns the CliVersion field if non-nil, zero value otherwise.

### GetCliVersionOk

`func (o *Workspace) GetCliVersionOk() (*string, bool)`

GetCliVersionOk returns a tuple with the CliVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliVersion

`func (o *Workspace) SetCliVersion(v string)`

SetCliVersion sets CliVersion field to given value.

### HasCliVersion

`func (o *Workspace) HasCliVersion() bool`

HasCliVersion returns a boolean if a field has been set.

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


### GetCreatedBy

`func (o *Workspace) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Workspace) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Workspace) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Workspace) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *Workspace) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Workspace) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Workspace) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


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

### GetDeletedAt

`func (o *Workspace) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Workspace) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Workspace) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Workspace) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Workspace) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Workspace) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Workspace) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Workspace) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *Workspace) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *Workspace) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *Workspace) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetDesiredState

`func (o *Workspace) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *Workspace) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *Workspace) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.


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


### GetNotices

`func (o *Workspace) GetNotices() map[string]interface{}`

GetNotices returns the Notices field if non-nil, zero value otherwise.

### GetNoticesOk

`func (o *Workspace) GetNoticesOk() (*map[string]interface{}, bool)`

GetNoticesOk returns a tuple with the Notices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotices

`func (o *Workspace) SetNotices(v map[string]interface{})`

SetNotices sets Notices field to given value.

### HasNotices

`func (o *Workspace) HasNotices() bool`

HasNotices returns a boolean if a field has been set.

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

### GetState

`func (o *Workspace) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Workspace) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Workspace) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Workspace) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *Workspace) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *Workspace) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *Workspace) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *Workspace) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

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

### GetUpdatedBy

`func (o *Workspace) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Workspace) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Workspace) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Workspace) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *Workspace) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Workspace) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Workspace) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


