# Quota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **int32** | Remaining limit | [optional] 
**Limit** | Pointer to **int32** | Max limit | [optional] 
**Used** | Pointer to **int32** | Exhausted limit | [optional] 

## Methods

### NewQuota

`func NewQuota() *Quota`

NewQuota instantiates a new Quota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaWithDefaults

`func NewQuotaWithDefaults() *Quota`

NewQuotaWithDefaults instantiates a new Quota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *Quota) GetAvailable() int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Quota) GetAvailableOk() (*int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Quota) SetAvailable(v int32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Quota) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetLimit

`func (o *Quota) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Quota) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Quota) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Quota) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetUsed

`func (o *Quota) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *Quota) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *Quota) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *Quota) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


