# ContactinformationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IAddressDefault** | **int32** | The index in the a_objAddress array (zero based index) representing the Address object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IPhoneDefault** | **int32** | The index in the a_objPhone array (zero based index) representing the Phone object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IEmailDefault** | **int32** | The index in the a_objEmail array (zero based index) representing the Email object that should become the default one.  You can leave the value to 0 if the array is empty. | 
**IWebsiteDefault** | **int32** | The index in the a_objWebsite array (zero based index) representing the Website object that should become the default one.  You can leave the value to 0 if the array is empty. | 

## Methods

### NewContactinformationsRequest

`func NewContactinformationsRequest(iAddressDefault int32, iPhoneDefault int32, iEmailDefault int32, iWebsiteDefault int32, ) *ContactinformationsRequest`

NewContactinformationsRequest instantiates a new ContactinformationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactinformationsRequestWithDefaults

`func NewContactinformationsRequestWithDefaults() *ContactinformationsRequest`

NewContactinformationsRequestWithDefaults instantiates a new ContactinformationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIAddressDefault

`func (o *ContactinformationsRequest) GetIAddressDefault() int32`

GetIAddressDefault returns the IAddressDefault field if non-nil, zero value otherwise.

### GetIAddressDefaultOk

`func (o *ContactinformationsRequest) GetIAddressDefaultOk() (*int32, bool)`

GetIAddressDefaultOk returns a tuple with the IAddressDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAddressDefault

`func (o *ContactinformationsRequest) SetIAddressDefault(v int32)`

SetIAddressDefault sets IAddressDefault field to given value.


### GetIPhoneDefault

`func (o *ContactinformationsRequest) GetIPhoneDefault() int32`

GetIPhoneDefault returns the IPhoneDefault field if non-nil, zero value otherwise.

### GetIPhoneDefaultOk

`func (o *ContactinformationsRequest) GetIPhoneDefaultOk() (*int32, bool)`

GetIPhoneDefaultOk returns a tuple with the IPhoneDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIPhoneDefault

`func (o *ContactinformationsRequest) SetIPhoneDefault(v int32)`

SetIPhoneDefault sets IPhoneDefault field to given value.


### GetIEmailDefault

`func (o *ContactinformationsRequest) GetIEmailDefault() int32`

GetIEmailDefault returns the IEmailDefault field if non-nil, zero value otherwise.

### GetIEmailDefaultOk

`func (o *ContactinformationsRequest) GetIEmailDefaultOk() (*int32, bool)`

GetIEmailDefaultOk returns a tuple with the IEmailDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEmailDefault

`func (o *ContactinformationsRequest) SetIEmailDefault(v int32)`

SetIEmailDefault sets IEmailDefault field to given value.


### GetIWebsiteDefault

`func (o *ContactinformationsRequest) GetIWebsiteDefault() int32`

GetIWebsiteDefault returns the IWebsiteDefault field if non-nil, zero value otherwise.

### GetIWebsiteDefaultOk

`func (o *ContactinformationsRequest) GetIWebsiteDefaultOk() (*int32, bool)`

GetIWebsiteDefaultOk returns a tuple with the IWebsiteDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIWebsiteDefault

`func (o *ContactinformationsRequest) SetIWebsiteDefault(v int32)`

SetIWebsiteDefault sets IWebsiteDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


