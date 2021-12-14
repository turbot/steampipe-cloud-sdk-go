# ListLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]LogRecord**](LogRecord.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListLogsResponse

`func NewListLogsResponse() *ListLogsResponse`

NewListLogsResponse instantiates a new ListLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogsResponseWithDefaults

`func NewListLogsResponseWithDefaults() *ListLogsResponse`

NewListLogsResponseWithDefaults instantiates a new ListLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListLogsResponse) GetItems() []LogRecord`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListLogsResponse) GetItemsOk() (*[]LogRecord, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListLogsResponse) SetItems(v []LogRecord)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListLogsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *ListLogsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListLogsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListLogsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListLogsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


