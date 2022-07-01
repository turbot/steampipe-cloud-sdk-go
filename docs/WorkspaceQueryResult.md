# WorkspaceQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]WorkspaceQueryResultColumn**](WorkspaceQueryResultColumn.md) |  | 
**Rows** | **[]map[string]interface{}** |  | 

## Methods

### NewWorkspaceQueryResult

`func NewWorkspaceQueryResult(columns []WorkspaceQueryResultColumn, rows []map[string]interface{}, ) *WorkspaceQueryResult`

NewWorkspaceQueryResult instantiates a new WorkspaceQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceQueryResultWithDefaults

`func NewWorkspaceQueryResultWithDefaults() *WorkspaceQueryResult`

NewWorkspaceQueryResultWithDefaults instantiates a new WorkspaceQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *WorkspaceQueryResult) GetColumns() []WorkspaceQueryResultColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *WorkspaceQueryResult) GetColumnsOk() (*[]WorkspaceQueryResultColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *WorkspaceQueryResult) SetColumns(v []WorkspaceQueryResultColumn)`

SetColumns sets Columns field to given value.


### GetRows

`func (o *WorkspaceQueryResult) GetRows() []map[string]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *WorkspaceQueryResult) GetRowsOk() (*[]map[string]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *WorkspaceQueryResult) SetRows(v []map[string]interface{})`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


