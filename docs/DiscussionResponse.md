# DiscussionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDiscussionID** | **int32** | The unique ID of the Discussion | 
**SDiscussionDescription** | **string** | The description of the Discussion | 
**BDiscussionClosed** | **bool** | Whether if it&#39;s an closed | 
**DtDiscussionLastread** | Pointer to **string** | The date the Discussion was last read | [optional] 
**IDiscussionmessageCount** | **int32** | The count of Attachment. | 
**IDiscussionmessageCountunread** | **int32** | The count of Attachment. | 
**ObjDiscussionconfiguration** | Pointer to **map[string]interface{}** | A Custom Discussionconfiguration Object | [optional] 

## Methods

### NewDiscussionResponse

`func NewDiscussionResponse(pkiDiscussionID int32, sDiscussionDescription string, bDiscussionClosed bool, iDiscussionmessageCount int32, iDiscussionmessageCountunread int32, ) *DiscussionResponse`

NewDiscussionResponse instantiates a new DiscussionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionResponseWithDefaults

`func NewDiscussionResponseWithDefaults() *DiscussionResponse`

NewDiscussionResponseWithDefaults instantiates a new DiscussionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDiscussionID

`func (o *DiscussionResponse) GetPkiDiscussionID() int32`

GetPkiDiscussionID returns the PkiDiscussionID field if non-nil, zero value otherwise.

### GetPkiDiscussionIDOk

`func (o *DiscussionResponse) GetPkiDiscussionIDOk() (*int32, bool)`

GetPkiDiscussionIDOk returns a tuple with the PkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDiscussionID

`func (o *DiscussionResponse) SetPkiDiscussionID(v int32)`

SetPkiDiscussionID sets PkiDiscussionID field to given value.


### GetSDiscussionDescription

`func (o *DiscussionResponse) GetSDiscussionDescription() string`

GetSDiscussionDescription returns the SDiscussionDescription field if non-nil, zero value otherwise.

### GetSDiscussionDescriptionOk

`func (o *DiscussionResponse) GetSDiscussionDescriptionOk() (*string, bool)`

GetSDiscussionDescriptionOk returns a tuple with the SDiscussionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDiscussionDescription

`func (o *DiscussionResponse) SetSDiscussionDescription(v string)`

SetSDiscussionDescription sets SDiscussionDescription field to given value.


### GetBDiscussionClosed

`func (o *DiscussionResponse) GetBDiscussionClosed() bool`

GetBDiscussionClosed returns the BDiscussionClosed field if non-nil, zero value otherwise.

### GetBDiscussionClosedOk

`func (o *DiscussionResponse) GetBDiscussionClosedOk() (*bool, bool)`

GetBDiscussionClosedOk returns a tuple with the BDiscussionClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDiscussionClosed

`func (o *DiscussionResponse) SetBDiscussionClosed(v bool)`

SetBDiscussionClosed sets BDiscussionClosed field to given value.


### GetDtDiscussionLastread

`func (o *DiscussionResponse) GetDtDiscussionLastread() string`

GetDtDiscussionLastread returns the DtDiscussionLastread field if non-nil, zero value otherwise.

### GetDtDiscussionLastreadOk

`func (o *DiscussionResponse) GetDtDiscussionLastreadOk() (*string, bool)`

GetDtDiscussionLastreadOk returns a tuple with the DtDiscussionLastread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDiscussionLastread

`func (o *DiscussionResponse) SetDtDiscussionLastread(v string)`

SetDtDiscussionLastread sets DtDiscussionLastread field to given value.

### HasDtDiscussionLastread

`func (o *DiscussionResponse) HasDtDiscussionLastread() bool`

HasDtDiscussionLastread returns a boolean if a field has been set.

### GetIDiscussionmessageCount

`func (o *DiscussionResponse) GetIDiscussionmessageCount() int32`

GetIDiscussionmessageCount returns the IDiscussionmessageCount field if non-nil, zero value otherwise.

### GetIDiscussionmessageCountOk

`func (o *DiscussionResponse) GetIDiscussionmessageCountOk() (*int32, bool)`

GetIDiscussionmessageCountOk returns a tuple with the IDiscussionmessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIDiscussionmessageCount

`func (o *DiscussionResponse) SetIDiscussionmessageCount(v int32)`

SetIDiscussionmessageCount sets IDiscussionmessageCount field to given value.


### GetIDiscussionmessageCountunread

`func (o *DiscussionResponse) GetIDiscussionmessageCountunread() int32`

GetIDiscussionmessageCountunread returns the IDiscussionmessageCountunread field if non-nil, zero value otherwise.

### GetIDiscussionmessageCountunreadOk

`func (o *DiscussionResponse) GetIDiscussionmessageCountunreadOk() (*int32, bool)`

GetIDiscussionmessageCountunreadOk returns a tuple with the IDiscussionmessageCountunread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIDiscussionmessageCountunread

`func (o *DiscussionResponse) SetIDiscussionmessageCountunread(v int32)`

SetIDiscussionmessageCountunread sets IDiscussionmessageCountunread field to given value.


### GetObjDiscussionconfiguration

`func (o *DiscussionResponse) GetObjDiscussionconfiguration() map[string]interface{}`

GetObjDiscussionconfiguration returns the ObjDiscussionconfiguration field if non-nil, zero value otherwise.

### GetObjDiscussionconfigurationOk

`func (o *DiscussionResponse) GetObjDiscussionconfigurationOk() (*map[string]interface{}, bool)`

GetObjDiscussionconfigurationOk returns a tuple with the ObjDiscussionconfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDiscussionconfiguration

`func (o *DiscussionResponse) SetObjDiscussionconfiguration(v map[string]interface{})`

SetObjDiscussionconfiguration sets ObjDiscussionconfiguration field to given value.

### HasObjDiscussionconfiguration

`func (o *DiscussionResponse) HasObjDiscussionconfiguration() bool`

HasObjDiscussionconfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


