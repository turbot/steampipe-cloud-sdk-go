# WorkspaceSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregator** | Pointer to [**WorkspaceAggregator**](WorkspaceAggregator.md) |  | [optional] 
**AggregatorId** | Pointer to **string** | The id of the aggregator if the schema is of type &#39;aggregator&#39;. | [optional] 
**Connection** | Pointer to [**Connection**](Connection.md) |  | [optional] 
**ConnectionId** | Pointer to **string** | The id of the connection if the schema is of type &#39;connection&#39;. | [optional] 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**Datatank** | Pointer to [**Datatank**](Datatank.md) |  | [optional] 
**DatatankId** | Pointer to **string** | The id of the datatank if the schema is of type &#39;datatank&#39;. | [optional] 
**DeletedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) |  | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**Id** | **string** | The unique identifier for the schema. | 
**IdentityId** | **string** | The unique identifier for an identity where the schema has been created. | 
**Name** | **string** | The name of the schema. | 
**Type** | Pointer to **string** | Type of schems can be one of datatank, aggregator or connection. | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceId** | **string** | The unique identifier for the workspace where the schema exists. | 

## Methods

### NewWorkspaceSchema

`func NewWorkspaceSchema(createdAt string, createdById string, deletedById string, id string, identityId string, name string, updatedById string, versionId int32, workspaceId string, ) *WorkspaceSchema`

NewWorkspaceSchema instantiates a new WorkspaceSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceSchemaWithDefaults

`func NewWorkspaceSchemaWithDefaults() *WorkspaceSchema`

NewWorkspaceSchemaWithDefaults instantiates a new WorkspaceSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregator

`func (o *WorkspaceSchema) GetAggregator() WorkspaceAggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *WorkspaceSchema) GetAggregatorOk() (*WorkspaceAggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *WorkspaceSchema) SetAggregator(v WorkspaceAggregator)`

SetAggregator sets Aggregator field to given value.

### HasAggregator

`func (o *WorkspaceSchema) HasAggregator() bool`

HasAggregator returns a boolean if a field has been set.

### GetAggregatorId

`func (o *WorkspaceSchema) GetAggregatorId() string`

GetAggregatorId returns the AggregatorId field if non-nil, zero value otherwise.

### GetAggregatorIdOk

`func (o *WorkspaceSchema) GetAggregatorIdOk() (*string, bool)`

GetAggregatorIdOk returns a tuple with the AggregatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatorId

`func (o *WorkspaceSchema) SetAggregatorId(v string)`

SetAggregatorId sets AggregatorId field to given value.

### HasAggregatorId

`func (o *WorkspaceSchema) HasAggregatorId() bool`

HasAggregatorId returns a boolean if a field has been set.

### GetConnection

`func (o *WorkspaceSchema) GetConnection() Connection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *WorkspaceSchema) GetConnectionOk() (*Connection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *WorkspaceSchema) SetConnection(v Connection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *WorkspaceSchema) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetConnectionId

`func (o *WorkspaceSchema) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *WorkspaceSchema) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *WorkspaceSchema) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *WorkspaceSchema) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkspaceSchema) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceSchema) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceSchema) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *WorkspaceSchema) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkspaceSchema) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkspaceSchema) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkspaceSchema) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *WorkspaceSchema) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkspaceSchema) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkspaceSchema) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDatatank

`func (o *WorkspaceSchema) GetDatatank() Datatank`

GetDatatank returns the Datatank field if non-nil, zero value otherwise.

### GetDatatankOk

`func (o *WorkspaceSchema) GetDatatankOk() (*Datatank, bool)`

GetDatatankOk returns a tuple with the Datatank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatank

`func (o *WorkspaceSchema) SetDatatank(v Datatank)`

SetDatatank sets Datatank field to given value.

### HasDatatank

`func (o *WorkspaceSchema) HasDatatank() bool`

HasDatatank returns a boolean if a field has been set.

### GetDatatankId

`func (o *WorkspaceSchema) GetDatatankId() string`

GetDatatankId returns the DatatankId field if non-nil, zero value otherwise.

### GetDatatankIdOk

`func (o *WorkspaceSchema) GetDatatankIdOk() (*string, bool)`

GetDatatankIdOk returns a tuple with the DatatankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatankId

`func (o *WorkspaceSchema) SetDatatankId(v string)`

SetDatatankId sets DatatankId field to given value.

### HasDatatankId

`func (o *WorkspaceSchema) HasDatatankId() bool`

HasDatatankId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *WorkspaceSchema) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WorkspaceSchema) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WorkspaceSchema) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *WorkspaceSchema) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *WorkspaceSchema) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *WorkspaceSchema) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *WorkspaceSchema) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *WorkspaceSchema) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *WorkspaceSchema) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *WorkspaceSchema) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *WorkspaceSchema) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetId

`func (o *WorkspaceSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceSchema) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *WorkspaceSchema) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *WorkspaceSchema) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *WorkspaceSchema) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetName

`func (o *WorkspaceSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceSchema) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WorkspaceSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkspaceSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkspaceSchema) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkspaceSchema) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkspaceSchema) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceSchema) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceSchema) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkspaceSchema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkspaceSchema) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkspaceSchema) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkspaceSchema) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkspaceSchema) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *WorkspaceSchema) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *WorkspaceSchema) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *WorkspaceSchema) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *WorkspaceSchema) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkspaceSchema) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkspaceSchema) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceId

`func (o *WorkspaceSchema) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceSchema) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceSchema) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


