# AttachmentRenameV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SAttachmentName** | **string** | The name of the Attachment | 
**SAttachmentCategory** | **string** | The attachment category | 
**BForceOverride** | Pointer to **bool** | Forces an override if the attachment name and category conflicts with another attachment. | [optional] 

## Methods

### NewAttachmentRenameV1Request

`func NewAttachmentRenameV1Request(sAttachmentName string, sAttachmentCategory string, ) *AttachmentRenameV1Request`

NewAttachmentRenameV1Request instantiates a new AttachmentRenameV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentRenameV1RequestWithDefaults

`func NewAttachmentRenameV1RequestWithDefaults() *AttachmentRenameV1Request`

NewAttachmentRenameV1RequestWithDefaults instantiates a new AttachmentRenameV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSAttachmentName

`func (o *AttachmentRenameV1Request) GetSAttachmentName() string`

GetSAttachmentName returns the SAttachmentName field if non-nil, zero value otherwise.

### GetSAttachmentNameOk

`func (o *AttachmentRenameV1Request) GetSAttachmentNameOk() (*string, bool)`

GetSAttachmentNameOk returns a tuple with the SAttachmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentName

`func (o *AttachmentRenameV1Request) SetSAttachmentName(v string)`

SetSAttachmentName sets SAttachmentName field to given value.


### GetSAttachmentCategory

`func (o *AttachmentRenameV1Request) GetSAttachmentCategory() string`

GetSAttachmentCategory returns the SAttachmentCategory field if non-nil, zero value otherwise.

### GetSAttachmentCategoryOk

`func (o *AttachmentRenameV1Request) GetSAttachmentCategoryOk() (*string, bool)`

GetSAttachmentCategoryOk returns a tuple with the SAttachmentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentCategory

`func (o *AttachmentRenameV1Request) SetSAttachmentCategory(v string)`

SetSAttachmentCategory sets SAttachmentCategory field to given value.


### GetBForceOverride

`func (o *AttachmentRenameV1Request) GetBForceOverride() bool`

GetBForceOverride returns the BForceOverride field if non-nil, zero value otherwise.

### GetBForceOverrideOk

`func (o *AttachmentRenameV1Request) GetBForceOverrideOk() (*bool, bool)`

GetBForceOverrideOk returns a tuple with the BForceOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBForceOverride

`func (o *AttachmentRenameV1Request) SetBForceOverride(v bool)`

SetBForceOverride sets BForceOverride field to given value.

### HasBForceOverride

`func (o *AttachmentRenameV1Request) HasBForceOverride() bool`

HasBForceOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


