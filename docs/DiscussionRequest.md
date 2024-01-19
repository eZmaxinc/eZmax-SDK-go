# DiscussionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDiscussionID** | Pointer to **int32** | The unique ID of the Discussion | [optional] 
**SDiscussionDescription** | **string** | The description of the Discussion | 
**BDiscussionClosed** | Pointer to **bool** | Whether if it&#39;s an closed | [optional] 

## Methods

### NewDiscussionRequest

`func NewDiscussionRequest(sDiscussionDescription string, ) *DiscussionRequest`

NewDiscussionRequest instantiates a new DiscussionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionRequestWithDefaults

`func NewDiscussionRequestWithDefaults() *DiscussionRequest`

NewDiscussionRequestWithDefaults instantiates a new DiscussionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDiscussionID

`func (o *DiscussionRequest) GetPkiDiscussionID() int32`

GetPkiDiscussionID returns the PkiDiscussionID field if non-nil, zero value otherwise.

### GetPkiDiscussionIDOk

`func (o *DiscussionRequest) GetPkiDiscussionIDOk() (*int32, bool)`

GetPkiDiscussionIDOk returns a tuple with the PkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDiscussionID

`func (o *DiscussionRequest) SetPkiDiscussionID(v int32)`

SetPkiDiscussionID sets PkiDiscussionID field to given value.

### HasPkiDiscussionID

`func (o *DiscussionRequest) HasPkiDiscussionID() bool`

HasPkiDiscussionID returns a boolean if a field has been set.

### GetSDiscussionDescription

`func (o *DiscussionRequest) GetSDiscussionDescription() string`

GetSDiscussionDescription returns the SDiscussionDescription field if non-nil, zero value otherwise.

### GetSDiscussionDescriptionOk

`func (o *DiscussionRequest) GetSDiscussionDescriptionOk() (*string, bool)`

GetSDiscussionDescriptionOk returns a tuple with the SDiscussionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDiscussionDescription

`func (o *DiscussionRequest) SetSDiscussionDescription(v string)`

SetSDiscussionDescription sets SDiscussionDescription field to given value.


### GetBDiscussionClosed

`func (o *DiscussionRequest) GetBDiscussionClosed() bool`

GetBDiscussionClosed returns the BDiscussionClosed field if non-nil, zero value otherwise.

### GetBDiscussionClosedOk

`func (o *DiscussionRequest) GetBDiscussionClosedOk() (*bool, bool)`

GetBDiscussionClosedOk returns a tuple with the BDiscussionClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDiscussionClosed

`func (o *DiscussionRequest) SetBDiscussionClosed(v bool)`

SetBDiscussionClosed sets BDiscussionClosed field to given value.

### HasBDiscussionClosed

`func (o *DiscussionRequest) HasBDiscussionClosed() bool`

HasBDiscussionClosed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


