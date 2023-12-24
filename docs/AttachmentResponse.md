# AttachmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiAttachmentID** | **int32** | The unique ID of the Attachment. | 
**FkiComputerID** | Pointer to **int32** | The unique ID of the Computer | [optional] 
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
**FkiInscriptionchecklistID** | Pointer to **int32** | The unique ID of the Inscriptionchecklist | [optional] 
**FkiFolderID** | Pointer to **int32** | The unique ID of the Folder | [optional] 
**FkiRejectedoffertopurchaseID** | Pointer to **int32** | The unique ID of the Rejectedoffertopurchase | [optional] 
**FkiDisclosureID** | Pointer to **int32** | The unique ID of the Disclosure | [optional] 
**FkiReconciliationID** | Pointer to **int32** | The unique ID of the Reconciliation | [optional] 
**FkiEzsigndocumentIDReference** | Pointer to **int32** | The unique ID of the Ezsigndocument | [optional] 
**EAttachmentDocumenttype** | [**FieldEAttachmentDocumenttype**](FieldEAttachmentDocumenttype.md) |  | 
**SAttachmentName** | **string** | The name of the Attachment | 
**EAttachmentPrivacy** | [**FieldEAttachmentPrivacy**](FieldEAttachmentPrivacy.md) |  | 
**FkiUserIDSpecific** | Pointer to **int32** | The unique ID of the User | [optional] 
**EAttachmentType** | [**FieldEAttachmentType**](FieldEAttachmentType.md) |  | 
**IAttachmentSize** | **int32** | The size of the Attachment | 
**IAttachmentEDMmoduleflag** | Pointer to **int32** | The edmmoduleflag of the Attachment | [optional] 
**SAttachmentMD5** | **string** | The md5 of the Attachment | 
**BAttachmentDeleted** | **bool** | Whether if it&#39;s deleted | 
**BAttachmentValid** | **bool** | Whether if it&#39;s valid | 
**EAttachmentVerified** | [**FieldEAttachmentVerified**](FieldEAttachmentVerified.md) |  | 
**TAttachmentRejectioncomment** | Pointer to **string** | The rejectioncomment of the Attachment | [optional] 
**FkiUserIDOwner** | Pointer to **int32** | The unique ID of the User | [optional] 
**ObjAudit** | Pointer to [**CommonAudit**](CommonAudit.md) |  | [optional] 

## Methods

### NewAttachmentResponse

`func NewAttachmentResponse(pkiAttachmentID int32, eAttachmentDocumenttype FieldEAttachmentDocumenttype, sAttachmentName string, eAttachmentPrivacy FieldEAttachmentPrivacy, eAttachmentType FieldEAttachmentType, iAttachmentSize int32, sAttachmentMD5 string, bAttachmentDeleted bool, bAttachmentValid bool, eAttachmentVerified FieldEAttachmentVerified, ) *AttachmentResponse`

NewAttachmentResponse instantiates a new AttachmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentResponseWithDefaults

`func NewAttachmentResponseWithDefaults() *AttachmentResponse`

NewAttachmentResponseWithDefaults instantiates a new AttachmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiAttachmentID

`func (o *AttachmentResponse) GetPkiAttachmentID() int32`

GetPkiAttachmentID returns the PkiAttachmentID field if non-nil, zero value otherwise.

### GetPkiAttachmentIDOk

`func (o *AttachmentResponse) GetPkiAttachmentIDOk() (*int32, bool)`

GetPkiAttachmentIDOk returns a tuple with the PkiAttachmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiAttachmentID

`func (o *AttachmentResponse) SetPkiAttachmentID(v int32)`

SetPkiAttachmentID sets PkiAttachmentID field to given value.


### GetFkiComputerID

`func (o *AttachmentResponse) GetFkiComputerID() int32`

GetFkiComputerID returns the FkiComputerID field if non-nil, zero value otherwise.

### GetFkiComputerIDOk

`func (o *AttachmentResponse) GetFkiComputerIDOk() (*int32, bool)`

GetFkiComputerIDOk returns a tuple with the FkiComputerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiComputerID

`func (o *AttachmentResponse) SetFkiComputerID(v int32)`

SetFkiComputerID sets FkiComputerID field to given value.

### HasFkiComputerID

`func (o *AttachmentResponse) HasFkiComputerID() bool`

HasFkiComputerID returns a boolean if a field has been set.

### GetFkiAdjustmentID

`func (o *AttachmentResponse) GetFkiAdjustmentID() int32`

GetFkiAdjustmentID returns the FkiAdjustmentID field if non-nil, zero value otherwise.

### GetFkiAdjustmentIDOk

`func (o *AttachmentResponse) GetFkiAdjustmentIDOk() (*int32, bool)`

GetFkiAdjustmentIDOk returns a tuple with the FkiAdjustmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAdjustmentID

`func (o *AttachmentResponse) SetFkiAdjustmentID(v int32)`

SetFkiAdjustmentID sets FkiAdjustmentID field to given value.

### HasFkiAdjustmentID

`func (o *AttachmentResponse) HasFkiAdjustmentID() bool`

HasFkiAdjustmentID returns a boolean if a field has been set.

### GetFkiAgentID

`func (o *AttachmentResponse) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *AttachmentResponse) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *AttachmentResponse) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *AttachmentResponse) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBankaccountID

`func (o *AttachmentResponse) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *AttachmentResponse) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *AttachmentResponse) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.

### HasFkiBankaccountID

`func (o *AttachmentResponse) HasFkiBankaccountID() bool`

HasFkiBankaccountID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *AttachmentResponse) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *AttachmentResponse) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *AttachmentResponse) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *AttachmentResponse) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiCommissionadvanceID

`func (o *AttachmentResponse) GetFkiCommissionadvanceID() int32`

GetFkiCommissionadvanceID returns the FkiCommissionadvanceID field if non-nil, zero value otherwise.

### GetFkiCommissionadvanceIDOk

`func (o *AttachmentResponse) GetFkiCommissionadvanceIDOk() (*int32, bool)`

GetFkiCommissionadvanceIDOk returns a tuple with the FkiCommissionadvanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommissionadvanceID

`func (o *AttachmentResponse) SetFkiCommissionadvanceID(v int32)`

SetFkiCommissionadvanceID sets FkiCommissionadvanceID field to given value.

### HasFkiCommissionadvanceID

`func (o *AttachmentResponse) HasFkiCommissionadvanceID() bool`

HasFkiCommissionadvanceID returns a boolean if a field has been set.

### GetFkiCommunicationID

`func (o *AttachmentResponse) GetFkiCommunicationID() int32`

GetFkiCommunicationID returns the FkiCommunicationID field if non-nil, zero value otherwise.

### GetFkiCommunicationIDOk

`func (o *AttachmentResponse) GetFkiCommunicationIDOk() (*int32, bool)`

GetFkiCommunicationIDOk returns a tuple with the FkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommunicationID

`func (o *AttachmentResponse) SetFkiCommunicationID(v int32)`

SetFkiCommunicationID sets FkiCommunicationID field to given value.

### HasFkiCommunicationID

`func (o *AttachmentResponse) HasFkiCommunicationID() bool`

HasFkiCommunicationID returns a boolean if a field has been set.

### GetFkiCustomerID

`func (o *AttachmentResponse) GetFkiCustomerID() int32`

GetFkiCustomerID returns the FkiCustomerID field if non-nil, zero value otherwise.

### GetFkiCustomerIDOk

`func (o *AttachmentResponse) GetFkiCustomerIDOk() (*int32, bool)`

GetFkiCustomerIDOk returns a tuple with the FkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerID

`func (o *AttachmentResponse) SetFkiCustomerID(v int32)`

SetFkiCustomerID sets FkiCustomerID field to given value.

### HasFkiCustomerID

`func (o *AttachmentResponse) HasFkiCustomerID() bool`

HasFkiCustomerID returns a boolean if a field has been set.

### GetFkiCustomertemplateID

`func (o *AttachmentResponse) GetFkiCustomertemplateID() int32`

GetFkiCustomertemplateID returns the FkiCustomertemplateID field if non-nil, zero value otherwise.

### GetFkiCustomertemplateIDOk

`func (o *AttachmentResponse) GetFkiCustomertemplateIDOk() (*int32, bool)`

GetFkiCustomertemplateIDOk returns a tuple with the FkiCustomertemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomertemplateID

`func (o *AttachmentResponse) SetFkiCustomertemplateID(v int32)`

SetFkiCustomertemplateID sets FkiCustomertemplateID field to given value.

### HasFkiCustomertemplateID

`func (o *AttachmentResponse) HasFkiCustomertemplateID() bool`

HasFkiCustomertemplateID returns a boolean if a field has been set.

### GetFkiDepositID

`func (o *AttachmentResponse) GetFkiDepositID() int32`

GetFkiDepositID returns the FkiDepositID field if non-nil, zero value otherwise.

### GetFkiDepositIDOk

`func (o *AttachmentResponse) GetFkiDepositIDOk() (*int32, bool)`

GetFkiDepositIDOk returns a tuple with the FkiDepositID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepositID

`func (o *AttachmentResponse) SetFkiDepositID(v int32)`

SetFkiDepositID sets FkiDepositID field to given value.

### HasFkiDepositID

`func (o *AttachmentResponse) HasFkiDepositID() bool`

HasFkiDepositID returns a boolean if a field has been set.

### GetFkiDeposittransitchequeID

`func (o *AttachmentResponse) GetFkiDeposittransitchequeID() int32`

GetFkiDeposittransitchequeID returns the FkiDeposittransitchequeID field if non-nil, zero value otherwise.

### GetFkiDeposittransitchequeIDOk

`func (o *AttachmentResponse) GetFkiDeposittransitchequeIDOk() (*int32, bool)`

GetFkiDeposittransitchequeIDOk returns a tuple with the FkiDeposittransitchequeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDeposittransitchequeID

`func (o *AttachmentResponse) SetFkiDeposittransitchequeID(v int32)`

SetFkiDeposittransitchequeID sets FkiDeposittransitchequeID field to given value.

### HasFkiDeposittransitchequeID

`func (o *AttachmentResponse) HasFkiDeposittransitchequeID() bool`

HasFkiDeposittransitchequeID returns a boolean if a field has been set.

### GetFkiElectronicfundstransferID

`func (o *AttachmentResponse) GetFkiElectronicfundstransferID() int32`

GetFkiElectronicfundstransferID returns the FkiElectronicfundstransferID field if non-nil, zero value otherwise.

### GetFkiElectronicfundstransferIDOk

`func (o *AttachmentResponse) GetFkiElectronicfundstransferIDOk() (*int32, bool)`

GetFkiElectronicfundstransferIDOk returns a tuple with the FkiElectronicfundstransferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiElectronicfundstransferID

`func (o *AttachmentResponse) SetFkiElectronicfundstransferID(v int32)`

SetFkiElectronicfundstransferID sets FkiElectronicfundstransferID field to given value.

### HasFkiElectronicfundstransferID

`func (o *AttachmentResponse) HasFkiElectronicfundstransferID() bool`

HasFkiElectronicfundstransferID returns a boolean if a field has been set.

### GetFkiEmployeeID

`func (o *AttachmentResponse) GetFkiEmployeeID() int32`

GetFkiEmployeeID returns the FkiEmployeeID field if non-nil, zero value otherwise.

### GetFkiEmployeeIDOk

`func (o *AttachmentResponse) GetFkiEmployeeIDOk() (*int32, bool)`

GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmployeeID

`func (o *AttachmentResponse) SetFkiEmployeeID(v int32)`

SetFkiEmployeeID sets FkiEmployeeID field to given value.

### HasFkiEmployeeID

`func (o *AttachmentResponse) HasFkiEmployeeID() bool`

HasFkiEmployeeID returns a boolean if a field has been set.

### GetFkiExternalbrokerID

`func (o *AttachmentResponse) GetFkiExternalbrokerID() int32`

GetFkiExternalbrokerID returns the FkiExternalbrokerID field if non-nil, zero value otherwise.

### GetFkiExternalbrokerIDOk

`func (o *AttachmentResponse) GetFkiExternalbrokerIDOk() (*int32, bool)`

GetFkiExternalbrokerIDOk returns a tuple with the FkiExternalbrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiExternalbrokerID

`func (o *AttachmentResponse) SetFkiExternalbrokerID(v int32)`

SetFkiExternalbrokerID sets FkiExternalbrokerID field to given value.

### HasFkiExternalbrokerID

`func (o *AttachmentResponse) HasFkiExternalbrokerID() bool`

HasFkiExternalbrokerID returns a boolean if a field has been set.

### GetFkiEzcomadvanceserverID

`func (o *AttachmentResponse) GetFkiEzcomadvanceserverID() int32`

GetFkiEzcomadvanceserverID returns the FkiEzcomadvanceserverID field if non-nil, zero value otherwise.

### GetFkiEzcomadvanceserverIDOk

`func (o *AttachmentResponse) GetFkiEzcomadvanceserverIDOk() (*int32, bool)`

GetFkiEzcomadvanceserverIDOk returns a tuple with the FkiEzcomadvanceserverID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomadvanceserverID

`func (o *AttachmentResponse) SetFkiEzcomadvanceserverID(v int32)`

SetFkiEzcomadvanceserverID sets FkiEzcomadvanceserverID field to given value.

### HasFkiEzcomadvanceserverID

`func (o *AttachmentResponse) HasFkiEzcomadvanceserverID() bool`

HasFkiEzcomadvanceserverID returns a boolean if a field has been set.

### GetFkiEzcomcompanyID

`func (o *AttachmentResponse) GetFkiEzcomcompanyID() int32`

GetFkiEzcomcompanyID returns the FkiEzcomcompanyID field if non-nil, zero value otherwise.

### GetFkiEzcomcompanyIDOk

`func (o *AttachmentResponse) GetFkiEzcomcompanyIDOk() (*int32, bool)`

GetFkiEzcomcompanyIDOk returns a tuple with the FkiEzcomcompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomcompanyID

`func (o *AttachmentResponse) SetFkiEzcomcompanyID(v int32)`

SetFkiEzcomcompanyID sets FkiEzcomcompanyID field to given value.

### HasFkiEzcomcompanyID

`func (o *AttachmentResponse) HasFkiEzcomcompanyID() bool`

HasFkiEzcomcompanyID returns a boolean if a field has been set.

### GetFkiEzsigndocumentID

`func (o *AttachmentResponse) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *AttachmentResponse) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *AttachmentResponse) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.

### HasFkiEzsigndocumentID

`func (o *AttachmentResponse) HasFkiEzsigndocumentID() bool`

HasFkiEzsigndocumentID returns a boolean if a field has been set.

### GetFkiGhacqcontractID

`func (o *AttachmentResponse) GetFkiGhacqcontractID() int32`

GetFkiGhacqcontractID returns the FkiGhacqcontractID field if non-nil, zero value otherwise.

### GetFkiGhacqcontractIDOk

`func (o *AttachmentResponse) GetFkiGhacqcontractIDOk() (*int32, bool)`

GetFkiGhacqcontractIDOk returns a tuple with the FkiGhacqcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGhacqcontractID

`func (o *AttachmentResponse) SetFkiGhacqcontractID(v int32)`

SetFkiGhacqcontractID sets FkiGhacqcontractID field to given value.

### HasFkiGhacqcontractID

`func (o *AttachmentResponse) HasFkiGhacqcontractID() bool`

HasFkiGhacqcontractID returns a boolean if a field has been set.

### GetFkiInscriptionID

`func (o *AttachmentResponse) GetFkiInscriptionID() int32`

GetFkiInscriptionID returns the FkiInscriptionID field if non-nil, zero value otherwise.

### GetFkiInscriptionIDOk

`func (o *AttachmentResponse) GetFkiInscriptionIDOk() (*int32, bool)`

GetFkiInscriptionIDOk returns a tuple with the FkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionID

`func (o *AttachmentResponse) SetFkiInscriptionID(v int32)`

SetFkiInscriptionID sets FkiInscriptionID field to given value.

### HasFkiInscriptionID

`func (o *AttachmentResponse) HasFkiInscriptionID() bool`

HasFkiInscriptionID returns a boolean if a field has been set.

### GetFkiInscriptiontempID

`func (o *AttachmentResponse) GetFkiInscriptiontempID() int32`

GetFkiInscriptiontempID returns the FkiInscriptiontempID field if non-nil, zero value otherwise.

### GetFkiInscriptiontempIDOk

`func (o *AttachmentResponse) GetFkiInscriptiontempIDOk() (*int32, bool)`

GetFkiInscriptiontempIDOk returns a tuple with the FkiInscriptiontempID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontempID

`func (o *AttachmentResponse) SetFkiInscriptiontempID(v int32)`

SetFkiInscriptiontempID sets FkiInscriptiontempID field to given value.

### HasFkiInscriptiontempID

`func (o *AttachmentResponse) HasFkiInscriptiontempID() bool`

HasFkiInscriptiontempID returns a boolean if a field has been set.

### GetFkiInscriptionnotauthenticatedID

`func (o *AttachmentResponse) GetFkiInscriptionnotauthenticatedID() int32`

GetFkiInscriptionnotauthenticatedID returns the FkiInscriptionnotauthenticatedID field if non-nil, zero value otherwise.

### GetFkiInscriptionnotauthenticatedIDOk

`func (o *AttachmentResponse) GetFkiInscriptionnotauthenticatedIDOk() (*int32, bool)`

GetFkiInscriptionnotauthenticatedIDOk returns a tuple with the FkiInscriptionnotauthenticatedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionnotauthenticatedID

`func (o *AttachmentResponse) SetFkiInscriptionnotauthenticatedID(v int32)`

SetFkiInscriptionnotauthenticatedID sets FkiInscriptionnotauthenticatedID field to given value.

### HasFkiInscriptionnotauthenticatedID

`func (o *AttachmentResponse) HasFkiInscriptionnotauthenticatedID() bool`

HasFkiInscriptionnotauthenticatedID returns a boolean if a field has been set.

### GetFkiInvoiceID

`func (o *AttachmentResponse) GetFkiInvoiceID() int32`

GetFkiInvoiceID returns the FkiInvoiceID field if non-nil, zero value otherwise.

### GetFkiInvoiceIDOk

`func (o *AttachmentResponse) GetFkiInvoiceIDOk() (*int32, bool)`

GetFkiInvoiceIDOk returns a tuple with the FkiInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInvoiceID

`func (o *AttachmentResponse) SetFkiInvoiceID(v int32)`

SetFkiInvoiceID sets FkiInvoiceID field to given value.

### HasFkiInvoiceID

`func (o *AttachmentResponse) HasFkiInvoiceID() bool`

HasFkiInvoiceID returns a boolean if a field has been set.

### GetFkiBuyercontractID

`func (o *AttachmentResponse) GetFkiBuyercontractID() int32`

GetFkiBuyercontractID returns the FkiBuyercontractID field if non-nil, zero value otherwise.

### GetFkiBuyercontractIDOk

`func (o *AttachmentResponse) GetFkiBuyercontractIDOk() (*int32, bool)`

GetFkiBuyercontractIDOk returns a tuple with the FkiBuyercontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBuyercontractID

`func (o *AttachmentResponse) SetFkiBuyercontractID(v int32)`

SetFkiBuyercontractID sets FkiBuyercontractID field to given value.

### HasFkiBuyercontractID

`func (o *AttachmentResponse) HasFkiBuyercontractID() bool`

HasFkiBuyercontractID returns a boolean if a field has been set.

### GetFkiFranchisebrokerID

`func (o *AttachmentResponse) GetFkiFranchisebrokerID() int32`

GetFkiFranchisebrokerID returns the FkiFranchisebrokerID field if non-nil, zero value otherwise.

### GetFkiFranchisebrokerIDOk

`func (o *AttachmentResponse) GetFkiFranchisebrokerIDOk() (*int32, bool)`

GetFkiFranchisebrokerIDOk returns a tuple with the FkiFranchisebrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisebrokerID

`func (o *AttachmentResponse) SetFkiFranchisebrokerID(v int32)`

SetFkiFranchisebrokerID sets FkiFranchisebrokerID field to given value.

### HasFkiFranchisebrokerID

`func (o *AttachmentResponse) HasFkiFranchisebrokerID() bool`

HasFkiFranchisebrokerID returns a boolean if a field has been set.

### GetFkiFranchiseagenceID

`func (o *AttachmentResponse) GetFkiFranchiseagenceID() int32`

GetFkiFranchiseagenceID returns the FkiFranchiseagenceID field if non-nil, zero value otherwise.

### GetFkiFranchiseagenceIDOk

`func (o *AttachmentResponse) GetFkiFranchiseagenceIDOk() (*int32, bool)`

GetFkiFranchiseagenceIDOk returns a tuple with the FkiFranchiseagenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseagenceID

`func (o *AttachmentResponse) SetFkiFranchiseagenceID(v int32)`

SetFkiFranchiseagenceID sets FkiFranchiseagenceID field to given value.

### HasFkiFranchiseagenceID

`func (o *AttachmentResponse) HasFkiFranchiseagenceID() bool`

HasFkiFranchiseagenceID returns a boolean if a field has been set.

### GetFkiFranchiseofficeID

`func (o *AttachmentResponse) GetFkiFranchiseofficeID() int32`

GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field if non-nil, zero value otherwise.

### GetFkiFranchiseofficeIDOk

`func (o *AttachmentResponse) GetFkiFranchiseofficeIDOk() (*int32, bool)`

GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseofficeID

`func (o *AttachmentResponse) SetFkiFranchiseofficeID(v int32)`

SetFkiFranchiseofficeID sets FkiFranchiseofficeID field to given value.

### HasFkiFranchiseofficeID

`func (o *AttachmentResponse) HasFkiFranchiseofficeID() bool`

HasFkiFranchiseofficeID returns a boolean if a field has been set.

### GetFkiFranchisefranchiseID

`func (o *AttachmentResponse) GetFkiFranchisefranchiseID() int32`

GetFkiFranchisefranchiseID returns the FkiFranchisefranchiseID field if non-nil, zero value otherwise.

### GetFkiFranchisefranchiseIDOk

`func (o *AttachmentResponse) GetFkiFranchisefranchiseIDOk() (*int32, bool)`

GetFkiFranchisefranchiseIDOk returns a tuple with the FkiFranchisefranchiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisefranchiseID

`func (o *AttachmentResponse) SetFkiFranchisefranchiseID(v int32)`

SetFkiFranchisefranchiseID sets FkiFranchisefranchiseID field to given value.

### HasFkiFranchisefranchiseID

`func (o *AttachmentResponse) HasFkiFranchisefranchiseID() bool`

HasFkiFranchisefranchiseID returns a boolean if a field has been set.

### GetFkiFranchisecomplaintID

`func (o *AttachmentResponse) GetFkiFranchisecomplaintID() int32`

GetFkiFranchisecomplaintID returns the FkiFranchisecomplaintID field if non-nil, zero value otherwise.

### GetFkiFranchisecomplaintIDOk

`func (o *AttachmentResponse) GetFkiFranchisecomplaintIDOk() (*int32, bool)`

GetFkiFranchisecomplaintIDOk returns a tuple with the FkiFranchisecomplaintID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisecomplaintID

`func (o *AttachmentResponse) SetFkiFranchisecomplaintID(v int32)`

SetFkiFranchisecomplaintID sets FkiFranchisecomplaintID field to given value.

### HasFkiFranchisecomplaintID

`func (o *AttachmentResponse) HasFkiFranchisecomplaintID() bool`

HasFkiFranchisecomplaintID returns a boolean if a field has been set.

### GetFkiLeadID

`func (o *AttachmentResponse) GetFkiLeadID() int32`

GetFkiLeadID returns the FkiLeadID field if non-nil, zero value otherwise.

### GetFkiLeadIDOk

`func (o *AttachmentResponse) GetFkiLeadIDOk() (*int32, bool)`

GetFkiLeadIDOk returns a tuple with the FkiLeadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLeadID

`func (o *AttachmentResponse) SetFkiLeadID(v int32)`

SetFkiLeadID sets FkiLeadID field to given value.

### HasFkiLeadID

`func (o *AttachmentResponse) HasFkiLeadID() bool`

HasFkiLeadID returns a boolean if a field has been set.

### GetFkiMarketingprogramID

`func (o *AttachmentResponse) GetFkiMarketingprogramID() int32`

GetFkiMarketingprogramID returns the FkiMarketingprogramID field if non-nil, zero value otherwise.

### GetFkiMarketingprogramIDOk

`func (o *AttachmentResponse) GetFkiMarketingprogramIDOk() (*int32, bool)`

GetFkiMarketingprogramIDOk returns a tuple with the FkiMarketingprogramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMarketingprogramID

`func (o *AttachmentResponse) SetFkiMarketingprogramID(v int32)`

SetFkiMarketingprogramID sets FkiMarketingprogramID field to given value.

### HasFkiMarketingprogramID

`func (o *AttachmentResponse) HasFkiMarketingprogramID() bool`

HasFkiMarketingprogramID returns a boolean if a field has been set.

### GetFkiMarketingfollowID

`func (o *AttachmentResponse) GetFkiMarketingfollowID() int32`

GetFkiMarketingfollowID returns the FkiMarketingfollowID field if non-nil, zero value otherwise.

### GetFkiMarketingfollowIDOk

`func (o *AttachmentResponse) GetFkiMarketingfollowIDOk() (*int32, bool)`

GetFkiMarketingfollowIDOk returns a tuple with the FkiMarketingfollowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMarketingfollowID

`func (o *AttachmentResponse) SetFkiMarketingfollowID(v int32)`

SetFkiMarketingfollowID sets FkiMarketingfollowID field to given value.

### HasFkiMarketingfollowID

`func (o *AttachmentResponse) HasFkiMarketingfollowID() bool`

HasFkiMarketingfollowID returns a boolean if a field has been set.

### GetFkiNotaryID

`func (o *AttachmentResponse) GetFkiNotaryID() int32`

GetFkiNotaryID returns the FkiNotaryID field if non-nil, zero value otherwise.

### GetFkiNotaryIDOk

`func (o *AttachmentResponse) GetFkiNotaryIDOk() (*int32, bool)`

GetFkiNotaryIDOk returns a tuple with the FkiNotaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotaryID

`func (o *AttachmentResponse) SetFkiNotaryID(v int32)`

SetFkiNotaryID sets FkiNotaryID field to given value.

### HasFkiNotaryID

`func (o *AttachmentResponse) HasFkiNotaryID() bool`

HasFkiNotaryID returns a boolean if a field has been set.

### GetFkiOfficetaxreportID

`func (o *AttachmentResponse) GetFkiOfficetaxreportID() int32`

GetFkiOfficetaxreportID returns the FkiOfficetaxreportID field if non-nil, zero value otherwise.

### GetFkiOfficetaxreportIDOk

`func (o *AttachmentResponse) GetFkiOfficetaxreportIDOk() (*int32, bool)`

GetFkiOfficetaxreportIDOk returns a tuple with the FkiOfficetaxreportID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOfficetaxreportID

`func (o *AttachmentResponse) SetFkiOfficetaxreportID(v int32)`

SetFkiOfficetaxreportID sets FkiOfficetaxreportID field to given value.

### HasFkiOfficetaxreportID

`func (o *AttachmentResponse) HasFkiOfficetaxreportID() bool`

HasFkiOfficetaxreportID returns a boolean if a field has been set.

### GetFkiOtherincomeID

`func (o *AttachmentResponse) GetFkiOtherincomeID() int32`

GetFkiOtherincomeID returns the FkiOtherincomeID field if non-nil, zero value otherwise.

### GetFkiOtherincomeIDOk

`func (o *AttachmentResponse) GetFkiOtherincomeIDOk() (*int32, bool)`

GetFkiOtherincomeIDOk returns a tuple with the FkiOtherincomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOtherincomeID

`func (o *AttachmentResponse) SetFkiOtherincomeID(v int32)`

SetFkiOtherincomeID sets FkiOtherincomeID field to given value.

### HasFkiOtherincomeID

`func (o *AttachmentResponse) HasFkiOtherincomeID() bool`

HasFkiOtherincomeID returns a boolean if a field has been set.

### GetFkiPaymentpreparationID

`func (o *AttachmentResponse) GetFkiPaymentpreparationID() int32`

GetFkiPaymentpreparationID returns the FkiPaymentpreparationID field if non-nil, zero value otherwise.

### GetFkiPaymentpreparationIDOk

`func (o *AttachmentResponse) GetFkiPaymentpreparationIDOk() (*int32, bool)`

GetFkiPaymentpreparationIDOk returns a tuple with the FkiPaymentpreparationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPaymentpreparationID

`func (o *AttachmentResponse) SetFkiPaymentpreparationID(v int32)`

SetFkiPaymentpreparationID sets FkiPaymentpreparationID field to given value.

### HasFkiPaymentpreparationID

`func (o *AttachmentResponse) HasFkiPaymentpreparationID() bool`

HasFkiPaymentpreparationID returns a boolean if a field has been set.

### GetFkiPurchaseID

`func (o *AttachmentResponse) GetFkiPurchaseID() int32`

GetFkiPurchaseID returns the FkiPurchaseID field if non-nil, zero value otherwise.

### GetFkiPurchaseIDOk

`func (o *AttachmentResponse) GetFkiPurchaseIDOk() (*int32, bool)`

GetFkiPurchaseIDOk returns a tuple with the FkiPurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPurchaseID

`func (o *AttachmentResponse) SetFkiPurchaseID(v int32)`

SetFkiPurchaseID sets FkiPurchaseID field to given value.

### HasFkiPurchaseID

`func (o *AttachmentResponse) HasFkiPurchaseID() bool`

HasFkiPurchaseID returns a boolean if a field has been set.

### GetFkiSalaryID

`func (o *AttachmentResponse) GetFkiSalaryID() int32`

GetFkiSalaryID returns the FkiSalaryID field if non-nil, zero value otherwise.

### GetFkiSalaryIDOk

`func (o *AttachmentResponse) GetFkiSalaryIDOk() (*int32, bool)`

GetFkiSalaryIDOk returns a tuple with the FkiSalaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSalaryID

`func (o *AttachmentResponse) SetFkiSalaryID(v int32)`

SetFkiSalaryID sets FkiSalaryID field to given value.

### HasFkiSalaryID

`func (o *AttachmentResponse) HasFkiSalaryID() bool`

HasFkiSalaryID returns a boolean if a field has been set.

### GetFkiSupplierID

`func (o *AttachmentResponse) GetFkiSupplierID() int32`

GetFkiSupplierID returns the FkiSupplierID field if non-nil, zero value otherwise.

### GetFkiSupplierIDOk

`func (o *AttachmentResponse) GetFkiSupplierIDOk() (*int32, bool)`

GetFkiSupplierIDOk returns a tuple with the FkiSupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSupplierID

`func (o *AttachmentResponse) SetFkiSupplierID(v int32)`

SetFkiSupplierID sets FkiSupplierID field to given value.

### HasFkiSupplierID

`func (o *AttachmentResponse) HasFkiSupplierID() bool`

HasFkiSupplierID returns a boolean if a field has been set.

### GetFkiTranqcontractID

`func (o *AttachmentResponse) GetFkiTranqcontractID() int32`

GetFkiTranqcontractID returns the FkiTranqcontractID field if non-nil, zero value otherwise.

### GetFkiTranqcontractIDOk

`func (o *AttachmentResponse) GetFkiTranqcontractIDOk() (*int32, bool)`

GetFkiTranqcontractIDOk returns a tuple with the FkiTranqcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTranqcontractID

`func (o *AttachmentResponse) SetFkiTranqcontractID(v int32)`

SetFkiTranqcontractID sets FkiTranqcontractID field to given value.

### HasFkiTranqcontractID

`func (o *AttachmentResponse) HasFkiTranqcontractID() bool`

HasFkiTranqcontractID returns a boolean if a field has been set.

### GetFkiTemplateID

`func (o *AttachmentResponse) GetFkiTemplateID() int32`

GetFkiTemplateID returns the FkiTemplateID field if non-nil, zero value otherwise.

### GetFkiTemplateIDOk

`func (o *AttachmentResponse) GetFkiTemplateIDOk() (*int32, bool)`

GetFkiTemplateIDOk returns a tuple with the FkiTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTemplateID

`func (o *AttachmentResponse) SetFkiTemplateID(v int32)`

SetFkiTemplateID sets FkiTemplateID field to given value.

### HasFkiTemplateID

`func (o *AttachmentResponse) HasFkiTemplateID() bool`

HasFkiTemplateID returns a boolean if a field has been set.

### GetFkiInscriptionchecklistID

`func (o *AttachmentResponse) GetFkiInscriptionchecklistID() int32`

GetFkiInscriptionchecklistID returns the FkiInscriptionchecklistID field if non-nil, zero value otherwise.

### GetFkiInscriptionchecklistIDOk

`func (o *AttachmentResponse) GetFkiInscriptionchecklistIDOk() (*int32, bool)`

GetFkiInscriptionchecklistIDOk returns a tuple with the FkiInscriptionchecklistID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionchecklistID

`func (o *AttachmentResponse) SetFkiInscriptionchecklistID(v int32)`

SetFkiInscriptionchecklistID sets FkiInscriptionchecklistID field to given value.

### HasFkiInscriptionchecklistID

`func (o *AttachmentResponse) HasFkiInscriptionchecklistID() bool`

HasFkiInscriptionchecklistID returns a boolean if a field has been set.

### GetFkiFolderID

`func (o *AttachmentResponse) GetFkiFolderID() int32`

GetFkiFolderID returns the FkiFolderID field if non-nil, zero value otherwise.

### GetFkiFolderIDOk

`func (o *AttachmentResponse) GetFkiFolderIDOk() (*int32, bool)`

GetFkiFolderIDOk returns a tuple with the FkiFolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFolderID

`func (o *AttachmentResponse) SetFkiFolderID(v int32)`

SetFkiFolderID sets FkiFolderID field to given value.

### HasFkiFolderID

`func (o *AttachmentResponse) HasFkiFolderID() bool`

HasFkiFolderID returns a boolean if a field has been set.

### GetFkiRejectedoffertopurchaseID

`func (o *AttachmentResponse) GetFkiRejectedoffertopurchaseID() int32`

GetFkiRejectedoffertopurchaseID returns the FkiRejectedoffertopurchaseID field if non-nil, zero value otherwise.

### GetFkiRejectedoffertopurchaseIDOk

`func (o *AttachmentResponse) GetFkiRejectedoffertopurchaseIDOk() (*int32, bool)`

GetFkiRejectedoffertopurchaseIDOk returns a tuple with the FkiRejectedoffertopurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiRejectedoffertopurchaseID

`func (o *AttachmentResponse) SetFkiRejectedoffertopurchaseID(v int32)`

SetFkiRejectedoffertopurchaseID sets FkiRejectedoffertopurchaseID field to given value.

### HasFkiRejectedoffertopurchaseID

`func (o *AttachmentResponse) HasFkiRejectedoffertopurchaseID() bool`

HasFkiRejectedoffertopurchaseID returns a boolean if a field has been set.

### GetFkiDisclosureID

`func (o *AttachmentResponse) GetFkiDisclosureID() int32`

GetFkiDisclosureID returns the FkiDisclosureID field if non-nil, zero value otherwise.

### GetFkiDisclosureIDOk

`func (o *AttachmentResponse) GetFkiDisclosureIDOk() (*int32, bool)`

GetFkiDisclosureIDOk returns a tuple with the FkiDisclosureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDisclosureID

`func (o *AttachmentResponse) SetFkiDisclosureID(v int32)`

SetFkiDisclosureID sets FkiDisclosureID field to given value.

### HasFkiDisclosureID

`func (o *AttachmentResponse) HasFkiDisclosureID() bool`

HasFkiDisclosureID returns a boolean if a field has been set.

### GetFkiReconciliationID

`func (o *AttachmentResponse) GetFkiReconciliationID() int32`

GetFkiReconciliationID returns the FkiReconciliationID field if non-nil, zero value otherwise.

### GetFkiReconciliationIDOk

`func (o *AttachmentResponse) GetFkiReconciliationIDOk() (*int32, bool)`

GetFkiReconciliationIDOk returns a tuple with the FkiReconciliationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiReconciliationID

`func (o *AttachmentResponse) SetFkiReconciliationID(v int32)`

SetFkiReconciliationID sets FkiReconciliationID field to given value.

### HasFkiReconciliationID

`func (o *AttachmentResponse) HasFkiReconciliationID() bool`

HasFkiReconciliationID returns a boolean if a field has been set.

### GetFkiEzsigndocumentIDReference

`func (o *AttachmentResponse) GetFkiEzsigndocumentIDReference() int32`

GetFkiEzsigndocumentIDReference returns the FkiEzsigndocumentIDReference field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDReferenceOk

`func (o *AttachmentResponse) GetFkiEzsigndocumentIDReferenceOk() (*int32, bool)`

GetFkiEzsigndocumentIDReferenceOk returns a tuple with the FkiEzsigndocumentIDReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentIDReference

`func (o *AttachmentResponse) SetFkiEzsigndocumentIDReference(v int32)`

SetFkiEzsigndocumentIDReference sets FkiEzsigndocumentIDReference field to given value.

### HasFkiEzsigndocumentIDReference

`func (o *AttachmentResponse) HasFkiEzsigndocumentIDReference() bool`

HasFkiEzsigndocumentIDReference returns a boolean if a field has been set.

### GetEAttachmentDocumenttype

`func (o *AttachmentResponse) GetEAttachmentDocumenttype() FieldEAttachmentDocumenttype`

GetEAttachmentDocumenttype returns the EAttachmentDocumenttype field if non-nil, zero value otherwise.

### GetEAttachmentDocumenttypeOk

`func (o *AttachmentResponse) GetEAttachmentDocumenttypeOk() (*FieldEAttachmentDocumenttype, bool)`

GetEAttachmentDocumenttypeOk returns a tuple with the EAttachmentDocumenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentDocumenttype

`func (o *AttachmentResponse) SetEAttachmentDocumenttype(v FieldEAttachmentDocumenttype)`

SetEAttachmentDocumenttype sets EAttachmentDocumenttype field to given value.


### GetSAttachmentName

`func (o *AttachmentResponse) GetSAttachmentName() string`

GetSAttachmentName returns the SAttachmentName field if non-nil, zero value otherwise.

### GetSAttachmentNameOk

`func (o *AttachmentResponse) GetSAttachmentNameOk() (*string, bool)`

GetSAttachmentNameOk returns a tuple with the SAttachmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentName

`func (o *AttachmentResponse) SetSAttachmentName(v string)`

SetSAttachmentName sets SAttachmentName field to given value.


### GetEAttachmentPrivacy

`func (o *AttachmentResponse) GetEAttachmentPrivacy() FieldEAttachmentPrivacy`

GetEAttachmentPrivacy returns the EAttachmentPrivacy field if non-nil, zero value otherwise.

### GetEAttachmentPrivacyOk

`func (o *AttachmentResponse) GetEAttachmentPrivacyOk() (*FieldEAttachmentPrivacy, bool)`

GetEAttachmentPrivacyOk returns a tuple with the EAttachmentPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentPrivacy

`func (o *AttachmentResponse) SetEAttachmentPrivacy(v FieldEAttachmentPrivacy)`

SetEAttachmentPrivacy sets EAttachmentPrivacy field to given value.


### GetFkiUserIDSpecific

`func (o *AttachmentResponse) GetFkiUserIDSpecific() int32`

GetFkiUserIDSpecific returns the FkiUserIDSpecific field if non-nil, zero value otherwise.

### GetFkiUserIDSpecificOk

`func (o *AttachmentResponse) GetFkiUserIDSpecificOk() (*int32, bool)`

GetFkiUserIDSpecificOk returns a tuple with the FkiUserIDSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDSpecific

`func (o *AttachmentResponse) SetFkiUserIDSpecific(v int32)`

SetFkiUserIDSpecific sets FkiUserIDSpecific field to given value.

### HasFkiUserIDSpecific

`func (o *AttachmentResponse) HasFkiUserIDSpecific() bool`

HasFkiUserIDSpecific returns a boolean if a field has been set.

### GetEAttachmentType

`func (o *AttachmentResponse) GetEAttachmentType() FieldEAttachmentType`

GetEAttachmentType returns the EAttachmentType field if non-nil, zero value otherwise.

### GetEAttachmentTypeOk

`func (o *AttachmentResponse) GetEAttachmentTypeOk() (*FieldEAttachmentType, bool)`

GetEAttachmentTypeOk returns a tuple with the EAttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentType

`func (o *AttachmentResponse) SetEAttachmentType(v FieldEAttachmentType)`

SetEAttachmentType sets EAttachmentType field to given value.


### GetIAttachmentSize

`func (o *AttachmentResponse) GetIAttachmentSize() int32`

GetIAttachmentSize returns the IAttachmentSize field if non-nil, zero value otherwise.

### GetIAttachmentSizeOk

`func (o *AttachmentResponse) GetIAttachmentSizeOk() (*int32, bool)`

GetIAttachmentSizeOk returns a tuple with the IAttachmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAttachmentSize

`func (o *AttachmentResponse) SetIAttachmentSize(v int32)`

SetIAttachmentSize sets IAttachmentSize field to given value.


### GetIAttachmentEDMmoduleflag

`func (o *AttachmentResponse) GetIAttachmentEDMmoduleflag() int32`

GetIAttachmentEDMmoduleflag returns the IAttachmentEDMmoduleflag field if non-nil, zero value otherwise.

### GetIAttachmentEDMmoduleflagOk

`func (o *AttachmentResponse) GetIAttachmentEDMmoduleflagOk() (*int32, bool)`

GetIAttachmentEDMmoduleflagOk returns a tuple with the IAttachmentEDMmoduleflag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAttachmentEDMmoduleflag

`func (o *AttachmentResponse) SetIAttachmentEDMmoduleflag(v int32)`

SetIAttachmentEDMmoduleflag sets IAttachmentEDMmoduleflag field to given value.

### HasIAttachmentEDMmoduleflag

`func (o *AttachmentResponse) HasIAttachmentEDMmoduleflag() bool`

HasIAttachmentEDMmoduleflag returns a boolean if a field has been set.

### GetSAttachmentMD5

`func (o *AttachmentResponse) GetSAttachmentMD5() string`

GetSAttachmentMD5 returns the SAttachmentMD5 field if non-nil, zero value otherwise.

### GetSAttachmentMD5Ok

`func (o *AttachmentResponse) GetSAttachmentMD5Ok() (*string, bool)`

GetSAttachmentMD5Ok returns a tuple with the SAttachmentMD5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentMD5

`func (o *AttachmentResponse) SetSAttachmentMD5(v string)`

SetSAttachmentMD5 sets SAttachmentMD5 field to given value.


### GetBAttachmentDeleted

`func (o *AttachmentResponse) GetBAttachmentDeleted() bool`

GetBAttachmentDeleted returns the BAttachmentDeleted field if non-nil, zero value otherwise.

### GetBAttachmentDeletedOk

`func (o *AttachmentResponse) GetBAttachmentDeletedOk() (*bool, bool)`

GetBAttachmentDeletedOk returns a tuple with the BAttachmentDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAttachmentDeleted

`func (o *AttachmentResponse) SetBAttachmentDeleted(v bool)`

SetBAttachmentDeleted sets BAttachmentDeleted field to given value.


### GetBAttachmentValid

`func (o *AttachmentResponse) GetBAttachmentValid() bool`

GetBAttachmentValid returns the BAttachmentValid field if non-nil, zero value otherwise.

### GetBAttachmentValidOk

`func (o *AttachmentResponse) GetBAttachmentValidOk() (*bool, bool)`

GetBAttachmentValidOk returns a tuple with the BAttachmentValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAttachmentValid

`func (o *AttachmentResponse) SetBAttachmentValid(v bool)`

SetBAttachmentValid sets BAttachmentValid field to given value.


### GetEAttachmentVerified

`func (o *AttachmentResponse) GetEAttachmentVerified() FieldEAttachmentVerified`

GetEAttachmentVerified returns the EAttachmentVerified field if non-nil, zero value otherwise.

### GetEAttachmentVerifiedOk

`func (o *AttachmentResponse) GetEAttachmentVerifiedOk() (*FieldEAttachmentVerified, bool)`

GetEAttachmentVerifiedOk returns a tuple with the EAttachmentVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentVerified

`func (o *AttachmentResponse) SetEAttachmentVerified(v FieldEAttachmentVerified)`

SetEAttachmentVerified sets EAttachmentVerified field to given value.


### GetTAttachmentRejectioncomment

`func (o *AttachmentResponse) GetTAttachmentRejectioncomment() string`

GetTAttachmentRejectioncomment returns the TAttachmentRejectioncomment field if non-nil, zero value otherwise.

### GetTAttachmentRejectioncommentOk

`func (o *AttachmentResponse) GetTAttachmentRejectioncommentOk() (*string, bool)`

GetTAttachmentRejectioncommentOk returns a tuple with the TAttachmentRejectioncomment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTAttachmentRejectioncomment

`func (o *AttachmentResponse) SetTAttachmentRejectioncomment(v string)`

SetTAttachmentRejectioncomment sets TAttachmentRejectioncomment field to given value.

### HasTAttachmentRejectioncomment

`func (o *AttachmentResponse) HasTAttachmentRejectioncomment() bool`

HasTAttachmentRejectioncomment returns a boolean if a field has been set.

### GetFkiUserIDOwner

`func (o *AttachmentResponse) GetFkiUserIDOwner() int32`

GetFkiUserIDOwner returns the FkiUserIDOwner field if non-nil, zero value otherwise.

### GetFkiUserIDOwnerOk

`func (o *AttachmentResponse) GetFkiUserIDOwnerOk() (*int32, bool)`

GetFkiUserIDOwnerOk returns a tuple with the FkiUserIDOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDOwner

`func (o *AttachmentResponse) SetFkiUserIDOwner(v int32)`

SetFkiUserIDOwner sets FkiUserIDOwner field to given value.

### HasFkiUserIDOwner

`func (o *AttachmentResponse) HasFkiUserIDOwner() bool`

HasFkiUserIDOwner returns a boolean if a field has been set.

### GetObjAudit

`func (o *AttachmentResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *AttachmentResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *AttachmentResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *AttachmentResponse) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


