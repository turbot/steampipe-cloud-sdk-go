# ListUserEmailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]UserEmail**](UserEmail.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListUserEmailsResponse

`func NewListUserEmailsResponse() *ListUserEmailsResponse`

NewListUserEmailsResponse instantiates a new ListUserEmailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserEmailsResponseWithDefaults

`func NewListUserEmailsResponseWithDefaults() *ListUserEmailsResponse`

NewListUserEmailsResponseWithDefaults instantiates a new ListUserEmailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListUserEmailsResponse) GetItems() []UserEmail`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListUserEmailsResponse) GetItemsOk() (*[]UserEmail, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListUserEmailsResponse) SetItems(v []UserEmail)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListUserEmailsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *ListUserEmailsResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListUserEmailsResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListUserEmailsResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListUserEmailsResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


