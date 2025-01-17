# DiscussionRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDiscussionID** | Pointer to **int32** | The unique ID of the Discussion | [optional] 
**SDiscussionDescription** | **string** | The description of the Discussion | 
**BDiscussionClosed** | Pointer to **bool** | Whether if it&#39;s an closed | [optional] 

## Methods

### NewDiscussionRequestCompound

`func NewDiscussionRequestCompound(sDiscussionDescription string, ) *DiscussionRequestCompound`

NewDiscussionRequestCompound instantiates a new DiscussionRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionRequestCompoundWithDefaults

`func NewDiscussionRequestCompoundWithDefaults() *DiscussionRequestCompound`

NewDiscussionRequestCompoundWithDefaults instantiates a new DiscussionRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDiscussionID

`func (o *DiscussionRequestCompound) GetPkiDiscussionID() int32`

GetPkiDiscussionID returns the PkiDiscussionID field if non-nil, zero value otherwise.

### GetPkiDiscussionIDOk

`func (o *DiscussionRequestCompound) GetPkiDiscussionIDOk() (*int32, bool)`

GetPkiDiscussionIDOk returns a tuple with the PkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDiscussionID

`func (o *DiscussionRequestCompound) SetPkiDiscussionID(v int32)`

SetPkiDiscussionID sets PkiDiscussionID field to given value.

### HasPkiDiscussionID

`func (o *DiscussionRequestCompound) HasPkiDiscussionID() bool`

HasPkiDiscussionID returns a boolean if a field has been set.

### GetSDiscussionDescription

`func (o *DiscussionRequestCompound) GetSDiscussionDescription() string`

GetSDiscussionDescription returns the SDiscussionDescription field if non-nil, zero value otherwise.

### GetSDiscussionDescriptionOk

`func (o *DiscussionRequestCompound) GetSDiscussionDescriptionOk() (*string, bool)`

GetSDiscussionDescriptionOk returns a tuple with the SDiscussionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDiscussionDescription

`func (o *DiscussionRequestCompound) SetSDiscussionDescription(v string)`

SetSDiscussionDescription sets SDiscussionDescription field to given value.


### GetBDiscussionClosed

`func (o *DiscussionRequestCompound) GetBDiscussionClosed() bool`

GetBDiscussionClosed returns the BDiscussionClosed field if non-nil, zero value otherwise.

### GetBDiscussionClosedOk

`func (o *DiscussionRequestCompound) GetBDiscussionClosedOk() (*bool, bool)`

GetBDiscussionClosedOk returns a tuple with the BDiscussionClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDiscussionClosed

`func (o *DiscussionRequestCompound) SetBDiscussionClosed(v bool)`

SetBDiscussionClosed sets BDiscussionClosed field to given value.

### HasBDiscussionClosed

`func (o *DiscussionRequestCompound) HasBDiscussionClosed() bool`

HasBDiscussionClosed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


