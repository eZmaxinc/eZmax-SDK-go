# DocumenttypechecklistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiDocumenttypechecklistID** | Pointer to **int32** | The unique ID of the Documenttypechecklist | [optional] 
**FkiDocumenttypeID** | **int32** | The unique ID of the Documenttype | 
**EDocumenttypechecklistStatus** | [**FieldEDocumenttypechecklistStatus**](FieldEDocumenttypechecklistStatus.md) |  | 
**TDocumenttypechecklistComment** | **string** | The comment of the Documenttypechecklist | 
**FkiAdjustmentID** | Pointer to **int32** | The unique ID of the Adjustment | [optional] 
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBankaccountID** | Pointer to **int32** | The unique ID of the Bankaccount | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**FkiCommissionadvanceID** | Pointer to **int32** | The unique ID of the Commissionadvance | [optional] 
**FkiCommunicationID** | Pointer to **int32** | The unique ID of the Communication. | [optional] 
**FkiCustomerID** | Pointer to **int32** | The unique ID of the Customer. | [optional] 
**FkiCustomertemplateID** | Pointer to **int32** | The unique ID of the Customertemplate | [optional] 
**FkiDepositID** | Pointer to **int32** | The unique ID of the Deposit | [optional] 
**FkiDeposittransitchequeID** | Pointer to **int32** | The unique ID of the Deposittransitcheque | [optional] 
**FkiElectronicfundstransferID** | Pointer to **int32** | The unique ID of the Electronicfundstransfer | [optional] 
**FkiEmployeeID** | Pointer to **int32** | The unique ID of the Employee. | [optional] 
**FkiExternalbrokerID** | Pointer to **int32** | The unique ID of the Externalbroker. | [optional] 
**FkiEzcomadvanceserverID** | Pointer to **int32** | The unique ID of the Ezcomadvanceserver | [optional] 
**FkiEzcomcompanyID** | Pointer to **int32** | The unique ID of the Ezcomcompany | [optional] 
**FkiEzsigndocumentID** | Pointer to **int32** | The unique ID of the Ezsigndocument | [optional] 
**FkiGhacqcontractID** | Pointer to **int32** | The unique ID of the Ghacqcontract | [optional] 
**FkiInscriptionID** | Pointer to **int32** | The unique ID of the Inscription. | [optional] 
**FkiInscriptiontempID** | Pointer to **int32** | The unique ID of the Inscriptiontemp | [optional] 
**FkiInscriptionnotauthenticatedID** | Pointer to **int32** | The unique ID of the Inscriptionnotauthenticated. | [optional] 
**FkiInvoiceID** | Pointer to **int32** | The unique ID of the Invoice. | [optional] 
**FkiBuyercontractID** | Pointer to **int32** | The unique ID of the Buyercontract | [optional] 
**FkiFranchisebrokerID** | Pointer to **int32** | The unique ID of the Franchisebroker | [optional] 
**FkiFranchiseagenceID** | Pointer to **int32** | The unique ID of the Franchiseagence | [optional] 
**FkiFranchiseofficeID** | Pointer to **int32** | The unique ID of the Franchisereoffice | [optional] 
**FkiFranchisefranchiseID** | Pointer to **int32** | The unique ID of the Franchisefranchise | [optional] 
**FkiFranchisecomplaintID** | Pointer to **int32** | The unique ID of the Franchisecomplaint | [optional] 
**FkiLeadID** | Pointer to **int32** | The unique ID of the Lead | [optional] 
**FkiMarketingprogramID** | Pointer to **int32** | The unique ID of the Marketingprogram | [optional] 
**FkiMarketingfollowID** | Pointer to **int32** | The unique ID of the Marketingfollow | [optional] 
**FkiNotaryID** | Pointer to **int32** | The unique ID of the Notary. | [optional] 
**FkiOfficetaxreportID** | Pointer to **int32** | The unique ID of the Officetaxreport | [optional] 
**FkiOtherincomeID** | Pointer to **int32** | The unique ID of the Otherincome | [optional] 
**FkiPaymentpreparationID** | Pointer to **int32** | The unique ID of the Paymentpreparation | [optional] 
**FkiPurchaseID** | Pointer to **int32** | The unique ID of the Purchase | [optional] 
**FkiSalaryID** | Pointer to **int32** | The unique ID of the Salary | [optional] 
**FkiSupplierID** | Pointer to **int32** | The unique ID of the Supplier. | [optional] 
**FkiTranqcontractID** | Pointer to **int32** | The unique ID of the Tranqcontract | [optional] 
**FkiTemplateID** | Pointer to **int32** | The unique ID of the Template | [optional] 

## Methods

### NewDocumenttypechecklistRequest

`func NewDocumenttypechecklistRequest(fkiDocumenttypeID int32, eDocumenttypechecklistStatus FieldEDocumenttypechecklistStatus, tDocumenttypechecklistComment string, ) *DocumenttypechecklistRequest`

NewDocumenttypechecklistRequest instantiates a new DocumenttypechecklistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumenttypechecklistRequestWithDefaults

`func NewDocumenttypechecklistRequestWithDefaults() *DocumenttypechecklistRequest`

NewDocumenttypechecklistRequestWithDefaults instantiates a new DocumenttypechecklistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiDocumenttypechecklistID

`func (o *DocumenttypechecklistRequest) GetPkiDocumenttypechecklistID() int32`

GetPkiDocumenttypechecklistID returns the PkiDocumenttypechecklistID field if non-nil, zero value otherwise.

### GetPkiDocumenttypechecklistIDOk

`func (o *DocumenttypechecklistRequest) GetPkiDocumenttypechecklistIDOk() (*int32, bool)`

GetPkiDocumenttypechecklistIDOk returns a tuple with the PkiDocumenttypechecklistID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiDocumenttypechecklistID

`func (o *DocumenttypechecklistRequest) SetPkiDocumenttypechecklistID(v int32)`

SetPkiDocumenttypechecklistID sets PkiDocumenttypechecklistID field to given value.

### HasPkiDocumenttypechecklistID

`func (o *DocumenttypechecklistRequest) HasPkiDocumenttypechecklistID() bool`

HasPkiDocumenttypechecklistID returns a boolean if a field has been set.

### GetFkiDocumenttypeID

`func (o *DocumenttypechecklistRequest) GetFkiDocumenttypeID() int32`

GetFkiDocumenttypeID returns the FkiDocumenttypeID field if non-nil, zero value otherwise.

### GetFkiDocumenttypeIDOk

`func (o *DocumenttypechecklistRequest) GetFkiDocumenttypeIDOk() (*int32, bool)`

GetFkiDocumenttypeIDOk returns a tuple with the FkiDocumenttypeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDocumenttypeID

`func (o *DocumenttypechecklistRequest) SetFkiDocumenttypeID(v int32)`

SetFkiDocumenttypeID sets FkiDocumenttypeID field to given value.


### GetEDocumenttypechecklistStatus

`func (o *DocumenttypechecklistRequest) GetEDocumenttypechecklistStatus() FieldEDocumenttypechecklistStatus`

GetEDocumenttypechecklistStatus returns the EDocumenttypechecklistStatus field if non-nil, zero value otherwise.

### GetEDocumenttypechecklistStatusOk

`func (o *DocumenttypechecklistRequest) GetEDocumenttypechecklistStatusOk() (*FieldEDocumenttypechecklistStatus, bool)`

GetEDocumenttypechecklistStatusOk returns a tuple with the EDocumenttypechecklistStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDocumenttypechecklistStatus

`func (o *DocumenttypechecklistRequest) SetEDocumenttypechecklistStatus(v FieldEDocumenttypechecklistStatus)`

SetEDocumenttypechecklistStatus sets EDocumenttypechecklistStatus field to given value.


### GetTDocumenttypechecklistComment

`func (o *DocumenttypechecklistRequest) GetTDocumenttypechecklistComment() string`

GetTDocumenttypechecklistComment returns the TDocumenttypechecklistComment field if non-nil, zero value otherwise.

### GetTDocumenttypechecklistCommentOk

`func (o *DocumenttypechecklistRequest) GetTDocumenttypechecklistCommentOk() (*string, bool)`

GetTDocumenttypechecklistCommentOk returns a tuple with the TDocumenttypechecklistComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTDocumenttypechecklistComment

`func (o *DocumenttypechecklistRequest) SetTDocumenttypechecklistComment(v string)`

SetTDocumenttypechecklistComment sets TDocumenttypechecklistComment field to given value.


### GetFkiAdjustmentID

`func (o *DocumenttypechecklistRequest) GetFkiAdjustmentID() int32`

GetFkiAdjustmentID returns the FkiAdjustmentID field if non-nil, zero value otherwise.

### GetFkiAdjustmentIDOk

`func (o *DocumenttypechecklistRequest) GetFkiAdjustmentIDOk() (*int32, bool)`

GetFkiAdjustmentIDOk returns a tuple with the FkiAdjustmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAdjustmentID

`func (o *DocumenttypechecklistRequest) SetFkiAdjustmentID(v int32)`

SetFkiAdjustmentID sets FkiAdjustmentID field to given value.

### HasFkiAdjustmentID

`func (o *DocumenttypechecklistRequest) HasFkiAdjustmentID() bool`

HasFkiAdjustmentID returns a boolean if a field has been set.

### GetFkiAgentID

`func (o *DocumenttypechecklistRequest) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *DocumenttypechecklistRequest) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *DocumenttypechecklistRequest) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *DocumenttypechecklistRequest) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBankaccountID

`func (o *DocumenttypechecklistRequest) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *DocumenttypechecklistRequest) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *DocumenttypechecklistRequest) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.

### HasFkiBankaccountID

`func (o *DocumenttypechecklistRequest) HasFkiBankaccountID() bool`

HasFkiBankaccountID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *DocumenttypechecklistRequest) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *DocumenttypechecklistRequest) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *DocumenttypechecklistRequest) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *DocumenttypechecklistRequest) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiCommissionadvanceID

`func (o *DocumenttypechecklistRequest) GetFkiCommissionadvanceID() int32`

GetFkiCommissionadvanceID returns the FkiCommissionadvanceID field if non-nil, zero value otherwise.

### GetFkiCommissionadvanceIDOk

`func (o *DocumenttypechecklistRequest) GetFkiCommissionadvanceIDOk() (*int32, bool)`

GetFkiCommissionadvanceIDOk returns a tuple with the FkiCommissionadvanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommissionadvanceID

`func (o *DocumenttypechecklistRequest) SetFkiCommissionadvanceID(v int32)`

SetFkiCommissionadvanceID sets FkiCommissionadvanceID field to given value.

### HasFkiCommissionadvanceID

`func (o *DocumenttypechecklistRequest) HasFkiCommissionadvanceID() bool`

HasFkiCommissionadvanceID returns a boolean if a field has been set.

### GetFkiCommunicationID

`func (o *DocumenttypechecklistRequest) GetFkiCommunicationID() int32`

GetFkiCommunicationID returns the FkiCommunicationID field if non-nil, zero value otherwise.

### GetFkiCommunicationIDOk

`func (o *DocumenttypechecklistRequest) GetFkiCommunicationIDOk() (*int32, bool)`

GetFkiCommunicationIDOk returns a tuple with the FkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommunicationID

`func (o *DocumenttypechecklistRequest) SetFkiCommunicationID(v int32)`

SetFkiCommunicationID sets FkiCommunicationID field to given value.

### HasFkiCommunicationID

`func (o *DocumenttypechecklistRequest) HasFkiCommunicationID() bool`

HasFkiCommunicationID returns a boolean if a field has been set.

### GetFkiCustomerID

`func (o *DocumenttypechecklistRequest) GetFkiCustomerID() int32`

GetFkiCustomerID returns the FkiCustomerID field if non-nil, zero value otherwise.

### GetFkiCustomerIDOk

`func (o *DocumenttypechecklistRequest) GetFkiCustomerIDOk() (*int32, bool)`

GetFkiCustomerIDOk returns a tuple with the FkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerID

`func (o *DocumenttypechecklistRequest) SetFkiCustomerID(v int32)`

SetFkiCustomerID sets FkiCustomerID field to given value.

### HasFkiCustomerID

`func (o *DocumenttypechecklistRequest) HasFkiCustomerID() bool`

HasFkiCustomerID returns a boolean if a field has been set.

### GetFkiCustomertemplateID

`func (o *DocumenttypechecklistRequest) GetFkiCustomertemplateID() int32`

GetFkiCustomertemplateID returns the FkiCustomertemplateID field if non-nil, zero value otherwise.

### GetFkiCustomertemplateIDOk

`func (o *DocumenttypechecklistRequest) GetFkiCustomertemplateIDOk() (*int32, bool)`

GetFkiCustomertemplateIDOk returns a tuple with the FkiCustomertemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomertemplateID

`func (o *DocumenttypechecklistRequest) SetFkiCustomertemplateID(v int32)`

SetFkiCustomertemplateID sets FkiCustomertemplateID field to given value.

### HasFkiCustomertemplateID

`func (o *DocumenttypechecklistRequest) HasFkiCustomertemplateID() bool`

HasFkiCustomertemplateID returns a boolean if a field has been set.

### GetFkiDepositID

`func (o *DocumenttypechecklistRequest) GetFkiDepositID() int32`

GetFkiDepositID returns the FkiDepositID field if non-nil, zero value otherwise.

### GetFkiDepositIDOk

`func (o *DocumenttypechecklistRequest) GetFkiDepositIDOk() (*int32, bool)`

GetFkiDepositIDOk returns a tuple with the FkiDepositID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepositID

`func (o *DocumenttypechecklistRequest) SetFkiDepositID(v int32)`

SetFkiDepositID sets FkiDepositID field to given value.

### HasFkiDepositID

`func (o *DocumenttypechecklistRequest) HasFkiDepositID() bool`

HasFkiDepositID returns a boolean if a field has been set.

### GetFkiDeposittransitchequeID

`func (o *DocumenttypechecklistRequest) GetFkiDeposittransitchequeID() int32`

GetFkiDeposittransitchequeID returns the FkiDeposittransitchequeID field if non-nil, zero value otherwise.

### GetFkiDeposittransitchequeIDOk

`func (o *DocumenttypechecklistRequest) GetFkiDeposittransitchequeIDOk() (*int32, bool)`

GetFkiDeposittransitchequeIDOk returns a tuple with the FkiDeposittransitchequeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDeposittransitchequeID

`func (o *DocumenttypechecklistRequest) SetFkiDeposittransitchequeID(v int32)`

SetFkiDeposittransitchequeID sets FkiDeposittransitchequeID field to given value.

### HasFkiDeposittransitchequeID

`func (o *DocumenttypechecklistRequest) HasFkiDeposittransitchequeID() bool`

HasFkiDeposittransitchequeID returns a boolean if a field has been set.

### GetFkiElectronicfundstransferID

`func (o *DocumenttypechecklistRequest) GetFkiElectronicfundstransferID() int32`

GetFkiElectronicfundstransferID returns the FkiElectronicfundstransferID field if non-nil, zero value otherwise.

### GetFkiElectronicfundstransferIDOk

`func (o *DocumenttypechecklistRequest) GetFkiElectronicfundstransferIDOk() (*int32, bool)`

GetFkiElectronicfundstransferIDOk returns a tuple with the FkiElectronicfundstransferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiElectronicfundstransferID

`func (o *DocumenttypechecklistRequest) SetFkiElectronicfundstransferID(v int32)`

SetFkiElectronicfundstransferID sets FkiElectronicfundstransferID field to given value.

### HasFkiElectronicfundstransferID

`func (o *DocumenttypechecklistRequest) HasFkiElectronicfundstransferID() bool`

HasFkiElectronicfundstransferID returns a boolean if a field has been set.

### GetFkiEmployeeID

`func (o *DocumenttypechecklistRequest) GetFkiEmployeeID() int32`

GetFkiEmployeeID returns the FkiEmployeeID field if non-nil, zero value otherwise.

### GetFkiEmployeeIDOk

`func (o *DocumenttypechecklistRequest) GetFkiEmployeeIDOk() (*int32, bool)`

GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmployeeID

`func (o *DocumenttypechecklistRequest) SetFkiEmployeeID(v int32)`

SetFkiEmployeeID sets FkiEmployeeID field to given value.

### HasFkiEmployeeID

`func (o *DocumenttypechecklistRequest) HasFkiEmployeeID() bool`

HasFkiEmployeeID returns a boolean if a field has been set.

### GetFkiExternalbrokerID

`func (o *DocumenttypechecklistRequest) GetFkiExternalbrokerID() int32`

GetFkiExternalbrokerID returns the FkiExternalbrokerID field if non-nil, zero value otherwise.

### GetFkiExternalbrokerIDOk

`func (o *DocumenttypechecklistRequest) GetFkiExternalbrokerIDOk() (*int32, bool)`

GetFkiExternalbrokerIDOk returns a tuple with the FkiExternalbrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiExternalbrokerID

`func (o *DocumenttypechecklistRequest) SetFkiExternalbrokerID(v int32)`

SetFkiExternalbrokerID sets FkiExternalbrokerID field to given value.

### HasFkiExternalbrokerID

`func (o *DocumenttypechecklistRequest) HasFkiExternalbrokerID() bool`

HasFkiExternalbrokerID returns a boolean if a field has been set.

### GetFkiEzcomadvanceserverID

`func (o *DocumenttypechecklistRequest) GetFkiEzcomadvanceserverID() int32`

GetFkiEzcomadvanceserverID returns the FkiEzcomadvanceserverID field if non-nil, zero value otherwise.

### GetFkiEzcomadvanceserverIDOk

`func (o *DocumenttypechecklistRequest) GetFkiEzcomadvanceserverIDOk() (*int32, bool)`

GetFkiEzcomadvanceserverIDOk returns a tuple with the FkiEzcomadvanceserverID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomadvanceserverID

`func (o *DocumenttypechecklistRequest) SetFkiEzcomadvanceserverID(v int32)`

SetFkiEzcomadvanceserverID sets FkiEzcomadvanceserverID field to given value.

### HasFkiEzcomadvanceserverID

`func (o *DocumenttypechecklistRequest) HasFkiEzcomadvanceserverID() bool`

HasFkiEzcomadvanceserverID returns a boolean if a field has been set.

### GetFkiEzcomcompanyID

`func (o *DocumenttypechecklistRequest) GetFkiEzcomcompanyID() int32`

GetFkiEzcomcompanyID returns the FkiEzcomcompanyID field if non-nil, zero value otherwise.

### GetFkiEzcomcompanyIDOk

`func (o *DocumenttypechecklistRequest) GetFkiEzcomcompanyIDOk() (*int32, bool)`

GetFkiEzcomcompanyIDOk returns a tuple with the FkiEzcomcompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomcompanyID

`func (o *DocumenttypechecklistRequest) SetFkiEzcomcompanyID(v int32)`

SetFkiEzcomcompanyID sets FkiEzcomcompanyID field to given value.

### HasFkiEzcomcompanyID

`func (o *DocumenttypechecklistRequest) HasFkiEzcomcompanyID() bool`

HasFkiEzcomcompanyID returns a boolean if a field has been set.

### GetFkiEzsigndocumentID

`func (o *DocumenttypechecklistRequest) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *DocumenttypechecklistRequest) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *DocumenttypechecklistRequest) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.

### HasFkiEzsigndocumentID

`func (o *DocumenttypechecklistRequest) HasFkiEzsigndocumentID() bool`

HasFkiEzsigndocumentID returns a boolean if a field has been set.

### GetFkiGhacqcontractID

`func (o *DocumenttypechecklistRequest) GetFkiGhacqcontractID() int32`

GetFkiGhacqcontractID returns the FkiGhacqcontractID field if non-nil, zero value otherwise.

### GetFkiGhacqcontractIDOk

`func (o *DocumenttypechecklistRequest) GetFkiGhacqcontractIDOk() (*int32, bool)`

GetFkiGhacqcontractIDOk returns a tuple with the FkiGhacqcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGhacqcontractID

`func (o *DocumenttypechecklistRequest) SetFkiGhacqcontractID(v int32)`

SetFkiGhacqcontractID sets FkiGhacqcontractID field to given value.

### HasFkiGhacqcontractID

`func (o *DocumenttypechecklistRequest) HasFkiGhacqcontractID() bool`

HasFkiGhacqcontractID returns a boolean if a field has been set.

### GetFkiInscriptionID

`func (o *DocumenttypechecklistRequest) GetFkiInscriptionID() int32`

GetFkiInscriptionID returns the FkiInscriptionID field if non-nil, zero value otherwise.

### GetFkiInscriptionIDOk

`func (o *DocumenttypechecklistRequest) GetFkiInscriptionIDOk() (*int32, bool)`

GetFkiInscriptionIDOk returns a tuple with the FkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionID

`func (o *DocumenttypechecklistRequest) SetFkiInscriptionID(v int32)`

SetFkiInscriptionID sets FkiInscriptionID field to given value.

### HasFkiInscriptionID

`func (o *DocumenttypechecklistRequest) HasFkiInscriptionID() bool`

HasFkiInscriptionID returns a boolean if a field has been set.

### GetFkiInscriptiontempID

`func (o *DocumenttypechecklistRequest) GetFkiInscriptiontempID() int32`

GetFkiInscriptiontempID returns the FkiInscriptiontempID field if non-nil, zero value otherwise.

### GetFkiInscriptiontempIDOk

`func (o *DocumenttypechecklistRequest) GetFkiInscriptiontempIDOk() (*int32, bool)`

GetFkiInscriptiontempIDOk returns a tuple with the FkiInscriptiontempID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontempID

`func (o *DocumenttypechecklistRequest) SetFkiInscriptiontempID(v int32)`

SetFkiInscriptiontempID sets FkiInscriptiontempID field to given value.

### HasFkiInscriptiontempID

`func (o *DocumenttypechecklistRequest) HasFkiInscriptiontempID() bool`

HasFkiInscriptiontempID returns a boolean if a field has been set.

### GetFkiInscriptionnotauthenticatedID

`func (o *DocumenttypechecklistRequest) GetFkiInscriptionnotauthenticatedID() int32`

GetFkiInscriptionnotauthenticatedID returns the FkiInscriptionnotauthenticatedID field if non-nil, zero value otherwise.

### GetFkiInscriptionnotauthenticatedIDOk

`func (o *DocumenttypechecklistRequest) GetFkiInscriptionnotauthenticatedIDOk() (*int32, bool)`

GetFkiInscriptionnotauthenticatedIDOk returns a tuple with the FkiInscriptionnotauthenticatedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionnotauthenticatedID

`func (o *DocumenttypechecklistRequest) SetFkiInscriptionnotauthenticatedID(v int32)`

SetFkiInscriptionnotauthenticatedID sets FkiInscriptionnotauthenticatedID field to given value.

### HasFkiInscriptionnotauthenticatedID

`func (o *DocumenttypechecklistRequest) HasFkiInscriptionnotauthenticatedID() bool`

HasFkiInscriptionnotauthenticatedID returns a boolean if a field has been set.

### GetFkiInvoiceID

`func (o *DocumenttypechecklistRequest) GetFkiInvoiceID() int32`

GetFkiInvoiceID returns the FkiInvoiceID field if non-nil, zero value otherwise.

### GetFkiInvoiceIDOk

`func (o *DocumenttypechecklistRequest) GetFkiInvoiceIDOk() (*int32, bool)`

GetFkiInvoiceIDOk returns a tuple with the FkiInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInvoiceID

`func (o *DocumenttypechecklistRequest) SetFkiInvoiceID(v int32)`

SetFkiInvoiceID sets FkiInvoiceID field to given value.

### HasFkiInvoiceID

`func (o *DocumenttypechecklistRequest) HasFkiInvoiceID() bool`

HasFkiInvoiceID returns a boolean if a field has been set.

### GetFkiBuyercontractID

`func (o *DocumenttypechecklistRequest) GetFkiBuyercontractID() int32`

GetFkiBuyercontractID returns the FkiBuyercontractID field if non-nil, zero value otherwise.

### GetFkiBuyercontractIDOk

`func (o *DocumenttypechecklistRequest) GetFkiBuyercontractIDOk() (*int32, bool)`

GetFkiBuyercontractIDOk returns a tuple with the FkiBuyercontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBuyercontractID

`func (o *DocumenttypechecklistRequest) SetFkiBuyercontractID(v int32)`

SetFkiBuyercontractID sets FkiBuyercontractID field to given value.

### HasFkiBuyercontractID

`func (o *DocumenttypechecklistRequest) HasFkiBuyercontractID() bool`

HasFkiBuyercontractID returns a boolean if a field has been set.

### GetFkiFranchisebrokerID

`func (o *DocumenttypechecklistRequest) GetFkiFranchisebrokerID() int32`

GetFkiFranchisebrokerID returns the FkiFranchisebrokerID field if non-nil, zero value otherwise.

### GetFkiFranchisebrokerIDOk

`func (o *DocumenttypechecklistRequest) GetFkiFranchisebrokerIDOk() (*int32, bool)`

GetFkiFranchisebrokerIDOk returns a tuple with the FkiFranchisebrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisebrokerID

`func (o *DocumenttypechecklistRequest) SetFkiFranchisebrokerID(v int32)`

SetFkiFranchisebrokerID sets FkiFranchisebrokerID field to given value.

### HasFkiFranchisebrokerID

`func (o *DocumenttypechecklistRequest) HasFkiFranchisebrokerID() bool`

HasFkiFranchisebrokerID returns a boolean if a field has been set.

### GetFkiFranchiseagenceID

`func (o *DocumenttypechecklistRequest) GetFkiFranchiseagenceID() int32`

GetFkiFranchiseagenceID returns the FkiFranchiseagenceID field if non-nil, zero value otherwise.

### GetFkiFranchiseagenceIDOk

`func (o *DocumenttypechecklistRequest) GetFkiFranchiseagenceIDOk() (*int32, bool)`

GetFkiFranchiseagenceIDOk returns a tuple with the FkiFranchiseagenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseagenceID

`func (o *DocumenttypechecklistRequest) SetFkiFranchiseagenceID(v int32)`

SetFkiFranchiseagenceID sets FkiFranchiseagenceID field to given value.

### HasFkiFranchiseagenceID

`func (o *DocumenttypechecklistRequest) HasFkiFranchiseagenceID() bool`

HasFkiFranchiseagenceID returns a boolean if a field has been set.

### GetFkiFranchiseofficeID

`func (o *DocumenttypechecklistRequest) GetFkiFranchiseofficeID() int32`

GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field if non-nil, zero value otherwise.

### GetFkiFranchiseofficeIDOk

`func (o *DocumenttypechecklistRequest) GetFkiFranchiseofficeIDOk() (*int32, bool)`

GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseofficeID

`func (o *DocumenttypechecklistRequest) SetFkiFranchiseofficeID(v int32)`

SetFkiFranchiseofficeID sets FkiFranchiseofficeID field to given value.

### HasFkiFranchiseofficeID

`func (o *DocumenttypechecklistRequest) HasFkiFranchiseofficeID() bool`

HasFkiFranchiseofficeID returns a boolean if a field has been set.

### GetFkiFranchisefranchiseID

`func (o *DocumenttypechecklistRequest) GetFkiFranchisefranchiseID() int32`

GetFkiFranchisefranchiseID returns the FkiFranchisefranchiseID field if non-nil, zero value otherwise.

### GetFkiFranchisefranchiseIDOk

`func (o *DocumenttypechecklistRequest) GetFkiFranchisefranchiseIDOk() (*int32, bool)`

GetFkiFranchisefranchiseIDOk returns a tuple with the FkiFranchisefranchiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisefranchiseID

`func (o *DocumenttypechecklistRequest) SetFkiFranchisefranchiseID(v int32)`

SetFkiFranchisefranchiseID sets FkiFranchisefranchiseID field to given value.

### HasFkiFranchisefranchiseID

`func (o *DocumenttypechecklistRequest) HasFkiFranchisefranchiseID() bool`

HasFkiFranchisefranchiseID returns a boolean if a field has been set.

### GetFkiFranchisecomplaintID

`func (o *DocumenttypechecklistRequest) GetFkiFranchisecomplaintID() int32`

GetFkiFranchisecomplaintID returns the FkiFranchisecomplaintID field if non-nil, zero value otherwise.

### GetFkiFranchisecomplaintIDOk

`func (o *DocumenttypechecklistRequest) GetFkiFranchisecomplaintIDOk() (*int32, bool)`

GetFkiFranchisecomplaintIDOk returns a tuple with the FkiFranchisecomplaintID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisecomplaintID

`func (o *DocumenttypechecklistRequest) SetFkiFranchisecomplaintID(v int32)`

SetFkiFranchisecomplaintID sets FkiFranchisecomplaintID field to given value.

### HasFkiFranchisecomplaintID

`func (o *DocumenttypechecklistRequest) HasFkiFranchisecomplaintID() bool`

HasFkiFranchisecomplaintID returns a boolean if a field has been set.

### GetFkiLeadID

`func (o *DocumenttypechecklistRequest) GetFkiLeadID() int32`

GetFkiLeadID returns the FkiLeadID field if non-nil, zero value otherwise.

### GetFkiLeadIDOk

`func (o *DocumenttypechecklistRequest) GetFkiLeadIDOk() (*int32, bool)`

GetFkiLeadIDOk returns a tuple with the FkiLeadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLeadID

`func (o *DocumenttypechecklistRequest) SetFkiLeadID(v int32)`

SetFkiLeadID sets FkiLeadID field to given value.

### HasFkiLeadID

`func (o *DocumenttypechecklistRequest) HasFkiLeadID() bool`

HasFkiLeadID returns a boolean if a field has been set.

### GetFkiMarketingprogramID

`func (o *DocumenttypechecklistRequest) GetFkiMarketingprogramID() int32`

GetFkiMarketingprogramID returns the FkiMarketingprogramID field if non-nil, zero value otherwise.

### GetFkiMarketingprogramIDOk

`func (o *DocumenttypechecklistRequest) GetFkiMarketingprogramIDOk() (*int32, bool)`

GetFkiMarketingprogramIDOk returns a tuple with the FkiMarketingprogramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMarketingprogramID

`func (o *DocumenttypechecklistRequest) SetFkiMarketingprogramID(v int32)`

SetFkiMarketingprogramID sets FkiMarketingprogramID field to given value.

### HasFkiMarketingprogramID

`func (o *DocumenttypechecklistRequest) HasFkiMarketingprogramID() bool`

HasFkiMarketingprogramID returns a boolean if a field has been set.

### GetFkiMarketingfollowID

`func (o *DocumenttypechecklistRequest) GetFkiMarketingfollowID() int32`

GetFkiMarketingfollowID returns the FkiMarketingfollowID field if non-nil, zero value otherwise.

### GetFkiMarketingfollowIDOk

`func (o *DocumenttypechecklistRequest) GetFkiMarketingfollowIDOk() (*int32, bool)`

GetFkiMarketingfollowIDOk returns a tuple with the FkiMarketingfollowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMarketingfollowID

`func (o *DocumenttypechecklistRequest) SetFkiMarketingfollowID(v int32)`

SetFkiMarketingfollowID sets FkiMarketingfollowID field to given value.

### HasFkiMarketingfollowID

`func (o *DocumenttypechecklistRequest) HasFkiMarketingfollowID() bool`

HasFkiMarketingfollowID returns a boolean if a field has been set.

### GetFkiNotaryID

`func (o *DocumenttypechecklistRequest) GetFkiNotaryID() int32`

GetFkiNotaryID returns the FkiNotaryID field if non-nil, zero value otherwise.

### GetFkiNotaryIDOk

`func (o *DocumenttypechecklistRequest) GetFkiNotaryIDOk() (*int32, bool)`

GetFkiNotaryIDOk returns a tuple with the FkiNotaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotaryID

`func (o *DocumenttypechecklistRequest) SetFkiNotaryID(v int32)`

SetFkiNotaryID sets FkiNotaryID field to given value.

### HasFkiNotaryID

`func (o *DocumenttypechecklistRequest) HasFkiNotaryID() bool`

HasFkiNotaryID returns a boolean if a field has been set.

### GetFkiOfficetaxreportID

`func (o *DocumenttypechecklistRequest) GetFkiOfficetaxreportID() int32`

GetFkiOfficetaxreportID returns the FkiOfficetaxreportID field if non-nil, zero value otherwise.

### GetFkiOfficetaxreportIDOk

`func (o *DocumenttypechecklistRequest) GetFkiOfficetaxreportIDOk() (*int32, bool)`

GetFkiOfficetaxreportIDOk returns a tuple with the FkiOfficetaxreportID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOfficetaxreportID

`func (o *DocumenttypechecklistRequest) SetFkiOfficetaxreportID(v int32)`

SetFkiOfficetaxreportID sets FkiOfficetaxreportID field to given value.

### HasFkiOfficetaxreportID

`func (o *DocumenttypechecklistRequest) HasFkiOfficetaxreportID() bool`

HasFkiOfficetaxreportID returns a boolean if a field has been set.

### GetFkiOtherincomeID

`func (o *DocumenttypechecklistRequest) GetFkiOtherincomeID() int32`

GetFkiOtherincomeID returns the FkiOtherincomeID field if non-nil, zero value otherwise.

### GetFkiOtherincomeIDOk

`func (o *DocumenttypechecklistRequest) GetFkiOtherincomeIDOk() (*int32, bool)`

GetFkiOtherincomeIDOk returns a tuple with the FkiOtherincomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOtherincomeID

`func (o *DocumenttypechecklistRequest) SetFkiOtherincomeID(v int32)`

SetFkiOtherincomeID sets FkiOtherincomeID field to given value.

### HasFkiOtherincomeID

`func (o *DocumenttypechecklistRequest) HasFkiOtherincomeID() bool`

HasFkiOtherincomeID returns a boolean if a field has been set.

### GetFkiPaymentpreparationID

`func (o *DocumenttypechecklistRequest) GetFkiPaymentpreparationID() int32`

GetFkiPaymentpreparationID returns the FkiPaymentpreparationID field if non-nil, zero value otherwise.

### GetFkiPaymentpreparationIDOk

`func (o *DocumenttypechecklistRequest) GetFkiPaymentpreparationIDOk() (*int32, bool)`

GetFkiPaymentpreparationIDOk returns a tuple with the FkiPaymentpreparationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPaymentpreparationID

`func (o *DocumenttypechecklistRequest) SetFkiPaymentpreparationID(v int32)`

SetFkiPaymentpreparationID sets FkiPaymentpreparationID field to given value.

### HasFkiPaymentpreparationID

`func (o *DocumenttypechecklistRequest) HasFkiPaymentpreparationID() bool`

HasFkiPaymentpreparationID returns a boolean if a field has been set.

### GetFkiPurchaseID

`func (o *DocumenttypechecklistRequest) GetFkiPurchaseID() int32`

GetFkiPurchaseID returns the FkiPurchaseID field if non-nil, zero value otherwise.

### GetFkiPurchaseIDOk

`func (o *DocumenttypechecklistRequest) GetFkiPurchaseIDOk() (*int32, bool)`

GetFkiPurchaseIDOk returns a tuple with the FkiPurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPurchaseID

`func (o *DocumenttypechecklistRequest) SetFkiPurchaseID(v int32)`

SetFkiPurchaseID sets FkiPurchaseID field to given value.

### HasFkiPurchaseID

`func (o *DocumenttypechecklistRequest) HasFkiPurchaseID() bool`

HasFkiPurchaseID returns a boolean if a field has been set.

### GetFkiSalaryID

`func (o *DocumenttypechecklistRequest) GetFkiSalaryID() int32`

GetFkiSalaryID returns the FkiSalaryID field if non-nil, zero value otherwise.

### GetFkiSalaryIDOk

`func (o *DocumenttypechecklistRequest) GetFkiSalaryIDOk() (*int32, bool)`

GetFkiSalaryIDOk returns a tuple with the FkiSalaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSalaryID

`func (o *DocumenttypechecklistRequest) SetFkiSalaryID(v int32)`

SetFkiSalaryID sets FkiSalaryID field to given value.

### HasFkiSalaryID

`func (o *DocumenttypechecklistRequest) HasFkiSalaryID() bool`

HasFkiSalaryID returns a boolean if a field has been set.

### GetFkiSupplierID

`func (o *DocumenttypechecklistRequest) GetFkiSupplierID() int32`

GetFkiSupplierID returns the FkiSupplierID field if non-nil, zero value otherwise.

### GetFkiSupplierIDOk

`func (o *DocumenttypechecklistRequest) GetFkiSupplierIDOk() (*int32, bool)`

GetFkiSupplierIDOk returns a tuple with the FkiSupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSupplierID

`func (o *DocumenttypechecklistRequest) SetFkiSupplierID(v int32)`

SetFkiSupplierID sets FkiSupplierID field to given value.

### HasFkiSupplierID

`func (o *DocumenttypechecklistRequest) HasFkiSupplierID() bool`

HasFkiSupplierID returns a boolean if a field has been set.

### GetFkiTranqcontractID

`func (o *DocumenttypechecklistRequest) GetFkiTranqcontractID() int32`

GetFkiTranqcontractID returns the FkiTranqcontractID field if non-nil, zero value otherwise.

### GetFkiTranqcontractIDOk

`func (o *DocumenttypechecklistRequest) GetFkiTranqcontractIDOk() (*int32, bool)`

GetFkiTranqcontractIDOk returns a tuple with the FkiTranqcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTranqcontractID

`func (o *DocumenttypechecklistRequest) SetFkiTranqcontractID(v int32)`

SetFkiTranqcontractID sets FkiTranqcontractID field to given value.

### HasFkiTranqcontractID

`func (o *DocumenttypechecklistRequest) HasFkiTranqcontractID() bool`

HasFkiTranqcontractID returns a boolean if a field has been set.

### GetFkiTemplateID

`func (o *DocumenttypechecklistRequest) GetFkiTemplateID() int32`

GetFkiTemplateID returns the FkiTemplateID field if non-nil, zero value otherwise.

### GetFkiTemplateIDOk

`func (o *DocumenttypechecklistRequest) GetFkiTemplateIDOk() (*int32, bool)`

GetFkiTemplateIDOk returns a tuple with the FkiTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTemplateID

`func (o *DocumenttypechecklistRequest) SetFkiTemplateID(v int32)`

SetFkiTemplateID sets FkiTemplateID field to given value.

### HasFkiTemplateID

`func (o *DocumenttypechecklistRequest) HasFkiTemplateID() bool`

HasFkiTemplateID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


