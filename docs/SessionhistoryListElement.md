# SessionhistoryListElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiSessionhistoryID** | **int32** | The unique ID of the Sessionhistory | 
**FkiComputerID** | Pointer to **int32** | The unique ID of the Computer | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**DtSessionhistoryFirsthit** | **string** | The first hit of the Sessionhistory | 
**DtSessionhistoryLasthit** | **string** | The last hit of the Sessionhistory | 
**ESessionhistoryEndby** | [**FieldESessionhistoryEndby**](FieldESessionhistoryEndby.md) |  | 
**SComputerDescription** | Pointer to **string** | The description of the Computer | [optional] 
**SSessionhistoryDuration** | **string** | The duration of the session | 
**SSessionhistoryIP** | **string** | Represent an IP address. | 
**SUserLoginname** | Pointer to **string** | The login name of the User. | [optional] 

## Methods

### NewSessionhistoryListElement

`func NewSessionhistoryListElement(pkiSessionhistoryID int32, dtSessionhistoryFirsthit string, dtSessionhistoryLasthit string, eSessionhistoryEndby FieldESessionhistoryEndby, sSessionhistoryDuration string, sSessionhistoryIP string, ) *SessionhistoryListElement`

NewSessionhistoryListElement instantiates a new SessionhistoryListElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionhistoryListElementWithDefaults

`func NewSessionhistoryListElementWithDefaults() *SessionhistoryListElement`

NewSessionhistoryListElementWithDefaults instantiates a new SessionhistoryListElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiSessionhistoryID

`func (o *SessionhistoryListElement) GetPkiSessionhistoryID() int32`

GetPkiSessionhistoryID returns the PkiSessionhistoryID field if non-nil, zero value otherwise.

### GetPkiSessionhistoryIDOk

`func (o *SessionhistoryListElement) GetPkiSessionhistoryIDOk() (*int32, bool)`

GetPkiSessionhistoryIDOk returns a tuple with the PkiSessionhistoryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiSessionhistoryID

`func (o *SessionhistoryListElement) SetPkiSessionhistoryID(v int32)`

SetPkiSessionhistoryID sets PkiSessionhistoryID field to given value.


### GetFkiComputerID

`func (o *SessionhistoryListElement) GetFkiComputerID() int32`

GetFkiComputerID returns the FkiComputerID field if non-nil, zero value otherwise.

### GetFkiComputerIDOk

`func (o *SessionhistoryListElement) GetFkiComputerIDOk() (*int32, bool)`

GetFkiComputerIDOk returns a tuple with the FkiComputerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiComputerID

`func (o *SessionhistoryListElement) SetFkiComputerID(v int32)`

SetFkiComputerID sets FkiComputerID field to given value.

### HasFkiComputerID

`func (o *SessionhistoryListElement) HasFkiComputerID() bool`

HasFkiComputerID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *SessionhistoryListElement) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *SessionhistoryListElement) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *SessionhistoryListElement) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *SessionhistoryListElement) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetDtSessionhistoryFirsthit

`func (o *SessionhistoryListElement) GetDtSessionhistoryFirsthit() string`

GetDtSessionhistoryFirsthit returns the DtSessionhistoryFirsthit field if non-nil, zero value otherwise.

### GetDtSessionhistoryFirsthitOk

`func (o *SessionhistoryListElement) GetDtSessionhistoryFirsthitOk() (*string, bool)`

GetDtSessionhistoryFirsthitOk returns a tuple with the DtSessionhistoryFirsthit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSessionhistoryFirsthit

`func (o *SessionhistoryListElement) SetDtSessionhistoryFirsthit(v string)`

SetDtSessionhistoryFirsthit sets DtSessionhistoryFirsthit field to given value.


### GetDtSessionhistoryLasthit

`func (o *SessionhistoryListElement) GetDtSessionhistoryLasthit() string`

GetDtSessionhistoryLasthit returns the DtSessionhistoryLasthit field if non-nil, zero value otherwise.

### GetDtSessionhistoryLasthitOk

`func (o *SessionhistoryListElement) GetDtSessionhistoryLasthitOk() (*string, bool)`

GetDtSessionhistoryLasthitOk returns a tuple with the DtSessionhistoryLasthit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtSessionhistoryLasthit

`func (o *SessionhistoryListElement) SetDtSessionhistoryLasthit(v string)`

SetDtSessionhistoryLasthit sets DtSessionhistoryLasthit field to given value.


### GetESessionhistoryEndby

`func (o *SessionhistoryListElement) GetESessionhistoryEndby() FieldESessionhistoryEndby`

GetESessionhistoryEndby returns the ESessionhistoryEndby field if non-nil, zero value otherwise.

### GetESessionhistoryEndbyOk

`func (o *SessionhistoryListElement) GetESessionhistoryEndbyOk() (*FieldESessionhistoryEndby, bool)`

GetESessionhistoryEndbyOk returns a tuple with the ESessionhistoryEndby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetESessionhistoryEndby

`func (o *SessionhistoryListElement) SetESessionhistoryEndby(v FieldESessionhistoryEndby)`

SetESessionhistoryEndby sets ESessionhistoryEndby field to given value.


### GetSComputerDescription

`func (o *SessionhistoryListElement) GetSComputerDescription() string`

GetSComputerDescription returns the SComputerDescription field if non-nil, zero value otherwise.

### GetSComputerDescriptionOk

`func (o *SessionhistoryListElement) GetSComputerDescriptionOk() (*string, bool)`

GetSComputerDescriptionOk returns a tuple with the SComputerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSComputerDescription

`func (o *SessionhistoryListElement) SetSComputerDescription(v string)`

SetSComputerDescription sets SComputerDescription field to given value.

### HasSComputerDescription

`func (o *SessionhistoryListElement) HasSComputerDescription() bool`

HasSComputerDescription returns a boolean if a field has been set.

### GetSSessionhistoryDuration

`func (o *SessionhistoryListElement) GetSSessionhistoryDuration() string`

GetSSessionhistoryDuration returns the SSessionhistoryDuration field if non-nil, zero value otherwise.

### GetSSessionhistoryDurationOk

`func (o *SessionhistoryListElement) GetSSessionhistoryDurationOk() (*string, bool)`

GetSSessionhistoryDurationOk returns a tuple with the SSessionhistoryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSessionhistoryDuration

`func (o *SessionhistoryListElement) SetSSessionhistoryDuration(v string)`

SetSSessionhistoryDuration sets SSessionhistoryDuration field to given value.


### GetSSessionhistoryIP

`func (o *SessionhistoryListElement) GetSSessionhistoryIP() string`

GetSSessionhistoryIP returns the SSessionhistoryIP field if non-nil, zero value otherwise.

### GetSSessionhistoryIPOk

`func (o *SessionhistoryListElement) GetSSessionhistoryIPOk() (*string, bool)`

GetSSessionhistoryIPOk returns a tuple with the SSessionhistoryIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSessionhistoryIP

`func (o *SessionhistoryListElement) SetSSessionhistoryIP(v string)`

SetSSessionhistoryIP sets SSessionhistoryIP field to given value.


### GetSUserLoginname

`func (o *SessionhistoryListElement) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *SessionhistoryListElement) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *SessionhistoryListElement) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.

### HasSUserLoginname

`func (o *SessionhistoryListElement) HasSUserLoginname() bool`

HasSUserLoginname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


