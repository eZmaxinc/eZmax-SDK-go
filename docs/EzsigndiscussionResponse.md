# EzsigndiscussionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzsigndiscussionID** | **int32** | The unique ID of the Ezsigndiscussion | 
**FkiEzsignpageID** | **int32** | The unique ID of the Ezsignpage | 
**FkiDiscussionID** | **int32** | The unique ID of the Discussion | 
**IEzsigndiscussionX** | **int32** | The x of the Ezsigndiscussion | 
**IEzsigndiscussionY** | **int32** | The y of the Ezsigndiscussion | 
**IEzsigndiscussionPagenumber** | **int32** | The page number in the Ezsigndocument for the Ezsigndiscussion | 
**ObjDiscussion** | [**DiscussionResponseCompound**](DiscussionResponseCompound.md) |  | 

## Methods

### NewEzsigndiscussionResponse

`func NewEzsigndiscussionResponse(pkiEzsigndiscussionID int32, fkiEzsignpageID int32, fkiDiscussionID int32, iEzsigndiscussionX int32, iEzsigndiscussionY int32, iEzsigndiscussionPagenumber int32, objDiscussion DiscussionResponseCompound, ) *EzsigndiscussionResponse`

NewEzsigndiscussionResponse instantiates a new EzsigndiscussionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndiscussionResponseWithDefaults

`func NewEzsigndiscussionResponseWithDefaults() *EzsigndiscussionResponse`

NewEzsigndiscussionResponseWithDefaults instantiates a new EzsigndiscussionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzsigndiscussionID

`func (o *EzsigndiscussionResponse) GetPkiEzsigndiscussionID() int32`

GetPkiEzsigndiscussionID returns the PkiEzsigndiscussionID field if non-nil, zero value otherwise.

### GetPkiEzsigndiscussionIDOk

`func (o *EzsigndiscussionResponse) GetPkiEzsigndiscussionIDOk() (*int32, bool)`

GetPkiEzsigndiscussionIDOk returns a tuple with the PkiEzsigndiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzsigndiscussionID

`func (o *EzsigndiscussionResponse) SetPkiEzsigndiscussionID(v int32)`

SetPkiEzsigndiscussionID sets PkiEzsigndiscussionID field to given value.


### GetFkiEzsignpageID

`func (o *EzsigndiscussionResponse) GetFkiEzsignpageID() int32`

GetFkiEzsignpageID returns the FkiEzsignpageID field if non-nil, zero value otherwise.

### GetFkiEzsignpageIDOk

`func (o *EzsigndiscussionResponse) GetFkiEzsignpageIDOk() (*int32, bool)`

GetFkiEzsignpageIDOk returns a tuple with the FkiEzsignpageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignpageID

`func (o *EzsigndiscussionResponse) SetFkiEzsignpageID(v int32)`

SetFkiEzsignpageID sets FkiEzsignpageID field to given value.


### GetFkiDiscussionID

`func (o *EzsigndiscussionResponse) GetFkiDiscussionID() int32`

GetFkiDiscussionID returns the FkiDiscussionID field if non-nil, zero value otherwise.

### GetFkiDiscussionIDOk

`func (o *EzsigndiscussionResponse) GetFkiDiscussionIDOk() (*int32, bool)`

GetFkiDiscussionIDOk returns a tuple with the FkiDiscussionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDiscussionID

`func (o *EzsigndiscussionResponse) SetFkiDiscussionID(v int32)`

SetFkiDiscussionID sets FkiDiscussionID field to given value.


### GetIEzsigndiscussionX

`func (o *EzsigndiscussionResponse) GetIEzsigndiscussionX() int32`

GetIEzsigndiscussionX returns the IEzsigndiscussionX field if non-nil, zero value otherwise.

### GetIEzsigndiscussionXOk

`func (o *EzsigndiscussionResponse) GetIEzsigndiscussionXOk() (*int32, bool)`

GetIEzsigndiscussionXOk returns a tuple with the IEzsigndiscussionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndiscussionX

`func (o *EzsigndiscussionResponse) SetIEzsigndiscussionX(v int32)`

SetIEzsigndiscussionX sets IEzsigndiscussionX field to given value.


### GetIEzsigndiscussionY

`func (o *EzsigndiscussionResponse) GetIEzsigndiscussionY() int32`

GetIEzsigndiscussionY returns the IEzsigndiscussionY field if non-nil, zero value otherwise.

### GetIEzsigndiscussionYOk

`func (o *EzsigndiscussionResponse) GetIEzsigndiscussionYOk() (*int32, bool)`

GetIEzsigndiscussionYOk returns a tuple with the IEzsigndiscussionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndiscussionY

`func (o *EzsigndiscussionResponse) SetIEzsigndiscussionY(v int32)`

SetIEzsigndiscussionY sets IEzsigndiscussionY field to given value.


### GetIEzsigndiscussionPagenumber

`func (o *EzsigndiscussionResponse) GetIEzsigndiscussionPagenumber() int32`

GetIEzsigndiscussionPagenumber returns the IEzsigndiscussionPagenumber field if non-nil, zero value otherwise.

### GetIEzsigndiscussionPagenumberOk

`func (o *EzsigndiscussionResponse) GetIEzsigndiscussionPagenumberOk() (*int32, bool)`

GetIEzsigndiscussionPagenumberOk returns a tuple with the IEzsigndiscussionPagenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzsigndiscussionPagenumber

`func (o *EzsigndiscussionResponse) SetIEzsigndiscussionPagenumber(v int32)`

SetIEzsigndiscussionPagenumber sets IEzsigndiscussionPagenumber field to given value.


### GetObjDiscussion

`func (o *EzsigndiscussionResponse) GetObjDiscussion() DiscussionResponseCompound`

GetObjDiscussion returns the ObjDiscussion field if non-nil, zero value otherwise.

### GetObjDiscussionOk

`func (o *EzsigndiscussionResponse) GetObjDiscussionOk() (*DiscussionResponseCompound, bool)`

GetObjDiscussionOk returns a tuple with the ObjDiscussion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDiscussion

`func (o *EzsigndiscussionResponse) SetObjDiscussion(v DiscussionResponseCompound)`

SetObjDiscussion sets ObjDiscussion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


