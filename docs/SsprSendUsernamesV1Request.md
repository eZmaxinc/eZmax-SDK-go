# SsprSendUsernamesV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PksCustomerCode** | **string** | The customer code assigned to your account | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**EUserTypeSSPR** | [**FieldEUserTypeSSPR**](FieldEUserTypeSSPR.md) |  | 
**SEmailAddress** | **string** | The email address. | 

## Methods

### NewSsprSendUsernamesV1Request

`func NewSsprSendUsernamesV1Request(pksCustomerCode string, fkiLanguageID int32, eUserTypeSSPR FieldEUserTypeSSPR, sEmailAddress string, ) *SsprSendUsernamesV1Request`

NewSsprSendUsernamesV1Request instantiates a new SsprSendUsernamesV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsprSendUsernamesV1RequestWithDefaults

`func NewSsprSendUsernamesV1RequestWithDefaults() *SsprSendUsernamesV1Request`

NewSsprSendUsernamesV1RequestWithDefaults instantiates a new SsprSendUsernamesV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPksCustomerCode

`func (o *SsprSendUsernamesV1Request) GetPksCustomerCode() string`

GetPksCustomerCode returns the PksCustomerCode field if non-nil, zero value otherwise.

### GetPksCustomerCodeOk

`func (o *SsprSendUsernamesV1Request) GetPksCustomerCodeOk() (*string, bool)`

GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksCustomerCode

`func (o *SsprSendUsernamesV1Request) SetPksCustomerCode(v string)`

SetPksCustomerCode sets PksCustomerCode field to given value.


### GetFkiLanguageID

`func (o *SsprSendUsernamesV1Request) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *SsprSendUsernamesV1Request) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *SsprSendUsernamesV1Request) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetEUserTypeSSPR

`func (o *SsprSendUsernamesV1Request) GetEUserTypeSSPR() FieldEUserTypeSSPR`

GetEUserTypeSSPR returns the EUserTypeSSPR field if non-nil, zero value otherwise.

### GetEUserTypeSSPROk

`func (o *SsprSendUsernamesV1Request) GetEUserTypeSSPROk() (*FieldEUserTypeSSPR, bool)`

GetEUserTypeSSPROk returns a tuple with the EUserTypeSSPR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUserTypeSSPR

`func (o *SsprSendUsernamesV1Request) SetEUserTypeSSPR(v FieldEUserTypeSSPR)`

SetEUserTypeSSPR sets EUserTypeSSPR field to given value.


### GetSEmailAddress

`func (o *SsprSendUsernamesV1Request) GetSEmailAddress() string`

GetSEmailAddress returns the SEmailAddress field if non-nil, zero value otherwise.

### GetSEmailAddressOk

`func (o *SsprSendUsernamesV1Request) GetSEmailAddressOk() (*string, bool)`

GetSEmailAddressOk returns a tuple with the SEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEmailAddress

`func (o *SsprSendUsernamesV1Request) SetSEmailAddress(v string)`

SetSEmailAddress sets SEmailAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


