# BillingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteAt** | Pointer to **string** | The time when a subscription will be cancelled in stripe. | [optional] 
**DisableAt** | Pointer to **string** | The time when the org will be suspended due to failed billing in ISO 8601 UTC. | [optional] 
**Status** | **string** | The billing status for the org in stripe i.e. whether its pending verification or verified. | 

## Methods

### NewBillingInfo

`func NewBillingInfo(status string, ) *BillingInfo`

NewBillingInfo instantiates a new BillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInfoWithDefaults

`func NewBillingInfoWithDefaults() *BillingInfo`

NewBillingInfoWithDefaults instantiates a new BillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteAt

`func (o *BillingInfo) GetDeleteAt() string`

GetDeleteAt returns the DeleteAt field if non-nil, zero value otherwise.

### GetDeleteAtOk

`func (o *BillingInfo) GetDeleteAtOk() (*string, bool)`

GetDeleteAtOk returns a tuple with the DeleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAt

`func (o *BillingInfo) SetDeleteAt(v string)`

SetDeleteAt sets DeleteAt field to given value.

### HasDeleteAt

`func (o *BillingInfo) HasDeleteAt() bool`

HasDeleteAt returns a boolean if a field has been set.

### GetDisableAt

`func (o *BillingInfo) GetDisableAt() string`

GetDisableAt returns the DisableAt field if non-nil, zero value otherwise.

### GetDisableAtOk

`func (o *BillingInfo) GetDisableAtOk() (*string, bool)`

GetDisableAtOk returns a tuple with the DisableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAt

`func (o *BillingInfo) SetDisableAt(v string)`

SetDisableAt sets DisableAt field to given value.

### HasDisableAt

`func (o *BillingInfo) HasDisableAt() bool`

HasDisableAt returns a boolean if a field has been set.

### GetStatus

`func (o *BillingInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingInfo) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


