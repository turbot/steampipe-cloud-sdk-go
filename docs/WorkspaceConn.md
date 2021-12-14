# WorkspaceConn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | Pointer to [**Connection**](Connection.md) |  | [optional] 
**ConnectionId** | **string** |  | 
**CreatedAt** | **string** |  | 
**Id** | **string** |  | 
**IdentityId** | **string** |  | 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VersionId** | **int32** |  | 
**Workspace** | Pointer to [**Workspace**](Workspace.md) |  | [optional] 
**WorkspaceId** | **string** |  | 

## Methods

### NewWorkspaceConn

`func NewWorkspaceConn(connectionId string, createdAt string, id string, identityId string, versionId int32, workspaceId string, ) *WorkspaceConn`

NewWorkspaceConn instantiates a new WorkspaceConn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceConnWithDefaults

`func NewWorkspaceConnWithDefaults() *WorkspaceConn`

NewWorkspaceConnWithDefaults instantiates a new WorkspaceConn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *WorkspaceConn) GetConnection() Connection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *WorkspaceConn) GetConnectionOk() (*Connection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *WorkspaceConn) SetConnection(v Connection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *WorkspaceConn) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetConnectionId

`func (o *WorkspaceConn) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *WorkspaceConn) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *WorkspaceConn) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetCreatedAt

`func (o *WorkspaceConn) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceConn) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceConn) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *WorkspaceConn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceConn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceConn) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *WorkspaceConn) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *WorkspaceConn) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *WorkspaceConn) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetUpdatedAt

`func (o *WorkspaceConn) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceConn) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceConn) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkspaceConn) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersionId

`func (o *WorkspaceConn) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkspaceConn) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkspaceConn) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspace

`func (o *WorkspaceConn) GetWorkspace() Workspace`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *WorkspaceConn) GetWorkspaceOk() (*Workspace, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *WorkspaceConn) SetWorkspace(v Workspace)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *WorkspaceConn) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *WorkspaceConn) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceConn) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceConn) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


