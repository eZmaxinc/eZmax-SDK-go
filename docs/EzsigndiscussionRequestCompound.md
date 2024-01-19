# EzsigndiscussionRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigndiscussionID** | Pointer to **int32** | The unique ID of the Ezsigndiscussion | [optional] 
**FkiEzsigndocumentID** | **int32** | The unique ID of the Ezsigndocument | 
**IEzsigndiscussionPagenumber** | **int32** | The page number in the Ezsigndocument for the Ezsigndiscussion | 
**IEzsigndiscussionX** | **int32** | The x of the Ezsigndiscussion | 
**IEzsigndiscussionY** | **int32** | The y of the Ezsigndiscussion | 
**ObjDiscussion** | [**DiscussionRequest**](DiscussionRequest.md) |  | 

## Methods

### NewEzsigndiscussionRequestCompound

`func NewEzsigndiscussionRequestCompound(fkiEzsigndocumentID int32, iEzsigndiscussionPagenumber int32, iEzsigndiscussionX int32, iEzsigndiscussionY int32, objDiscussion DiscussionRequest, ) *EzsigndiscussionRequestCompound`

NewEzsigndiscussionRequestCompound instantiates a new EzsigndiscussionRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndiscussionRequestCompoundWithDefaults

`func NewEzsigndiscussionRequestCompoundWithDefaults() *EzsigndiscussionRequestCompound`

NewEzsigndiscussionRequestCompoundWithDefaults instantiates a new EzsigndiscussionRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndiscussionID

`func (o *EzsigndiscussionRequestCompound) GetPkiEzsigndiscussionID() int32`

GetPkiEzsigndiscussionID returns the PkiEzsigndiscussionID field if non-nil, zero value otherwise.

### GetPkiEzsigndiscussionIDOk

`func (o *EzsigndiscussionRequestCompound) GetPkiEzsigndiscussionIDOk() (*int32, bool)`

GetPkiEzsigndiscussionIDOk returns a tuple with the PkiEzsigndiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndiscussionID

`func (o *EzsigndiscussionRequestCompound) SetPkiEzsigndiscussionID(v int32)`

SetPkiEzsigndiscussionID sets PkiEzsigndiscussionID field to given value.

### HasPkiEzsigndiscussionID

`func (o *EzsigndiscussionRequestCompound) HasPkiEzsigndiscussionID() bool`

HasPkiEzsigndiscussionID returns a boolean if a field has been set.

### GetFkiEzsigndocumentID

`func (o *EzsigndiscussionRequestCompound) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *EzsigndiscussionRequestCompound) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *EzsigndiscussionRequestCompound) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.


### GetIEzsigndiscussionPagenumber

`func (o *EzsigndiscussionRequestCompound) GetIEzsigndiscussionPagenumber() int32`

GetIEzsigndiscussionPagenumber returns the IEzsigndiscussionPagenumber field if non-nil, zero value otherwise.

### GetIEzsigndiscussionPagenumberOk

`func (o *EzsigndiscussionRequestCompound) GetIEzsigndiscussionPagenumberOk() (*int32, bool)`

GetIEzsigndiscussionPagenumberOk returns a tuple with the IEzsigndiscussionPagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndiscussionPagenumber

`func (o *EzsigndiscussionRequestCompound) SetIEzsigndiscussionPagenumber(v int32)`

SetIEzsigndiscussionPagenumber sets IEzsigndiscussionPagenumber field to given value.


### GetIEzsigndiscussionX

`func (o *EzsigndiscussionRequestCompound) GetIEzsigndiscussionX() int32`

GetIEzsigndiscussionX returns the IEzsigndiscussionX field if non-nil, zero value otherwise.

### GetIEzsigndiscussionXOk

`func (o *EzsigndiscussionRequestCompound) GetIEzsigndiscussionXOk() (*int32, bool)`

GetIEzsigndiscussionXOk returns a tuple with the IEzsigndiscussionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndiscussionX

`func (o *EzsigndiscussionRequestCompound) SetIEzsigndiscussionX(v int32)`

SetIEzsigndiscussionX sets IEzsigndiscussionX field to given value.


### GetIEzsigndiscussionY

`func (o *EzsigndiscussionRequestCompound) GetIEzsigndiscussionY() int32`

GetIEzsigndiscussionY returns the IEzsigndiscussionY field if non-nil, zero value otherwise.

### GetIEzsigndiscussionYOk

`func (o *EzsigndiscussionRequestCompound) GetIEzsigndiscussionYOk() (*int32, bool)`

GetIEzsigndiscussionYOk returns a tuple with the IEzsigndiscussionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndiscussionY

`func (o *EzsigndiscussionRequestCompound) SetIEzsigndiscussionY(v int32)`

SetIEzsigndiscussionY sets IEzsigndiscussionY field to given value.


### GetObjDiscussion

`func (o *EzsigndiscussionRequestCompound) GetObjDiscussion() DiscussionRequest`

GetObjDiscussion returns the ObjDiscussion field if non-nil, zero value otherwise.

### GetObjDiscussionOk

`func (o *EzsigndiscussionRequestCompound) GetObjDiscussionOk() (*DiscussionRequest, bool)`

GetObjDiscussionOk returns a tuple with the ObjDiscussion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDiscussion

`func (o *EzsigndiscussionRequestCompound) SetObjDiscussion(v DiscussionRequest)`

SetObjDiscussion sets ObjDiscussion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


