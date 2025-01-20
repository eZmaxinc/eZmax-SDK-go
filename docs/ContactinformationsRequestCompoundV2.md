# ContactinformationsRequestCompoundV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EContactinformationsType** | [**FieldEContactinformationsType**](FieldEContactinformationsType.md) |  | 
**IAddressDefault** | **int32** | The index in the a_objAddress array (zero based index) representing the Address object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IPhoneDefault** | **int32** | The index in the a_objPhone array (zero based index) representing the Phone object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IEmailDefault** | **int32** | The index in the a_objEmail array (zero based index) representing the Email object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IWebsiteDefault** | **int32** | The index in the a_objWebsite array (zero based index) representing the Website object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**AObjAddress** | [**[]AddressRequest**](AddressRequest.md) |  | 
**AObjPhone** | [**[]PhoneRequest**](PhoneRequest.md) |  | 
**AObjEmail** | [**[]EmailRequest**](EmailRequest.md) |  | 
**AObjWebsite** | [**[]WebsiteRequest**](WebsiteRequest.md) |  | 

## Methods

### NewContactinformationsRequestCompoundV2

`func NewContactinformationsRequestCompoundV2(eContactinformationsType FieldEContactinformationsType, iAddressDefault int32, iPhoneDefault int32, iEmailDefault int32, iWebsiteDefault int32, aObjAddress []AddressRequestCompound, aObjPhone []PhoneRequestCompound, aObjEmail []EmailRequestCompound, aObjWebsite []WebsiteRequestCompound, ) *ContactinformationsRequestCompoundV2`

NewContactinformationsRequestCompoundV2 instantiates a new ContactinformationsRequestCompoundV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactinformationsRequestCompoundV2WithDefaults

`func NewContactinformationsRequestCompoundV2WithDefaults() *ContactinformationsRequestCompoundV2`

NewContactinformationsRequestCompoundV2WithDefaults instantiates a new ContactinformationsRequestCompoundV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEContactinformationsType

`func (o *ContactinformationsRequestCompoundV2) GetEContactinformationsType() FieldEContactinformationsType`

GetEContactinformationsType returns the EContactinformationsType field if non-nil, zero value otherwise.

### GetEContactinformationsTypeOk

`func (o *ContactinformationsRequestCompoundV2) GetEContactinformationsTypeOk() (*FieldEContactinformationsType, bool)`

GetEContactinformationsTypeOk returns a tuple with the EContactinformationsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEContactinformationsType

`func (o *ContactinformationsRequestCompoundV2) SetEContactinformationsType(v FieldEContactinformationsType)`

SetEContactinformationsType sets EContactinformationsType field to given value.


### GetIAddressDefault

`func (o *ContactinformationsRequestCompoundV2) GetIAddressDefault() int32`

GetIAddressDefault returns the IAddressDefault field if non-nil, zero value otherwise.

### GetIAddressDefaultOk

`func (o *ContactinformationsRequestCompoundV2) GetIAddressDefaultOk() (*int32, bool)`

GetIAddressDefaultOk returns a tuple with the IAddressDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAddressDefault

`func (o *ContactinformationsRequestCompoundV2) SetIAddressDefault(v int32)`

SetIAddressDefault sets IAddressDefault field to given value.


### GetIPhoneDefault

`func (o *ContactinformationsRequestCompoundV2) GetIPhoneDefault() int32`

GetIPhoneDefault returns the IPhoneDefault field if non-nil, zero value otherwise.

### GetIPhoneDefaultOk

`func (o *ContactinformationsRequestCompoundV2) GetIPhoneDefaultOk() (*int32, bool)`

GetIPhoneDefaultOk returns a tuple with the IPhoneDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPhoneDefault

`func (o *ContactinformationsRequestCompoundV2) SetIPhoneDefault(v int32)`

SetIPhoneDefault sets IPhoneDefault field to given value.


### GetIEmailDefault

`func (o *ContactinformationsRequestCompoundV2) GetIEmailDefault() int32`

GetIEmailDefault returns the IEmailDefault field if non-nil, zero value otherwise.

### GetIEmailDefaultOk

`func (o *ContactinformationsRequestCompoundV2) GetIEmailDefaultOk() (*int32, bool)`

GetIEmailDefaultOk returns a tuple with the IEmailDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEmailDefault

`func (o *ContactinformationsRequestCompoundV2) SetIEmailDefault(v int32)`

SetIEmailDefault sets IEmailDefault field to given value.


### GetIWebsiteDefault

`func (o *ContactinformationsRequestCompoundV2) GetIWebsiteDefault() int32`

GetIWebsiteDefault returns the IWebsiteDefault field if non-nil, zero value otherwise.

### GetIWebsiteDefaultOk

`func (o *ContactinformationsRequestCompoundV2) GetIWebsiteDefaultOk() (*int32, bool)`

GetIWebsiteDefaultOk returns a tuple with the IWebsiteDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIWebsiteDefault

`func (o *ContactinformationsRequestCompoundV2) SetIWebsiteDefault(v int32)`

SetIWebsiteDefault sets IWebsiteDefault field to given value.


### GetAObjAddress

`func (o *ContactinformationsRequestCompoundV2) GetAObjAddress() []AddressRequestCompound`

GetAObjAddress returns the AObjAddress field if non-nil, zero value otherwise.

### GetAObjAddressOk

`func (o *ContactinformationsRequestCompoundV2) GetAObjAddressOk() (*[]AddressRequestCompound, bool)`

GetAObjAddressOk returns a tuple with the AObjAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAddress

`func (o *ContactinformationsRequestCompoundV2) SetAObjAddress(v []AddressRequestCompound)`

SetAObjAddress sets AObjAddress field to given value.


### GetAObjPhone

`func (o *ContactinformationsRequestCompoundV2) GetAObjPhone() []PhoneRequestCompound`

GetAObjPhone returns the AObjPhone field if non-nil, zero value otherwise.

### GetAObjPhoneOk

`func (o *ContactinformationsRequestCompoundV2) GetAObjPhoneOk() (*[]PhoneRequestCompound, bool)`

GetAObjPhoneOk returns a tuple with the AObjPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjPhone

`func (o *ContactinformationsRequestCompoundV2) SetAObjPhone(v []PhoneRequestCompound)`

SetAObjPhone sets AObjPhone field to given value.


### GetAObjEmail

`func (o *ContactinformationsRequestCompoundV2) GetAObjEmail() []EmailRequestCompound`

GetAObjEmail returns the AObjEmail field if non-nil, zero value otherwise.

### GetAObjEmailOk

`func (o *ContactinformationsRequestCompoundV2) GetAObjEmailOk() (*[]EmailRequestCompound, bool)`

GetAObjEmailOk returns a tuple with the AObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEmail

`func (o *ContactinformationsRequestCompoundV2) SetAObjEmail(v []EmailRequestCompound)`

SetAObjEmail sets AObjEmail field to given value.


### GetAObjWebsite

`func (o *ContactinformationsRequestCompoundV2) GetAObjWebsite() []WebsiteRequestCompound`

GetAObjWebsite returns the AObjWebsite field if non-nil, zero value otherwise.

### GetAObjWebsiteOk

`func (o *ContactinformationsRequestCompoundV2) GetAObjWebsiteOk() (*[]WebsiteRequestCompound, bool)`

GetAObjWebsiteOk returns a tuple with the AObjWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjWebsite

`func (o *ContactinformationsRequestCompoundV2) SetAObjWebsite(v []WebsiteRequestCompound)`

SetAObjWebsite sets AObjWebsite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


