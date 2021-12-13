# TypesWorkspaceConn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | Pointer to [**TypesConnection**](TypesConnection.md) |  | [optional] 
**ConnectionId** | **string** |  | 
**CreatedAt** | **string** |  | 
**Id** | **string** |  | 
**IdentityId** | **string** |  | 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VersionId** | **int32** |  | 
**Workspace** | Pointer to [**TypesWorkspace**](TypesWorkspace.md) |  | [optional] 
**WorkspaceId** | **string** |  | 

## Methods

### NewTypesWorkspaceConn

`func NewTypesWorkspaceConn(connectionId string, createdAt string, id string, identityId string, versionId int32, workspaceId string, ) *TypesWorkspaceConn`

NewTypesWorkspaceConn instantiates a new TypesWorkspaceConn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesWorkspaceConnWithDefaults

`func NewTypesWorkspaceConnWithDefaults() *TypesWorkspaceConn`

NewTypesWorkspaceConnWithDefaults instantiates a new TypesWorkspaceConn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *TypesWorkspaceConn) GetConnection() TypesConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *TypesWorkspaceConn) GetConnectionOk() (*TypesConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *TypesWorkspaceConn) SetConnection(v TypesConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *TypesWorkspaceConn) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetConnectionId

`func (o *TypesWorkspaceConn) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *TypesWorkspaceConn) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *TypesWorkspaceConn) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetCreatedAt

`func (o *TypesWorkspaceConn) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TypesWorkspaceConn) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TypesWorkspaceConn) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *TypesWorkspaceConn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypesWorkspaceConn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypesWorkspaceConn) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *TypesWorkspaceConn) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *TypesWorkspaceConn) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *TypesWorkspaceConn) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetUpdatedAt

`func (o *TypesWorkspaceConn) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TypesWorkspaceConn) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TypesWorkspaceConn) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TypesWorkspaceConn) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersionId

`func (o *TypesWorkspaceConn) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TypesWorkspaceConn) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TypesWorkspaceConn) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspace

`func (o *TypesWorkspaceConn) GetWorkspace() TypesWorkspace`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *TypesWorkspaceConn) GetWorkspaceOk() (*TypesWorkspace, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *TypesWorkspaceConn) SetWorkspace(v TypesWorkspace)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *TypesWorkspaceConn) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *TypesWorkspaceConn) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *TypesWorkspaceConn) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *TypesWorkspaceConn) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


