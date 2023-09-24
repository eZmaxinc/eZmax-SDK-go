# NotificationtestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiNotificationtestID** | **int32** | The unique ID of the Notificationtest | 
**ObjNotificationtestName** | [**MultilingualNotificationtestName**](MultilingualNotificationtestName.md) |  | 
**FkiNotificationsubsectionID** | **int32** | The unique ID of the Notificationsubsection | 
**SNotificationtestFunction** | **string** | The function name of the Notificationtest | 
**SNotificationtestNameX** | **string** | The name of the Notificationtest in the language of the requester | 

## Methods

### NewNotificationtestResponse

`func NewNotificationtestResponse(pkiNotificationtestID int32, objNotificationtestName MultilingualNotificationtestName, fkiNotificationsubsectionID int32, sNotificationtestFunction string, sNotificationtestNameX string, ) *NotificationtestResponse`

NewNotificationtestResponse instantiates a new NotificationtestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationtestResponseWithDefaults

`func NewNotificationtestResponseWithDefaults() *NotificationtestResponse`

NewNotificationtestResponseWithDefaults instantiates a new NotificationtestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiNotificationtestID

`func (o *NotificationtestResponse) GetPkiNotificationtestID() int32`

GetPkiNotificationtestID returns the PkiNotificationtestID field if non-nil, zero value otherwise.

### GetPkiNotificationtestIDOk

`func (o *NotificationtestResponse) GetPkiNotificationtestIDOk() (*int32, bool)`

GetPkiNotificationtestIDOk returns a tuple with the PkiNotificationtestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiNotificationtestID

`func (o *NotificationtestResponse) SetPkiNotificationtestID(v int32)`

SetPkiNotificationtestID sets PkiNotificationtestID field to given value.


### GetObjNotificationtestName

`func (o *NotificationtestResponse) GetObjNotificationtestName() MultilingualNotificationtestName`

GetObjNotificationtestName returns the ObjNotificationtestName field if non-nil, zero value otherwise.

### GetObjNotificationtestNameOk

`func (o *NotificationtestResponse) GetObjNotificationtestNameOk() (*MultilingualNotificationtestName, bool)`

GetObjNotificationtestNameOk returns a tuple with the ObjNotificationtestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjNotificationtestName

`func (o *NotificationtestResponse) SetObjNotificationtestName(v MultilingualNotificationtestName)`

SetObjNotificationtestName sets ObjNotificationtestName field to given value.


### GetFkiNotificationsubsectionID

`func (o *NotificationtestResponse) GetFkiNotificationsubsectionID() int32`

GetFkiNotificationsubsectionID returns the FkiNotificationsubsectionID field if non-nil, zero value otherwise.

### GetFkiNotificationsubsectionIDOk

`func (o *NotificationtestResponse) GetFkiNotificationsubsectionIDOk() (*int32, bool)`

GetFkiNotificationsubsectionIDOk returns a tuple with the FkiNotificationsubsectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotificationsubsectionID

`func (o *NotificationtestResponse) SetFkiNotificationsubsectionID(v int32)`

SetFkiNotificationsubsectionID sets FkiNotificationsubsectionID field to given value.


### GetSNotificationtestFunction

`func (o *NotificationtestResponse) GetSNotificationtestFunction() string`

GetSNotificationtestFunction returns the SNotificationtestFunction field if non-nil, zero value otherwise.

### GetSNotificationtestFunctionOk

`func (o *NotificationtestResponse) GetSNotificationtestFunctionOk() (*string, bool)`

GetSNotificationtestFunctionOk returns a tuple with the SNotificationtestFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNotificationtestFunction

`func (o *NotificationtestResponse) SetSNotificationtestFunction(v string)`

SetSNotificationtestFunction sets SNotificationtestFunction field to given value.


### GetSNotificationtestNameX

`func (o *NotificationtestResponse) GetSNotificationtestNameX() string`

GetSNotificationtestNameX returns the SNotificationtestNameX field if non-nil, zero value otherwise.

### GetSNotificationtestNameXOk

`func (o *NotificationtestResponse) GetSNotificationtestNameXOk() (*string, bool)`

GetSNotificationtestNameXOk returns a tuple with the SNotificationtestNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNotificationtestNameX

`func (o *NotificationtestResponse) SetSNotificationtestNameX(v string)`

SetSNotificationtestNameX sets SNotificationtestNameX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


