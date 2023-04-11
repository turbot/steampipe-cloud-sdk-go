# CreateWorkspaceAggregatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | **[]string** | The connections that are a part of the aggregator. | 
**Handle** | **string** | The handle of the aggregator. | 
**Plugin** | **string** | The plugin which this aggregator belongs to. | 

## Methods

### NewCreateWorkspaceAggregatorRequest

`func NewCreateWorkspaceAggregatorRequest(connections []string, handle string, plugin string, ) *CreateWorkspaceAggregatorRequest`

NewCreateWorkspaceAggregatorRequest instantiates a new CreateWorkspaceAggregatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkspaceAggregatorRequestWithDefaults

`func NewCreateWorkspaceAggregatorRequestWithDefaults() *CreateWorkspaceAggregatorRequest`

NewCreateWorkspaceAggregatorRequestWithDefaults instantiates a new CreateWorkspaceAggregatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *CreateWorkspaceAggregatorRequest) GetConnections() []string`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *CreateWorkspaceAggregatorRequest) GetConnectionsOk() (*[]string, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *CreateWorkspaceAggregatorRequest) SetConnections(v []string)`

SetConnections sets Connections field to given value.


### GetHandle

`func (o *CreateWorkspaceAggregatorRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateWorkspaceAggregatorRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateWorkspaceAggregatorRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetPlugin

`func (o *CreateWorkspaceAggregatorRequest) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *CreateWorkspaceAggregatorRequest) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *CreateWorkspaceAggregatorRequest) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


