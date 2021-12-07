# TypesUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**Handle** | **string** |  | 
**Id** | **string** |  | 
**PreviewAccessMode** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**VersionId** | **int32** |  | 

## Methods

### NewTypesUser

`func NewTypesUser(createdAt string, email string, handle string, id string, status string, versionId int32, ) *TypesUser`

NewTypesUser instantiates a new TypesUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesUserWithDefaults

`func NewTypesUserWithDefaults() *TypesUser`

NewTypesUserWithDefaults instantiates a new TypesUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *TypesUser) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *TypesUser) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *TypesUser) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *TypesUser) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TypesUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TypesUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TypesUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisplayName

`func (o *TypesUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TypesUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TypesUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *TypesUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *TypesUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TypesUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TypesUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetHandle

`func (o *TypesUser) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *TypesUser) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *TypesUser) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *TypesUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TypesUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TypesUser) SetId(v string)`

SetId sets Id field to given value.


### GetPreviewAccessMode

`func (o *TypesUser) GetPreviewAccessMode() string`

GetPreviewAccessMode returns the PreviewAccessMode field if non-nil, zero value otherwise.

### GetPreviewAccessModeOk

`func (o *TypesUser) GetPreviewAccessModeOk() (*string, bool)`

GetPreviewAccessModeOk returns a tuple with the PreviewAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewAccessMode

`func (o *TypesUser) SetPreviewAccessMode(v string)`

SetPreviewAccessMode sets PreviewAccessMode field to given value.

### HasPreviewAccessMode

`func (o *TypesUser) HasPreviewAccessMode() bool`

HasPreviewAccessMode returns a boolean if a field has been set.

### GetStatus

`func (o *TypesUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesUser) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *TypesUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TypesUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TypesUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TypesUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *TypesUser) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TypesUser) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TypesUser) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TypesUser) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVersionId

`func (o *TypesUser) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TypesUser) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TypesUser) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


