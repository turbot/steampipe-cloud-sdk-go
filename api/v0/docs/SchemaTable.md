# SchemaTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]SchemaTableColumn**](SchemaTableColumn.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewSchemaTable

`func NewSchemaTable(columns []SchemaTableColumn, name string, ) *SchemaTable`

NewSchemaTable instantiates a new SchemaTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaTableWithDefaults

`func NewSchemaTableWithDefaults() *SchemaTable`

NewSchemaTableWithDefaults instantiates a new SchemaTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *SchemaTable) GetColumns() []SchemaTableColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *SchemaTable) GetColumnsOk() (*[]SchemaTableColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *SchemaTable) SetColumns(v []SchemaTableColumn)`

SetColumns sets Columns field to given value.


### GetDescription

`func (o *SchemaTable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaTable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaTable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaTable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SchemaTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaTable) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


