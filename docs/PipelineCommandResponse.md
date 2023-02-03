# PipelineCommandResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** |  | 
**ProcessId** | **string** |  | 

## Methods

### NewPipelineCommandResponse

`func NewPipelineCommandResponse(command string, processId string, ) *PipelineCommandResponse`

NewPipelineCommandResponse instantiates a new PipelineCommandResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineCommandResponseWithDefaults

`func NewPipelineCommandResponseWithDefaults() *PipelineCommandResponse`

NewPipelineCommandResponseWithDefaults instantiates a new PipelineCommandResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *PipelineCommandResponse) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PipelineCommandResponse) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PipelineCommandResponse) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetProcessId

`func (o *PipelineCommandResponse) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *PipelineCommandResponse) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *PipelineCommandResponse) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


