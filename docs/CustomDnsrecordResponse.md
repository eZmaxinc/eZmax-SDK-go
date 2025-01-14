# CustomDnsrecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EDnsrecordType** | **string** | The type of the Dnsrecord | 
**EDnsrecordValidation** | **string** | The validation of the Dnsrecord | 
**SDnsrecordName** | **string** | The name of the Dnsrecord | 
**SDnsrecordValue** | Pointer to **string** | The value of the Dnsrecord | [optional] 
**SDnsrecordExpectedvalue** | Pointer to **string** | The expected value of the Dnsrecord | [optional] 
**BDnsrecordMustMatch** | **bool** | Whether the Dnsrecord must match or not | 

## Methods

### NewCustomDnsrecordResponse

`func NewCustomDnsrecordResponse(eDnsrecordType string, eDnsrecordValidation string, sDnsrecordName string, bDnsrecordMustMatch bool, ) *CustomDnsrecordResponse`

NewCustomDnsrecordResponse instantiates a new CustomDnsrecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDnsrecordResponseWithDefaults

`func NewCustomDnsrecordResponseWithDefaults() *CustomDnsrecordResponse`

NewCustomDnsrecordResponseWithDefaults instantiates a new CustomDnsrecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEDnsrecordType

`func (o *CustomDnsrecordResponse) GetEDnsrecordType() string`

GetEDnsrecordType returns the EDnsrecordType field if non-nil, zero value otherwise.

### GetEDnsrecordTypeOk

`func (o *CustomDnsrecordResponse) GetEDnsrecordTypeOk() (*string, bool)`

GetEDnsrecordTypeOk returns a tuple with the EDnsrecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDnsrecordType

`func (o *CustomDnsrecordResponse) SetEDnsrecordType(v string)`

SetEDnsrecordType sets EDnsrecordType field to given value.


### GetEDnsrecordValidation

`func (o *CustomDnsrecordResponse) GetEDnsrecordValidation() string`

GetEDnsrecordValidation returns the EDnsrecordValidation field if non-nil, zero value otherwise.

### GetEDnsrecordValidationOk

`func (o *CustomDnsrecordResponse) GetEDnsrecordValidationOk() (*string, bool)`

GetEDnsrecordValidationOk returns a tuple with the EDnsrecordValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDnsrecordValidation

`func (o *CustomDnsrecordResponse) SetEDnsrecordValidation(v string)`

SetEDnsrecordValidation sets EDnsrecordValidation field to given value.


### GetSDnsrecordName

`func (o *CustomDnsrecordResponse) GetSDnsrecordName() string`

GetSDnsrecordName returns the SDnsrecordName field if non-nil, zero value otherwise.

### GetSDnsrecordNameOk

`func (o *CustomDnsrecordResponse) GetSDnsrecordNameOk() (*string, bool)`

GetSDnsrecordNameOk returns a tuple with the SDnsrecordName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDnsrecordName

`func (o *CustomDnsrecordResponse) SetSDnsrecordName(v string)`

SetSDnsrecordName sets SDnsrecordName field to given value.


### GetSDnsrecordValue

`func (o *CustomDnsrecordResponse) GetSDnsrecordValue() string`

GetSDnsrecordValue returns the SDnsrecordValue field if non-nil, zero value otherwise.

### GetSDnsrecordValueOk

`func (o *CustomDnsrecordResponse) GetSDnsrecordValueOk() (*string, bool)`

GetSDnsrecordValueOk returns a tuple with the SDnsrecordValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDnsrecordValue

`func (o *CustomDnsrecordResponse) SetSDnsrecordValue(v string)`

SetSDnsrecordValue sets SDnsrecordValue field to given value.

### HasSDnsrecordValue

`func (o *CustomDnsrecordResponse) HasSDnsrecordValue() bool`

HasSDnsrecordValue returns a boolean if a field has been set.

### GetSDnsrecordExpectedvalue

`func (o *CustomDnsrecordResponse) GetSDnsrecordExpectedvalue() string`

GetSDnsrecordExpectedvalue returns the SDnsrecordExpectedvalue field if non-nil, zero value otherwise.

### GetSDnsrecordExpectedvalueOk

`func (o *CustomDnsrecordResponse) GetSDnsrecordExpectedvalueOk() (*string, bool)`

GetSDnsrecordExpectedvalueOk returns a tuple with the SDnsrecordExpectedvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDnsrecordExpectedvalue

`func (o *CustomDnsrecordResponse) SetSDnsrecordExpectedvalue(v string)`

SetSDnsrecordExpectedvalue sets SDnsrecordExpectedvalue field to given value.

### HasSDnsrecordExpectedvalue

`func (o *CustomDnsrecordResponse) HasSDnsrecordExpectedvalue() bool`

HasSDnsrecordExpectedvalue returns a boolean if a field has been set.

### GetBDnsrecordMustMatch

`func (o *CustomDnsrecordResponse) GetBDnsrecordMustMatch() bool`

GetBDnsrecordMustMatch returns the BDnsrecordMustMatch field if non-nil, zero value otherwise.

### GetBDnsrecordMustMatchOk

`func (o *CustomDnsrecordResponse) GetBDnsrecordMustMatchOk() (*bool, bool)`

GetBDnsrecordMustMatchOk returns a tuple with the BDnsrecordMustMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBDnsrecordMustMatch

`func (o *CustomDnsrecordResponse) SetBDnsrecordMustMatch(v bool)`

SetBDnsrecordMustMatch sets BDnsrecordMustMatch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


