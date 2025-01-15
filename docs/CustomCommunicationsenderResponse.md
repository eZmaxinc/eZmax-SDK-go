# CustomCommunicationsenderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiMailboxsharedID** | Pointer to **int32** | The unique ID of the Mailboxshared | [optional] 
**FkiPhonelinesharedID** | Pointer to **int32** | The unique ID of the Phonelineshared | [optional] 
**ECommunicationsenderObjecttype** | **string** |  | 
**ObjContactName** | [**CustomContactNameResponse**](CustomContactNameResponse.md) |  | 
**ObjEmail** | Pointer to [**EmailResponse**](EmailResponse.md) | An Email Object and children to create a complete structure | [optional] 
**ObjPhoneFax** | Pointer to [**PhoneResponseCompound**](PhoneResponseCompound.md) |  | [optional] 
**ObjPhoneSMS** | Pointer to [**PhoneResponseCompound**](PhoneResponseCompound.md) |  | [optional] 

## Methods

### NewCustomCommunicationsenderResponse

`func NewCustomCommunicationsenderResponse(eCommunicationsenderObjecttype string, objContactName CustomContactNameResponse, ) *CustomCommunicationsenderResponse`

NewCustomCommunicationsenderResponse instantiates a new CustomCommunicationsenderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomCommunicationsenderResponseWithDefaults

`func NewCustomCommunicationsenderResponseWithDefaults() *CustomCommunicationsenderResponse`

NewCustomCommunicationsenderResponseWithDefaults instantiates a new CustomCommunicationsenderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiAgentID

`func (o *CustomCommunicationsenderResponse) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *CustomCommunicationsenderResponse) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *CustomCommunicationsenderResponse) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *CustomCommunicationsenderResponse) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *CustomCommunicationsenderResponse) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *CustomCommunicationsenderResponse) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *CustomCommunicationsenderResponse) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *CustomCommunicationsenderResponse) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *CustomCommunicationsenderResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *CustomCommunicationsenderResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *CustomCommunicationsenderResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *CustomCommunicationsenderResponse) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiMailboxsharedID

`func (o *CustomCommunicationsenderResponse) GetFkiMailboxsharedID() int32`

GetFkiMailboxsharedID returns the FkiMailboxsharedID field if non-nil, zero value otherwise.

### GetFkiMailboxsharedIDOk

`func (o *CustomCommunicationsenderResponse) GetFkiMailboxsharedIDOk() (*int32, bool)`

GetFkiMailboxsharedIDOk returns a tuple with the FkiMailboxsharedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMailboxsharedID

`func (o *CustomCommunicationsenderResponse) SetFkiMailboxsharedID(v int32)`

SetFkiMailboxsharedID sets FkiMailboxsharedID field to given value.

### HasFkiMailboxsharedID

`func (o *CustomCommunicationsenderResponse) HasFkiMailboxsharedID() bool`

HasFkiMailboxsharedID returns a boolean if a field has been set.

### GetFkiPhonelinesharedID

`func (o *CustomCommunicationsenderResponse) GetFkiPhonelinesharedID() int32`

GetFkiPhonelinesharedID returns the FkiPhonelinesharedID field if non-nil, zero value otherwise.

### GetFkiPhonelinesharedIDOk

`func (o *CustomCommunicationsenderResponse) GetFkiPhonelinesharedIDOk() (*int32, bool)`

GetFkiPhonelinesharedIDOk returns a tuple with the FkiPhonelinesharedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPhonelinesharedID

`func (o *CustomCommunicationsenderResponse) SetFkiPhonelinesharedID(v int32)`

SetFkiPhonelinesharedID sets FkiPhonelinesharedID field to given value.

### HasFkiPhonelinesharedID

`func (o *CustomCommunicationsenderResponse) HasFkiPhonelinesharedID() bool`

HasFkiPhonelinesharedID returns a boolean if a field has been set.

### GetECommunicationsenderObjecttype

`func (o *CustomCommunicationsenderResponse) GetECommunicationsenderObjecttype() string`

GetECommunicationsenderObjecttype returns the ECommunicationsenderObjecttype field if non-nil, zero value otherwise.

### GetECommunicationsenderObjecttypeOk

`func (o *CustomCommunicationsenderResponse) GetECommunicationsenderObjecttypeOk() (*string, bool)`

GetECommunicationsenderObjecttypeOk returns a tuple with the ECommunicationsenderObjecttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationsenderObjecttype

`func (o *CustomCommunicationsenderResponse) SetECommunicationsenderObjecttype(v string)`

SetECommunicationsenderObjecttype sets ECommunicationsenderObjecttype field to given value.


### GetObjContactName

`func (o *CustomCommunicationsenderResponse) GetObjContactName() CustomContactNameResponse`

GetObjContactName returns the ObjContactName field if non-nil, zero value otherwise.

### GetObjContactNameOk

`func (o *CustomCommunicationsenderResponse) GetObjContactNameOk() (*CustomContactNameResponse, bool)`

GetObjContactNameOk returns a tuple with the ObjContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactName

`func (o *CustomCommunicationsenderResponse) SetObjContactName(v CustomContactNameResponse)`

SetObjContactName sets ObjContactName field to given value.


### GetObjEmail

`func (o *CustomCommunicationsenderResponse) GetObjEmail() EmailResponse`

GetObjEmail returns the ObjEmail field if non-nil, zero value otherwise.

### GetObjEmailOk

`func (o *CustomCommunicationsenderResponse) GetObjEmailOk() (*EmailResponse, bool)`

GetObjEmailOk returns a tuple with the ObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmail

`func (o *CustomCommunicationsenderResponse) SetObjEmail(v EmailResponse)`

SetObjEmail sets ObjEmail field to given value.

### HasObjEmail

`func (o *CustomCommunicationsenderResponse) HasObjEmail() bool`

HasObjEmail returns a boolean if a field has been set.

### GetObjPhoneFax

`func (o *CustomCommunicationsenderResponse) GetObjPhoneFax() PhoneResponseCompound`

GetObjPhoneFax returns the ObjPhoneFax field if non-nil, zero value otherwise.

### GetObjPhoneFaxOk

`func (o *CustomCommunicationsenderResponse) GetObjPhoneFaxOk() (*PhoneResponseCompound, bool)`

GetObjPhoneFaxOk returns a tuple with the ObjPhoneFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhoneFax

`func (o *CustomCommunicationsenderResponse) SetObjPhoneFax(v PhoneResponseCompound)`

SetObjPhoneFax sets ObjPhoneFax field to given value.

### HasObjPhoneFax

`func (o *CustomCommunicationsenderResponse) HasObjPhoneFax() bool`

HasObjPhoneFax returns a boolean if a field has been set.

### GetObjPhoneSMS

`func (o *CustomCommunicationsenderResponse) GetObjPhoneSMS() PhoneResponseCompound`

GetObjPhoneSMS returns the ObjPhoneSMS field if non-nil, zero value otherwise.

### GetObjPhoneSMSOk

`func (o *CustomCommunicationsenderResponse) GetObjPhoneSMSOk() (*PhoneResponseCompound, bool)`

GetObjPhoneSMSOk returns a tuple with the ObjPhoneSMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhoneSMS

`func (o *CustomCommunicationsenderResponse) SetObjPhoneSMS(v PhoneResponseCompound)`

SetObjPhoneSMS sets ObjPhoneSMS field to given value.

### HasObjPhoneSMS

`func (o *CustomCommunicationsenderResponse) HasObjPhoneSMS() bool`

HasObjPhoneSMS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


