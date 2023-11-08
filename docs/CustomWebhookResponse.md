# CustomWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiWebhookID** | **int32** | The unique ID of the Webhook | 
**SWebhookDescription** | **string** | The description of the Webhook | 
**FkiEzsignfoldertypeID** | Pointer to **int32** | The unique ID of the Ezsignfoldertype. | [optional] 
**SEzsignfoldertypeNameX** | Pointer to **string** | The name of the Ezsignfoldertype in the language of the requester | [optional] 
**EWebhookModule** | [**FieldEWebhookModule**](FieldEWebhookModule.md) |  | 
**EWebhookEzsignevent** | Pointer to [**FieldEWebhookEzsignevent**](FieldEWebhookEzsignevent.md) |  | [optional] 
**EWebhookManagementevent** | Pointer to [**FieldEWebhookManagementevent**](FieldEWebhookManagementevent.md) |  | [optional] 
**SWebhookUrl** | **string** | The URL of the Webhook callback | 
**SWebhookEmailfailed** | **string** | The email that will receive the Webhook in case all attempts fail | 
**BWebhookIsactive** | **bool** | Whether the Webhook is active or not | 
**BWebhookSkipsslvalidation** | **bool** | Wheter the server&#39;s SSL certificate should be validated or not. Not recommended to skip for production use | 
**PksCustomerCode** | **string** | The customer code assigned to your account | 
**BWebhookTest** | **bool** | Wheter the webhook received is a manual test or a real event | 

## Methods

### NewCustomWebhookResponse

`func NewCustomWebhookResponse(pkiWebhookID int32, sWebhookDescription string, eWebhookModule FieldEWebhookModule, sWebhookUrl string, sWebhookEmailfailed string, bWebhookIsactive bool, bWebhookSkipsslvalidation bool, pksCustomerCode string, bWebhookTest bool, ) *CustomWebhookResponse`

NewCustomWebhookResponse instantiates a new CustomWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomWebhookResponseWithDefaults

`func NewCustomWebhookResponseWithDefaults() *CustomWebhookResponse`

NewCustomWebhookResponseWithDefaults instantiates a new CustomWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiWebhookID

`func (o *CustomWebhookResponse) GetPkiWebhookID() int32`

GetPkiWebhookID returns the PkiWebhookID field if non-nil, zero value otherwise.

### GetPkiWebhookIDOk

`func (o *CustomWebhookResponse) GetPkiWebhookIDOk() (*int32, bool)`

GetPkiWebhookIDOk returns a tuple with the PkiWebhookID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiWebhookID

`func (o *CustomWebhookResponse) SetPkiWebhookID(v int32)`

SetPkiWebhookID sets PkiWebhookID field to given value.


### GetSWebhookDescription

`func (o *CustomWebhookResponse) GetSWebhookDescription() string`

GetSWebhookDescription returns the SWebhookDescription field if non-nil, zero value otherwise.

### GetSWebhookDescriptionOk

`func (o *CustomWebhookResponse) GetSWebhookDescriptionOk() (*string, bool)`

GetSWebhookDescriptionOk returns a tuple with the SWebhookDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookDescription

`func (o *CustomWebhookResponse) SetSWebhookDescription(v string)`

SetSWebhookDescription sets SWebhookDescription field to given value.


### GetFkiEzsignfoldertypeID

`func (o *CustomWebhookResponse) GetFkiEzsignfoldertypeID() int32`

GetFkiEzsignfoldertypeID returns the FkiEzsignfoldertypeID field if non-nil, zero value otherwise.

### GetFkiEzsignfoldertypeIDOk

`func (o *CustomWebhookResponse) GetFkiEzsignfoldertypeIDOk() (*int32, bool)`

GetFkiEzsignfoldertypeIDOk returns a tuple with the FkiEzsignfoldertypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignfoldertypeID

`func (o *CustomWebhookResponse) SetFkiEzsignfoldertypeID(v int32)`

SetFkiEzsignfoldertypeID sets FkiEzsignfoldertypeID field to given value.

### HasFkiEzsignfoldertypeID

`func (o *CustomWebhookResponse) HasFkiEzsignfoldertypeID() bool`

HasFkiEzsignfoldertypeID returns a boolean if a field has been set.

### GetSEzsignfoldertypeNameX

`func (o *CustomWebhookResponse) GetSEzsignfoldertypeNameX() string`

GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field if non-nil, zero value otherwise.

### GetSEzsignfoldertypeNameXOk

`func (o *CustomWebhookResponse) GetSEzsignfoldertypeNameXOk() (*string, bool)`

GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSEzsignfoldertypeNameX

`func (o *CustomWebhookResponse) SetSEzsignfoldertypeNameX(v string)`

SetSEzsignfoldertypeNameX sets SEzsignfoldertypeNameX field to given value.

### HasSEzsignfoldertypeNameX

`func (o *CustomWebhookResponse) HasSEzsignfoldertypeNameX() bool`

HasSEzsignfoldertypeNameX returns a boolean if a field has been set.

### GetEWebhookModule

`func (o *CustomWebhookResponse) GetEWebhookModule() FieldEWebhookModule`

GetEWebhookModule returns the EWebhookModule field if non-nil, zero value otherwise.

### GetEWebhookModuleOk

`func (o *CustomWebhookResponse) GetEWebhookModuleOk() (*FieldEWebhookModule, bool)`

GetEWebhookModuleOk returns a tuple with the EWebhookModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookModule

`func (o *CustomWebhookResponse) SetEWebhookModule(v FieldEWebhookModule)`

SetEWebhookModule sets EWebhookModule field to given value.


### GetEWebhookEzsignevent

`func (o *CustomWebhookResponse) GetEWebhookEzsignevent() FieldEWebhookEzsignevent`

GetEWebhookEzsignevent returns the EWebhookEzsignevent field if non-nil, zero value otherwise.

### GetEWebhookEzsigneventOk

`func (o *CustomWebhookResponse) GetEWebhookEzsigneventOk() (*FieldEWebhookEzsignevent, bool)`

GetEWebhookEzsigneventOk returns a tuple with the EWebhookEzsignevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookEzsignevent

`func (o *CustomWebhookResponse) SetEWebhookEzsignevent(v FieldEWebhookEzsignevent)`

SetEWebhookEzsignevent sets EWebhookEzsignevent field to given value.

### HasEWebhookEzsignevent

`func (o *CustomWebhookResponse) HasEWebhookEzsignevent() bool`

HasEWebhookEzsignevent returns a boolean if a field has been set.

### GetEWebhookManagementevent

`func (o *CustomWebhookResponse) GetEWebhookManagementevent() FieldEWebhookManagementevent`

GetEWebhookManagementevent returns the EWebhookManagementevent field if non-nil, zero value otherwise.

### GetEWebhookManagementeventOk

`func (o *CustomWebhookResponse) GetEWebhookManagementeventOk() (*FieldEWebhookManagementevent, bool)`

GetEWebhookManagementeventOk returns a tuple with the EWebhookManagementevent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEWebhookManagementevent

`func (o *CustomWebhookResponse) SetEWebhookManagementevent(v FieldEWebhookManagementevent)`

SetEWebhookManagementevent sets EWebhookManagementevent field to given value.

### HasEWebhookManagementevent

`func (o *CustomWebhookResponse) HasEWebhookManagementevent() bool`

HasEWebhookManagementevent returns a boolean if a field has been set.

### GetSWebhookUrl

`func (o *CustomWebhookResponse) GetSWebhookUrl() string`

GetSWebhookUrl returns the SWebhookUrl field if non-nil, zero value otherwise.

### GetSWebhookUrlOk

`func (o *CustomWebhookResponse) GetSWebhookUrlOk() (*string, bool)`

GetSWebhookUrlOk returns a tuple with the SWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookUrl

`func (o *CustomWebhookResponse) SetSWebhookUrl(v string)`

SetSWebhookUrl sets SWebhookUrl field to given value.


### GetSWebhookEmailfailed

`func (o *CustomWebhookResponse) GetSWebhookEmailfailed() string`

GetSWebhookEmailfailed returns the SWebhookEmailfailed field if non-nil, zero value otherwise.

### GetSWebhookEmailfailedOk

`func (o *CustomWebhookResponse) GetSWebhookEmailfailedOk() (*string, bool)`

GetSWebhookEmailfailedOk returns a tuple with the SWebhookEmailfailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSWebhookEmailfailed

`func (o *CustomWebhookResponse) SetSWebhookEmailfailed(v string)`

SetSWebhookEmailfailed sets SWebhookEmailfailed field to given value.


### GetBWebhookIsactive

`func (o *CustomWebhookResponse) GetBWebhookIsactive() bool`

GetBWebhookIsactive returns the BWebhookIsactive field if non-nil, zero value otherwise.

### GetBWebhookIsactiveOk

`func (o *CustomWebhookResponse) GetBWebhookIsactiveOk() (*bool, bool)`

GetBWebhookIsactiveOk returns a tuple with the BWebhookIsactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWebhookIsactive

`func (o *CustomWebhookResponse) SetBWebhookIsactive(v bool)`

SetBWebhookIsactive sets BWebhookIsactive field to given value.


### GetBWebhookSkipsslvalidation

`func (o *CustomWebhookResponse) GetBWebhookSkipsslvalidation() bool`

GetBWebhookSkipsslvalidation returns the BWebhookSkipsslvalidation field if non-nil, zero value otherwise.

### GetBWebhookSkipsslvalidationOk

`func (o *CustomWebhookResponse) GetBWebhookSkipsslvalidationOk() (*bool, bool)`

GetBWebhookSkipsslvalidationOk returns a tuple with the BWebhookSkipsslvalidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWebhookSkipsslvalidation

`func (o *CustomWebhookResponse) SetBWebhookSkipsslvalidation(v bool)`

SetBWebhookSkipsslvalidation sets BWebhookSkipsslvalidation field to given value.


### GetPksCustomerCode

`func (o *CustomWebhookResponse) GetPksCustomerCode() string`

GetPksCustomerCode returns the PksCustomerCode field if non-nil, zero value otherwise.

### GetPksCustomerCodeOk

`func (o *CustomWebhookResponse) GetPksCustomerCodeOk() (*string, bool)`

GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPksCustomerCode

`func (o *CustomWebhookResponse) SetPksCustomerCode(v string)`

SetPksCustomerCode sets PksCustomerCode field to given value.


### GetBWebhookTest

`func (o *CustomWebhookResponse) GetBWebhookTest() bool`

GetBWebhookTest returns the BWebhookTest field if non-nil, zero value otherwise.

### GetBWebhookTestOk

`func (o *CustomWebhookResponse) GetBWebhookTestOk() (*bool, bool)`

GetBWebhookTestOk returns a tuple with the BWebhookTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBWebhookTest

`func (o *CustomWebhookResponse) SetBWebhookTest(v bool)`

SetBWebhookTest sets BWebhookTest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


