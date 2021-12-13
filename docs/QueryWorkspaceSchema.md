# QueryWorkspaceSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** |  | 
**Schemas** | [**[]QuerySchemaInfo**](QuerySchemaInfo.md) |  | 

## Methods

### NewQueryWorkspaceSchema

`func NewQueryWorkspaceSchema(duration int32, schemas []QuerySchemaInfo, ) *QueryWorkspaceSchema`

NewQueryWorkspaceSchema instantiates a new QueryWorkspaceSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWorkspaceSchemaWithDefaults

`func NewQueryWorkspaceSchemaWithDefaults() *QueryWorkspaceSchema`

NewQueryWorkspaceSchemaWithDefaults instantiates a new QueryWorkspaceSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *QueryWorkspaceSchema) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *QueryWorkspaceSchema) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *QueryWorkspaceSchema) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetSchemas

`func (o *QueryWorkspaceSchema) GetSchemas() []QuerySchemaInfo`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *QueryWorkspaceSchema) GetSchemasOk() (*[]QuerySchemaInfo, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *QueryWorkspaceSchema) SetSchemas(v []QuerySchemaInfo)`

SetSchemas sets Schemas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


