# CreateConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Handle** | **string** |  | 
**Plugin** | **string** | Type   string                 &#x60;json:\&quot;type\&quot; binding:\&quot;required,steampipe_connection_type\&quot;&#x60; | 

## Methods

### NewCreateConnectionRequest

`func NewCreateConnectionRequest(handle string, plugin string, ) *CreateConnectionRequest`

NewCreateConnectionRequest instantiates a new CreateConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectionRequestWithDefaults

`func NewCreateConnectionRequestWithDefaults() *CreateConnectionRequest`

NewCreateConnectionRequestWithDefaults instantiates a new CreateConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateConnectionRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateConnectionRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateConnectionRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateConnectionRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetHandle

`func (o *CreateConnectionRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateConnectionRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateConnectionRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetPlugin

`func (o *CreateConnectionRequest) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *CreateConnectionRequest) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *CreateConnectionRequest) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


