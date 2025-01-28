# ContactinformationsRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IAddressDefault** | **int32** | The index in the a_objAddress array (zero based index) representing the Address object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IPhoneDefault** | **int32** | The index in the a_objPhone array (zero based index) representing the Phone object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IEmailDefault** | **int32** | The index in the a_objEmail array (zero based index) representing the Email object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IWebsiteDefault** | **int32** | The index in the a_objWebsite array (zero based index) representing the Website object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**AObjAddress** | [**[]AddressRequestCompound**](AddressRequestCompound.md) |  | 
**AObjPhone** | [**[]PhoneRequestCompound**](PhoneRequestCompound.md) |  | 
**AObjEmail** | [**[]EmailRequestCompound**](EmailRequestCompound.md) |  | 
**AObjWebsite** | [**[]WebsiteRequestCompound**](WebsiteRequestCompound.md) |  | 

## Methods

### NewContactinformationsRequestCompound

`func NewContactinformationsRequestCompound(iAddressDefault int32, iPhoneDefault int32, iEmailDefault int32, iWebsiteDefault int32, aObjAddress []AddressRequestCompound, aObjPhone []PhoneRequestCompound, aObjEmail []EmailRequestCompound, aObjWebsite []WebsiteRequestCompound, ) *ContactinformationsRequestCompound`

NewContactinformationsRequestCompound instantiates a new ContactinformationsRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactinformationsRequestCompoundWithDefaults

`func NewContactinformationsRequestCompoundWithDefaults() *ContactinformationsRequestCompound`

NewContactinformationsRequestCompoundWithDefaults instantiates a new ContactinformationsRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIAddressDefault

`func (o *ContactinformationsRequestCompound) GetIAddressDefault() int32`

GetIAddressDefault returns the IAddressDefault field if non-nil, zero value otherwise.

### GetIAddressDefaultOk

`func (o *ContactinformationsRequestCompound) GetIAddressDefaultOk() (*int32, bool)`

GetIAddressDefaultOk returns a tuple with the IAddressDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAddressDefault

`func (o *ContactinformationsRequestCompound) SetIAddressDefault(v int32)`

SetIAddressDefault sets IAddressDefault field to given value.


### GetIPhoneDefault

`func (o *ContactinformationsRequestCompound) GetIPhoneDefault() int32`

GetIPhoneDefault returns the IPhoneDefault field if non-nil, zero value otherwise.

### GetIPhoneDefaultOk

`func (o *ContactinformationsRequestCompound) GetIPhoneDefaultOk() (*int32, bool)`

GetIPhoneDefaultOk returns a tuple with the IPhoneDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPhoneDefault

`func (o *ContactinformationsRequestCompound) SetIPhoneDefault(v int32)`

SetIPhoneDefault sets IPhoneDefault field to given value.


### GetIEmailDefault

`func (o *ContactinformationsRequestCompound) GetIEmailDefault() int32`

GetIEmailDefault returns the IEmailDefault field if non-nil, zero value otherwise.

### GetIEmailDefaultOk

`func (o *ContactinformationsRequestCompound) GetIEmailDefaultOk() (*int32, bool)`

GetIEmailDefaultOk returns a tuple with the IEmailDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEmailDefault

`func (o *ContactinformationsRequestCompound) SetIEmailDefault(v int32)`

SetIEmailDefault sets IEmailDefault field to given value.


### GetIWebsiteDefault

`func (o *ContactinformationsRequestCompound) GetIWebsiteDefault() int32`

GetIWebsiteDefault returns the IWebsiteDefault field if non-nil, zero value otherwise.

### GetIWebsiteDefaultOk

`func (o *ContactinformationsRequestCompound) GetIWebsiteDefaultOk() (*int32, bool)`

GetIWebsiteDefaultOk returns a tuple with the IWebsiteDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIWebsiteDefault

`func (o *ContactinformationsRequestCompound) SetIWebsiteDefault(v int32)`

SetIWebsiteDefault sets IWebsiteDefault field to given value.


### GetAObjAddress

`func (o *ContactinformationsRequestCompound) GetAObjAddress() []AddressRequestCompound`

GetAObjAddress returns the AObjAddress field if non-nil, zero value otherwise.

### GetAObjAddressOk

`func (o *ContactinformationsRequestCompound) GetAObjAddressOk() (*[]AddressRequestCompound, bool)`

GetAObjAddressOk returns a tuple with the AObjAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAddress

`func (o *ContactinformationsRequestCompound) SetAObjAddress(v []AddressRequestCompound)`

SetAObjAddress sets AObjAddress field to given value.


### GetAObjPhone

`func (o *ContactinformationsRequestCompound) GetAObjPhone() []PhoneRequestCompound`

GetAObjPhone returns the AObjPhone field if non-nil, zero value otherwise.

### GetAObjPhoneOk

`func (o *ContactinformationsRequestCompound) GetAObjPhoneOk() (*[]PhoneRequestCompound, bool)`

GetAObjPhoneOk returns a tuple with the AObjPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjPhone

`func (o *ContactinformationsRequestCompound) SetAObjPhone(v []PhoneRequestCompound)`

SetAObjPhone sets AObjPhone field to given value.


### GetAObjEmail

`func (o *ContactinformationsRequestCompound) GetAObjEmail() []EmailRequestCompound`

GetAObjEmail returns the AObjEmail field if non-nil, zero value otherwise.

### GetAObjEmailOk

`func (o *ContactinformationsRequestCompound) GetAObjEmailOk() (*[]EmailRequestCompound, bool)`

GetAObjEmailOk returns a tuple with the AObjEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjEmail

`func (o *ContactinformationsRequestCompound) SetAObjEmail(v []EmailRequestCompound)`

SetAObjEmail sets AObjEmail field to given value.


### GetAObjWebsite

`func (o *ContactinformationsRequestCompound) GetAObjWebsite() []WebsiteRequestCompound`

GetAObjWebsite returns the AObjWebsite field if non-nil, zero value otherwise.

### GetAObjWebsiteOk

`func (o *ContactinformationsRequestCompound) GetAObjWebsiteOk() (*[]WebsiteRequestCompound, bool)`

GetAObjWebsiteOk returns a tuple with the AObjWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjWebsite

`func (o *ContactinformationsRequestCompound) SetAObjWebsite(v []WebsiteRequestCompound)`

SetAObjWebsite sets AObjWebsite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


