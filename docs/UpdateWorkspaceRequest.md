# UpdateWorkspaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredState** | Pointer to **string** | paused is not yet supported for Workspace | [optional] 
**Handle** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateWorkspaceRequest

`func NewUpdateWorkspaceRequest() *UpdateWorkspaceRequest`

NewUpdateWorkspaceRequest instantiates a new UpdateWorkspaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkspaceRequestWithDefaults

`func NewUpdateWorkspaceRequestWithDefaults() *UpdateWorkspaceRequest`

NewUpdateWorkspaceRequestWithDefaults instantiates a new UpdateWorkspaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredState

`func (o *UpdateWorkspaceRequest) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *UpdateWorkspaceRequest) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *UpdateWorkspaceRequest) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *UpdateWorkspaceRequest) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetHandle

`func (o *UpdateWorkspaceRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UpdateWorkspaceRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UpdateWorkspaceRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UpdateWorkspaceRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


