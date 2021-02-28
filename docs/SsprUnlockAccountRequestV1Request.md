# SsprUnlockAccountRequestV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PksCustomerCode** | **string** | The customer code assigned to your account | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**EUserTypeSSPR** | [**FieldEUserTypeSSPR**](FieldEUserTypeSSPR.md) |  | 
**SEmailAddress** | Pointer to **string** | The email address. | [optional] 
**SUserLoginname** | Pointer to **string** | The Login name of the User. | [optional] 

## Methods

### NewSsprUnlockAccountRequestV1Request

`func NewSsprUnlockAccountRequestV1Request(pksCustomerCode string, fkiLanguageID int32, eUserTypeSSPR FieldEUserTypeSSPR, ) *SsprUnlockAccountRequestV1Request`

NewSsprUnlockAccountRequestV1Request instantiates a new SsprUnlockAccountRequestV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsprUnlockAccountRequestV1RequestWithDefaults

`func NewSsprUnlockAccountRequestV1RequestWithDefaults() *SsprUnlockAccountRequestV1Request`

NewSsprUnlockAccountRequestV1RequestWithDefaults instantiates a new SsprUnlockAccountRequestV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPksCustomerCode

`func (o *SsprUnlockAccountRequestV1Request) GetPksCustomerCode() string`

GetPksCustomerCode returns the PksCustomerCode field if non-nil, zero value otherwise.

### GetPksCustomerCodeOk

`func (o *SsprUnlockAccountRequestV1Request) GetPksCustomerCodeOk() (*string, bool)`

GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksCustomerCode

`func (o *SsprUnlockAccountRequestV1Request) SetPksCustomerCode(v string)`

SetPksCustomerCode sets PksCustomerCode field to given value.


### GetFkiLanguageID

`func (o *SsprUnlockAccountRequestV1Request) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *SsprUnlockAccountRequestV1Request) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *SsprUnlockAccountRequestV1Request) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetEUserTypeSSPR

`func (o *SsprUnlockAccountRequestV1Request) GetEUserTypeSSPR() FieldEUserTypeSSPR`

GetEUserTypeSSPR returns the EUserTypeSSPR field if non-nil, zero value otherwise.

### GetEUserTypeSSPROk

`func (o *SsprUnlockAccountRequestV1Request) GetEUserTypeSSPROk() (*FieldEUserTypeSSPR, bool)`

GetEUserTypeSSPROk returns a tuple with the EUserTypeSSPR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserTypeSSPR

`func (o *SsprUnlockAccountRequestV1Request) SetEUserTypeSSPR(v FieldEUserTypeSSPR)`

SetEUserTypeSSPR sets EUserTypeSSPR field to given value.


### GetSEmailAddress

`func (o *SsprUnlockAccountRequestV1Request) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *SsprUnlockAccountRequestV1Request) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *SsprUnlockAccountRequestV1Request) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.

### HasSEmailAddress

`func (o *SsprUnlockAccountRequestV1Request) HasSEmailAddress() bool`

HasSEmailAddress returns a boolean if a field has been set.

### GetSUserLoginname

`func (o *SsprUnlockAccountRequestV1Request) GetSUserLoginname() string`

GetSUserLoginname returns the SUserLoginname field if non-nil, zero value otherwise.

### GetSUserLoginnameOk

`func (o *SsprUnlockAccountRequestV1Request) GetSUserLoginnameOk() (*string, bool)`

GetSUserLoginnameOk returns a tuple with the SUserLoginname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginname

`func (o *SsprUnlockAccountRequestV1Request) SetSUserLoginname(v string)`

SetSUserLoginname sets SUserLoginname field to given value.

### HasSUserLoginname

`func (o *SsprUnlockAccountRequestV1Request) HasSUserLoginname() bool`

HasSUserLoginname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


