# CustomEzmaxcustomeruserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjEzmaxcustomer** | [**CustomEzmaxcustomerResponse**](CustomEzmaxcustomerResponse.md) |  | 
**FkiContacttitleID** | **int32** | The unique ID of the Contacttitle.  Valid values:  |Value|Description| |-|-| |1|Ms.| |2|Mr.| |4|(Blank)| |5|Me (For Notaries)| | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**ObjEmail** | Pointer to [**EmailResponseCompound**](EmailResponseCompound.md) |  | [optional] 
**ObjPhone** | Pointer to [**PhoneResponseCompound**](PhoneResponseCompound.md) |  | [optional] 
**SEzmaxcustomeruserFirstname** | **string** | The First name of the Ezmaxcustomeruser | 
**SEzmaxcustomeruserLastname** | **string** | The First name of the Ezmaxcustomeruser | 

## Methods

### NewCustomEzmaxcustomeruserResponse

`func NewCustomEzmaxcustomeruserResponse(objEzmaxcustomer CustomEzmaxcustomerResponse, fkiContacttitleID int32, fkiLanguageID int32, sEzmaxcustomeruserFirstname string, sEzmaxcustomeruserLastname string, ) *CustomEzmaxcustomeruserResponse`

NewCustomEzmaxcustomeruserResponse instantiates a new CustomEzmaxcustomeruserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzmaxcustomeruserResponseWithDefaults

`func NewCustomEzmaxcustomeruserResponseWithDefaults() *CustomEzmaxcustomeruserResponse`

NewCustomEzmaxcustomeruserResponseWithDefaults instantiates a new CustomEzmaxcustomeruserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjEzmaxcustomer

`func (o *CustomEzmaxcustomeruserResponse) GetObjEzmaxcustomer() CustomEzmaxcustomerResponse`

GetObjEzmaxcustomer returns the ObjEzmaxcustomer field if non-nil, zero value otherwise.

### GetObjEzmaxcustomerOk

`func (o *CustomEzmaxcustomeruserResponse) GetObjEzmaxcustomerOk() (*CustomEzmaxcustomerResponse, bool)`

GetObjEzmaxcustomerOk returns a tuple with the ObjEzmaxcustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEzmaxcustomer

`func (o *CustomEzmaxcustomeruserResponse) SetObjEzmaxcustomer(v CustomEzmaxcustomerResponse)`

SetObjEzmaxcustomer sets ObjEzmaxcustomer field to given value.


### GetFkiContacttitleID

`func (o *CustomEzmaxcustomeruserResponse) GetFkiContacttitleID() int32`

GetFkiContacttitleID returns the FkiContacttitleID field if non-nil, zero value otherwise.

### GetFkiContacttitleIDOk

`func (o *CustomEzmaxcustomeruserResponse) GetFkiContacttitleIDOk() (*int32, bool)`

GetFkiContacttitleIDOk returns a tuple with the FkiContacttitleID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContacttitleID

`func (o *CustomEzmaxcustomeruserResponse) SetFkiContacttitleID(v int32)`

SetFkiContacttitleID sets FkiContacttitleID field to given value.


### GetFkiLanguageID

`func (o *CustomEzmaxcustomeruserResponse) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *CustomEzmaxcustomeruserResponse) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *CustomEzmaxcustomeruserResponse) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetObjEmail

`func (o *CustomEzmaxcustomeruserResponse) GetObjEmail() EmailResponseCompound`

GetObjEmail returns the ObjEmail field if non-nil, zero value otherwise.

### GetObjEmailOk

`func (o *CustomEzmaxcustomeruserResponse) GetObjEmailOk() (*EmailResponseCompound, bool)`

GetObjEmailOk returns a tuple with the ObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmail

`func (o *CustomEzmaxcustomeruserResponse) SetObjEmail(v EmailResponseCompound)`

SetObjEmail sets ObjEmail field to given value.

### HasObjEmail

`func (o *CustomEzmaxcustomeruserResponse) HasObjEmail() bool`

HasObjEmail returns a boolean if a field has been set.

### GetObjPhone

`func (o *CustomEzmaxcustomeruserResponse) GetObjPhone() PhoneResponseCompound`

GetObjPhone returns the ObjPhone field if non-nil, zero value otherwise.

### GetObjPhoneOk

`func (o *CustomEzmaxcustomeruserResponse) GetObjPhoneOk() (*PhoneResponseCompound, bool)`

GetObjPhoneOk returns a tuple with the ObjPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhone

`func (o *CustomEzmaxcustomeruserResponse) SetObjPhone(v PhoneResponseCompound)`

SetObjPhone sets ObjPhone field to given value.

### HasObjPhone

`func (o *CustomEzmaxcustomeruserResponse) HasObjPhone() bool`

HasObjPhone returns a boolean if a field has been set.

### GetSEzmaxcustomeruserFirstname

`func (o *CustomEzmaxcustomeruserResponse) GetSEzmaxcustomeruserFirstname() string`

GetSEzmaxcustomeruserFirstname returns the SEzmaxcustomeruserFirstname field if non-nil, zero value otherwise.

### GetSEzmaxcustomeruserFirstnameOk

`func (o *CustomEzmaxcustomeruserResponse) GetSEzmaxcustomeruserFirstnameOk() (*string, bool)`

GetSEzmaxcustomeruserFirstnameOk returns a tuple with the SEzmaxcustomeruserFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxcustomeruserFirstname

`func (o *CustomEzmaxcustomeruserResponse) SetSEzmaxcustomeruserFirstname(v string)`

SetSEzmaxcustomeruserFirstname sets SEzmaxcustomeruserFirstname field to given value.


### GetSEzmaxcustomeruserLastname

`func (o *CustomEzmaxcustomeruserResponse) GetSEzmaxcustomeruserLastname() string`

GetSEzmaxcustomeruserLastname returns the SEzmaxcustomeruserLastname field if non-nil, zero value otherwise.

### GetSEzmaxcustomeruserLastnameOk

`func (o *CustomEzmaxcustomeruserResponse) GetSEzmaxcustomeruserLastnameOk() (*string, bool)`

GetSEzmaxcustomeruserLastnameOk returns a tuple with the SEzmaxcustomeruserLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxcustomeruserLastname

`func (o *CustomEzmaxcustomeruserResponse) SetSEzmaxcustomeruserLastname(v string)`

SetSEzmaxcustomeruserLastname sets SEzmaxcustomeruserLastname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


