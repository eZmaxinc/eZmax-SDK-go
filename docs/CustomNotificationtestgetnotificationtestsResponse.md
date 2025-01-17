# CustomNotificationtestgetnotificationtestsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiNotificationtestID** | **int32** | The unique ID of the Notificationtest | 
**ObjNotificationtestName** | [**MultilingualNotificationtestName**](MultilingualNotificationtestName.md) |  | 
**FkiNotificationsubsectionID** | **int32** | The unique ID of the Notificationsubsection | 
**SNotificationtestFunction** | **string** | The function name of the Notificationtest | 
**SNotificationtestNameX** | **string** | The name of the Notificationtest in the language of the requester | 
**ENotificationpreferenceStatus** | [**FieldENotificationpreferenceStatus**](FieldENotificationpreferenceStatus.md) |  | 
**INotificationtest** | **int32** | The number of elements returned by the Notificationtest | 

## Methods

### NewCustomNotificationtestgetnotificationtestsResponse

`func NewCustomNotificationtestgetnotificationtestsResponse(pkiNotificationtestID int32, objNotificationtestName MultilingualNotificationtestName, fkiNotificationsubsectionID int32, sNotificationtestFunction string, sNotificationtestNameX string, eNotificationpreferenceStatus FieldENotificationpreferenceStatus, iNotificationtest int32, ) *CustomNotificationtestgetnotificationtestsResponse`

NewCustomNotificationtestgetnotificationtestsResponse instantiates a new CustomNotificationtestgetnotificationtestsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomNotificationtestgetnotificationtestsResponseWithDefaults

`func NewCustomNotificationtestgetnotificationtestsResponseWithDefaults() *CustomNotificationtestgetnotificationtestsResponse`

NewCustomNotificationtestgetnotificationtestsResponseWithDefaults instantiates a new CustomNotificationtestgetnotificationtestsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiNotificationtestID

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetPkiNotificationtestID() int32`

GetPkiNotificationtestID returns the PkiNotificationtestID field if non-nil, zero value otherwise.

### GetPkiNotificationtestIDOk

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetPkiNotificationtestIDOk() (*int32, bool)`

GetPkiNotificationtestIDOk returns a tuple with the PkiNotificationtestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiNotificationtestID

`func (o *CustomNotificationtestgetnotificationtestsResponse) SetPkiNotificationtestID(v int32)`

SetPkiNotificationtestID sets PkiNotificationtestID field to given value.


### GetObjNotificationtestName

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetObjNotificationtestName() MultilingualNotificationtestName`

GetObjNotificationtestName returns the ObjNotificationtestName field if non-nil, zero value otherwise.

### GetObjNotificationtestNameOk

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetObjNotificationtestNameOk() (*MultilingualNotificationtestName, bool)`

GetObjNotificationtestNameOk returns a tuple with the ObjNotificationtestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjNotificationtestName

`func (o *CustomNotificationtestgetnotificationtestsResponse) SetObjNotificationtestName(v MultilingualNotificationtestName)`

SetObjNotificationtestName sets ObjNotificationtestName field to given value.


### GetFkiNotificationsubsectionID

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetFkiNotificationsubsectionID() int32`

GetFkiNotificationsubsectionID returns the FkiNotificationsubsectionID field if non-nil, zero value otherwise.

### GetFkiNotificationsubsectionIDOk

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetFkiNotificationsubsectionIDOk() (*int32, bool)`

GetFkiNotificationsubsectionIDOk returns a tuple with the FkiNotificationsubsectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotificationsubsectionID

`func (o *CustomNotificationtestgetnotificationtestsResponse) SetFkiNotificationsubsectionID(v int32)`

SetFkiNotificationsubsectionID sets FkiNotificationsubsectionID field to given value.


### GetSNotificationtestFunction

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetSNotificationtestFunction() string`

GetSNotificationtestFunction returns the SNotificationtestFunction field if non-nil, zero value otherwise.

### GetSNotificationtestFunctionOk

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetSNotificationtestFunctionOk() (*string, bool)`

GetSNotificationtestFunctionOk returns a tuple with the SNotificationtestFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNotificationtestFunction

`func (o *CustomNotificationtestgetnotificationtestsResponse) SetSNotificationtestFunction(v string)`

SetSNotificationtestFunction sets SNotificationtestFunction field to given value.


### GetSNotificationtestNameX

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetSNotificationtestNameX() string`

GetSNotificationtestNameX returns the SNotificationtestNameX field if non-nil, zero value otherwise.

### GetSNotificationtestNameXOk

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetSNotificationtestNameXOk() (*string, bool)`

GetSNotificationtestNameXOk returns a tuple with the SNotificationtestNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNotificationtestNameX

`func (o *CustomNotificationtestgetnotificationtestsResponse) SetSNotificationtestNameX(v string)`

SetSNotificationtestNameX sets SNotificationtestNameX field to given value.


### GetENotificationpreferenceStatus

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetENotificationpreferenceStatus() FieldENotificationpreferenceStatus`

GetENotificationpreferenceStatus returns the ENotificationpreferenceStatus field if non-nil, zero value otherwise.

### GetENotificationpreferenceStatusOk

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetENotificationpreferenceStatusOk() (*FieldENotificationpreferenceStatus, bool)`

GetENotificationpreferenceStatusOk returns a tuple with the ENotificationpreferenceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENotificationpreferenceStatus

`func (o *CustomNotificationtestgetnotificationtestsResponse) SetENotificationpreferenceStatus(v FieldENotificationpreferenceStatus)`

SetENotificationpreferenceStatus sets ENotificationpreferenceStatus field to given value.


### GetINotificationtest

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetINotificationtest() int32`

GetINotificationtest returns the INotificationtest field if non-nil, zero value otherwise.

### GetINotificationtestOk

`func (o *CustomNotificationtestgetnotificationtestsResponse) GetINotificationtestOk() (*int32, bool)`

GetINotificationtestOk returns a tuple with the INotificationtest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetINotificationtest

`func (o *CustomNotificationtestgetnotificationtestsResponse) SetINotificationtest(v int32)`

SetINotificationtest sets INotificationtest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


