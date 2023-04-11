# UpdatePipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **interface{}** |  | [optional] 
**DesiredState** | Pointer to **string** | The desired state of the pipeline. | [optional] 
**Frequency** | Pointer to [**PipelineFrequency**](PipelineFrequency.md) |  | [optional] 
**Tags** | Pointer to **interface{}** |  | [optional] 
**Title** | Pointer to **string** | The title of the pipeline. | [optional] 

## Methods

### NewUpdatePipelineRequest

`func NewUpdatePipelineRequest() *UpdatePipelineRequest`

NewUpdatePipelineRequest instantiates a new UpdatePipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePipelineRequestWithDefaults

`func NewUpdatePipelineRequestWithDefaults() *UpdatePipelineRequest`

NewUpdatePipelineRequestWithDefaults instantiates a new UpdatePipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *UpdatePipelineRequest) GetArgs() interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *UpdatePipelineRequest) GetArgsOk() (*interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *UpdatePipelineRequest) SetArgs(v interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *UpdatePipelineRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *UpdatePipelineRequest) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *UpdatePipelineRequest) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetDesiredState

`func (o *UpdatePipelineRequest) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *UpdatePipelineRequest) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *UpdatePipelineRequest) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *UpdatePipelineRequest) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetFrequency

`func (o *UpdatePipelineRequest) GetFrequency() PipelineFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *UpdatePipelineRequest) GetFrequencyOk() (*PipelineFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *UpdatePipelineRequest) SetFrequency(v PipelineFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *UpdatePipelineRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetTags

`func (o *UpdatePipelineRequest) GetTags() interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdatePipelineRequest) GetTagsOk() (*interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdatePipelineRequest) SetTags(v interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdatePipelineRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdatePipelineRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdatePipelineRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetTitle

`func (o *UpdatePipelineRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdatePipelineRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdatePipelineRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdatePipelineRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


