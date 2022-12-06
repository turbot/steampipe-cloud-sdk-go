# WorkspaceSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DashboardName** | **string** | The mod-prefixed name of the dashboard this snapshot belongs to. | 
**DashboardTitle** | **string** | The title of the dashboard this snapshot belongs to. | 
**ExpiresAt** | Pointer to **string** | The time when the snapshot will expire. | [optional] 
**Id** | **string** | The unique identifier for the snapshot. | 
**IdentityId** | **string** | The unique identifier for the identity that the snapshot belongs to. | 
**Inputs** | Pointer to **interface{}** |  | [optional] 
**SchemaVersion** | **string** | The schema version of the underlying snapshot. | 
**State** | Pointer to **string** | The current state of the snapshot. | [optional] 
**Tags** | Pointer to **interface{}** |  | [optional] 
**Title** | Pointer to **string** | The title of the snapshot. | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**Visibility** | Pointer to **string** | The visibility of the snapshot. | [optional] 
**WorkspaceId** | **string** | The unique identifier for the workspace that the snapshot belongs to. | 

## Methods

### NewWorkspaceSnapshot

`func NewWorkspaceSnapshot(createdAt string, createdById string, dashboardName string, dashboardTitle string, id string, identityId string, schemaVersion string, updatedById string, versionId int32, workspaceId string, ) *WorkspaceSnapshot`

NewWorkspaceSnapshot instantiates a new WorkspaceSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceSnapshotWithDefaults

`func NewWorkspaceSnapshotWithDefaults() *WorkspaceSnapshot`

NewWorkspaceSnapshotWithDefaults instantiates a new WorkspaceSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WorkspaceSnapshot) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceSnapshot) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceSnapshot) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *WorkspaceSnapshot) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkspaceSnapshot) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkspaceSnapshot) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkspaceSnapshot) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *WorkspaceSnapshot) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkspaceSnapshot) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkspaceSnapshot) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDashboardName

`func (o *WorkspaceSnapshot) GetDashboardName() string`

GetDashboardName returns the DashboardName field if non-nil, zero value otherwise.

### GetDashboardNameOk

`func (o *WorkspaceSnapshot) GetDashboardNameOk() (*string, bool)`

GetDashboardNameOk returns a tuple with the DashboardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardName

`func (o *WorkspaceSnapshot) SetDashboardName(v string)`

SetDashboardName sets DashboardName field to given value.


### GetDashboardTitle

`func (o *WorkspaceSnapshot) GetDashboardTitle() string`

GetDashboardTitle returns the DashboardTitle field if non-nil, zero value otherwise.

### GetDashboardTitleOk

`func (o *WorkspaceSnapshot) GetDashboardTitleOk() (*string, bool)`

GetDashboardTitleOk returns a tuple with the DashboardTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardTitle

`func (o *WorkspaceSnapshot) SetDashboardTitle(v string)`

SetDashboardTitle sets DashboardTitle field to given value.


### GetExpiresAt

`func (o *WorkspaceSnapshot) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *WorkspaceSnapshot) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *WorkspaceSnapshot) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *WorkspaceSnapshot) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *WorkspaceSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceSnapshot) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *WorkspaceSnapshot) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *WorkspaceSnapshot) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *WorkspaceSnapshot) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetInputs

`func (o *WorkspaceSnapshot) GetInputs() interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *WorkspaceSnapshot) GetInputsOk() (*interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *WorkspaceSnapshot) SetInputs(v interface{})`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *WorkspaceSnapshot) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### SetInputsNil

`func (o *WorkspaceSnapshot) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *WorkspaceSnapshot) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil
### GetSchemaVersion

`func (o *WorkspaceSnapshot) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *WorkspaceSnapshot) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *WorkspaceSnapshot) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetState

`func (o *WorkspaceSnapshot) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkspaceSnapshot) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkspaceSnapshot) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WorkspaceSnapshot) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *WorkspaceSnapshot) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkspaceSnapshot) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkspaceSnapshot) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkspaceSnapshot) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *WorkspaceSnapshot) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *WorkspaceSnapshot) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTitle

`func (o *WorkspaceSnapshot) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkspaceSnapshot) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkspaceSnapshot) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkspaceSnapshot) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkspaceSnapshot) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceSnapshot) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceSnapshot) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkspaceSnapshot) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkspaceSnapshot) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkspaceSnapshot) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkspaceSnapshot) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkspaceSnapshot) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *WorkspaceSnapshot) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *WorkspaceSnapshot) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *WorkspaceSnapshot) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *WorkspaceSnapshot) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkspaceSnapshot) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkspaceSnapshot) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetVisibility

`func (o *WorkspaceSnapshot) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *WorkspaceSnapshot) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *WorkspaceSnapshot) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *WorkspaceSnapshot) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *WorkspaceSnapshot) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceSnapshot) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceSnapshot) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


