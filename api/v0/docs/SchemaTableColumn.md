# SchemaTableColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewSchemaTableColumn

`func NewSchemaTableColumn(dataType string, name string, ) *SchemaTableColumn`

NewSchemaTableColumn instantiates a new SchemaTableColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaTableColumnWithDefaults

`func NewSchemaTableColumnWithDefaults() *SchemaTableColumn`

NewSchemaTableColumnWithDefaults instantiates a new SchemaTableColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *SchemaTableColumn) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *SchemaTableColumn) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *SchemaTableColumn) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetDescription

`func (o *SchemaTableColumn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaTableColumn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaTableColumn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaTableColumn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SchemaTableColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaTableColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaTableColumn) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


