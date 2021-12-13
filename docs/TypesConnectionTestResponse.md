# TypesConnectionTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]TypesConnectionTestRow**](TypesConnectionTestRow.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewTypesConnectionTestResponse

`func NewTypesConnectionTestResponse() *TypesConnectionTestResponse`

NewTypesConnectionTestResponse instantiates a new TypesConnectionTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesConnectionTestResponseWithDefaults

`func NewTypesConnectionTestResponseWithDefaults() *TypesConnectionTestResponse`

NewTypesConnectionTestResponseWithDefaults instantiates a new TypesConnectionTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TypesConnectionTestResponse) GetItems() []TypesConnectionTestRow`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TypesConnectionTestResponse) GetItemsOk() (*[]TypesConnectionTestRow, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TypesConnectionTestResponse) SetItems(v []TypesConnectionTestRow)`

SetItems sets Items field to given value.

### HasItems

`func (o *TypesConnectionTestResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *TypesConnectionTestResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *TypesConnectionTestResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *TypesConnectionTestResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *TypesConnectionTestResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


