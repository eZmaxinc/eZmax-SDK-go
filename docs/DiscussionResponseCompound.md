# DiscussionResponseCompound

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
**AObjDiscussionmembership** | [**[]DiscussionmembershipResponseCompound**](DiscussionmembershipResponseCompound.md) |  | 
**AObjDiscussionmessage** | [**[]DiscussionmessageResponseCompound**](DiscussionmessageResponseCompound.md) |  | 

## Methods

### NewDiscussionResponseCompound

`func NewDiscussionResponseCompound(pkiDiscussionID int32, sDiscussionDescription string, bDiscussionClosed bool, iDiscussionmessageCount int32, iDiscussionmessageCountunread int32, aObjDiscussionmembership []DiscussionmembershipResponseCompound, aObjDiscussionmessage []DiscussionmessageResponseCompound, ) *DiscussionResponseCompound`

NewDiscussionResponseCompound instantiates a new DiscussionResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscussionResponseCompoundWithDefaults

`func NewDiscussionResponseCompoundWithDefaults() *DiscussionResponseCompound`

NewDiscussionResponseCompoundWithDefaults instantiates a new DiscussionResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDiscussionID

`func (o *DiscussionResponseCompound) GetPkiDiscussionID() int32`

GetPkiDiscussionID returns the PkiDiscussionID field if non-nil, zero value otherwise.

### GetPkiDiscussionIDOk

`func (o *DiscussionResponseCompound) GetPkiDiscussionIDOk() (*int32, bool)`

GetPkiDiscussionIDOk returns a tuple with the PkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDiscussionID

`func (o *DiscussionResponseCompound) SetPkiDiscussionID(v int32)`

SetPkiDiscussionID sets PkiDiscussionID field to given value.


### GetSDiscussionDescription

`func (o *DiscussionResponseCompound) GetSDiscussionDescription() string`

GetSDiscussionDescription returns the SDiscussionDescription field if non-nil, zero value otherwise.

### GetSDiscussionDescriptionOk

`func (o *DiscussionResponseCompound) GetSDiscussionDescriptionOk() (*string, bool)`

GetSDiscussionDescriptionOk returns a tuple with the SDiscussionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDiscussionDescription

`func (o *DiscussionResponseCompound) SetSDiscussionDescription(v string)`

SetSDiscussionDescription sets SDiscussionDescription field to given value.


### GetBDiscussionClosed

`func (o *DiscussionResponseCompound) GetBDiscussionClosed() bool`

GetBDiscussionClosed returns the BDiscussionClosed field if non-nil, zero value otherwise.

### GetBDiscussionClosedOk

`func (o *DiscussionResponseCompound) GetBDiscussionClosedOk() (*bool, bool)`

GetBDiscussionClosedOk returns a tuple with the BDiscussionClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDiscussionClosed

`func (o *DiscussionResponseCompound) SetBDiscussionClosed(v bool)`

SetBDiscussionClosed sets BDiscussionClosed field to given value.


### GetDtDiscussionLastread

`func (o *DiscussionResponseCompound) GetDtDiscussionLastread() string`

GetDtDiscussionLastread returns the DtDiscussionLastread field if non-nil, zero value otherwise.

### GetDtDiscussionLastreadOk

`func (o *DiscussionResponseCompound) GetDtDiscussionLastreadOk() (*string, bool)`

GetDtDiscussionLastreadOk returns a tuple with the DtDiscussionLastread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtDiscussionLastread

`func (o *DiscussionResponseCompound) SetDtDiscussionLastread(v string)`

SetDtDiscussionLastread sets DtDiscussionLastread field to given value.

### HasDtDiscussionLastread

`func (o *DiscussionResponseCompound) HasDtDiscussionLastread() bool`

HasDtDiscussionLastread returns a boolean if a field has been set.

### GetIDiscussionmessageCount

`func (o *DiscussionResponseCompound) GetIDiscussionmessageCount() int32`

GetIDiscussionmessageCount returns the IDiscussionmessageCount field if non-nil, zero value otherwise.

### GetIDiscussionmessageCountOk

`func (o *DiscussionResponseCompound) GetIDiscussionmessageCountOk() (*int32, bool)`

GetIDiscussionmessageCountOk returns a tuple with the IDiscussionmessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIDiscussionmessageCount

`func (o *DiscussionResponseCompound) SetIDiscussionmessageCount(v int32)`

SetIDiscussionmessageCount sets IDiscussionmessageCount field to given value.


### GetIDiscussionmessageCountunread

`func (o *DiscussionResponseCompound) GetIDiscussionmessageCountunread() int32`

GetIDiscussionmessageCountunread returns the IDiscussionmessageCountunread field if non-nil, zero value otherwise.

### GetIDiscussionmessageCountunreadOk

`func (o *DiscussionResponseCompound) GetIDiscussionmessageCountunreadOk() (*int32, bool)`

GetIDiscussionmessageCountunreadOk returns a tuple with the IDiscussionmessageCountunread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIDiscussionmessageCountunread

`func (o *DiscussionResponseCompound) SetIDiscussionmessageCountunread(v int32)`

SetIDiscussionmessageCountunread sets IDiscussionmessageCountunread field to given value.


### GetObjDiscussionconfiguration

`func (o *DiscussionResponseCompound) GetObjDiscussionconfiguration() map[string]interface{}`

GetObjDiscussionconfiguration returns the ObjDiscussionconfiguration field if non-nil, zero value otherwise.

### GetObjDiscussionconfigurationOk

`func (o *DiscussionResponseCompound) GetObjDiscussionconfigurationOk() (*map[string]interface{}, bool)`

GetObjDiscussionconfigurationOk returns a tuple with the ObjDiscussionconfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDiscussionconfiguration

`func (o *DiscussionResponseCompound) SetObjDiscussionconfiguration(v map[string]interface{})`

SetObjDiscussionconfiguration sets ObjDiscussionconfiguration field to given value.

### HasObjDiscussionconfiguration

`func (o *DiscussionResponseCompound) HasObjDiscussionconfiguration() bool`

HasObjDiscussionconfiguration returns a boolean if a field has been set.

### GetAObjDiscussionmembership

`func (o *DiscussionResponseCompound) GetAObjDiscussionmembership() []DiscussionmembershipResponseCompound`

GetAObjDiscussionmembership returns the AObjDiscussionmembership field if non-nil, zero value otherwise.

### GetAObjDiscussionmembershipOk

`func (o *DiscussionResponseCompound) GetAObjDiscussionmembershipOk() (*[]DiscussionmembershipResponseCompound, bool)`

GetAObjDiscussionmembershipOk returns a tuple with the AObjDiscussionmembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjDiscussionmembership

`func (o *DiscussionResponseCompound) SetAObjDiscussionmembership(v []DiscussionmembershipResponseCompound)`

SetAObjDiscussionmembership sets AObjDiscussionmembership field to given value.


### GetAObjDiscussionmessage

`func (o *DiscussionResponseCompound) GetAObjDiscussionmessage() []DiscussionmessageResponseCompound`

GetAObjDiscussionmessage returns the AObjDiscussionmessage field if non-nil, zero value otherwise.

### GetAObjDiscussionmessageOk

`func (o *DiscussionResponseCompound) GetAObjDiscussionmessageOk() (*[]DiscussionmessageResponseCompound, bool)`

GetAObjDiscussionmessageOk returns a tuple with the AObjDiscussionmessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjDiscussionmessage

`func (o *DiscussionResponseCompound) SetAObjDiscussionmessage(v []DiscussionmessageResponseCompound)`

SetAObjDiscussionmessage sets AObjDiscussionmessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


