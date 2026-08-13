# CustomEzmaxpartnerproductSubscribe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PksEzmaxcustomerCode** | **string** | The Ezmaxcustomer code | 
**SInfrastructureenvironmenttypeDescription** | **string** | The environment type Description | 
**SCompanyName1** | **string** | The Name of the Company in French | 
**SCompanyName2** | **string** | The Name of the Company in English | 
**FkiSystemconfigurationtypeID** | **int32** | The unique ID of the Systemconfigurationtype | 
**SSystemconfigurationtypeDescription1** | **string** | The description of the Systemconfigurationtype in the language of the requester | 
**SSystemconfigurationtypeDescription2** | **string** | The description of the Systemconfigurationtype in the language of the requester | 
**FkiEzmaxpartnerID** | **int32** | The unique ID of the Ezmaxpartner | 
**SEzmaxpartnerName1** | **string** | The name of the Ezmaxpartner in french | 
**SEzmaxpartnerName2** | **string** | The name of the Ezmaxpartner in english | 
**FkiEzmaxpartnerproductID** | **int32** | The unique ID of the Ezmaxpartnerproduct | 
**SEzmaxpartnerproductName1** | **string** | The name1 of the Ezmaxpartnerproduct | 
**SEzmaxpartnerproductName2** | **string** | The name2 of the Ezmaxpartnerproduct | 
**FkiEzmaxpartnerproductstageID** | **int32** | The unique ID of the Ezmaxpartnerproductstage | 
**SEzmaxpartnerproductstageCode** | **string** | The code of the sEzmaxpartnerproductstage | 
**SUserLoginName** | **string** | The login name of the User. | 
**SUserFirstName** | **string** | The first name of the user | 
**SUserLastName** | **string** | The last name of the user | 
**FkiUserID** | **int32** | The unique ID of the User | 
**FkiLanguageID** | **int32** | The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English| | 
**ObjAddress** | Pointer to [**AddressRequestCompound**](AddressRequestCompound.md) |  | [optional] 
**Objphone** | Pointer to [**PhoneRequestCompoundV2**](PhoneRequestCompoundV2.md) |  | [optional] 
**ObjEmail** | Pointer to [**EmailRequestCompound**](EmailRequestCompound.md) |  | [optional] 

## Methods

### NewCustomEzmaxpartnerproductSubscribe

`func NewCustomEzmaxpartnerproductSubscribe(pksEzmaxcustomerCode string, sInfrastructureenvironmenttypeDescription string, sCompanyName1 string, sCompanyName2 string, fkiSystemconfigurationtypeID int32, sSystemconfigurationtypeDescription1 string, sSystemconfigurationtypeDescription2 string, fkiEzmaxpartnerID int32, sEzmaxpartnerName1 string, sEzmaxpartnerName2 string, fkiEzmaxpartnerproductID int32, sEzmaxpartnerproductName1 string, sEzmaxpartnerproductName2 string, fkiEzmaxpartnerproductstageID int32, sEzmaxpartnerproductstageCode string, sUserLoginName string, sUserFirstName string, sUserLastName string, fkiUserID int32, fkiLanguageID int32, ) *CustomEzmaxpartnerproductSubscribe`

NewCustomEzmaxpartnerproductSubscribe instantiates a new CustomEzmaxpartnerproductSubscribe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEzmaxpartnerproductSubscribeWithDefaults

`func NewCustomEzmaxpartnerproductSubscribeWithDefaults() *CustomEzmaxpartnerproductSubscribe`

NewCustomEzmaxpartnerproductSubscribeWithDefaults instantiates a new CustomEzmaxpartnerproductSubscribe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPksEzmaxcustomerCode

`func (o *CustomEzmaxpartnerproductSubscribe) GetPksEzmaxcustomerCode() string`

GetPksEzmaxcustomerCode returns the PksEzmaxcustomerCode field if non-nil, zero value otherwise.

### GetPksEzmaxcustomerCodeOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetPksEzmaxcustomerCodeOk() (*string, bool)`

GetPksEzmaxcustomerCodeOk returns a tuple with the PksEzmaxcustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksEzmaxcustomerCode

`func (o *CustomEzmaxpartnerproductSubscribe) SetPksEzmaxcustomerCode(v string)`

SetPksEzmaxcustomerCode sets PksEzmaxcustomerCode field to given value.


### GetSInfrastructureenvironmenttypeDescription

`func (o *CustomEzmaxpartnerproductSubscribe) GetSInfrastructureenvironmenttypeDescription() string`

GetSInfrastructureenvironmenttypeDescription returns the SInfrastructureenvironmenttypeDescription field if non-nil, zero value otherwise.

### GetSInfrastructureenvironmenttypeDescriptionOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetSInfrastructureenvironmenttypeDescriptionOk() (*string, bool)`

GetSInfrastructureenvironmenttypeDescriptionOk returns a tuple with the SInfrastructureenvironmenttypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSInfrastructureenvironmenttypeDescription

`func (o *CustomEzmaxpartnerproductSubscribe) SetSInfrastructureenvironmenttypeDescription(v string)`

SetSInfrastructureenvironmenttypeDescription sets SInfrastructureenvironmenttypeDescription field to given value.


### GetSCompanyName1

`func (o *CustomEzmaxpartnerproductSubscribe) GetSCompanyName1() string`

GetSCompanyName1 returns the SCompanyName1 field if non-nil, zero value otherwise.

### GetSCompanyName1Ok

`func (o *CustomEzmaxpartnerproductSubscribe) GetSCompanyName1Ok() (*string, bool)`

GetSCompanyName1Ok returns a tuple with the SCompanyName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyName1

`func (o *CustomEzmaxpartnerproductSubscribe) SetSCompanyName1(v string)`

SetSCompanyName1 sets SCompanyName1 field to given value.


### GetSCompanyName2

`func (o *CustomEzmaxpartnerproductSubscribe) GetSCompanyName2() string`

GetSCompanyName2 returns the SCompanyName2 field if non-nil, zero value otherwise.

### GetSCompanyName2Ok

`func (o *CustomEzmaxpartnerproductSubscribe) GetSCompanyName2Ok() (*string, bool)`

GetSCompanyName2Ok returns a tuple with the SCompanyName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCompanyName2

`func (o *CustomEzmaxpartnerproductSubscribe) SetSCompanyName2(v string)`

SetSCompanyName2 sets SCompanyName2 field to given value.


### GetFkiSystemconfigurationtypeID

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiSystemconfigurationtypeID() int32`

GetFkiSystemconfigurationtypeID returns the FkiSystemconfigurationtypeID field if non-nil, zero value otherwise.

### GetFkiSystemconfigurationtypeIDOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiSystemconfigurationtypeIDOk() (*int32, bool)`

GetFkiSystemconfigurationtypeIDOk returns a tuple with the FkiSystemconfigurationtypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSystemconfigurationtypeID

`func (o *CustomEzmaxpartnerproductSubscribe) SetFkiSystemconfigurationtypeID(v int32)`

SetFkiSystemconfigurationtypeID sets FkiSystemconfigurationtypeID field to given value.


### GetSSystemconfigurationtypeDescription1

`func (o *CustomEzmaxpartnerproductSubscribe) GetSSystemconfigurationtypeDescription1() string`

GetSSystemconfigurationtypeDescription1 returns the SSystemconfigurationtypeDescription1 field if non-nil, zero value otherwise.

### GetSSystemconfigurationtypeDescription1Ok

`func (o *CustomEzmaxpartnerproductSubscribe) GetSSystemconfigurationtypeDescription1Ok() (*string, bool)`

GetSSystemconfigurationtypeDescription1Ok returns a tuple with the SSystemconfigurationtypeDescription1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSystemconfigurationtypeDescription1

`func (o *CustomEzmaxpartnerproductSubscribe) SetSSystemconfigurationtypeDescription1(v string)`

SetSSystemconfigurationtypeDescription1 sets SSystemconfigurationtypeDescription1 field to given value.


### GetSSystemconfigurationtypeDescription2

`func (o *CustomEzmaxpartnerproductSubscribe) GetSSystemconfigurationtypeDescription2() string`

GetSSystemconfigurationtypeDescription2 returns the SSystemconfigurationtypeDescription2 field if non-nil, zero value otherwise.

### GetSSystemconfigurationtypeDescription2Ok

`func (o *CustomEzmaxpartnerproductSubscribe) GetSSystemconfigurationtypeDescription2Ok() (*string, bool)`

GetSSystemconfigurationtypeDescription2Ok returns a tuple with the SSystemconfigurationtypeDescription2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSystemconfigurationtypeDescription2

`func (o *CustomEzmaxpartnerproductSubscribe) SetSSystemconfigurationtypeDescription2(v string)`

SetSSystemconfigurationtypeDescription2 sets SSystemconfigurationtypeDescription2 field to given value.


### GetFkiEzmaxpartnerID

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiEzmaxpartnerID() int32`

GetFkiEzmaxpartnerID returns the FkiEzmaxpartnerID field if non-nil, zero value otherwise.

### GetFkiEzmaxpartnerIDOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiEzmaxpartnerIDOk() (*int32, bool)`

GetFkiEzmaxpartnerIDOk returns a tuple with the FkiEzmaxpartnerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpartnerID

`func (o *CustomEzmaxpartnerproductSubscribe) SetFkiEzmaxpartnerID(v int32)`

SetFkiEzmaxpartnerID sets FkiEzmaxpartnerID field to given value.


### GetSEzmaxpartnerName1

`func (o *CustomEzmaxpartnerproductSubscribe) GetSEzmaxpartnerName1() string`

GetSEzmaxpartnerName1 returns the SEzmaxpartnerName1 field if non-nil, zero value otherwise.

### GetSEzmaxpartnerName1Ok

`func (o *CustomEzmaxpartnerproductSubscribe) GetSEzmaxpartnerName1Ok() (*string, bool)`

GetSEzmaxpartnerName1Ok returns a tuple with the SEzmaxpartnerName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerName1

`func (o *CustomEzmaxpartnerproductSubscribe) SetSEzmaxpartnerName1(v string)`

SetSEzmaxpartnerName1 sets SEzmaxpartnerName1 field to given value.


### GetSEzmaxpartnerName2

`func (o *CustomEzmaxpartnerproductSubscribe) GetSEzmaxpartnerName2() string`

GetSEzmaxpartnerName2 returns the SEzmaxpartnerName2 field if non-nil, zero value otherwise.

### GetSEzmaxpartnerName2Ok

`func (o *CustomEzmaxpartnerproductSubscribe) GetSEzmaxpartnerName2Ok() (*string, bool)`

GetSEzmaxpartnerName2Ok returns a tuple with the SEzmaxpartnerName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerName2

`func (o *CustomEzmaxpartnerproductSubscribe) SetSEzmaxpartnerName2(v string)`

SetSEzmaxpartnerName2 sets SEzmaxpartnerName2 field to given value.


### GetFkiEzmaxpartnerproductID

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiEzmaxpartnerproductID() int32`

GetFkiEzmaxpartnerproductID returns the FkiEzmaxpartnerproductID field if non-nil, zero value otherwise.

### GetFkiEzmaxpartnerproductIDOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiEzmaxpartnerproductIDOk() (*int32, bool)`

GetFkiEzmaxpartnerproductIDOk returns a tuple with the FkiEzmaxpartnerproductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpartnerproductID

`func (o *CustomEzmaxpartnerproductSubscribe) SetFkiEzmaxpartnerproductID(v int32)`

SetFkiEzmaxpartnerproductID sets FkiEzmaxpartnerproductID field to given value.


### GetSEzmaxpartnerproductName1

`func (o *CustomEzmaxpartnerproductSubscribe) GetSEzmaxpartnerproductName1() string`

GetSEzmaxpartnerproductName1 returns the SEzmaxpartnerproductName1 field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductName1Ok

`func (o *CustomEzmaxpartnerproductSubscribe) GetSEzmaxpartnerproductName1Ok() (*string, bool)`

GetSEzmaxpartnerproductName1Ok returns a tuple with the SEzmaxpartnerproductName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductName1

`func (o *CustomEzmaxpartnerproductSubscribe) SetSEzmaxpartnerproductName1(v string)`

SetSEzmaxpartnerproductName1 sets SEzmaxpartnerproductName1 field to given value.


### GetSEzmaxpartnerproductName2

`func (o *CustomEzmaxpartnerproductSubscribe) GetSEzmaxpartnerproductName2() string`

GetSEzmaxpartnerproductName2 returns the SEzmaxpartnerproductName2 field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductName2Ok

`func (o *CustomEzmaxpartnerproductSubscribe) GetSEzmaxpartnerproductName2Ok() (*string, bool)`

GetSEzmaxpartnerproductName2Ok returns a tuple with the SEzmaxpartnerproductName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductName2

`func (o *CustomEzmaxpartnerproductSubscribe) SetSEzmaxpartnerproductName2(v string)`

SetSEzmaxpartnerproductName2 sets SEzmaxpartnerproductName2 field to given value.


### GetFkiEzmaxpartnerproductstageID

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiEzmaxpartnerproductstageID() int32`

GetFkiEzmaxpartnerproductstageID returns the FkiEzmaxpartnerproductstageID field if non-nil, zero value otherwise.

### GetFkiEzmaxpartnerproductstageIDOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiEzmaxpartnerproductstageIDOk() (*int32, bool)`

GetFkiEzmaxpartnerproductstageIDOk returns a tuple with the FkiEzmaxpartnerproductstageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpartnerproductstageID

`func (o *CustomEzmaxpartnerproductSubscribe) SetFkiEzmaxpartnerproductstageID(v int32)`

SetFkiEzmaxpartnerproductstageID sets FkiEzmaxpartnerproductstageID field to given value.


### GetSEzmaxpartnerproductstageCode

`func (o *CustomEzmaxpartnerproductSubscribe) GetSEzmaxpartnerproductstageCode() string`

GetSEzmaxpartnerproductstageCode returns the SEzmaxpartnerproductstageCode field if non-nil, zero value otherwise.

### GetSEzmaxpartnerproductstageCodeOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetSEzmaxpartnerproductstageCodeOk() (*string, bool)`

GetSEzmaxpartnerproductstageCodeOk returns a tuple with the SEzmaxpartnerproductstageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzmaxpartnerproductstageCode

`func (o *CustomEzmaxpartnerproductSubscribe) SetSEzmaxpartnerproductstageCode(v string)`

SetSEzmaxpartnerproductstageCode sets SEzmaxpartnerproductstageCode field to given value.


### GetSUserLoginName

`func (o *CustomEzmaxpartnerproductSubscribe) GetSUserLoginName() string`

GetSUserLoginName returns the SUserLoginName field if non-nil, zero value otherwise.

### GetSUserLoginNameOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetSUserLoginNameOk() (*string, bool)`

GetSUserLoginNameOk returns a tuple with the SUserLoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLoginName

`func (o *CustomEzmaxpartnerproductSubscribe) SetSUserLoginName(v string)`

SetSUserLoginName sets SUserLoginName field to given value.


### GetSUserFirstName

`func (o *CustomEzmaxpartnerproductSubscribe) GetSUserFirstName() string`

GetSUserFirstName returns the SUserFirstName field if non-nil, zero value otherwise.

### GetSUserFirstNameOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetSUserFirstNameOk() (*string, bool)`

GetSUserFirstNameOk returns a tuple with the SUserFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserFirstName

`func (o *CustomEzmaxpartnerproductSubscribe) SetSUserFirstName(v string)`

SetSUserFirstName sets SUserFirstName field to given value.


### GetSUserLastName

`func (o *CustomEzmaxpartnerproductSubscribe) GetSUserLastName() string`

GetSUserLastName returns the SUserLastName field if non-nil, zero value otherwise.

### GetSUserLastNameOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetSUserLastNameOk() (*string, bool)`

GetSUserLastNameOk returns a tuple with the SUserLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSUserLastName

`func (o *CustomEzmaxpartnerproductSubscribe) SetSUserLastName(v string)`

SetSUserLastName sets SUserLastName field to given value.


### GetFkiUserID

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *CustomEzmaxpartnerproductSubscribe) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.


### GetFkiLanguageID

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiLanguageID() int32`

GetFkiLanguageID returns the FkiLanguageID field if non-nil, zero value otherwise.

### GetFkiLanguageIDOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetFkiLanguageIDOk() (*int32, bool)`

GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLanguageID

`func (o *CustomEzmaxpartnerproductSubscribe) SetFkiLanguageID(v int32)`

SetFkiLanguageID sets FkiLanguageID field to given value.


### GetObjAddress

`func (o *CustomEzmaxpartnerproductSubscribe) GetObjAddress() AddressRequestCompound`

GetObjAddress returns the ObjAddress field if non-nil, zero value otherwise.

### GetObjAddressOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetObjAddressOk() (*AddressRequestCompound, bool)`

GetObjAddressOk returns a tuple with the ObjAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAddress

`func (o *CustomEzmaxpartnerproductSubscribe) SetObjAddress(v AddressRequestCompound)`

SetObjAddress sets ObjAddress field to given value.

### HasObjAddress

`func (o *CustomEzmaxpartnerproductSubscribe) HasObjAddress() bool`

HasObjAddress returns a boolean if a field has been set.

### GetObjphone

`func (o *CustomEzmaxpartnerproductSubscribe) GetObjphone() PhoneRequestCompoundV2`

GetObjphone returns the Objphone field if non-nil, zero value otherwise.

### GetObjphoneOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetObjphoneOk() (*PhoneRequestCompoundV2, bool)`

GetObjphoneOk returns a tuple with the Objphone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjphone

`func (o *CustomEzmaxpartnerproductSubscribe) SetObjphone(v PhoneRequestCompoundV2)`

SetObjphone sets Objphone field to given value.

### HasObjphone

`func (o *CustomEzmaxpartnerproductSubscribe) HasObjphone() bool`

HasObjphone returns a boolean if a field has been set.

### GetObjEmail

`func (o *CustomEzmaxpartnerproductSubscribe) GetObjEmail() EmailRequestCompound`

GetObjEmail returns the ObjEmail field if non-nil, zero value otherwise.

### GetObjEmailOk

`func (o *CustomEzmaxpartnerproductSubscribe) GetObjEmailOk() (*EmailRequestCompound, bool)`

GetObjEmailOk returns a tuple with the ObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmail

`func (o *CustomEzmaxpartnerproductSubscribe) SetObjEmail(v EmailRequestCompound)`

SetObjEmail sets ObjEmail field to given value.

### HasObjEmail

`func (o *CustomEzmaxpartnerproductSubscribe) HasObjEmail() bool`

HasObjEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


