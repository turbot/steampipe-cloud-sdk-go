# WorkspaceSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** |  | 
**Schemas** | [**[]SchemaInfo**](SchemaInfo.md) |  | 

## Methods

### NewWorkspaceSchema

`func NewWorkspaceSchema(duration int32, schemas []SchemaInfo, ) *WorkspaceSchema`

NewWorkspaceSchema instantiates a new WorkspaceSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceSchemaWithDefaults

`func NewWorkspaceSchemaWithDefaults() *WorkspaceSchema`

NewWorkspaceSchemaWithDefaults instantiates a new WorkspaceSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *WorkspaceSchema) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *WorkspaceSchema) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *WorkspaceSchema) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetSchemas

`func (o *WorkspaceSchema) GetSchemas() []SchemaInfo`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *WorkspaceSchema) GetSchemasOk() (*[]SchemaInfo, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *WorkspaceSchema) SetSchemas(v []SchemaInfo)`

SetSchemas sets Schemas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


