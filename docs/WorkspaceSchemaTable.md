# WorkspaceSchemaTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]WorkspaceSchemaTableColumn**](WorkspaceSchemaTableColumn.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewWorkspaceSchemaTable

`func NewWorkspaceSchemaTable(columns []WorkspaceSchemaTableColumn, name string, ) *WorkspaceSchemaTable`

NewWorkspaceSchemaTable instantiates a new WorkspaceSchemaTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceSchemaTableWithDefaults

`func NewWorkspaceSchemaTableWithDefaults() *WorkspaceSchemaTable`

NewWorkspaceSchemaTableWithDefaults instantiates a new WorkspaceSchemaTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *WorkspaceSchemaTable) GetColumns() []WorkspaceSchemaTableColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *WorkspaceSchemaTable) GetColumnsOk() (*[]WorkspaceSchemaTableColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *WorkspaceSchemaTable) SetColumns(v []WorkspaceSchemaTableColumn)`

SetColumns sets Columns field to given value.


### GetDescription

`func (o *WorkspaceSchemaTable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceSchemaTable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceSchemaTable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkspaceSchemaTable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceSchemaTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceSchemaTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceSchemaTable) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


