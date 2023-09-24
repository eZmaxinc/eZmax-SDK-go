# CommunicationattachmentResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCommunicationattachmentID** | **int32** | The unique ID of the Communicationattachment | 
**FkiAttachmentID** | Pointer to **int32** | The unique ID of the Attachment. | [optional] 
**FkiInvoiceID** | Pointer to **int32** | The unique ID of the Invoice. | [optional] 
**FkiSalarypreparationID** | Pointer to **int32** | The unique ID of the Salarypreparation. | [optional] 
**SCommunicationattachmentName** | **string** | The name of the Communicationattachment | 
**SDownloadUrl** | Pointer to **string** | The Url to the requested document.  Url will expire after 3 hours. | [optional] 

## Methods

### NewCommunicationattachmentResponseCompound

`func NewCommunicationattachmentResponseCompound(pkiCommunicationattachmentID int32, sCommunicationattachmentName string, ) *CommunicationattachmentResponseCompound`

NewCommunicationattachmentResponseCompound instantiates a new CommunicationattachmentResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationattachmentResponseCompoundWithDefaults

`func NewCommunicationattachmentResponseCompoundWithDefaults() *CommunicationattachmentResponseCompound`

NewCommunicationattachmentResponseCompoundWithDefaults instantiates a new CommunicationattachmentResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCommunicationattachmentID

`func (o *CommunicationattachmentResponseCompound) GetPkiCommunicationattachmentID() int32`

GetPkiCommunicationattachmentID returns the PkiCommunicationattachmentID field if non-nil, zero value otherwise.

### GetPkiCommunicationattachmentIDOk

`func (o *CommunicationattachmentResponseCompound) GetPkiCommunicationattachmentIDOk() (*int32, bool)`

GetPkiCommunicationattachmentIDOk returns a tuple with the PkiCommunicationattachmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCommunicationattachmentID

`func (o *CommunicationattachmentResponseCompound) SetPkiCommunicationattachmentID(v int32)`

SetPkiCommunicationattachmentID sets PkiCommunicationattachmentID field to given value.


### GetFkiAttachmentID

`func (o *CommunicationattachmentResponseCompound) GetFkiAttachmentID() int32`

GetFkiAttachmentID returns the FkiAttachmentID field if non-nil, zero value otherwise.

### GetFkiAttachmentIDOk

`func (o *CommunicationattachmentResponseCompound) GetFkiAttachmentIDOk() (*int32, bool)`

GetFkiAttachmentIDOk returns a tuple with the FkiAttachmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAttachmentID

`func (o *CommunicationattachmentResponseCompound) SetFkiAttachmentID(v int32)`

SetFkiAttachmentID sets FkiAttachmentID field to given value.

### HasFkiAttachmentID

`func (o *CommunicationattachmentResponseCompound) HasFkiAttachmentID() bool`

HasFkiAttachmentID returns a boolean if a field has been set.

### GetFkiInvoiceID

`func (o *CommunicationattachmentResponseCompound) GetFkiInvoiceID() int32`

GetFkiInvoiceID returns the FkiInvoiceID field if non-nil, zero value otherwise.

### GetFkiInvoiceIDOk

`func (o *CommunicationattachmentResponseCompound) GetFkiInvoiceIDOk() (*int32, bool)`

GetFkiInvoiceIDOk returns a tuple with the FkiInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInvoiceID

`func (o *CommunicationattachmentResponseCompound) SetFkiInvoiceID(v int32)`

SetFkiInvoiceID sets FkiInvoiceID field to given value.

### HasFkiInvoiceID

`func (o *CommunicationattachmentResponseCompound) HasFkiInvoiceID() bool`

HasFkiInvoiceID returns a boolean if a field has been set.

### GetFkiSalarypreparationID

`func (o *CommunicationattachmentResponseCompound) GetFkiSalarypreparationID() int32`

GetFkiSalarypreparationID returns the FkiSalarypreparationID field if non-nil, zero value otherwise.

### GetFkiSalarypreparationIDOk

`func (o *CommunicationattachmentResponseCompound) GetFkiSalarypreparationIDOk() (*int32, bool)`

GetFkiSalarypreparationIDOk returns a tuple with the FkiSalarypreparationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSalarypreparationID

`func (o *CommunicationattachmentResponseCompound) SetFkiSalarypreparationID(v int32)`

SetFkiSalarypreparationID sets FkiSalarypreparationID field to given value.

### HasFkiSalarypreparationID

`func (o *CommunicationattachmentResponseCompound) HasFkiSalarypreparationID() bool`

HasFkiSalarypreparationID returns a boolean if a field has been set.

### GetSCommunicationattachmentName

`func (o *CommunicationattachmentResponseCompound) GetSCommunicationattachmentName() string`

GetSCommunicationattachmentName returns the SCommunicationattachmentName field if non-nil, zero value otherwise.

### GetSCommunicationattachmentNameOk

`func (o *CommunicationattachmentResponseCompound) GetSCommunicationattachmentNameOk() (*string, bool)`

GetSCommunicationattachmentNameOk returns a tuple with the SCommunicationattachmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCommunicationattachmentName

`func (o *CommunicationattachmentResponseCompound) SetSCommunicationattachmentName(v string)`

SetSCommunicationattachmentName sets SCommunicationattachmentName field to given value.


### GetSDownloadUrl

`func (o *CommunicationattachmentResponseCompound) GetSDownloadUrl() string`

GetSDownloadUrl returns the SDownloadUrl field if non-nil, zero value otherwise.

### GetSDownloadUrlOk

`func (o *CommunicationattachmentResponseCompound) GetSDownloadUrlOk() (*string, bool)`

GetSDownloadUrlOk returns a tuple with the SDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDownloadUrl

`func (o *CommunicationattachmentResponseCompound) SetSDownloadUrl(v string)`

SetSDownloadUrl sets SDownloadUrl field to given value.

### HasSDownloadUrl

`func (o *CommunicationattachmentResponseCompound) HasSDownloadUrl() bool`

HasSDownloadUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


