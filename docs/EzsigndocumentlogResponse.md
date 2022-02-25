# EzsigndocumentlogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FkiUserID** | **NullableInt32** | The unique ID of the User | 
**FkiEzsignsignerID** | **NullableInt32** | The unique ID of the Ezsignsigner | 
**DtEzsigndocumentlogDatetime** | **string** | The date and time at which the event was logged | 
**EEzsigndocumentlogType** | [**FieldEEzsigndocumentlogType**](FieldEEzsigndocumentlogType.md) |  | 
**SEzsigndocumentlogDetail** | **string** | The detail of the Ezsigndocumentlog | 
**SEzsigndocumentlogLastname** | **string** | The last name of the User or Ezsignsigner | 
**SEzsigndocumentlogFirstname** | **string** | The first name of the User or Ezsignsigner | 
**SEzsigndocumentlogIP** | **string** | Represent an IP address. | 

## Methods

### NewEzsigndocumentlogResponse

`func NewEzsigndocumentlogResponse(fkiUserID NullableInt32, fkiEzsignsignerID NullableInt32, dtEzsigndocumentlogDatetime string, eEzsigndocumentlogType FieldEEzsigndocumentlogType, sEzsigndocumentlogDetail string, sEzsigndocumentlogLastname string, sEzsigndocumentlogFirstname string, sEzsigndocumentlogIP string, ) *EzsigndocumentlogResponse`

NewEzsigndocumentlogResponse instantiates a new EzsigndocumentlogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsigndocumentlogResponseWithDefaults

`func NewEzsigndocumentlogResponseWithDefaults() *EzsigndocumentlogResponse`

NewEzsigndocumentlogResponseWithDefaults instantiates a new EzsigndocumentlogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFkiUserID

`func (o *EzsigndocumentlogResponse) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzsigndocumentlogResponse) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzsigndocumentlogResponse) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### SetFkiUserIDNil

`func (o *EzsigndocumentlogResponse) SetFkiUserIDNil(b bool)`

 SetFkiUserIDNil sets the value for FkiUserID to be an explicit nil

### UnsetFkiUserID
`func (o *EzsigndocumentlogResponse) UnsetFkiUserID()`

UnsetFkiUserID ensures that no value is present for FkiUserID, not even an explicit nil
### GetFkiEzsignsignerID

`func (o *EzsigndocumentlogResponse) GetFkiEzsignsignerID() int32`

GetFkiEzsignsignerID returns the FkiEzsignsignerID field if non-nil, zero value otherwise.

### GetFkiEzsignsignerIDOk

`func (o *EzsigndocumentlogResponse) GetFkiEzsignsignerIDOk() (*int32, bool)`

GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignerID

`func (o *EzsigndocumentlogResponse) SetFkiEzsignsignerID(v int32)`

SetFkiEzsignsignerID sets FkiEzsignsignerID field to given value.


### SetFkiEzsignsignerIDNil

`func (o *EzsigndocumentlogResponse) SetFkiEzsignsignerIDNil(b bool)`

 SetFkiEzsignsignerIDNil sets the value for FkiEzsignsignerID to be an explicit nil

### UnsetFkiEzsignsignerID
`func (o *EzsigndocumentlogResponse) UnsetFkiEzsignsignerID()`

UnsetFkiEzsignsignerID ensures that no value is present for FkiEzsignsignerID, not even an explicit nil
### GetDtEzsigndocumentlogDatetime

`func (o *EzsigndocumentlogResponse) GetDtEzsigndocumentlogDatetime() string`

GetDtEzsigndocumentlogDatetime returns the DtEzsigndocumentlogDatetime field if non-nil, zero value otherwise.

### GetDtEzsigndocumentlogDatetimeOk

`func (o *EzsigndocumentlogResponse) GetDtEzsigndocumentlogDatetimeOk() (*string, bool)`

GetDtEzsigndocumentlogDatetimeOk returns a tuple with the DtEzsigndocumentlogDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzsigndocumentlogDatetime

`func (o *EzsigndocumentlogResponse) SetDtEzsigndocumentlogDatetime(v string)`

SetDtEzsigndocumentlogDatetime sets DtEzsigndocumentlogDatetime field to given value.


### GetEEzsigndocumentlogType

`func (o *EzsigndocumentlogResponse) GetEEzsigndocumentlogType() FieldEEzsigndocumentlogType`

GetEEzsigndocumentlogType returns the EEzsigndocumentlogType field if non-nil, zero value otherwise.

### GetEEzsigndocumentlogTypeOk

`func (o *EzsigndocumentlogResponse) GetEEzsigndocumentlogTypeOk() (*FieldEEzsigndocumentlogType, bool)`

GetEEzsigndocumentlogTypeOk returns a tuple with the EEzsigndocumentlogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzsigndocumentlogType

`func (o *EzsigndocumentlogResponse) SetEEzsigndocumentlogType(v FieldEEzsigndocumentlogType)`

SetEEzsigndocumentlogType sets EEzsigndocumentlogType field to given value.


### GetSEzsigndocumentlogDetail

`func (o *EzsigndocumentlogResponse) GetSEzsigndocumentlogDetail() string`

GetSEzsigndocumentlogDetail returns the SEzsigndocumentlogDetail field if non-nil, zero value otherwise.

### GetSEzsigndocumentlogDetailOk

`func (o *EzsigndocumentlogResponse) GetSEzsigndocumentlogDetailOk() (*string, bool)`

GetSEzsigndocumentlogDetailOk returns a tuple with the SEzsigndocumentlogDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentlogDetail

`func (o *EzsigndocumentlogResponse) SetSEzsigndocumentlogDetail(v string)`

SetSEzsigndocumentlogDetail sets SEzsigndocumentlogDetail field to given value.


### GetSEzsigndocumentlogLastname

`func (o *EzsigndocumentlogResponse) GetSEzsigndocumentlogLastname() string`

GetSEzsigndocumentlogLastname returns the SEzsigndocumentlogLastname field if non-nil, zero value otherwise.

### GetSEzsigndocumentlogLastnameOk

`func (o *EzsigndocumentlogResponse) GetSEzsigndocumentlogLastnameOk() (*string, bool)`

GetSEzsigndocumentlogLastnameOk returns a tuple with the SEzsigndocumentlogLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentlogLastname

`func (o *EzsigndocumentlogResponse) SetSEzsigndocumentlogLastname(v string)`

SetSEzsigndocumentlogLastname sets SEzsigndocumentlogLastname field to given value.


### GetSEzsigndocumentlogFirstname

`func (o *EzsigndocumentlogResponse) GetSEzsigndocumentlogFirstname() string`

GetSEzsigndocumentlogFirstname returns the SEzsigndocumentlogFirstname field if non-nil, zero value otherwise.

### GetSEzsigndocumentlogFirstnameOk

`func (o *EzsigndocumentlogResponse) GetSEzsigndocumentlogFirstnameOk() (*string, bool)`

GetSEzsigndocumentlogFirstnameOk returns a tuple with the SEzsigndocumentlogFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentlogFirstname

`func (o *EzsigndocumentlogResponse) SetSEzsigndocumentlogFirstname(v string)`

SetSEzsigndocumentlogFirstname sets SEzsigndocumentlogFirstname field to given value.


### GetSEzsigndocumentlogIP

`func (o *EzsigndocumentlogResponse) GetSEzsigndocumentlogIP() string`

GetSEzsigndocumentlogIP returns the SEzsigndocumentlogIP field if non-nil, zero value otherwise.

### GetSEzsigndocumentlogIPOk

`func (o *EzsigndocumentlogResponse) GetSEzsigndocumentlogIPOk() (*string, bool)`

GetSEzsigndocumentlogIPOk returns a tuple with the SEzsigndocumentlogIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsigndocumentlogIP

`func (o *EzsigndocumentlogResponse) SetSEzsigndocumentlogIP(v string)`

SetSEzsigndocumentlogIP sets SEzsigndocumentlogIP field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


