# CreatePipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | **map[string]interface{}** | A map of arguments to be passed to be pipeline. | 
**Frequency** | **map[string]interface{}** | The frequency at which the pipeline will run. | 
**Pipeline** | **string** | The name of the pipeline to be executed. | 
**Tags** | Pointer to **map[string]interface{}** | The tags for this pipeline. | [optional] 
**Title** | **string** | The title of the pipeline. | 

## Methods

### NewCreatePipelineRequest

`func NewCreatePipelineRequest(args map[string]interface{}, frequency map[string]interface{}, pipeline string, title string, ) *CreatePipelineRequest`

NewCreatePipelineRequest instantiates a new CreatePipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePipelineRequestWithDefaults

`func NewCreatePipelineRequestWithDefaults() *CreatePipelineRequest`

NewCreatePipelineRequestWithDefaults instantiates a new CreatePipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *CreatePipelineRequest) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *CreatePipelineRequest) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *CreatePipelineRequest) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.


### GetFrequency

`func (o *CreatePipelineRequest) GetFrequency() map[string]interface{}`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CreatePipelineRequest) GetFrequencyOk() (*map[string]interface{}, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CreatePipelineRequest) SetFrequency(v map[string]interface{})`

SetFrequency sets Frequency field to given value.


### GetPipeline

`func (o *CreatePipelineRequest) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *CreatePipelineRequest) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *CreatePipelineRequest) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.


### GetTags

`func (o *CreatePipelineRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreatePipelineRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreatePipelineRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreatePipelineRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *CreatePipelineRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreatePipelineRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreatePipelineRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


