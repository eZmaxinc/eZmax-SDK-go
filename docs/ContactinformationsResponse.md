# ContactinformationsResponse

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

## Methods

### NewContactinformationsResponse

`func NewContactinformationsResponse(pkiContactinformationsID int32, eContactinformationsType FieldEContactinformationsType, ) *ContactinformationsResponse`

NewContactinformationsResponse instantiates a new ContactinformationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactinformationsResponseWithDefaults

`func NewContactinformationsResponseWithDefaults() *ContactinformationsResponse`

NewContactinformationsResponseWithDefaults instantiates a new ContactinformationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiContactinformationsID

`func (o *ContactinformationsResponse) GetPkiContactinformationsID() int32`

GetPkiContactinformationsID returns the PkiContactinformationsID field if non-nil, zero value otherwise.

### GetPkiContactinformationsIDOk

`func (o *ContactinformationsResponse) GetPkiContactinformationsIDOk() (*int32, bool)`

GetPkiContactinformationsIDOk returns a tuple with the PkiContactinformationsID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiContactinformationsID

`func (o *ContactinformationsResponse) SetPkiContactinformationsID(v int32)`

SetPkiContactinformationsID sets PkiContactinformationsID field to given value.


### GetFkiAddressIDDefault

`func (o *ContactinformationsResponse) GetFkiAddressIDDefault() int32`

GetFkiAddressIDDefault returns the FkiAddressIDDefault field if non-nil, zero value otherwise.

### GetFkiAddressIDDefaultOk

`func (o *ContactinformationsResponse) GetFkiAddressIDDefaultOk() (*int32, bool)`

GetFkiAddressIDDefaultOk returns a tuple with the FkiAddressIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAddressIDDefault

`func (o *ContactinformationsResponse) SetFkiAddressIDDefault(v int32)`

SetFkiAddressIDDefault sets FkiAddressIDDefault field to given value.

### HasFkiAddressIDDefault

`func (o *ContactinformationsResponse) HasFkiAddressIDDefault() bool`

HasFkiAddressIDDefault returns a boolean if a field has been set.

### GetFkiPhoneIDDefault

`func (o *ContactinformationsResponse) GetFkiPhoneIDDefault() int32`

GetFkiPhoneIDDefault returns the FkiPhoneIDDefault field if non-nil, zero value otherwise.

### GetFkiPhoneIDDefaultOk

`func (o *ContactinformationsResponse) GetFkiPhoneIDDefaultOk() (*int32, bool)`

GetFkiPhoneIDDefaultOk returns a tuple with the FkiPhoneIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPhoneIDDefault

`func (o *ContactinformationsResponse) SetFkiPhoneIDDefault(v int32)`

SetFkiPhoneIDDefault sets FkiPhoneIDDefault field to given value.

### HasFkiPhoneIDDefault

`func (o *ContactinformationsResponse) HasFkiPhoneIDDefault() bool`

HasFkiPhoneIDDefault returns a boolean if a field has been set.

### GetFkiEmailIDDefault

`func (o *ContactinformationsResponse) GetFkiEmailIDDefault() int32`

GetFkiEmailIDDefault returns the FkiEmailIDDefault field if non-nil, zero value otherwise.

### GetFkiEmailIDDefaultOk

`func (o *ContactinformationsResponse) GetFkiEmailIDDefaultOk() (*int32, bool)`

GetFkiEmailIDDefaultOk returns a tuple with the FkiEmailIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmailIDDefault

`func (o *ContactinformationsResponse) SetFkiEmailIDDefault(v int32)`

SetFkiEmailIDDefault sets FkiEmailIDDefault field to given value.

### HasFkiEmailIDDefault

`func (o *ContactinformationsResponse) HasFkiEmailIDDefault() bool`

HasFkiEmailIDDefault returns a boolean if a field has been set.

### GetFkiWebsiteIDDefault

`func (o *ContactinformationsResponse) GetFkiWebsiteIDDefault() int32`

GetFkiWebsiteIDDefault returns the FkiWebsiteIDDefault field if non-nil, zero value otherwise.

### GetFkiWebsiteIDDefaultOk

`func (o *ContactinformationsResponse) GetFkiWebsiteIDDefaultOk() (*int32, bool)`

GetFkiWebsiteIDDefaultOk returns a tuple with the FkiWebsiteIDDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiWebsiteIDDefault

`func (o *ContactinformationsResponse) SetFkiWebsiteIDDefault(v int32)`

SetFkiWebsiteIDDefault sets FkiWebsiteIDDefault field to given value.

### HasFkiWebsiteIDDefault

`func (o *ContactinformationsResponse) HasFkiWebsiteIDDefault() bool`

HasFkiWebsiteIDDefault returns a boolean if a field has been set.

### GetEContactinformationsType

`func (o *ContactinformationsResponse) GetEContactinformationsType() FieldEContactinformationsType`

GetEContactinformationsType returns the EContactinformationsType field if non-nil, zero value otherwise.

### GetEContactinformationsTypeOk

`func (o *ContactinformationsResponse) GetEContactinformationsTypeOk() (*FieldEContactinformationsType, bool)`

GetEContactinformationsTypeOk returns a tuple with the EContactinformationsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEContactinformationsType

`func (o *ContactinformationsResponse) SetEContactinformationsType(v FieldEContactinformationsType)`

SetEContactinformationsType sets EContactinformationsType field to given value.


### GetSContactinformationsUrl

`func (o *ContactinformationsResponse) GetSContactinformationsUrl() string`

GetSContactinformationsUrl returns the SContactinformationsUrl field if non-nil, zero value otherwise.

### GetSContactinformationsUrlOk

`func (o *ContactinformationsResponse) GetSContactinformationsUrlOk() (*string, bool)`

GetSContactinformationsUrlOk returns a tuple with the SContactinformationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSContactinformationsUrl

`func (o *ContactinformationsResponse) SetSContactinformationsUrl(v string)`

SetSContactinformationsUrl sets SContactinformationsUrl field to given value.

### HasSContactinformationsUrl

`func (o *ContactinformationsResponse) HasSContactinformationsUrl() bool`

HasSContactinformationsUrl returns a boolean if a field has been set.

### GetObjAddressDefault

`func (o *ContactinformationsResponse) GetObjAddressDefault() AddressResponse`

GetObjAddressDefault returns the ObjAddressDefault field if non-nil, zero value otherwise.

### GetObjAddressDefaultOk

`func (o *ContactinformationsResponse) GetObjAddressDefaultOk() (*AddressResponse, bool)`

GetObjAddressDefaultOk returns a tuple with the ObjAddressDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAddressDefault

`func (o *ContactinformationsResponse) SetObjAddressDefault(v AddressResponse)`

SetObjAddressDefault sets ObjAddressDefault field to given value.

### HasObjAddressDefault

`func (o *ContactinformationsResponse) HasObjAddressDefault() bool`

HasObjAddressDefault returns a boolean if a field has been set.

### GetObjPhoneDefault

`func (o *ContactinformationsResponse) GetObjPhoneDefault() PhoneResponseCompound`

GetObjPhoneDefault returns the ObjPhoneDefault field if non-nil, zero value otherwise.

### GetObjPhoneDefaultOk

`func (o *ContactinformationsResponse) GetObjPhoneDefaultOk() (*PhoneResponseCompound, bool)`

GetObjPhoneDefaultOk returns a tuple with the ObjPhoneDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjPhoneDefault

`func (o *ContactinformationsResponse) SetObjPhoneDefault(v PhoneResponseCompound)`

SetObjPhoneDefault sets ObjPhoneDefault field to given value.

### HasObjPhoneDefault

`func (o *ContactinformationsResponse) HasObjPhoneDefault() bool`

HasObjPhoneDefault returns a boolean if a field has been set.

### GetObjEmailDefault

`func (o *ContactinformationsResponse) GetObjEmailDefault() EmailResponse`

GetObjEmailDefault returns the ObjEmailDefault field if non-nil, zero value otherwise.

### GetObjEmailDefaultOk

`func (o *ContactinformationsResponse) GetObjEmailDefaultOk() (*EmailResponse, bool)`

GetObjEmailDefaultOk returns a tuple with the ObjEmailDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjEmailDefault

`func (o *ContactinformationsResponse) SetObjEmailDefault(v EmailResponse)`

SetObjEmailDefault sets ObjEmailDefault field to given value.

### HasObjEmailDefault

`func (o *ContactinformationsResponse) HasObjEmailDefault() bool`

HasObjEmailDefault returns a boolean if a field has been set.

### GetObjWebsiteDefault

`func (o *ContactinformationsResponse) GetObjWebsiteDefault() WebsiteResponse`

GetObjWebsiteDefault returns the ObjWebsiteDefault field if non-nil, zero value otherwise.

### GetObjWebsiteDefaultOk

`func (o *ContactinformationsResponse) GetObjWebsiteDefaultOk() (*WebsiteResponse, bool)`

GetObjWebsiteDefaultOk returns a tuple with the ObjWebsiteDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjWebsiteDefault

`func (o *ContactinformationsResponse) SetObjWebsiteDefault(v WebsiteResponse)`

SetObjWebsiteDefault sets ObjWebsiteDefault field to given value.

### HasObjWebsiteDefault

`func (o *ContactinformationsResponse) HasObjWebsiteDefault() bool`

HasObjWebsiteDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


