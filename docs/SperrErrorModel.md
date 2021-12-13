# SperrErrorModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]SperrErrorDetailModel**](SperrErrorDetailModel.md) |  | [optional] 
**Id** | **string** |  | 
**Instance** | Pointer to **string** |  | [optional] 
**Status** | **int32** |  | 
**Title** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewSperrErrorModel

`func NewSperrErrorModel(id string, status int32, title string, type_ string, ) *SperrErrorModel`

NewSperrErrorModel instantiates a new SperrErrorModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSperrErrorModelWithDefaults

`func NewSperrErrorModelWithDefaults() *SperrErrorModel`

NewSperrErrorModelWithDefaults instantiates a new SperrErrorModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *SperrErrorModel) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *SperrErrorModel) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *SperrErrorModel) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *SperrErrorModel) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetErrors

`func (o *SperrErrorModel) GetErrors() []SperrErrorDetailModel`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SperrErrorModel) GetErrorsOk() (*[]SperrErrorDetailModel, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SperrErrorModel) SetErrors(v []SperrErrorDetailModel)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SperrErrorModel) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetId

`func (o *SperrErrorModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SperrErrorModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SperrErrorModel) SetId(v string)`

SetId sets Id field to given value.


### GetInstance

`func (o *SperrErrorModel) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *SperrErrorModel) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *SperrErrorModel) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *SperrErrorModel) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetStatus

`func (o *SperrErrorModel) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SperrErrorModel) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SperrErrorModel) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *SperrErrorModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SperrErrorModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SperrErrorModel) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *SperrErrorModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SperrErrorModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SperrErrorModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


