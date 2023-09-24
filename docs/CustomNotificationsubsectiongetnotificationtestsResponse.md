# CustomNotificationsubsectiongetnotificationtestsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiNotificationsubsectionID** | **int32** | The unique ID of the Notificationsubsection | 
**FkiNotificationsectionID** | **int32** | The unique ID of the Notificationsection | 
**ObjNotificationsubsectionName** | Pointer to [**MultilingualNotificationsubsectionName**](MultilingualNotificationsubsectionName.md) |  | [optional] 
**SNotificationsectionNameX** | Pointer to **string** | The name of the Notificationsection in the language of the requester | [optional] 
**SNotificationsubsectionNameX** | **string** | The name of the Notificationsubsection in the language of the requester | 
**AObjNotificationtest** | [**[]CustomNotificationtestgetnotificationtestsResponse**](CustomNotificationtestgetnotificationtestsResponse.md) |  | 

## Methods

### NewCustomNotificationsubsectiongetnotificationtestsResponse

`func NewCustomNotificationsubsectiongetnotificationtestsResponse(pkiNotificationsubsectionID int32, fkiNotificationsectionID int32, sNotificationsubsectionNameX string, aObjNotificationtest []CustomNotificationtestgetnotificationtestsResponse, ) *CustomNotificationsubsectiongetnotificationtestsResponse`

NewCustomNotificationsubsectiongetnotificationtestsResponse instantiates a new CustomNotificationsubsectiongetnotificationtestsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomNotificationsubsectiongetnotificationtestsResponseWithDefaults

`func NewCustomNotificationsubsectiongetnotificationtestsResponseWithDefaults() *CustomNotificationsubsectiongetnotificationtestsResponse`

NewCustomNotificationsubsectiongetnotificationtestsResponseWithDefaults instantiates a new CustomNotificationsubsectiongetnotificationtestsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiNotificationsubsectionID

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetPkiNotificationsubsectionID() int32`

GetPkiNotificationsubsectionID returns the PkiNotificationsubsectionID field if non-nil, zero value otherwise.

### GetPkiNotificationsubsectionIDOk

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetPkiNotificationsubsectionIDOk() (*int32, bool)`

GetPkiNotificationsubsectionIDOk returns a tuple with the PkiNotificationsubsectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiNotificationsubsectionID

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetPkiNotificationsubsectionID(v int32)`

SetPkiNotificationsubsectionID sets PkiNotificationsubsectionID field to given value.


### GetFkiNotificationsectionID

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetFkiNotificationsectionID() int32`

GetFkiNotificationsectionID returns the FkiNotificationsectionID field if non-nil, zero value otherwise.

### GetFkiNotificationsectionIDOk

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetFkiNotificationsectionIDOk() (*int32, bool)`

GetFkiNotificationsectionIDOk returns a tuple with the FkiNotificationsectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotificationsectionID

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetFkiNotificationsectionID(v int32)`

SetFkiNotificationsectionID sets FkiNotificationsectionID field to given value.


### GetObjNotificationsubsectionName

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetObjNotificationsubsectionName() MultilingualNotificationsubsectionName`

GetObjNotificationsubsectionName returns the ObjNotificationsubsectionName field if non-nil, zero value otherwise.

### GetObjNotificationsubsectionNameOk

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetObjNotificationsubsectionNameOk() (*MultilingualNotificationsubsectionName, bool)`

GetObjNotificationsubsectionNameOk returns a tuple with the ObjNotificationsubsectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjNotificationsubsectionName

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetObjNotificationsubsectionName(v MultilingualNotificationsubsectionName)`

SetObjNotificationsubsectionName sets ObjNotificationsubsectionName field to given value.

### HasObjNotificationsubsectionName

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) HasObjNotificationsubsectionName() bool`

HasObjNotificationsubsectionName returns a boolean if a field has been set.

### GetSNotificationsectionNameX

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetSNotificationsectionNameX() string`

GetSNotificationsectionNameX returns the SNotificationsectionNameX field if non-nil, zero value otherwise.

### GetSNotificationsectionNameXOk

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetSNotificationsectionNameXOk() (*string, bool)`

GetSNotificationsectionNameXOk returns a tuple with the SNotificationsectionNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNotificationsectionNameX

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetSNotificationsectionNameX(v string)`

SetSNotificationsectionNameX sets SNotificationsectionNameX field to given value.

### HasSNotificationsectionNameX

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) HasSNotificationsectionNameX() bool`

HasSNotificationsectionNameX returns a boolean if a field has been set.

### GetSNotificationsubsectionNameX

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetSNotificationsubsectionNameX() string`

GetSNotificationsubsectionNameX returns the SNotificationsubsectionNameX field if non-nil, zero value otherwise.

### GetSNotificationsubsectionNameXOk

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetSNotificationsubsectionNameXOk() (*string, bool)`

GetSNotificationsubsectionNameXOk returns a tuple with the SNotificationsubsectionNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNotificationsubsectionNameX

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetSNotificationsubsectionNameX(v string)`

SetSNotificationsubsectionNameX sets SNotificationsubsectionNameX field to given value.


### GetAObjNotificationtest

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetAObjNotificationtest() []CustomNotificationtestgetnotificationtestsResponse`

GetAObjNotificationtest returns the AObjNotificationtest field if non-nil, zero value otherwise.

### GetAObjNotificationtestOk

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) GetAObjNotificationtestOk() (*[]CustomNotificationtestgetnotificationtestsResponse, bool)`

GetAObjNotificationtestOk returns a tuple with the AObjNotificationtest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjNotificationtest

`func (o *CustomNotificationsubsectiongetnotificationtestsResponse) SetAObjNotificationtest(v []CustomNotificationtestgetnotificationtestsResponse)`

SetAObjNotificationtest sets AObjNotificationtest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


