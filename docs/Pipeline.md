# Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DeletedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) |  | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**DesiredState** | **string** | DesiredState is user editable state of the pipeline, there are only 3 options: enabled, disabled or paused | 
**Frequency** | [**PipelineFrequency**](PipelineFrequency.md) |  | 
**Id** | **string** | The unique identifier of the pipeline. | 
**IdentityId** | **string** | We need the IdentityID to be in the JSON field because I&#39;m sending the pipeline definition across the Temporal Workflow (see Datatank migration workflow) | 
**LastProcess** | Pointer to [**SpProcess**](SpProcess.md) |  | [optional] 
**LastProcessId** | Pointer to **string** | The id of the last process that was run for the pipeline. | [optional] 
**NextRunAt** | Pointer to **string** | The time when the pipeline is next scheduled to run in ISO 8601 UTC. | [optional] 
**Pipeline** | **string** | The name of the pipeline to be executed. | 
**State** | **string** | State is the system state of the pipeline. User should not be able to modify this value | 
**StateReason** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **interface{}** |  | [optional] 
**Title** | Pointer to **string** | The title of the pipeline. | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceId** | Pointer to **string** | The unique identifier of the workspace on which the pipeline is created if any. | [optional] 

## Methods

### NewPipeline

`func NewPipeline(createdAt string, createdById string, deletedById string, desiredState string, frequency PipelineFrequency, id string, identityId string, pipeline string, state string, updatedById string, versionId int32, ) *Pipeline`

NewPipeline instantiates a new Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineWithDefaults

`func NewPipelineWithDefaults() *Pipeline`

NewPipelineWithDefaults instantiates a new Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *Pipeline) GetArgs() interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *Pipeline) GetArgsOk() (*interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *Pipeline) SetArgs(v interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *Pipeline) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *Pipeline) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *Pipeline) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetCreatedAt

`func (o *Pipeline) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Pipeline) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Pipeline) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Pipeline) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Pipeline) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Pipeline) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Pipeline) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *Pipeline) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Pipeline) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Pipeline) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDeletedAt

`func (o *Pipeline) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Pipeline) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Pipeline) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Pipeline) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Pipeline) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Pipeline) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Pipeline) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Pipeline) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *Pipeline) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *Pipeline) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *Pipeline) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetDesiredState

`func (o *Pipeline) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *Pipeline) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *Pipeline) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.


### GetFrequency

`func (o *Pipeline) GetFrequency() PipelineFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *Pipeline) GetFrequencyOk() (*PipelineFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *Pipeline) SetFrequency(v PipelineFrequency)`

SetFrequency sets Frequency field to given value.


### GetId

`func (o *Pipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pipeline) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *Pipeline) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *Pipeline) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *Pipeline) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetLastProcess

`func (o *Pipeline) GetLastProcess() SpProcess`

GetLastProcess returns the LastProcess field if non-nil, zero value otherwise.

### GetLastProcessOk

`func (o *Pipeline) GetLastProcessOk() (*SpProcess, bool)`

GetLastProcessOk returns a tuple with the LastProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcess

`func (o *Pipeline) SetLastProcess(v SpProcess)`

SetLastProcess sets LastProcess field to given value.

### HasLastProcess

`func (o *Pipeline) HasLastProcess() bool`

HasLastProcess returns a boolean if a field has been set.

### GetLastProcessId

`func (o *Pipeline) GetLastProcessId() string`

GetLastProcessId returns the LastProcessId field if non-nil, zero value otherwise.

### GetLastProcessIdOk

`func (o *Pipeline) GetLastProcessIdOk() (*string, bool)`

GetLastProcessIdOk returns a tuple with the LastProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessId

`func (o *Pipeline) SetLastProcessId(v string)`

SetLastProcessId sets LastProcessId field to given value.

### HasLastProcessId

`func (o *Pipeline) HasLastProcessId() bool`

HasLastProcessId returns a boolean if a field has been set.

### GetNextRunAt

`func (o *Pipeline) GetNextRunAt() string`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *Pipeline) GetNextRunAtOk() (*string, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *Pipeline) SetNextRunAt(v string)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *Pipeline) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### GetPipeline

`func (o *Pipeline) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *Pipeline) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *Pipeline) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.


### GetState

`func (o *Pipeline) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Pipeline) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Pipeline) SetState(v string)`

SetState sets State field to given value.


### GetStateReason

`func (o *Pipeline) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *Pipeline) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *Pipeline) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *Pipeline) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetTags

`func (o *Pipeline) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Pipeline) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Pipeline) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Pipeline) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Pipeline) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Pipeline) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTitle

`func (o *Pipeline) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Pipeline) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Pipeline) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Pipeline) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Pipeline) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Pipeline) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Pipeline) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Pipeline) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Pipeline) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Pipeline) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Pipeline) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Pipeline) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *Pipeline) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Pipeline) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Pipeline) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *Pipeline) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Pipeline) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Pipeline) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceId

`func (o *Pipeline) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Pipeline) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Pipeline) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *Pipeline) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


