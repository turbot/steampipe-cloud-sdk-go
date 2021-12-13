# TypesLogRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorAvatarUrl** | **string** |  | 
**ActorDisplayName** | **string** |  | 
**ActorHandle** | **string** |  | 
**ActorId** | **string** |  | 
**CreatedAt** | **string** |  | 
**Duration** | Pointer to **int32** |  | [optional] 
**Id** | **string** |  | 
**LogTimestamp** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**WorkspaceHandle** | **string** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewTypesLogRecord

`func NewTypesLogRecord(actorAvatarUrl string, actorDisplayName string, actorHandle string, actorId string, createdAt string, id string, workspaceHandle string, workspaceId string, ) *TypesLogRecord`

NewTypesLogRecord instantiates a new TypesLogRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesLogRecordWithDefaults

`func NewTypesLogRecordWithDefaults() *TypesLogRecord`

NewTypesLogRecordWithDefaults instantiates a new TypesLogRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorAvatarUrl

`func (o *TypesLogRecord) GetActorAvatarUrl() string`

GetActorAvatarUrl returns the ActorAvatarUrl field if non-nil, zero value otherwise.

### GetActorAvatarUrlOk

`func (o *TypesLogRecord) GetActorAvatarUrlOk() (*string, bool)`

GetActorAvatarUrlOk returns a tuple with the ActorAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorAvatarUrl

`func (o *TypesLogRecord) SetActorAvatarUrl(v string)`

SetActorAvatarUrl sets ActorAvatarUrl field to given value.


### GetActorDisplayName

`func (o *TypesLogRecord) GetActorDisplayName() string`

GetActorDisplayName returns the ActorDisplayName field if non-nil, zero value otherwise.

### GetActorDisplayNameOk

`func (o *TypesLogRecord) GetActorDisplayNameOk() (*string, bool)`

GetActorDisplayNameOk returns a tuple with the ActorDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorDisplayName

`func (o *TypesLogRecord) SetActorDisplayName(v string)`

SetActorDisplayName sets ActorDisplayName field to given value.


### GetActorHandle

`func (o *TypesLogRecord) GetActorHandle() string`

GetActorHandle returns the ActorHandle field if non-nil, zero value otherwise.

### GetActorHandleOk

`func (o *TypesLogRecord) GetActorHandleOk() (*string, bool)`

GetActorHandleOk returns a tuple with the ActorHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorHandle

`func (o *TypesLogRecord) SetActorHandle(v string)`

SetActorHandle sets ActorHandle field to given value.


### GetActorId

`func (o *TypesLogRecord) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *TypesLogRecord) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *TypesLogRecord) SetActorId(v string)`

SetActorId sets ActorId field to given value.


### GetCreatedAt

`func (o *TypesLogRecord) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TypesLogRecord) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TypesLogRecord) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDuration

`func (o *TypesLogRecord) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TypesLogRecord) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TypesLogRecord) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TypesLogRecord) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *TypesLogRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypesLogRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypesLogRecord) SetId(v string)`

SetId sets Id field to given value.


### GetLogTimestamp

`func (o *TypesLogRecord) GetLogTimestamp() string`

GetLogTimestamp returns the LogTimestamp field if non-nil, zero value otherwise.

### GetLogTimestampOk

`func (o *TypesLogRecord) GetLogTimestampOk() (*string, bool)`

GetLogTimestampOk returns a tuple with the LogTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTimestamp

`func (o *TypesLogRecord) SetLogTimestamp(v string)`

SetLogTimestamp sets LogTimestamp field to given value.

### HasLogTimestamp

`func (o *TypesLogRecord) HasLogTimestamp() bool`

HasLogTimestamp returns a boolean if a field has been set.

### GetQuery

`func (o *TypesLogRecord) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TypesLogRecord) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TypesLogRecord) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *TypesLogRecord) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetWorkspaceHandle

`func (o *TypesLogRecord) GetWorkspaceHandle() string`

GetWorkspaceHandle returns the WorkspaceHandle field if non-nil, zero value otherwise.

### GetWorkspaceHandleOk

`func (o *TypesLogRecord) GetWorkspaceHandleOk() (*string, bool)`

GetWorkspaceHandleOk returns a tuple with the WorkspaceHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceHandle

`func (o *TypesLogRecord) SetWorkspaceHandle(v string)`

SetWorkspaceHandle sets WorkspaceHandle field to given value.


### GetWorkspaceId

`func (o *TypesLogRecord) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *TypesLogRecord) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *TypesLogRecord) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


