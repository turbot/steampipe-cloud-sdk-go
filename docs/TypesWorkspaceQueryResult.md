# TypesWorkspaceQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]TypesWorkspaceQueryResultColumn**](TypesWorkspaceQueryResultColumn.md) |  | 
**Rows** | **[][]map[string]interface{}** |  | 

## Methods

### NewTypesWorkspaceQueryResult

`func NewTypesWorkspaceQueryResult(columns []TypesWorkspaceQueryResultColumn, rows [][]map[string]interface{}, ) *TypesWorkspaceQueryResult`

NewTypesWorkspaceQueryResult instantiates a new TypesWorkspaceQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesWorkspaceQueryResultWithDefaults

`func NewTypesWorkspaceQueryResultWithDefaults() *TypesWorkspaceQueryResult`

NewTypesWorkspaceQueryResultWithDefaults instantiates a new TypesWorkspaceQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *TypesWorkspaceQueryResult) GetColumns() []TypesWorkspaceQueryResultColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TypesWorkspaceQueryResult) GetColumnsOk() (*[]TypesWorkspaceQueryResultColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TypesWorkspaceQueryResult) SetColumns(v []TypesWorkspaceQueryResultColumn)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *TypesWorkspaceQueryResult) GetRows() [][]map[string]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *TypesWorkspaceQueryResult) GetRowsOk() (*[][]map[string]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *TypesWorkspaceQueryResult) SetRows(v [][]map[string]interface{})`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


