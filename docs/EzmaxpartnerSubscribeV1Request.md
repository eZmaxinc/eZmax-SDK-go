# EzmaxpartnerSubscribeV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PksEzmaxcustomerCode** | Pointer to **string** | The Ezmaxcustomer code | [optional] 
**SInfrastructureenvironmenttypeDescription** | Pointer to **string** | The environment type Description | [optional] 
**SCompanyName1** | Pointer to **string** | The Name of the Company in French | [optional] 
**SCompanyName2** | Pointer to **string** | The Name of the Company in English | [optional] 
**FkiSystemconfigurationtypeID** | Pointer to **int32** | The unique ID of the Systemconfigurationtype | [optional] 
**SSystemconfigurationtypeDescription1** | Pointer to **string** | The description of the Systemconfigurationtype in the language of the requester | [optional] 
**SSystemconfigurationtypeDescription2** | Pointer to **string** | The description of the Systemconfigurationtype in the language of the requester | [optional] 
**FkiEzmaxpartnerID** | Pointer to **int32** | The unique ID of the Ezmaxpartner | [optional] 
**SEzmaxpartnerName1** | Pointer to **string** | The name of the Ezmaxpartner in french | [optional] 
**SEzmaxpartnerName2** | Pointer to **string** | The name of the Ezmaxpartner in english | [optional] 
**FkiEzmaxpartnerproductID** | Pointer to **int32** | The unique ID of the Ezmaxpartnerproduct | [optional] 
**SEzmaxpartnerproductName1** | Pointer to **string** | The name1 of the Ezmaxpartnerproduct | [optional] 
**SEzmaxpartnerproductName2** | Pointer to **string** | The name2 of the Ezmaxpartnerproduct | [optional] 
**FkiEzmaxpartnerproductstageID** | Pointer to **int32** | The unique ID of the Ezmaxpartnerproductstage | [optional] 
**SEzmaxpartnerproductstageCode** | Pointer to **string** | The code of the sEzmaxpartnerproductstage | [optional] 
**SUserLoginName** | Pointer to **string** | The login name of the User. | [optional] 
**SUserFirstName** | Pointer to **string** | The first name of the user | [optional] 
**SUserLastName** | Pointer to **string** | The last name of the user | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiLanguageID** | Pointer to **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | [optional] 
**ObjAddress** | Pointer to [**AddressRequestCompound**](AddressRequestCompound.md) |  | [optional] 
**Objphone** | Pointer to [**PhoneRequestCompoundV2**](PhoneRequestCompoundV2.md) |  | [optional] 
**ObjEmail** | Pointer to [**EmailRequestCompound**](EmailRequestCompound.md) |  | [optional] 

## Methods

### NewEzmaxpartnerSubscribeV1Request

`func NewEzmaxpartnerSubscribeV1Request() *EzmaxpartnerSubscribeV1Request`

NewEzmaxpartnerSubscribeV1Request instantiates a new EzmaxpartnerSubscribeV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxpartnerSubscribeV1RequestWithDefaults

`func NewEzmaxpartnerSubscribeV1RequestWithDefaults() *EzmaxpartnerSubscribeV1Request`

NewEzmaxpartnerSubscribeV1RequestWithDefaults instantiates a new EzmaxpartnerSubscribeV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPksEzmaxcustomerCode

`func (o *EzmaxpartnerSubscribeV1Request) GetPksEzmaxcustomerCode() string`

GetPksEzmaxcustomerCode returns the PksEzmaxcustomerCode field if non-nil, zero value otherwise.

### GetPksEzmaxcustomerCodeOk

`func (o *EzmaxpartnerSubscribeV1Request) GetPksEzmaxcustomerCodeOk() (*string, bool)`

GetPksEzmaxcustomerCodeOk returns a tuple with the PksEzmaxcustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksEzmaxcustomerCode

`func (o *EzmaxpartnerSubscribeV1Request) SetPksEzmaxcustomerCode(v string)`

SetPksEzmaxcustomerCode sets PksEzmaxcustomerCode field to given value.

### HasPksEzmaxcustomerCode

`func (o *EzmaxpartnerSubscribeV1Request) HasPksEzmaxcustomerCode() bool`

HasPksEzmaxcustomerCode returns a boolean if a field has been set.

### GetSInfrastructureenvironmenttypeDescription

`func (o *EzmaxpartnerSubscribeV1Request) GetSInfrastructureenvironmenttypeDescription() string`

GetSInfrastructureenvironmenttypeDescription returns the SInfrastructureenvironmenttypeDescription field if non-nil, zero value otherwise.

### GetSInfrastructureenvironmenttypeDescriptionOk

`func (o *EzmaxpartnerSubscribeV1Request) GetSInfrastructureenvironmenttypeDescriptionOk() (*string, bool)`

GetSInfrastructureenvironmenttypeDescriptionOk returns a tuple with the SInfrastructureenvironmenttypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInfrastructureenvironmenttypeDescription

`func (o *EzmaxpartnerSubscribeV1Request) SetSInfrastructureenvironmenttypeDescription(v string)`

SetSInfrastructureenvironmenttypeDescription sets SInfrastructureenvironmenttypeDescription field to given value.

### HasSInfrastructureenvironmenttypeDescription

`func (o *EzmaxpartnerSubscribeV1Request) HasSInfrastructureenvironmenttypeDescription() bool`

HasSInfrastructureenvironmenttypeDescription returns a boolean if a field has been set.

### GetSCompanyName1

`func (o *EzmaxpartnerSubscribeV1Request) GetSCompanyName1() string`

GetSCompanyName1 returns the SCompanyName1 field if non-nil, zero value otherwise.

### GetSCompanyName1Ok

`func (o *EzmaxpartnerSubscribeV1Request) GetSCompanyName1Ok() (*string, bool)`

GetSCompanyName1Ok returns a tuple with the SCompanyName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyName1

`func (o *EzmaxpartnerSubscribeV1Request) SetSCompanyName1(v string)`

SetSCompanyName1 sets SCompanyName1 field to given value.

### HasSCompanyName1

`func (o *EzmaxpartnerSubscribeV1Request) HasSCompanyName1() bool`

HasSCompanyName1 returns a boolean if a field has been set.

### GetSCompanyName2

`func (o *EzmaxpartnerSubscribeV1Request) GetSCompanyName2() string`

GetSCompanyName2 returns the SCompanyName2 field if non-nil, zero value otherwise.

### GetSCompanyName2Ok

`func (o *EzmaxpartnerSubscribeV1Request) GetSCompanyName2Ok() (*string, bool)`

GetSCompanyName2Ok returns a tuple with the SCompanyName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyName2

`func (o *EzmaxpartnerSubscribeV1Request) SetSCompanyName2(v string)`

SetSCompanyName2 sets SCompanyName2 field to given value.

### HasSCompanyName2

`func (o *EzmaxpartnerSubscribeV1Request) HasSCompanyName2() bool`

HasSCompanyName2 returns a boolean if a field has been set.

### GetFkiSystemconfigurationtypeID

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiSystemconfigurationtypeID() int32`

GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field if non-nil, zero value otherwise.

### GetFkiSystemconfigurationtypeIDOk

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiSystemconfigurationtypeIDOk() (*int32, bool)`

GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSystemconfigurationtypeID

`func (o *EzmaxpartnerSubscribeV1Request) SetFkiSystemconfigurationtypeID(v int32)`

SetFkiSystemconfigurationtypeID sets FkiSystemconfigurationtypeID field to given value.

### HasFkiSystemconfigurationtypeID

`func (o *EzmaxpartnerSubscribeV1Request) HasFkiSystemconfigurationtypeID() bool`

HasFkiSystemconfigurationtypeID returns a boolean if a field has been set.

### GetSSystemconfigurationtypeDescription1

`func (o *EzmaxpartnerSubscribeV1Request) GetSSystemconfigurationtypeDescription1() string`

GetSSystemconfigurationtypeDescription1 returns the SSystemconfigurationtypeDescription1 field if non-nil, zero value otherwise.

### GetSSystemconfigurationtypeDescription1Ok

`func (o *EzmaxpartnerSubscribeV1Request) GetSSystemconfigurationtypeDescription1Ok() (*string, bool)`

GetSSystemconfigurationtypeDescription1Ok returns a tuple with the SSystemconfigurationtypeDescription1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSystemconfigurationtypeDescription1

`func (o *EzmaxpartnerSubscribeV1Request) SetSSystemconfigurationtypeDescription1(v string)`

SetSSystemconfigurationtypeDescription1 sets SSystemconfigurationtypeDescription1 field to given value.

### HasSSystemconfigurationtypeDescription1

`func (o *EzmaxpartnerSubscribeV1Request) HasSSystemconfigurationtypeDescription1() bool`

HasSSystemconfigurationtypeDescription1 returns a boolean if a field has been set.

### GetSSystemconfigurationtypeDescription2

`func (o *EzmaxpartnerSubscribeV1Request) GetSSystemconfigurationtypeDescription2() string`

GetSSystemconfigurationtypeDescription2 returns the SSystemconfigurationtypeDescription2 field if non-nil, zero value otherwise.

### GetSSystemconfigurationtypeDescription2Ok

`func (o *EzmaxpartnerSubscribeV1Request) GetSSystemconfigurationtypeDescription2Ok() (*string, bool)`

GetSSystemconfigurationtypeDescription2Ok returns a tuple with the SSystemconfigurationtypeDescription2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSystemconfigurationtypeDescription2

`func (o *EzmaxpartnerSubscribeV1Request) SetSSystemconfigurationtypeDescription2(v string)`

SetSSystemconfigurationtypeDescription2 sets SSystemconfigurationtypeDescription2 field to given value.

### HasSSystemconfigurationtypeDescription2

`func (o *EzmaxpartnerSubscribeV1Request) HasSSystemconfigurationtypeDescription2() bool`

HasSSystemconfigurationtypeDescription2 returns a boolean if a field has been set.

### GetFkiEzmaxpartnerID

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiEzmaxpartnerID() int32`

GetFkiEzmaxpartnerID returns the FkiEzmaxpartnerID field if non-nil, zero value otherwise.

### GetFkiEzmaxpartnerIDOk

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiEzmaxpartnerIDOk() (*int32, bool)`

GetFkiEzmaxpartnerIDOk returns a tuple with the FkiEzmaxpartnerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpartnerID

`func (o *EzmaxpartnerSubscribeV1Request) SetFkiEzmaxpartnerID(v int32)`

SetFkiEzmaxpartnerID sets FkiEzmaxpartnerID field to given value.

### HasFkiEzmaxpartnerID

`func (o *EzmaxpartnerSubscribeV1Request) HasFkiEzmaxpartnerID() bool`

HasFkiEzmaxpartnerID returns a boolean if a field has been set.

### GetSEzmaxpartnerName1

`func (o *EzmaxpartnerSubscribeV1Request) GetSEzmaxpartnerName1() string`

GetSEzmaxpartnerName1 returns the SEzmaxpartnerName1 field if non-nil, zero value otherwise.

### GetSEzmaxpartnerName1Ok

`func (o *EzmaxpartnerSubscribeV1Request) GetSEzmaxpartnerName1Ok() (*string, bool)`

GetSEzmaxpartnerName1Ok returns a tuple with the SEzmaxpartnerName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerName1

`func (o *EzmaxpartnerSubscribeV1Request) SetSEzmaxpartnerName1(v string)`

SetSEzmaxpartnerName1 sets SEzmaxpartnerName1 field to given value.

### HasSEzmaxpartnerName1

`func (o *EzmaxpartnerSubscribeV1Request) HasSEzmaxpartnerName1() bool`

HasSEzmaxpartnerName1 returns a boolean if a field has been set.

### GetSEzmaxpartnerName2

`func (o *EzmaxpartnerSubscribeV1Request) GetSEzmaxpartnerName2() string`

GetSEzmaxpartnerName2 returns the SEzmaxpartnerName2 field if non-nil, zero value otherwise.

### GetSEzmaxpartnerName2Ok

`func (o *EzmaxpartnerSubscribeV1Request) GetSEzmaxpartnerName2Ok() (*string, bool)`

GetSEzmaxpartnerName2Ok returns a tuple with the SEzmaxpartnerName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerName2

`func (o *EzmaxpartnerSubscribeV1Request) SetSEzmaxpartnerName2(v string)`

SetSEzmaxpartnerName2 sets SEzmaxpartnerName2 field to given value.

### HasSEzmaxpartnerName2

`func (o *EzmaxpartnerSubscribeV1Request) HasSEzmaxpartnerName2() bool`

HasSEzmaxpartnerName2 returns a boolean if a field has been set.

### GetFkiEzmaxpartnerproductID

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiEzmaxpartnerproductID() int32`

GetFkiEzmaxpartnerproductID returns the FkiEzmaxpartnerproductID field if non-nil, zero value otherwise.

### GetFkiEzmaxpartnerproductIDOk

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiEzmaxpartnerproductIDOk() (*int32, bool)`

GetFkiEzmaxpartnerproductIDOk returns a tuple with the FkiEzmaxpartnerproductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpartnerproductID

`func (o *EzmaxpartnerSubscribeV1Request) SetFkiEzmaxpartnerproductID(v int32)`

SetFkiEzmaxpartnerproductID sets FkiEzmaxpartnerproductID field to given value.

### HasFkiEzmaxpartnerproductID

`func (o *EzmaxpartnerSubscribeV1Request) HasFkiEzmaxpartnerproductID() bool`

HasFkiEzmaxpartnerproductID returns a boolean if a field has been set.

### GetSEzmaxpartnerproductName1

`func (o *EzmaxpartnerSubscribeV1Request) GetSEzmaxpartnerproductName1() string`

GetSEzmaxpartnerproductName1 returns the SEzmaxpartnerproductName1 field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductName1Ok

`func (o *EzmaxpartnerSubscribeV1Request) GetSEzmaxpartnerproductName1Ok() (*string, bool)`

GetSEzmaxpartnerproductName1Ok returns a tuple with the SEzmaxpartnerproductName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductName1

`func (o *EzmaxpartnerSubscribeV1Request) SetSEzmaxpartnerproductName1(v string)`

SetSEzmaxpartnerproductName1 sets SEzmaxpartnerproductName1 field to given value.

### HasSEzmaxpartnerproductName1

`func (o *EzmaxpartnerSubscribeV1Request) HasSEzmaxpartnerproductName1() bool`

HasSEzmaxpartnerproductName1 returns a boolean if a field has been set.

### GetSEzmaxpartnerproductName2

`func (o *EzmaxpartnerSubscribeV1Request) GetSEzmaxpartnerproductName2() string`

GetSEzmaxpartnerproductName2 returns the SEzmaxpartnerproductName2 field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductName2Ok

`func (o *EzmaxpartnerSubscribeV1Request) GetSEzmaxpartnerproductName2Ok() (*string, bool)`

GetSEzmaxpartnerproductName2Ok returns a tuple with the SEzmaxpartnerproductName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductName2

`func (o *EzmaxpartnerSubscribeV1Request) SetSEzmaxpartnerproductName2(v string)`

SetSEzmaxpartnerproductName2 sets SEzmaxpartnerproductName2 field to given value.

### HasSEzmaxpartnerproductName2

`func (o *EzmaxpartnerSubscribeV1Request) HasSEzmaxpartnerproductName2() bool`

HasSEzmaxpartnerproductName2 returns a boolean if a field has been set.

### GetFkiEzmaxpartnerproductstageID

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiEzmaxpartnerproductstageID() int32`

GetFkiEzmaxpartnerproductstageID returns the FkiEzmaxpartnerproductstageID field if non-nil, zero value otherwise.

### GetFkiEzmaxpartnerproductstageIDOk

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiEzmaxpartnerproductstageIDOk() (*int32, bool)`

GetFkiEzmaxpartnerproductstageIDOk returns a tuple with the FkiEzmaxpartnerproductstageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpartnerproductstageID

`func (o *EzmaxpartnerSubscribeV1Request) SetFkiEzmaxpartnerproductstageID(v int32)`

SetFkiEzmaxpartnerproductstageID sets FkiEzmaxpartnerproductstageID field to given value.

### HasFkiEzmaxpartnerproductstageID

`func (o *EzmaxpartnerSubscribeV1Request) HasFkiEzmaxpartnerproductstageID() bool`

HasFkiEzmaxpartnerproductstageID returns a boolean if a field has been set.

### GetSEzmaxpartnerproductstageCode

`func (o *EzmaxpartnerSubscribeV1Request) GetSEzmaxpartnerproductstageCode() string`

GetSEzmaxpartnerproductstageCode returns the SEzmaxpartnerproductstageCode field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductstageCodeOk

`func (o *EzmaxpartnerSubscribeV1Request) GetSEzmaxpartnerproductstageCodeOk() (*string, bool)`

GetSEzmaxpartnerproductstageCodeOk returns a tuple with the SEzmaxpartnerproductstageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductstageCode

`func (o *EzmaxpartnerSubscribeV1Request) SetSEzmaxpartnerproductstageCode(v string)`

SetSEzmaxpartnerproductstageCode sets SEzmaxpartnerproductstageCode field to given value.

### HasSEzmaxpartnerproductstageCode

`func (o *EzmaxpartnerSubscribeV1Request) HasSEzmaxpartnerproductstageCode() bool`

HasSEzmaxpartnerproductstageCode returns a boolean if a field has been set.

### GetSUserLoginName

`func (o *EzmaxpartnerSubscribeV1Request) GetSUserLoginName() string`

GetSUserLoginName returns the SUserLoginName field if non-nil, zero value otherwise.

### GetSUserLoginNameOk

`func (o *EzmaxpartnerSubscribeV1Request) GetSUserLoginNameOk() (*string, bool)`

GetSUserLoginNameOk returns a tuple with the SUserLoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginName

`func (o *EzmaxpartnerSubscribeV1Request) SetSUserLoginName(v string)`

SetSUserLoginName sets SUserLoginName field to given value.

### HasSUserLoginName

`func (o *EzmaxpartnerSubscribeV1Request) HasSUserLoginName() bool`

HasSUserLoginName returns a boolean if a field has been set.

### GetSUserFirstName

`func (o *EzmaxpartnerSubscribeV1Request) GetSUserFirstName() string`

GetSUserFirstName returns the SUserFirstName field if non-nil, zero value otherwise.

### GetSUserFirstNameOk

`func (o *EzmaxpartnerSubscribeV1Request) GetSUserFirstNameOk() (*string, bool)`

GetSUserFirstNameOk returns a tuple with the SUserFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstName

`func (o *EzmaxpartnerSubscribeV1Request) SetSUserFirstName(v string)`

SetSUserFirstName sets SUserFirstName field to given value.

### HasSUserFirstName

`func (o *EzmaxpartnerSubscribeV1Request) HasSUserFirstName() bool`

HasSUserFirstName returns a boolean if a field has been set.

### GetSUserLastName

`func (o *EzmaxpartnerSubscribeV1Request) GetSUserLastName() string`

GetSUserLastName returns the SUserLastName field if non-nil, zero value otherwise.

### GetSUserLastNameOk

`func (o *EzmaxpartnerSubscribeV1Request) GetSUserLastNameOk() (*string, bool)`

GetSUserLastNameOk returns a tuple with the SUserLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastName

`func (o *EzmaxpartnerSubscribeV1Request) SetSUserLastName(v string)`

SetSUserLastName sets SUserLastName field to given value.

### HasSUserLastName

`func (o *EzmaxpartnerSubscribeV1Request) HasSUserLastName() bool`

HasSUserLastName returns a boolean if a field has been set.

### GetFkiUserID

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *EzmaxpartnerSubscribeV1Request) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *EzmaxpartnerSubscribeV1Request) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiLanguageID

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *EzmaxpartnerSubscribeV1Request) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *EzmaxpartnerSubscribeV1Request) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.

### HasFkiLanguageID

`func (o *EzmaxpartnerSubscribeV1Request) HasFkiLanguageID() bool`

HasFkiLanguageID returns a boolean if a field has been set.

### GetObjAddress

`func (o *EzmaxpartnerSubscribeV1Request) GetObjAddress() AddressRequestCompound`

GetObjAddress returns the ObjAddress field if non-nil, zero value otherwise.

### GetObjAddressOk

`func (o *EzmaxpartnerSubscribeV1Request) GetObjAddressOk() (*AddressRequestCompound, bool)`

GetObjAddressOk returns a tuple with the ObjAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAddress

`func (o *EzmaxpartnerSubscribeV1Request) SetObjAddress(v AddressRequestCompound)`

SetObjAddress sets ObjAddress field to given value.

### HasObjAddress

`func (o *EzmaxpartnerSubscribeV1Request) HasObjAddress() bool`

HasObjAddress returns a boolean if a field has been set.

### GetObjphone

`func (o *EzmaxpartnerSubscribeV1Request) GetObjphone() PhoneRequestCompoundV2`

GetObjphone returns the Objphone field if non-nil, zero value otherwise.

### GetObjphoneOk

`func (o *EzmaxpartnerSubscribeV1Request) GetObjphoneOk() (*PhoneRequestCompoundV2, bool)`

GetObjphoneOk returns a tuple with the Objphone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjphone

`func (o *EzmaxpartnerSubscribeV1Request) SetObjphone(v PhoneRequestCompoundV2)`

SetObjphone sets Objphone field to given value.

### HasObjphone

`func (o *EzmaxpartnerSubscribeV1Request) HasObjphone() bool`

HasObjphone returns a boolean if a field has been set.

### GetObjEmail

`func (o *EzmaxpartnerSubscribeV1Request) GetObjEmail() EmailRequestCompound`

GetObjEmail returns the ObjEmail field if non-nil, zero value otherwise.

### GetObjEmailOk

`func (o *EzmaxpartnerSubscribeV1Request) GetObjEmailOk() (*EmailRequestCompound, bool)`

GetObjEmailOk returns a tuple with the ObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmail

`func (o *EzmaxpartnerSubscribeV1Request) SetObjEmail(v EmailRequestCompound)`

SetObjEmail sets ObjEmail field to given value.

### HasObjEmail

`func (o *EzmaxpartnerSubscribeV1Request) HasObjEmail() bool`

HasObjEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


