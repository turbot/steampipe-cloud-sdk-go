# TestConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Plugin** | **string** |  | 

## Methods

### NewTestConnectionRequest

`func NewTestConnectionRequest(plugin string, ) *TestConnectionRequest`

NewTestConnectionRequest instantiates a new TestConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestConnectionRequestWithDefaults

`func NewTestConnectionRequestWithDefaults() *TestConnectionRequest`

NewTestConnectionRequestWithDefaults instantiates a new TestConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *TestConnectionRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TestConnectionRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TestConnectionRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *TestConnectionRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPlugin

`func (o *TestConnectionRequest) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *TestConnectionRequest) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *TestConnectionRequest) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


