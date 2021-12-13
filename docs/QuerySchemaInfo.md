# QuerySchemaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Tables** | [**[]QuerySchemaTable**](QuerySchemaTable.md) |  | 

## Methods

### NewQuerySchemaInfo

`func NewQuerySchemaInfo(name string, tables []QuerySchemaTable, ) *QuerySchemaInfo`

NewQuerySchemaInfo instantiates a new QuerySchemaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySchemaInfoWithDefaults

`func NewQuerySchemaInfoWithDefaults() *QuerySchemaInfo`

NewQuerySchemaInfoWithDefaults instantiates a new QuerySchemaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *QuerySchemaInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuerySchemaInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuerySchemaInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuerySchemaInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *QuerySchemaInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuerySchemaInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuerySchemaInfo) SetName(v string)`

SetName sets Name field to given value.


### GetTables

`func (o *QuerySchemaInfo) GetTables() []QuerySchemaTable`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *QuerySchemaInfo) GetTablesOk() (*[]QuerySchemaTable, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *QuerySchemaInfo) SetTables(v []QuerySchemaTable)`

SetTables sets Tables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


