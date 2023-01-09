# ErrorDetailModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorDetailModel

`func NewErrorDetailModel() *ErrorDetailModel`

NewErrorDetailModel instantiates a new ErrorDetailModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailModelWithDefaults

`func NewErrorDetailModelWithDefaults() *ErrorDetailModel`

NewErrorDetailModelWithDefaults instantiates a new ErrorDetailModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ErrorDetailModel) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ErrorDetailModel) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ErrorDetailModel) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ErrorDetailModel) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorDetailModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDetailModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDetailModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorDetailModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


