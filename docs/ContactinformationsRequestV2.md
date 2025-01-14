# ContactinformationsRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EContactinformationsType** | [**FieldEContactinformationsType**](FieldEContactinformationsType.md) |  | 
**IAddressDefault** | **int32** | The index in the a_objAddress array (zero based index) representing the Address object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IPhoneDefault** | **int32** | The index in the a_objPhone array (zero based index) representing the Phone object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IEmailDefault** | **int32** | The index in the a_objEmail array (zero based index) representing the Email object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IWebsiteDefault** | **int32** | The index in the a_objWebsite array (zero based index) representing the Website object that should become the default one.  You can leave the value to 0 if the array is empty. | 

## Methods

### NewContactinformationsRequestV2

`func NewContactinformationsRequestV2(eContactinformationsType FieldEContactinformationsType, iAddressDefault int32, iPhoneDefault int32, iEmailDefault int32, iWebsiteDefault int32, ) *ContactinformationsRequestV2`

NewContactinformationsRequestV2 instantiates a new ContactinformationsRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactinformationsRequestV2WithDefaults

`func NewContactinformationsRequestV2WithDefaults() *ContactinformationsRequestV2`

NewContactinformationsRequestV2WithDefaults instantiates a new ContactinformationsRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEContactinformationsType

`func (o *ContactinformationsRequestV2) GetEContactinformationsType() FieldEContactinformationsType`

GetEContactinformationsType returns the EContactinformationsType field if non-nil, zero value otherwise.

### GetEContactinformationsTypeOk

`func (o *ContactinformationsRequestV2) GetEContactinformationsTypeOk() (*FieldEContactinformationsType, bool)`

GetEContactinformationsTypeOk returns a tuple with the EContactinformationsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEContactinformationsType

`func (o *ContactinformationsRequestV2) SetEContactinformationsType(v FieldEContactinformationsType)`

SetEContactinformationsType sets EContactinformationsType field to given value.


### GetIAddressDefault

`func (o *ContactinformationsRequestV2) GetIAddressDefault() int32`

GetIAddressDefault returns the IAddressDefault field if non-nil, zero value otherwise.

### GetIAddressDefaultOk

`func (o *ContactinformationsRequestV2) GetIAddressDefaultOk() (*int32, bool)`

GetIAddressDefaultOk returns a tuple with the IAddressDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAddressDefault

`func (o *ContactinformationsRequestV2) SetIAddressDefault(v int32)`

SetIAddressDefault sets IAddressDefault field to given value.


### GetIPhoneDefault

`func (o *ContactinformationsRequestV2) GetIPhoneDefault() int32`

GetIPhoneDefault returns the IPhoneDefault field if non-nil, zero value otherwise.

### GetIPhoneDefaultOk

`func (o *ContactinformationsRequestV2) GetIPhoneDefaultOk() (*int32, bool)`

GetIPhoneDefaultOk returns a tuple with the IPhoneDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPhoneDefault

`func (o *ContactinformationsRequestV2) SetIPhoneDefault(v int32)`

SetIPhoneDefault sets IPhoneDefault field to given value.


### GetIEmailDefault

`func (o *ContactinformationsRequestV2) GetIEmailDefault() int32`

GetIEmailDefault returns the IEmailDefault field if non-nil, zero value otherwise.

### GetIEmailDefaultOk

`func (o *ContactinformationsRequestV2) GetIEmailDefaultOk() (*int32, bool)`

GetIEmailDefaultOk returns a tuple with the IEmailDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEmailDefault

`func (o *ContactinformationsRequestV2) SetIEmailDefault(v int32)`

SetIEmailDefault sets IEmailDefault field to given value.


### GetIWebsiteDefault

`func (o *ContactinformationsRequestV2) GetIWebsiteDefault() int32`

GetIWebsiteDefault returns the IWebsiteDefault field if non-nil, zero value otherwise.

### GetIWebsiteDefaultOk

`func (o *ContactinformationsRequestV2) GetIWebsiteDefaultOk() (*int32, bool)`

GetIWebsiteDefaultOk returns a tuple with the IWebsiteDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIWebsiteDefault

`func (o *ContactinformationsRequestV2) SetIWebsiteDefault(v int32)`

SetIWebsiteDefault sets IWebsiteDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


