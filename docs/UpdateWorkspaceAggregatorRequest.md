# UpdateWorkspaceAggregatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to **[]string** | The connections that are a part of the aggregator. | [optional] 
**Handle** | Pointer to **string** | The handle of the aggregator. | [optional] 

## Methods

### NewUpdateWorkspaceAggregatorRequest

`func NewUpdateWorkspaceAggregatorRequest() *UpdateWorkspaceAggregatorRequest`

NewUpdateWorkspaceAggregatorRequest instantiates a new UpdateWorkspaceAggregatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkspaceAggregatorRequestWithDefaults

`func NewUpdateWorkspaceAggregatorRequestWithDefaults() *UpdateWorkspaceAggregatorRequest`

NewUpdateWorkspaceAggregatorRequestWithDefaults instantiates a new UpdateWorkspaceAggregatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *UpdateWorkspaceAggregatorRequest) GetConnections() []string`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *UpdateWorkspaceAggregatorRequest) GetConnectionsOk() (*[]string, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *UpdateWorkspaceAggregatorRequest) SetConnections(v []string)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *UpdateWorkspaceAggregatorRequest) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetHandle

`func (o *UpdateWorkspaceAggregatorRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UpdateWorkspaceAggregatorRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UpdateWorkspaceAggregatorRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UpdateWorkspaceAggregatorRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


