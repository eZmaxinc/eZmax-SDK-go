# CommunicationrecipientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiCommunicationrecipientID** | Pointer to **int32** | The unique ID of the Communicationrecipient. | [optional] 
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**FkiContactID** | Pointer to **int32** | The unique ID of the Contact | [optional] 
**FkiCustomerID** | Pointer to **int32** | The unique ID of the Customer. | [optional] 
**FkiEmployeeID** | Pointer to **int32** | The unique ID of the Employee. | [optional] 
**FkiAssistantID** | Pointer to **int32** | The unique ID of the Assistant. | [optional] 
**FkiExternalbrokerID** | Pointer to **int32** | The unique ID of the Externalbroker. | [optional] 
**FkiEzsignsignerID** | Pointer to **int32** | The unique ID of the Ezsignsigner | [optional] 
**FkiNotaryID** | Pointer to **int32** | The unique ID of the Notary. | [optional] 
**FkiSupplierID** | Pointer to **int32** | The unique ID of the Supplier. | [optional] 
**FkiUserID** | Pointer to **int32** | The unique ID of the User | [optional] 
**FkiMailboxsharedID** | Pointer to **int32** | The unique ID of the Mailboxshared | [optional] 
**FkiPhonelinesharedID** | Pointer to **int32** | The unique ID of the Phonelineshared | [optional] 
**ECommunicationrecipientType** | Pointer to [**FieldECommunicationrecipientType**](FieldECommunicationrecipientType.md) |  | [optional] 

## Methods

### NewCommunicationrecipientRequest

`func NewCommunicationrecipientRequest() *CommunicationrecipientRequest`

NewCommunicationrecipientRequest instantiates a new CommunicationrecipientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationrecipientRequestWithDefaults

`func NewCommunicationrecipientRequestWithDefaults() *CommunicationrecipientRequest`

NewCommunicationrecipientRequestWithDefaults instantiates a new CommunicationrecipientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiCommunicationrecipientID

`func (o *CommunicationrecipientRequest) GetPkiCommunicationrecipientID() int32`

GetPkiCommunicationrecipientID returns the PkiCommunicationrecipientID field if non-nil, zero value otherwise.

### GetPkiCommunicationrecipientIDOk

`func (o *CommunicationrecipientRequest) GetPkiCommunicationrecipientIDOk() (*int32, bool)`

GetPkiCommunicationrecipientIDOk returns a tuple with the PkiCommunicationrecipientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiCommunicationrecipientID

`func (o *CommunicationrecipientRequest) SetPkiCommunicationrecipientID(v int32)`

SetPkiCommunicationrecipientID sets PkiCommunicationrecipientID field to given value.

### HasPkiCommunicationrecipientID

`func (o *CommunicationrecipientRequest) HasPkiCommunicationrecipientID() bool`

HasPkiCommunicationrecipientID returns a boolean if a field has been set.

### GetFkiAgentID

`func (o *CommunicationrecipientRequest) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *CommunicationrecipientRequest) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *CommunicationrecipientRequest) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *CommunicationrecipientRequest) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *CommunicationrecipientRequest) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *CommunicationrecipientRequest) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *CommunicationrecipientRequest) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *CommunicationrecipientRequest) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiContactID

`func (o *CommunicationrecipientRequest) GetFkiContactID() int32`

GetFkiContactID returns the FkiContactID field if non-nil, zero value otherwise.

### GetFkiContactIDOk

`func (o *CommunicationrecipientRequest) GetFkiContactIDOk() (*int32, bool)`

GetFkiContactIDOk returns a tuple with the FkiContactID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiContactID

`func (o *CommunicationrecipientRequest) SetFkiContactID(v int32)`

SetFkiContactID sets FkiContactID field to given value.

### HasFkiContactID

`func (o *CommunicationrecipientRequest) HasFkiContactID() bool`

HasFkiContactID returns a boolean if a field has been set.

### GetFkiCustomerID

`func (o *CommunicationrecipientRequest) GetFkiCustomerID() int32`

GetFkiCustomerID returns the FkiCustomerID field if non-nil, zero value otherwise.

### GetFkiCustomerIDOk

`func (o *CommunicationrecipientRequest) GetFkiCustomerIDOk() (*int32, bool)`

GetFkiCustomerIDOk returns a tuple with the FkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerID

`func (o *CommunicationrecipientRequest) SetFkiCustomerID(v int32)`

SetFkiCustomerID sets FkiCustomerID field to given value.

### HasFkiCustomerID

`func (o *CommunicationrecipientRequest) HasFkiCustomerID() bool`

HasFkiCustomerID returns a boolean if a field has been set.

### GetFkiEmployeeID

`func (o *CommunicationrecipientRequest) GetFkiEmployeeID() int32`

GetFkiEmployeeID returns the FkiEmployeeID field if non-nil, zero value otherwise.

### GetFkiEmployeeIDOk

`func (o *CommunicationrecipientRequest) GetFkiEmployeeIDOk() (*int32, bool)`

GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmployeeID

`func (o *CommunicationrecipientRequest) SetFkiEmployeeID(v int32)`

SetFkiEmployeeID sets FkiEmployeeID field to given value.

### HasFkiEmployeeID

`func (o *CommunicationrecipientRequest) HasFkiEmployeeID() bool`

HasFkiEmployeeID returns a boolean if a field has been set.

### GetFkiAssistantID

`func (o *CommunicationrecipientRequest) GetFkiAssistantID() int32`

GetFkiAssistantID returns the FkiAssistantID field if non-nil, zero value otherwise.

### GetFkiAssistantIDOk

`func (o *CommunicationrecipientRequest) GetFkiAssistantIDOk() (*int32, bool)`

GetFkiAssistantIDOk returns a tuple with the FkiAssistantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAssistantID

`func (o *CommunicationrecipientRequest) SetFkiAssistantID(v int32)`

SetFkiAssistantID sets FkiAssistantID field to given value.

### HasFkiAssistantID

`func (o *CommunicationrecipientRequest) HasFkiAssistantID() bool`

HasFkiAssistantID returns a boolean if a field has been set.

### GetFkiExternalbrokerID

`func (o *CommunicationrecipientRequest) GetFkiExternalbrokerID() int32`

GetFkiExternalbrokerID returns the FkiExternalbrokerID field if non-nil, zero value otherwise.

### GetFkiExternalbrokerIDOk

`func (o *CommunicationrecipientRequest) GetFkiExternalbrokerIDOk() (*int32, bool)`

GetFkiExternalbrokerIDOk returns a tuple with the FkiExternalbrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiExternalbrokerID

`func (o *CommunicationrecipientRequest) SetFkiExternalbrokerID(v int32)`

SetFkiExternalbrokerID sets FkiExternalbrokerID field to given value.

### HasFkiExternalbrokerID

`func (o *CommunicationrecipientRequest) HasFkiExternalbrokerID() bool`

HasFkiExternalbrokerID returns a boolean if a field has been set.

### GetFkiEzsignsignerID

`func (o *CommunicationrecipientRequest) GetFkiEzsignsignerID() int32`

GetFkiEzsignsignerID returns the FkiEzsignsignerID field if non-nil, zero value otherwise.

### GetFkiEzsignsignerIDOk

`func (o *CommunicationrecipientRequest) GetFkiEzsignsignerIDOk() (*int32, bool)`

GetFkiEzsignsignerIDOk returns a tuple with the FkiEzsignsignerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsignsignerID

`func (o *CommunicationrecipientRequest) SetFkiEzsignsignerID(v int32)`

SetFkiEzsignsignerID sets FkiEzsignsignerID field to given value.

### HasFkiEzsignsignerID

`func (o *CommunicationrecipientRequest) HasFkiEzsignsignerID() bool`

HasFkiEzsignsignerID returns a boolean if a field has been set.

### GetFkiNotaryID

`func (o *CommunicationrecipientRequest) GetFkiNotaryID() int32`

GetFkiNotaryID returns the FkiNotaryID field if non-nil, zero value otherwise.

### GetFkiNotaryIDOk

`func (o *CommunicationrecipientRequest) GetFkiNotaryIDOk() (*int32, bool)`

GetFkiNotaryIDOk returns a tuple with the FkiNotaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotaryID

`func (o *CommunicationrecipientRequest) SetFkiNotaryID(v int32)`

SetFkiNotaryID sets FkiNotaryID field to given value.

### HasFkiNotaryID

`func (o *CommunicationrecipientRequest) HasFkiNotaryID() bool`

HasFkiNotaryID returns a boolean if a field has been set.

### GetFkiSupplierID

`func (o *CommunicationrecipientRequest) GetFkiSupplierID() int32`

GetFkiSupplierID returns the FkiSupplierID field if non-nil, zero value otherwise.

### GetFkiSupplierIDOk

`func (o *CommunicationrecipientRequest) GetFkiSupplierIDOk() (*int32, bool)`

GetFkiSupplierIDOk returns a tuple with the FkiSupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSupplierID

`func (o *CommunicationrecipientRequest) SetFkiSupplierID(v int32)`

SetFkiSupplierID sets FkiSupplierID field to given value.

### HasFkiSupplierID

`func (o *CommunicationrecipientRequest) HasFkiSupplierID() bool`

HasFkiSupplierID returns a boolean if a field has been set.

### GetFkiUserID

`func (o *CommunicationrecipientRequest) GetFkiUserID() int32`

GetFkiUserID returns the FkiUserID field if non-nil, zero value otherwise.

### GetFkiUserIDOk

`func (o *CommunicationrecipientRequest) GetFkiUserIDOk() (*int32, bool)`

GetFkiUserIDOk returns a tuple with the FkiUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserID

`func (o *CommunicationrecipientRequest) SetFkiUserID(v int32)`

SetFkiUserID sets FkiUserID field to given value.

### HasFkiUserID

`func (o *CommunicationrecipientRequest) HasFkiUserID() bool`

HasFkiUserID returns a boolean if a field has been set.

### GetFkiMailboxsharedID

`func (o *CommunicationrecipientRequest) GetFkiMailboxsharedID() int32`

GetFkiMailboxsharedID returns the FkiMailboxsharedID field if non-nil, zero value otherwise.

### GetFkiMailboxsharedIDOk

`func (o *CommunicationrecipientRequest) GetFkiMailboxsharedIDOk() (*int32, bool)`

GetFkiMailboxsharedIDOk returns a tuple with the FkiMailboxsharedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMailboxsharedID

`func (o *CommunicationrecipientRequest) SetFkiMailboxsharedID(v int32)`

SetFkiMailboxsharedID sets FkiMailboxsharedID field to given value.

### HasFkiMailboxsharedID

`func (o *CommunicationrecipientRequest) HasFkiMailboxsharedID() bool`

HasFkiMailboxsharedID returns a boolean if a field has been set.

### GetFkiPhonelinesharedID

`func (o *CommunicationrecipientRequest) GetFkiPhonelinesharedID() int32`

GetFkiPhonelinesharedID returns the FkiPhonelinesharedID field if non-nil, zero value otherwise.

### GetFkiPhonelinesharedIDOk

`func (o *CommunicationrecipientRequest) GetFkiPhonelinesharedIDOk() (*int32, bool)`

GetFkiPhonelinesharedIDOk returns a tuple with the FkiPhonelinesharedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPhonelinesharedID

`func (o *CommunicationrecipientRequest) SetFkiPhonelinesharedID(v int32)`

SetFkiPhonelinesharedID sets FkiPhonelinesharedID field to given value.

### HasFkiPhonelinesharedID

`func (o *CommunicationrecipientRequest) HasFkiPhonelinesharedID() bool`

HasFkiPhonelinesharedID returns a boolean if a field has been set.

### GetECommunicationrecipientType

`func (o *CommunicationrecipientRequest) GetECommunicationrecipientType() FieldECommunicationrecipientType`

GetECommunicationrecipientType returns the ECommunicationrecipientType field if non-nil, zero value otherwise.

### GetECommunicationrecipientTypeOk

`func (o *CommunicationrecipientRequest) GetECommunicationrecipientTypeOk() (*FieldECommunicationrecipientType, bool)`

GetECommunicationrecipientTypeOk returns a tuple with the ECommunicationrecipientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetECommunicationrecipientType

`func (o *CommunicationrecipientRequest) SetECommunicationrecipientType(v FieldECommunicationrecipientType)`

SetECommunicationrecipientType sets ECommunicationrecipientType field to given value.

### HasECommunicationrecipientType

`func (o *CommunicationrecipientRequest) HasECommunicationrecipientType() bool`

HasECommunicationrecipientType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


