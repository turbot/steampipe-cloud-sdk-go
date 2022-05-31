# WorkspaceMod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**CreatedAt** | **string** |  | 
**Details** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**IdentityId** | **string** | The unique identifier for an identity where the workspace mod has been install. | 
**InstalledVersion** | Pointer to **string** |  | [optional] 
**ModConstraint** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Workspace** | Pointer to [**Workspace**](Workspace.md) |  | [optional] 
**WorkspaceId** | **string** |  | 

## Methods

### NewWorkspaceMod

`func NewWorkspaceMod(createdAt string, id string, identityId string, workspaceId string, ) *WorkspaceMod`

NewWorkspaceMod instantiates a new WorkspaceMod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceModWithDefaults

`func NewWorkspaceModWithDefaults() *WorkspaceMod`

NewWorkspaceModWithDefaults instantiates a new WorkspaceMod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *WorkspaceMod) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *WorkspaceMod) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *WorkspaceMod) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *WorkspaceMod) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkspaceMod) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceMod) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceMod) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDetails

`func (o *WorkspaceMod) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *WorkspaceMod) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *WorkspaceMod) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *WorkspaceMod) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *WorkspaceMod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceMod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceMod) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *WorkspaceMod) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *WorkspaceMod) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *WorkspaceMod) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetInstalledVersion

`func (o *WorkspaceMod) GetInstalledVersion() string`

GetInstalledVersion returns the InstalledVersion field if non-nil, zero value otherwise.

### GetInstalledVersionOk

`func (o *WorkspaceMod) GetInstalledVersionOk() (*string, bool)`

GetInstalledVersionOk returns a tuple with the InstalledVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledVersion

`func (o *WorkspaceMod) SetInstalledVersion(v string)`

SetInstalledVersion sets InstalledVersion field to given value.

### HasInstalledVersion

`func (o *WorkspaceMod) HasInstalledVersion() bool`

HasInstalledVersion returns a boolean if a field has been set.

### GetModConstraint

`func (o *WorkspaceMod) GetModConstraint() string`

GetModConstraint returns the ModConstraint field if non-nil, zero value otherwise.

### GetModConstraintOk

`func (o *WorkspaceMod) GetModConstraintOk() (*string, bool)`

GetModConstraintOk returns a tuple with the ModConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModConstraint

`func (o *WorkspaceMod) SetModConstraint(v string)`

SetModConstraint sets ModConstraint field to given value.

### HasModConstraint

`func (o *WorkspaceMod) HasModConstraint() bool`

HasModConstraint returns a boolean if a field has been set.

### GetPath

`func (o *WorkspaceMod) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WorkspaceMod) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WorkspaceMod) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WorkspaceMod) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetState

`func (o *WorkspaceMod) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkspaceMod) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkspaceMod) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *WorkspaceMod) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkspaceMod) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceMod) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceMod) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkspaceMod) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkspace

`func (o *WorkspaceMod) GetWorkspace() Workspace`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *WorkspaceMod) GetWorkspaceOk() (*Workspace, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *WorkspaceMod) SetWorkspace(v Workspace)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *WorkspaceMod) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *WorkspaceMod) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceMod) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceMod) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


