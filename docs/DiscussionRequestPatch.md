# DiscussionRequestPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SDiscussionDescription** | Pointer to **string** | The description of the Discussion | [optional] 
**BDiscussionClosed** | Pointer to **bool** | Whether if it&#39;s an closed | [optional] 

## Methods

### NewDiscussionRequestPatch

`func NewDiscussionRequestPatch() *DiscussionRequestPatch`

NewDiscussionRequestPatch instantiates a new DiscussionRequestPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionRequestPatchWithDefaults

`func NewDiscussionRequestPatchWithDefaults() *DiscussionRequestPatch`

NewDiscussionRequestPatchWithDefaults instantiates a new DiscussionRequestPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSDiscussionDescription

`func (o *DiscussionRequestPatch) GetSDiscussionDescription() string`

GetSDiscussionDescription returns the SDiscussionDescription field if non-nil, zero value otherwise.

### GetSDiscussionDescriptionOk

`func (o *DiscussionRequestPatch) GetSDiscussionDescriptionOk() (*string, bool)`

GetSDiscussionDescriptionOk returns a tuple with the SDiscussionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDiscussionDescription

`func (o *DiscussionRequestPatch) SetSDiscussionDescription(v string)`

SetSDiscussionDescription sets SDiscussionDescription field to given value.

### HasSDiscussionDescription

`func (o *DiscussionRequestPatch) HasSDiscussionDescription() bool`

HasSDiscussionDescription returns a boolean if a field has been set.

### GetBDiscussionClosed

`func (o *DiscussionRequestPatch) GetBDiscussionClosed() bool`

GetBDiscussionClosed returns the BDiscussionClosed field if non-nil, zero value otherwise.

### GetBDiscussionClosedOk

`func (o *DiscussionRequestPatch) GetBDiscussionClosedOk() (*bool, bool)`

GetBDiscussionClosedOk returns a tuple with the BDiscussionClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDiscussionClosed

`func (o *DiscussionRequestPatch) SetBDiscussionClosed(v bool)`

SetBDiscussionClosed sets BDiscussionClosed field to given value.

### HasBDiscussionClosed

`func (o *DiscussionRequestPatch) HasBDiscussionClosed() bool`

HasBDiscussionClosed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


