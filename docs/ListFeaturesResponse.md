# ListFeaturesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Feature**](Feature.md) |  | [optional] 

## Methods

### NewListFeaturesResponse

`func NewListFeaturesResponse() *ListFeaturesResponse`

NewListFeaturesResponse instantiates a new ListFeaturesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFeaturesResponseWithDefaults

`func NewListFeaturesResponseWithDefaults() *ListFeaturesResponse`

NewListFeaturesResponseWithDefaults instantiates a new ListFeaturesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListFeaturesResponse) GetItems() []Feature`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListFeaturesResponse) GetItemsOk() (*[]Feature, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListFeaturesResponse) SetItems(v []Feature)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListFeaturesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


