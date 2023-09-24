# EzsigndocumentlogResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiEzsignsignerID** | Pointer to **int32** | The unique ID of the Ezsignsigner | [optional] 
**DtEzsigndocumentlogDatetime** | **string** | The date and time at which the event was logged | 
**EEzsigndocumentlogType** | [**FieldEEzsigndocumentlogType**](FieldEEzsigndocumentlogType.md) |  | 
**SEzsigndocumentlogDetail** | **string** | The detail of the Ezsigndocumentlog | 
**SEzsigndocumentlogLastname** | **string** | The last name of the User or Ezsignsigner | 
**SEzsigndocumentlogFirstname** | **string** | The first name of the User or Ezsignsigner | 
**SEzsigndocumentlogIP** | **string** | Represent an IP address. | 

## Methods

### NewEzsigndocumentlogResponseCompound

`func NewEzsigndocumentlogResponseCompound(dtEzsigndocumentlogDatetime string, eEzsigndocumentlogType FieldEEzsigndocumentlogType, sEzsigndocumentlogDetail string, sEzsigndocumentlogLastname string, sEzsigndocumentlogFirstname string, sEzsigndocumentlogIP string, ) *EzsigndocumentlogResponseCompound`

NewEzsigndocumentlogResponseCompound instantiates a new EzsigndocumentlogResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentlogResponseCompoundWithDefaults

`func NewEzsigndocumentlogResponseCompoundWithDefaults() *EzsigndocumentlogResponseCompound`

NewEzsigndocumentlogResponseCompoundWithDefaults instantiates a new EzsigndocumentlogResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiUserID

`func (o *EzsigndocumentlogResponseCompound) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsigndocumentlogResponseCompound) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsigndocumentlogResponseCompound) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzsigndocumentlogResponseCompound) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiEzsignsignerID

`func (o *EzsigndocumentlogResponseCompound) GetFkiEzsignsignerID() int32`

GetFkiEzsignsignerID returns the FkiEzsignsignerID field if non-nil, zero value otherwise.

### GetFkiEzsignsignerIDOk

`func (o *EzsigndocumentlogResponseCompound) GetFkiEzsignsignerIDOk() (*int32, bool)`

GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignerID

`func (o *EzsigndocumentlogResponseCompound) SetFkiEzsignsignerID(v int32)`

SetFkiEzsignsignerID sets FkiEzsignsignerID field to given value.

### HasFkiEzsignsignerID

`func (o *EzsigndocumentlogResponseCompound) HasFkiEzsignsignerID() bool`

HasFkiEzsignsignerID returns a boolean if a field has been set.

### GetDtEzsigndocumentlogDatetime

`func (o *EzsigndocumentlogResponseCompound) GetDtEzsigndocumentlogDatetime() string`

GetDtEzsigndocumentlogDatetime returns the DtEzsigndocumentlogDatetime field if non-nil, zero value otherwise.

### GetDtEzsigndocumentlogDatetimeOk

`func (o *EzsigndocumentlogResponseCompound) GetDtEzsigndocumentlogDatetimeOk() (*string, bool)`

GetDtEzsigndocumentlogDatetimeOk returns a tuple with the DtEzsigndocumentlogDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentlogDatetime

`func (o *EzsigndocumentlogResponseCompound) SetDtEzsigndocumentlogDatetime(v string)`

SetDtEzsigndocumentlogDatetime sets DtEzsigndocumentlogDatetime field to given value.


### GetEEzsigndocumentlogType

`func (o *EzsigndocumentlogResponseCompound) GetEEzsigndocumentlogType() FieldEEzsigndocumentlogType`

GetEEzsigndocumentlogType returns the EEzsigndocumentlogType field if non-nil, zero value otherwise.

### GetEEzsigndocumentlogTypeOk

`func (o *EzsigndocumentlogResponseCompound) GetEEzsigndocumentlogTypeOk() (*FieldEEzsigndocumentlogType, bool)`

GetEEzsigndocumentlogTypeOk returns a tuple with the EEzsigndocumentlogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentlogType

`func (o *EzsigndocumentlogResponseCompound) SetEEzsigndocumentlogType(v FieldEEzsigndocumentlogType)`

SetEEzsigndocumentlogType sets EEzsigndocumentlogType field to given value.


### GetSEzsigndocumentlogDetail

`func (o *EzsigndocumentlogResponseCompound) GetSEzsigndocumentlogDetail() string`

GetSEzsigndocumentlogDetail returns the SEzsigndocumentlogDetail field if non-nil, zero value otherwise.

### GetSEzsigndocumentlogDetailOk

`func (o *EzsigndocumentlogResponseCompound) GetSEzsigndocumentlogDetailOk() (*string, bool)`

GetSEzsigndocumentlogDetailOk returns a tuple with the SEzsigndocumentlogDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentlogDetail

`func (o *EzsigndocumentlogResponseCompound) SetSEzsigndocumentlogDetail(v string)`

SetSEzsigndocumentlogDetail sets SEzsigndocumentlogDetail field to given value.


### GetSEzsigndocumentlogLastname

`func (o *EzsigndocumentlogResponseCompound) GetSEzsigndocumentlogLastname() string`

GetSEzsigndocumentlogLastname returns the SEzsigndocumentlogLastname field if non-nil, zero value otherwise.

### GetSEzsigndocumentlogLastnameOk

`func (o *EzsigndocumentlogResponseCompound) GetSEzsigndocumentlogLastnameOk() (*string, bool)`

GetSEzsigndocumentlogLastnameOk returns a tuple with the SEzsigndocumentlogLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentlogLastname

`func (o *EzsigndocumentlogResponseCompound) SetSEzsigndocumentlogLastname(v string)`

SetSEzsigndocumentlogLastname sets SEzsigndocumentlogLastname field to given value.


### GetSEzsigndocumentlogFirstname

`func (o *EzsigndocumentlogResponseCompound) GetSEzsigndocumentlogFirstname() string`

GetSEzsigndocumentlogFirstname returns the SEzsigndocumentlogFirstname field if non-nil, zero value otherwise.

### GetSEzsigndocumentlogFirstnameOk

`func (o *EzsigndocumentlogResponseCompound) GetSEzsigndocumentlogFirstnameOk() (*string, bool)`

GetSEzsigndocumentlogFirstnameOk returns a tuple with the SEzsigndocumentlogFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentlogFirstname

`func (o *EzsigndocumentlogResponseCompound) SetSEzsigndocumentlogFirstname(v string)`

SetSEzsigndocumentlogFirstname sets SEzsigndocumentlogFirstname field to given value.


### GetSEzsigndocumentlogIP

`func (o *EzsigndocumentlogResponseCompound) GetSEzsigndocumentlogIP() string`

GetSEzsigndocumentlogIP returns the SEzsigndocumentlogIP field if non-nil, zero value otherwise.

### GetSEzsigndocumentlogIPOk

`func (o *EzsigndocumentlogResponseCompound) GetSEzsigndocumentlogIPOk() (*string, bool)`

GetSEzsigndocumentlogIPOk returns a tuple with the SEzsigndocumentlogIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentlogIP

`func (o *EzsigndocumentlogResponseCompound) SetSEzsigndocumentlogIP(v string)`

SetSEzsigndocumentlogIP sets SEzsigndocumentlogIP field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


