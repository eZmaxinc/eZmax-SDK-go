# ContactinformationsResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiContactinformationsID** | **int32** | The unique ID of the Contactinformations | 
**FkiAddressIDDefault** | Pointer to **int32** | The unique ID of the Address | [optional] 
**FkiPhoneIDDefault** | Pointer to **int32** | The unique ID of the Phone. | [optional] 
**FkiEmailIDDefault** | Pointer to **int32** | The unique ID of the Email | [optional] 
**FkiWebsiteIDDefault** | Pointer to **int32** | The unique ID of the Website Default | [optional] 
**EContactinformationsType** | [**FieldEContactinformationsType**](FieldEContactinformationsType.md) |  | 
**SContactinformationsUrl** | Pointer to **string** | The url of the Contactinformations | [optional] 
**ObjAddressDefault** | Pointer to [**AddressResponse**](AddressResponse.md) | An Address Object and children to create a complete structure | [optional] 
**ObjPhoneDefault** | Pointer to [**PhoneResponseCompound**](PhoneResponseCompound.md) |  | [optional] 
**ObjEmailDefault** | Pointer to [**EmailResponse**](EmailResponse.md) | An Email Object and children to create a complete structure | [optional] 
**ObjWebsiteDefault** | Pointer to [**WebsiteResponse**](WebsiteResponse.md) | A Website Object and children to create a complete structure | [optional] 
**AObjAddress** | [**[]AddressResponse**](AddressResponse.md) |  | 
**AObjPhone** | [**[]PhoneResponseCompound**](PhoneResponseCompound.md) |  | 
**AObjEmail** | [**[]EmailResponse**](EmailResponse.md) |  | 
**AObjWebsite** | [**[]WebsiteResponse**](WebsiteResponse.md) |  | 

## Methods

### NewContactinformationsResponseCompound

`func NewContactinformationsResponseCompound(pkiContactinformationsID int32, eContactinformationsType FieldEContactinformationsType, aObjAddress []AddressResponseCompound, aObjPhone []PhoneResponseCompound, aObjEmail []EmailResponseCompound, aObjWebsite []WebsiteResponseCompound, ) *ContactinformationsResponseCompound`

NewContactinformationsResponseCompound instantiates a new ContactinformationsResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactinformationsResponseCompoundWithDefaults

`func NewContactinformationsResponseCompoundWithDefaults() *ContactinformationsResponseCompound`

NewContactinformationsResponseCompoundWithDefaults instantiates a new ContactinformationsResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiContactinformationsID

`func (o *ContactinformationsResponseCompound) GetPkiContactinformationsID() int32`

GetPkiContactinformationsID returns the PkiContactinformationsID field if non-nil, zero value otherwise.

### GetPkiContactinformationsIDOk

`func (o *ContactinformationsResponseCompound) GetPkiContactinformationsIDOk() (*int32, bool)`

GetPkiContactinformationsIDOk returns a tuple with the PkiContactinformationsID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiContactinformationsID

`func (o *ContactinformationsResponseCompound) SetPkiContactinformationsID(v int32)`

SetPkiContactinformationsID sets PkiContactinformationsID field to given value.


### GetFkiAddressIDDefault

`func (o *ContactinformationsResponseCompound) GetFkiAddressIDDefault() int32`

GetFkiAddressIDDefault returns the FkiAddressIDDefault field if non-nil, zero value otherwise.

### GetFkiAddressIDDefaultOk

`func (o *ContactinformationsResponseCompound) GetFkiAddressIDDefaultOk() (*int32, bool)`

GetFkiAddressIDDefaultOk returns a tuple with the FkiAddressIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAddressIDDefault

`func (o *ContactinformationsResponseCompound) SetFkiAddressIDDefault(v int32)`

SetFkiAddressIDDefault sets FkiAddressIDDefault field to given value.

### HasFkiAddressIDDefault

`func (o *ContactinformationsResponseCompound) HasFkiAddressIDDefault() bool`

HasFkiAddressIDDefault returns a boolean if a field has been set.

### GetFkiPhoneIDDefault

`func (o *ContactinformationsResponseCompound) GetFkiPhoneIDDefault() int32`

GetFkiPhoneIDDefault returns the FkiPhoneIDDefault field if non-nil, zero value otherwise.

### GetFkiPhoneIDDefaultOk

`func (o *ContactinformationsResponseCompound) GetFkiPhoneIDDefaultOk() (*int32, bool)`

GetFkiPhoneIDDefaultOk returns a tuple with the FkiPhoneIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPhoneIDDefault

`func (o *ContactinformationsResponseCompound) SetFkiPhoneIDDefault(v int32)`

SetFkiPhoneIDDefault sets FkiPhoneIDDefault field to given value.

### HasFkiPhoneIDDefault

`func (o *ContactinformationsResponseCompound) HasFkiPhoneIDDefault() bool`

HasFkiPhoneIDDefault returns a boolean if a field has been set.

### GetFkiEmailIDDefault

`func (o *ContactinformationsResponseCompound) GetFkiEmailIDDefault() int32`

GetFkiEmailIDDefault returns the FkiEmailIDDefault field if non-nil, zero value otherwise.

### GetFkiEmailIDDefaultOk

`func (o *ContactinformationsResponseCompound) GetFkiEmailIDDefaultOk() (*int32, bool)`

GetFkiEmailIDDefaultOk returns a tuple with the FkiEmailIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmailIDDefault

`func (o *ContactinformationsResponseCompound) SetFkiEmailIDDefault(v int32)`

SetFkiEmailIDDefault sets FkiEmailIDDefault field to given value.

### HasFkiEmailIDDefault

`func (o *ContactinformationsResponseCompound) HasFkiEmailIDDefault() bool`

HasFkiEmailIDDefault returns a boolean if a field has been set.

### GetFkiWebsiteIDDefault

`func (o *ContactinformationsResponseCompound) GetFkiWebsiteIDDefault() int32`

GetFkiWebsiteIDDefault returns the FkiWebsiteIDDefault field if non-nil, zero value otherwise.

### GetFkiWebsiteIDDefaultOk

`func (o *ContactinformationsResponseCompound) GetFkiWebsiteIDDefaultOk() (*int32, bool)`

GetFkiWebsiteIDDefaultOk returns a tuple with the FkiWebsiteIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiWebsiteIDDefault

`func (o *ContactinformationsResponseCompound) SetFkiWebsiteIDDefault(v int32)`

SetFkiWebsiteIDDefault sets FkiWebsiteIDDefault field to given value.

### HasFkiWebsiteIDDefault

`func (o *ContactinformationsResponseCompound) HasFkiWebsiteIDDefault() bool`

HasFkiWebsiteIDDefault returns a boolean if a field has been set.

### GetEContactinformationsType

`func (o *ContactinformationsResponseCompound) GetEContactinformationsType() FieldEContactinformationsType`

GetEContactinformationsType returns the EContactinformationsType field if non-nil, zero value otherwise.

### GetEContactinformationsTypeOk

`func (o *ContactinformationsResponseCompound) GetEContactinformationsTypeOk() (*FieldEContactinformationsType, bool)`

GetEContactinformationsTypeOk returns a tuple with the EContactinformationsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEContactinformationsType

`func (o *ContactinformationsResponseCompound) SetEContactinformationsType(v FieldEContactinformationsType)`

SetEContactinformationsType sets EContactinformationsType field to given value.


### GetSContactinformationsUrl

`func (o *ContactinformationsResponseCompound) GetSContactinformationsUrl() string`

GetSContactinformationsUrl returns the SContactinformationsUrl field if non-nil, zero value otherwise.

### GetSContactinformationsUrlOk

`func (o *ContactinformationsResponseCompound) GetSContactinformationsUrlOk() (*string, bool)`

GetSContactinformationsUrlOk returns a tuple with the SContactinformationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactinformationsUrl

`func (o *ContactinformationsResponseCompound) SetSContactinformationsUrl(v string)`

SetSContactinformationsUrl sets SContactinformationsUrl field to given value.

### HasSContactinformationsUrl

`func (o *ContactinformationsResponseCompound) HasSContactinformationsUrl() bool`

HasSContactinformationsUrl returns a boolean if a field has been set.

### GetObjAddressDefault

`func (o *ContactinformationsResponseCompound) GetObjAddressDefault() AddressResponse`

GetObjAddressDefault returns the ObjAddressDefault field if non-nil, zero value otherwise.

### GetObjAddressDefaultOk

`func (o *ContactinformationsResponseCompound) GetObjAddressDefaultOk() (*AddressResponse, bool)`

GetObjAddressDefaultOk returns a tuple with the ObjAddressDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAddressDefault

`func (o *ContactinformationsResponseCompound) SetObjAddressDefault(v AddressResponse)`

SetObjAddressDefault sets ObjAddressDefault field to given value.

### HasObjAddressDefault

`func (o *ContactinformationsResponseCompound) HasObjAddressDefault() bool`

HasObjAddressDefault returns a boolean if a field has been set.

### GetObjPhoneDefault

`func (o *ContactinformationsResponseCompound) GetObjPhoneDefault() PhoneResponseCompound`

GetObjPhoneDefault returns the ObjPhoneDefault field if non-nil, zero value otherwise.

### GetObjPhoneDefaultOk

`func (o *ContactinformationsResponseCompound) GetObjPhoneDefaultOk() (*PhoneResponseCompound, bool)`

GetObjPhoneDefaultOk returns a tuple with the ObjPhoneDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhoneDefault

`func (o *ContactinformationsResponseCompound) SetObjPhoneDefault(v PhoneResponseCompound)`

SetObjPhoneDefault sets ObjPhoneDefault field to given value.

### HasObjPhoneDefault

`func (o *ContactinformationsResponseCompound) HasObjPhoneDefault() bool`

HasObjPhoneDefault returns a boolean if a field has been set.

### GetObjEmailDefault

`func (o *ContactinformationsResponseCompound) GetObjEmailDefault() EmailResponse`

GetObjEmailDefault returns the ObjEmailDefault field if non-nil, zero value otherwise.

### GetObjEmailDefaultOk

`func (o *ContactinformationsResponseCompound) GetObjEmailDefaultOk() (*EmailResponse, bool)`

GetObjEmailDefaultOk returns a tuple with the ObjEmailDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmailDefault

`func (o *ContactinformationsResponseCompound) SetObjEmailDefault(v EmailResponse)`

SetObjEmailDefault sets ObjEmailDefault field to given value.

### HasObjEmailDefault

`func (o *ContactinformationsResponseCompound) HasObjEmailDefault() bool`

HasObjEmailDefault returns a boolean if a field has been set.

### GetObjWebsiteDefault

`func (o *ContactinformationsResponseCompound) GetObjWebsiteDefault() WebsiteResponse`

GetObjWebsiteDefault returns the ObjWebsiteDefault field if non-nil, zero value otherwise.

### GetObjWebsiteDefaultOk

`func (o *ContactinformationsResponseCompound) GetObjWebsiteDefaultOk() (*WebsiteResponse, bool)`

GetObjWebsiteDefaultOk returns a tuple with the ObjWebsiteDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebsiteDefault

`func (o *ContactinformationsResponseCompound) SetObjWebsiteDefault(v WebsiteResponse)`

SetObjWebsiteDefault sets ObjWebsiteDefault field to given value.

### HasObjWebsiteDefault

`func (o *ContactinformationsResponseCompound) HasObjWebsiteDefault() bool`

HasObjWebsiteDefault returns a boolean if a field has been set.

### GetAObjAddress

`func (o *ContactinformationsResponseCompound) GetAObjAddress() []AddressResponseCompound`

GetAObjAddress returns the AObjAddress field if non-nil, zero value otherwise.

### GetAObjAddressOk

`func (o *ContactinformationsResponseCompound) GetAObjAddressOk() (*[]AddressResponseCompound, bool)`

GetAObjAddressOk returns a tuple with the AObjAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAddress

`func (o *ContactinformationsResponseCompound) SetAObjAddress(v []AddressResponseCompound)`

SetAObjAddress sets AObjAddress field to given value.


### GetAObjPhone

`func (o *ContactinformationsResponseCompound) GetAObjPhone() []PhoneResponseCompound`

GetAObjPhone returns the AObjPhone field if non-nil, zero value otherwise.

### GetAObjPhoneOk

`func (o *ContactinformationsResponseCompound) GetAObjPhoneOk() (*[]PhoneResponseCompound, bool)`

GetAObjPhoneOk returns a tuple with the AObjPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjPhone

`func (o *ContactinformationsResponseCompound) SetAObjPhone(v []PhoneResponseCompound)`

SetAObjPhone sets AObjPhone field to given value.


### GetAObjEmail

`func (o *ContactinformationsResponseCompound) GetAObjEmail() []EmailResponseCompound`

GetAObjEmail returns the AObjEmail field if non-nil, zero value otherwise.

### GetAObjEmailOk

`func (o *ContactinformationsResponseCompound) GetAObjEmailOk() (*[]EmailResponseCompound, bool)`

GetAObjEmailOk returns a tuple with the AObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEmail

`func (o *ContactinformationsResponseCompound) SetAObjEmail(v []EmailResponseCompound)`

SetAObjEmail sets AObjEmail field to given value.


### GetAObjWebsite

`func (o *ContactinformationsResponseCompound) GetAObjWebsite() []WebsiteResponseCompound`

GetAObjWebsite returns the AObjWebsite field if non-nil, zero value otherwise.

### GetAObjWebsiteOk

`func (o *ContactinformationsResponseCompound) GetAObjWebsiteOk() (*[]WebsiteResponseCompound, bool)`

GetAObjWebsiteOk returns a tuple with the AObjWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjWebsite

`func (o *ContactinformationsResponseCompound) SetAObjWebsite(v []WebsiteResponseCompound)`

SetAObjWebsite sets AObjWebsite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


