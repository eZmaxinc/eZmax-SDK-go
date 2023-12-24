# CustomAttachmentResponse

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
**ObjAttachmentProof** | Pointer to [**AttachmentResponseCompound**](AttachmentResponseCompound.md) |  | [optional] 
**ObjAttachmentProofdocument** | Pointer to [**AttachmentResponseCompound**](AttachmentResponseCompound.md) |  | [optional] 
**AObjAttachmentAttachment** | Pointer to [**[]AttachmentResponseCompound**](AttachmentResponseCompound.md) |  | [optional] 
**AObjAttachmentVersion** | Pointer to [**[]AttachmentResponseCompound**](AttachmentResponseCompound.md) |  | [optional] 

## Methods

### NewCustomAttachmentResponse

`func NewCustomAttachmentResponse(pkiAttachmentID int32, eAttachmentDocumenttype FieldEAttachmentDocumenttype, sAttachmentName string, eAttachmentPrivacy FieldEAttachmentPrivacy, eAttachmentType FieldEAttachmentType, iAttachmentSize int32, sAttachmentMD5 string, bAttachmentDeleted bool, bAttachmentValid bool, eAttachmentVerified FieldEAttachmentVerified, ) *CustomAttachmentResponse`

NewCustomAttachmentResponse instantiates a new CustomAttachmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttachmentResponseWithDefaults

`func NewCustomAttachmentResponseWithDefaults() *CustomAttachmentResponse`

NewCustomAttachmentResponseWithDefaults instantiates a new CustomAttachmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiAttachmentID

`func (o *CustomAttachmentResponse) GetPkiAttachmentID() int32`

GetPkiAttachmentID returns the PkiAttachmentID field if non-nil, zero value otherwise.

### GetPkiAttachmentIDOk

`func (o *CustomAttachmentResponse) GetPkiAttachmentIDOk() (*int32, bool)`

GetPkiAttachmentIDOk returns a tuple with the PkiAttachmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiAttachmentID

`func (o *CustomAttachmentResponse) SetPkiAttachmentID(v int32)`

SetPkiAttachmentID sets PkiAttachmentID field to given value.


### GetFkiComputerID

`func (o *CustomAttachmentResponse) GetFkiComputerID() int32`

GetFkiComputerID returns the FkiComputerID field if non-nil, zero value otherwise.

### GetFkiComputerIDOk

`func (o *CustomAttachmentResponse) GetFkiComputerIDOk() (*int32, bool)`

GetFkiComputerIDOk returns a tuple with the FkiComputerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiComputerID

`func (o *CustomAttachmentResponse) SetFkiComputerID(v int32)`

SetFkiComputerID sets FkiComputerID field to given value.

### HasFkiComputerID

`func (o *CustomAttachmentResponse) HasFkiComputerID() bool`

HasFkiComputerID returns a boolean if a field has been set.

### GetFkiAdjustmentID

`func (o *CustomAttachmentResponse) GetFkiAdjustmentID() int32`

GetFkiAdjustmentID returns the FkiAdjustmentID field if non-nil, zero value otherwise.

### GetFkiAdjustmentIDOk

`func (o *CustomAttachmentResponse) GetFkiAdjustmentIDOk() (*int32, bool)`

GetFkiAdjustmentIDOk returns a tuple with the FkiAdjustmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAdjustmentID

`func (o *CustomAttachmentResponse) SetFkiAdjustmentID(v int32)`

SetFkiAdjustmentID sets FkiAdjustmentID field to given value.

### HasFkiAdjustmentID

`func (o *CustomAttachmentResponse) HasFkiAdjustmentID() bool`

HasFkiAdjustmentID returns a boolean if a field has been set.

### GetFkiAgentID

`func (o *CustomAttachmentResponse) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *CustomAttachmentResponse) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *CustomAttachmentResponse) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *CustomAttachmentResponse) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBankaccountID

`func (o *CustomAttachmentResponse) GetFkiBankaccountID() int32`

GetFkiBankaccountID returns the FkiBankaccountID field if non-nil, zero value otherwise.

### GetFkiBankaccountIDOk

`func (o *CustomAttachmentResponse) GetFkiBankaccountIDOk() (*int32, bool)`

GetFkiBankaccountIDOk returns a tuple with the FkiBankaccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBankaccountID

`func (o *CustomAttachmentResponse) SetFkiBankaccountID(v int32)`

SetFkiBankaccountID sets FkiBankaccountID field to given value.

### HasFkiBankaccountID

`func (o *CustomAttachmentResponse) HasFkiBankaccountID() bool`

HasFkiBankaccountID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *CustomAttachmentResponse) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *CustomAttachmentResponse) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *CustomAttachmentResponse) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *CustomAttachmentResponse) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetFkiCommissionadvanceID

`func (o *CustomAttachmentResponse) GetFkiCommissionadvanceID() int32`

GetFkiCommissionadvanceID returns the FkiCommissionadvanceID field if non-nil, zero value otherwise.

### GetFkiCommissionadvanceIDOk

`func (o *CustomAttachmentResponse) GetFkiCommissionadvanceIDOk() (*int32, bool)`

GetFkiCommissionadvanceIDOk returns a tuple with the FkiCommissionadvanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommissionadvanceID

`func (o *CustomAttachmentResponse) SetFkiCommissionadvanceID(v int32)`

SetFkiCommissionadvanceID sets FkiCommissionadvanceID field to given value.

### HasFkiCommissionadvanceID

`func (o *CustomAttachmentResponse) HasFkiCommissionadvanceID() bool`

HasFkiCommissionadvanceID returns a boolean if a field has been set.

### GetFkiCommunicationID

`func (o *CustomAttachmentResponse) GetFkiCommunicationID() int32`

GetFkiCommunicationID returns the FkiCommunicationID field if non-nil, zero value otherwise.

### GetFkiCommunicationIDOk

`func (o *CustomAttachmentResponse) GetFkiCommunicationIDOk() (*int32, bool)`

GetFkiCommunicationIDOk returns a tuple with the FkiCommunicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCommunicationID

`func (o *CustomAttachmentResponse) SetFkiCommunicationID(v int32)`

SetFkiCommunicationID sets FkiCommunicationID field to given value.

### HasFkiCommunicationID

`func (o *CustomAttachmentResponse) HasFkiCommunicationID() bool`

HasFkiCommunicationID returns a boolean if a field has been set.

### GetFkiCustomerID

`func (o *CustomAttachmentResponse) GetFkiCustomerID() int32`

GetFkiCustomerID returns the FkiCustomerID field if non-nil, zero value otherwise.

### GetFkiCustomerIDOk

`func (o *CustomAttachmentResponse) GetFkiCustomerIDOk() (*int32, bool)`

GetFkiCustomerIDOk returns a tuple with the FkiCustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomerID

`func (o *CustomAttachmentResponse) SetFkiCustomerID(v int32)`

SetFkiCustomerID sets FkiCustomerID field to given value.

### HasFkiCustomerID

`func (o *CustomAttachmentResponse) HasFkiCustomerID() bool`

HasFkiCustomerID returns a boolean if a field has been set.

### GetFkiCustomertemplateID

`func (o *CustomAttachmentResponse) GetFkiCustomertemplateID() int32`

GetFkiCustomertemplateID returns the FkiCustomertemplateID field if non-nil, zero value otherwise.

### GetFkiCustomertemplateIDOk

`func (o *CustomAttachmentResponse) GetFkiCustomertemplateIDOk() (*int32, bool)`

GetFkiCustomertemplateIDOk returns a tuple with the FkiCustomertemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiCustomertemplateID

`func (o *CustomAttachmentResponse) SetFkiCustomertemplateID(v int32)`

SetFkiCustomertemplateID sets FkiCustomertemplateID field to given value.

### HasFkiCustomertemplateID

`func (o *CustomAttachmentResponse) HasFkiCustomertemplateID() bool`

HasFkiCustomertemplateID returns a boolean if a field has been set.

### GetFkiDepositID

`func (o *CustomAttachmentResponse) GetFkiDepositID() int32`

GetFkiDepositID returns the FkiDepositID field if non-nil, zero value otherwise.

### GetFkiDepositIDOk

`func (o *CustomAttachmentResponse) GetFkiDepositIDOk() (*int32, bool)`

GetFkiDepositIDOk returns a tuple with the FkiDepositID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDepositID

`func (o *CustomAttachmentResponse) SetFkiDepositID(v int32)`

SetFkiDepositID sets FkiDepositID field to given value.

### HasFkiDepositID

`func (o *CustomAttachmentResponse) HasFkiDepositID() bool`

HasFkiDepositID returns a boolean if a field has been set.

### GetFkiDeposittransitchequeID

`func (o *CustomAttachmentResponse) GetFkiDeposittransitchequeID() int32`

GetFkiDeposittransitchequeID returns the FkiDeposittransitchequeID field if non-nil, zero value otherwise.

### GetFkiDeposittransitchequeIDOk

`func (o *CustomAttachmentResponse) GetFkiDeposittransitchequeIDOk() (*int32, bool)`

GetFkiDeposittransitchequeIDOk returns a tuple with the FkiDeposittransitchequeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDeposittransitchequeID

`func (o *CustomAttachmentResponse) SetFkiDeposittransitchequeID(v int32)`

SetFkiDeposittransitchequeID sets FkiDeposittransitchequeID field to given value.

### HasFkiDeposittransitchequeID

`func (o *CustomAttachmentResponse) HasFkiDeposittransitchequeID() bool`

HasFkiDeposittransitchequeID returns a boolean if a field has been set.

### GetFkiElectronicfundstransferID

`func (o *CustomAttachmentResponse) GetFkiElectronicfundstransferID() int32`

GetFkiElectronicfundstransferID returns the FkiElectronicfundstransferID field if non-nil, zero value otherwise.

### GetFkiElectronicfundstransferIDOk

`func (o *CustomAttachmentResponse) GetFkiElectronicfundstransferIDOk() (*int32, bool)`

GetFkiElectronicfundstransferIDOk returns a tuple with the FkiElectronicfundstransferID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiElectronicfundstransferID

`func (o *CustomAttachmentResponse) SetFkiElectronicfundstransferID(v int32)`

SetFkiElectronicfundstransferID sets FkiElectronicfundstransferID field to given value.

### HasFkiElectronicfundstransferID

`func (o *CustomAttachmentResponse) HasFkiElectronicfundstransferID() bool`

HasFkiElectronicfundstransferID returns a boolean if a field has been set.

### GetFkiEmployeeID

`func (o *CustomAttachmentResponse) GetFkiEmployeeID() int32`

GetFkiEmployeeID returns the FkiEmployeeID field if non-nil, zero value otherwise.

### GetFkiEmployeeIDOk

`func (o *CustomAttachmentResponse) GetFkiEmployeeIDOk() (*int32, bool)`

GetFkiEmployeeIDOk returns a tuple with the FkiEmployeeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEmployeeID

`func (o *CustomAttachmentResponse) SetFkiEmployeeID(v int32)`

SetFkiEmployeeID sets FkiEmployeeID field to given value.

### HasFkiEmployeeID

`func (o *CustomAttachmentResponse) HasFkiEmployeeID() bool`

HasFkiEmployeeID returns a boolean if a field has been set.

### GetFkiExternalbrokerID

`func (o *CustomAttachmentResponse) GetFkiExternalbrokerID() int32`

GetFkiExternalbrokerID returns the FkiExternalbrokerID field if non-nil, zero value otherwise.

### GetFkiExternalbrokerIDOk

`func (o *CustomAttachmentResponse) GetFkiExternalbrokerIDOk() (*int32, bool)`

GetFkiExternalbrokerIDOk returns a tuple with the FkiExternalbrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiExternalbrokerID

`func (o *CustomAttachmentResponse) SetFkiExternalbrokerID(v int32)`

SetFkiExternalbrokerID sets FkiExternalbrokerID field to given value.

### HasFkiExternalbrokerID

`func (o *CustomAttachmentResponse) HasFkiExternalbrokerID() bool`

HasFkiExternalbrokerID returns a boolean if a field has been set.

### GetFkiEzcomadvanceserverID

`func (o *CustomAttachmentResponse) GetFkiEzcomadvanceserverID() int32`

GetFkiEzcomadvanceserverID returns the FkiEzcomadvanceserverID field if non-nil, zero value otherwise.

### GetFkiEzcomadvanceserverIDOk

`func (o *CustomAttachmentResponse) GetFkiEzcomadvanceserverIDOk() (*int32, bool)`

GetFkiEzcomadvanceserverIDOk returns a tuple with the FkiEzcomadvanceserverID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomadvanceserverID

`func (o *CustomAttachmentResponse) SetFkiEzcomadvanceserverID(v int32)`

SetFkiEzcomadvanceserverID sets FkiEzcomadvanceserverID field to given value.

### HasFkiEzcomadvanceserverID

`func (o *CustomAttachmentResponse) HasFkiEzcomadvanceserverID() bool`

HasFkiEzcomadvanceserverID returns a boolean if a field has been set.

### GetFkiEzcomcompanyID

`func (o *CustomAttachmentResponse) GetFkiEzcomcompanyID() int32`

GetFkiEzcomcompanyID returns the FkiEzcomcompanyID field if non-nil, zero value otherwise.

### GetFkiEzcomcompanyIDOk

`func (o *CustomAttachmentResponse) GetFkiEzcomcompanyIDOk() (*int32, bool)`

GetFkiEzcomcompanyIDOk returns a tuple with the FkiEzcomcompanyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzcomcompanyID

`func (o *CustomAttachmentResponse) SetFkiEzcomcompanyID(v int32)`

SetFkiEzcomcompanyID sets FkiEzcomcompanyID field to given value.

### HasFkiEzcomcompanyID

`func (o *CustomAttachmentResponse) HasFkiEzcomcompanyID() bool`

HasFkiEzcomcompanyID returns a boolean if a field has been set.

### GetFkiEzsigndocumentID

`func (o *CustomAttachmentResponse) GetFkiEzsigndocumentID() int32`

GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDOk

`func (o *CustomAttachmentResponse) GetFkiEzsigndocumentIDOk() (*int32, bool)`

GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentID

`func (o *CustomAttachmentResponse) SetFkiEzsigndocumentID(v int32)`

SetFkiEzsigndocumentID sets FkiEzsigndocumentID field to given value.

### HasFkiEzsigndocumentID

`func (o *CustomAttachmentResponse) HasFkiEzsigndocumentID() bool`

HasFkiEzsigndocumentID returns a boolean if a field has been set.

### GetFkiGhacqcontractID

`func (o *CustomAttachmentResponse) GetFkiGhacqcontractID() int32`

GetFkiGhacqcontractID returns the FkiGhacqcontractID field if non-nil, zero value otherwise.

### GetFkiGhacqcontractIDOk

`func (o *CustomAttachmentResponse) GetFkiGhacqcontractIDOk() (*int32, bool)`

GetFkiGhacqcontractIDOk returns a tuple with the FkiGhacqcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiGhacqcontractID

`func (o *CustomAttachmentResponse) SetFkiGhacqcontractID(v int32)`

SetFkiGhacqcontractID sets FkiGhacqcontractID field to given value.

### HasFkiGhacqcontractID

`func (o *CustomAttachmentResponse) HasFkiGhacqcontractID() bool`

HasFkiGhacqcontractID returns a boolean if a field has been set.

### GetFkiInscriptionID

`func (o *CustomAttachmentResponse) GetFkiInscriptionID() int32`

GetFkiInscriptionID returns the FkiInscriptionID field if non-nil, zero value otherwise.

### GetFkiInscriptionIDOk

`func (o *CustomAttachmentResponse) GetFkiInscriptionIDOk() (*int32, bool)`

GetFkiInscriptionIDOk returns a tuple with the FkiInscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionID

`func (o *CustomAttachmentResponse) SetFkiInscriptionID(v int32)`

SetFkiInscriptionID sets FkiInscriptionID field to given value.

### HasFkiInscriptionID

`func (o *CustomAttachmentResponse) HasFkiInscriptionID() bool`

HasFkiInscriptionID returns a boolean if a field has been set.

### GetFkiInscriptiontempID

`func (o *CustomAttachmentResponse) GetFkiInscriptiontempID() int32`

GetFkiInscriptiontempID returns the FkiInscriptiontempID field if non-nil, zero value otherwise.

### GetFkiInscriptiontempIDOk

`func (o *CustomAttachmentResponse) GetFkiInscriptiontempIDOk() (*int32, bool)`

GetFkiInscriptiontempIDOk returns a tuple with the FkiInscriptiontempID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptiontempID

`func (o *CustomAttachmentResponse) SetFkiInscriptiontempID(v int32)`

SetFkiInscriptiontempID sets FkiInscriptiontempID field to given value.

### HasFkiInscriptiontempID

`func (o *CustomAttachmentResponse) HasFkiInscriptiontempID() bool`

HasFkiInscriptiontempID returns a boolean if a field has been set.

### GetFkiInscriptionnotauthenticatedID

`func (o *CustomAttachmentResponse) GetFkiInscriptionnotauthenticatedID() int32`

GetFkiInscriptionnotauthenticatedID returns the FkiInscriptionnotauthenticatedID field if non-nil, zero value otherwise.

### GetFkiInscriptionnotauthenticatedIDOk

`func (o *CustomAttachmentResponse) GetFkiInscriptionnotauthenticatedIDOk() (*int32, bool)`

GetFkiInscriptionnotauthenticatedIDOk returns a tuple with the FkiInscriptionnotauthenticatedID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionnotauthenticatedID

`func (o *CustomAttachmentResponse) SetFkiInscriptionnotauthenticatedID(v int32)`

SetFkiInscriptionnotauthenticatedID sets FkiInscriptionnotauthenticatedID field to given value.

### HasFkiInscriptionnotauthenticatedID

`func (o *CustomAttachmentResponse) HasFkiInscriptionnotauthenticatedID() bool`

HasFkiInscriptionnotauthenticatedID returns a boolean if a field has been set.

### GetFkiInvoiceID

`func (o *CustomAttachmentResponse) GetFkiInvoiceID() int32`

GetFkiInvoiceID returns the FkiInvoiceID field if non-nil, zero value otherwise.

### GetFkiInvoiceIDOk

`func (o *CustomAttachmentResponse) GetFkiInvoiceIDOk() (*int32, bool)`

GetFkiInvoiceIDOk returns a tuple with the FkiInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInvoiceID

`func (o *CustomAttachmentResponse) SetFkiInvoiceID(v int32)`

SetFkiInvoiceID sets FkiInvoiceID field to given value.

### HasFkiInvoiceID

`func (o *CustomAttachmentResponse) HasFkiInvoiceID() bool`

HasFkiInvoiceID returns a boolean if a field has been set.

### GetFkiBuyercontractID

`func (o *CustomAttachmentResponse) GetFkiBuyercontractID() int32`

GetFkiBuyercontractID returns the FkiBuyercontractID field if non-nil, zero value otherwise.

### GetFkiBuyercontractIDOk

`func (o *CustomAttachmentResponse) GetFkiBuyercontractIDOk() (*int32, bool)`

GetFkiBuyercontractIDOk returns a tuple with the FkiBuyercontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBuyercontractID

`func (o *CustomAttachmentResponse) SetFkiBuyercontractID(v int32)`

SetFkiBuyercontractID sets FkiBuyercontractID field to given value.

### HasFkiBuyercontractID

`func (o *CustomAttachmentResponse) HasFkiBuyercontractID() bool`

HasFkiBuyercontractID returns a boolean if a field has been set.

### GetFkiFranchisebrokerID

`func (o *CustomAttachmentResponse) GetFkiFranchisebrokerID() int32`

GetFkiFranchisebrokerID returns the FkiFranchisebrokerID field if non-nil, zero value otherwise.

### GetFkiFranchisebrokerIDOk

`func (o *CustomAttachmentResponse) GetFkiFranchisebrokerIDOk() (*int32, bool)`

GetFkiFranchisebrokerIDOk returns a tuple with the FkiFranchisebrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisebrokerID

`func (o *CustomAttachmentResponse) SetFkiFranchisebrokerID(v int32)`

SetFkiFranchisebrokerID sets FkiFranchisebrokerID field to given value.

### HasFkiFranchisebrokerID

`func (o *CustomAttachmentResponse) HasFkiFranchisebrokerID() bool`

HasFkiFranchisebrokerID returns a boolean if a field has been set.

### GetFkiFranchiseagenceID

`func (o *CustomAttachmentResponse) GetFkiFranchiseagenceID() int32`

GetFkiFranchiseagenceID returns the FkiFranchiseagenceID field if non-nil, zero value otherwise.

### GetFkiFranchiseagenceIDOk

`func (o *CustomAttachmentResponse) GetFkiFranchiseagenceIDOk() (*int32, bool)`

GetFkiFranchiseagenceIDOk returns a tuple with the FkiFranchiseagenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseagenceID

`func (o *CustomAttachmentResponse) SetFkiFranchiseagenceID(v int32)`

SetFkiFranchiseagenceID sets FkiFranchiseagenceID field to given value.

### HasFkiFranchiseagenceID

`func (o *CustomAttachmentResponse) HasFkiFranchiseagenceID() bool`

HasFkiFranchiseagenceID returns a boolean if a field has been set.

### GetFkiFranchiseofficeID

`func (o *CustomAttachmentResponse) GetFkiFranchiseofficeID() int32`

GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field if non-nil, zero value otherwise.

### GetFkiFranchiseofficeIDOk

`func (o *CustomAttachmentResponse) GetFkiFranchiseofficeIDOk() (*int32, bool)`

GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseofficeID

`func (o *CustomAttachmentResponse) SetFkiFranchiseofficeID(v int32)`

SetFkiFranchiseofficeID sets FkiFranchiseofficeID field to given value.

### HasFkiFranchiseofficeID

`func (o *CustomAttachmentResponse) HasFkiFranchiseofficeID() bool`

HasFkiFranchiseofficeID returns a boolean if a field has been set.

### GetFkiFranchisefranchiseID

`func (o *CustomAttachmentResponse) GetFkiFranchisefranchiseID() int32`

GetFkiFranchisefranchiseID returns the FkiFranchisefranchiseID field if non-nil, zero value otherwise.

### GetFkiFranchisefranchiseIDOk

`func (o *CustomAttachmentResponse) GetFkiFranchisefranchiseIDOk() (*int32, bool)`

GetFkiFranchisefranchiseIDOk returns a tuple with the FkiFranchisefranchiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisefranchiseID

`func (o *CustomAttachmentResponse) SetFkiFranchisefranchiseID(v int32)`

SetFkiFranchisefranchiseID sets FkiFranchisefranchiseID field to given value.

### HasFkiFranchisefranchiseID

`func (o *CustomAttachmentResponse) HasFkiFranchisefranchiseID() bool`

HasFkiFranchisefranchiseID returns a boolean if a field has been set.

### GetFkiFranchisecomplaintID

`func (o *CustomAttachmentResponse) GetFkiFranchisecomplaintID() int32`

GetFkiFranchisecomplaintID returns the FkiFranchisecomplaintID field if non-nil, zero value otherwise.

### GetFkiFranchisecomplaintIDOk

`func (o *CustomAttachmentResponse) GetFkiFranchisecomplaintIDOk() (*int32, bool)`

GetFkiFranchisecomplaintIDOk returns a tuple with the FkiFranchisecomplaintID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisecomplaintID

`func (o *CustomAttachmentResponse) SetFkiFranchisecomplaintID(v int32)`

SetFkiFranchisecomplaintID sets FkiFranchisecomplaintID field to given value.

### HasFkiFranchisecomplaintID

`func (o *CustomAttachmentResponse) HasFkiFranchisecomplaintID() bool`

HasFkiFranchisecomplaintID returns a boolean if a field has been set.

### GetFkiLeadID

`func (o *CustomAttachmentResponse) GetFkiLeadID() int32`

GetFkiLeadID returns the FkiLeadID field if non-nil, zero value otherwise.

### GetFkiLeadIDOk

`func (o *CustomAttachmentResponse) GetFkiLeadIDOk() (*int32, bool)`

GetFkiLeadIDOk returns a tuple with the FkiLeadID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiLeadID

`func (o *CustomAttachmentResponse) SetFkiLeadID(v int32)`

SetFkiLeadID sets FkiLeadID field to given value.

### HasFkiLeadID

`func (o *CustomAttachmentResponse) HasFkiLeadID() bool`

HasFkiLeadID returns a boolean if a field has been set.

### GetFkiMarketingprogramID

`func (o *CustomAttachmentResponse) GetFkiMarketingprogramID() int32`

GetFkiMarketingprogramID returns the FkiMarketingprogramID field if non-nil, zero value otherwise.

### GetFkiMarketingprogramIDOk

`func (o *CustomAttachmentResponse) GetFkiMarketingprogramIDOk() (*int32, bool)`

GetFkiMarketingprogramIDOk returns a tuple with the FkiMarketingprogramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMarketingprogramID

`func (o *CustomAttachmentResponse) SetFkiMarketingprogramID(v int32)`

SetFkiMarketingprogramID sets FkiMarketingprogramID field to given value.

### HasFkiMarketingprogramID

`func (o *CustomAttachmentResponse) HasFkiMarketingprogramID() bool`

HasFkiMarketingprogramID returns a boolean if a field has been set.

### GetFkiMarketingfollowID

`func (o *CustomAttachmentResponse) GetFkiMarketingfollowID() int32`

GetFkiMarketingfollowID returns the FkiMarketingfollowID field if non-nil, zero value otherwise.

### GetFkiMarketingfollowIDOk

`func (o *CustomAttachmentResponse) GetFkiMarketingfollowIDOk() (*int32, bool)`

GetFkiMarketingfollowIDOk returns a tuple with the FkiMarketingfollowID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiMarketingfollowID

`func (o *CustomAttachmentResponse) SetFkiMarketingfollowID(v int32)`

SetFkiMarketingfollowID sets FkiMarketingfollowID field to given value.

### HasFkiMarketingfollowID

`func (o *CustomAttachmentResponse) HasFkiMarketingfollowID() bool`

HasFkiMarketingfollowID returns a boolean if a field has been set.

### GetFkiNotaryID

`func (o *CustomAttachmentResponse) GetFkiNotaryID() int32`

GetFkiNotaryID returns the FkiNotaryID field if non-nil, zero value otherwise.

### GetFkiNotaryIDOk

`func (o *CustomAttachmentResponse) GetFkiNotaryIDOk() (*int32, bool)`

GetFkiNotaryIDOk returns a tuple with the FkiNotaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiNotaryID

`func (o *CustomAttachmentResponse) SetFkiNotaryID(v int32)`

SetFkiNotaryID sets FkiNotaryID field to given value.

### HasFkiNotaryID

`func (o *CustomAttachmentResponse) HasFkiNotaryID() bool`

HasFkiNotaryID returns a boolean if a field has been set.

### GetFkiOfficetaxreportID

`func (o *CustomAttachmentResponse) GetFkiOfficetaxreportID() int32`

GetFkiOfficetaxreportID returns the FkiOfficetaxreportID field if non-nil, zero value otherwise.

### GetFkiOfficetaxreportIDOk

`func (o *CustomAttachmentResponse) GetFkiOfficetaxreportIDOk() (*int32, bool)`

GetFkiOfficetaxreportIDOk returns a tuple with the FkiOfficetaxreportID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOfficetaxreportID

`func (o *CustomAttachmentResponse) SetFkiOfficetaxreportID(v int32)`

SetFkiOfficetaxreportID sets FkiOfficetaxreportID field to given value.

### HasFkiOfficetaxreportID

`func (o *CustomAttachmentResponse) HasFkiOfficetaxreportID() bool`

HasFkiOfficetaxreportID returns a boolean if a field has been set.

### GetFkiOtherincomeID

`func (o *CustomAttachmentResponse) GetFkiOtherincomeID() int32`

GetFkiOtherincomeID returns the FkiOtherincomeID field if non-nil, zero value otherwise.

### GetFkiOtherincomeIDOk

`func (o *CustomAttachmentResponse) GetFkiOtherincomeIDOk() (*int32, bool)`

GetFkiOtherincomeIDOk returns a tuple with the FkiOtherincomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiOtherincomeID

`func (o *CustomAttachmentResponse) SetFkiOtherincomeID(v int32)`

SetFkiOtherincomeID sets FkiOtherincomeID field to given value.

### HasFkiOtherincomeID

`func (o *CustomAttachmentResponse) HasFkiOtherincomeID() bool`

HasFkiOtherincomeID returns a boolean if a field has been set.

### GetFkiPaymentpreparationID

`func (o *CustomAttachmentResponse) GetFkiPaymentpreparationID() int32`

GetFkiPaymentpreparationID returns the FkiPaymentpreparationID field if non-nil, zero value otherwise.

### GetFkiPaymentpreparationIDOk

`func (o *CustomAttachmentResponse) GetFkiPaymentpreparationIDOk() (*int32, bool)`

GetFkiPaymentpreparationIDOk returns a tuple with the FkiPaymentpreparationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPaymentpreparationID

`func (o *CustomAttachmentResponse) SetFkiPaymentpreparationID(v int32)`

SetFkiPaymentpreparationID sets FkiPaymentpreparationID field to given value.

### HasFkiPaymentpreparationID

`func (o *CustomAttachmentResponse) HasFkiPaymentpreparationID() bool`

HasFkiPaymentpreparationID returns a boolean if a field has been set.

### GetFkiPurchaseID

`func (o *CustomAttachmentResponse) GetFkiPurchaseID() int32`

GetFkiPurchaseID returns the FkiPurchaseID field if non-nil, zero value otherwise.

### GetFkiPurchaseIDOk

`func (o *CustomAttachmentResponse) GetFkiPurchaseIDOk() (*int32, bool)`

GetFkiPurchaseIDOk returns a tuple with the FkiPurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPurchaseID

`func (o *CustomAttachmentResponse) SetFkiPurchaseID(v int32)`

SetFkiPurchaseID sets FkiPurchaseID field to given value.

### HasFkiPurchaseID

`func (o *CustomAttachmentResponse) HasFkiPurchaseID() bool`

HasFkiPurchaseID returns a boolean if a field has been set.

### GetFkiSalaryID

`func (o *CustomAttachmentResponse) GetFkiSalaryID() int32`

GetFkiSalaryID returns the FkiSalaryID field if non-nil, zero value otherwise.

### GetFkiSalaryIDOk

`func (o *CustomAttachmentResponse) GetFkiSalaryIDOk() (*int32, bool)`

GetFkiSalaryIDOk returns a tuple with the FkiSalaryID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSalaryID

`func (o *CustomAttachmentResponse) SetFkiSalaryID(v int32)`

SetFkiSalaryID sets FkiSalaryID field to given value.

### HasFkiSalaryID

`func (o *CustomAttachmentResponse) HasFkiSalaryID() bool`

HasFkiSalaryID returns a boolean if a field has been set.

### GetFkiSupplierID

`func (o *CustomAttachmentResponse) GetFkiSupplierID() int32`

GetFkiSupplierID returns the FkiSupplierID field if non-nil, zero value otherwise.

### GetFkiSupplierIDOk

`func (o *CustomAttachmentResponse) GetFkiSupplierIDOk() (*int32, bool)`

GetFkiSupplierIDOk returns a tuple with the FkiSupplierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiSupplierID

`func (o *CustomAttachmentResponse) SetFkiSupplierID(v int32)`

SetFkiSupplierID sets FkiSupplierID field to given value.

### HasFkiSupplierID

`func (o *CustomAttachmentResponse) HasFkiSupplierID() bool`

HasFkiSupplierID returns a boolean if a field has been set.

### GetFkiTranqcontractID

`func (o *CustomAttachmentResponse) GetFkiTranqcontractID() int32`

GetFkiTranqcontractID returns the FkiTranqcontractID field if non-nil, zero value otherwise.

### GetFkiTranqcontractIDOk

`func (o *CustomAttachmentResponse) GetFkiTranqcontractIDOk() (*int32, bool)`

GetFkiTranqcontractIDOk returns a tuple with the FkiTranqcontractID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTranqcontractID

`func (o *CustomAttachmentResponse) SetFkiTranqcontractID(v int32)`

SetFkiTranqcontractID sets FkiTranqcontractID field to given value.

### HasFkiTranqcontractID

`func (o *CustomAttachmentResponse) HasFkiTranqcontractID() bool`

HasFkiTranqcontractID returns a boolean if a field has been set.

### GetFkiTemplateID

`func (o *CustomAttachmentResponse) GetFkiTemplateID() int32`

GetFkiTemplateID returns the FkiTemplateID field if non-nil, zero value otherwise.

### GetFkiTemplateIDOk

`func (o *CustomAttachmentResponse) GetFkiTemplateIDOk() (*int32, bool)`

GetFkiTemplateIDOk returns a tuple with the FkiTemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiTemplateID

`func (o *CustomAttachmentResponse) SetFkiTemplateID(v int32)`

SetFkiTemplateID sets FkiTemplateID field to given value.

### HasFkiTemplateID

`func (o *CustomAttachmentResponse) HasFkiTemplateID() bool`

HasFkiTemplateID returns a boolean if a field has been set.

### GetFkiInscriptionchecklistID

`func (o *CustomAttachmentResponse) GetFkiInscriptionchecklistID() int32`

GetFkiInscriptionchecklistID returns the FkiInscriptionchecklistID field if non-nil, zero value otherwise.

### GetFkiInscriptionchecklistIDOk

`func (o *CustomAttachmentResponse) GetFkiInscriptionchecklistIDOk() (*int32, bool)`

GetFkiInscriptionchecklistIDOk returns a tuple with the FkiInscriptionchecklistID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiInscriptionchecklistID

`func (o *CustomAttachmentResponse) SetFkiInscriptionchecklistID(v int32)`

SetFkiInscriptionchecklistID sets FkiInscriptionchecklistID field to given value.

### HasFkiInscriptionchecklistID

`func (o *CustomAttachmentResponse) HasFkiInscriptionchecklistID() bool`

HasFkiInscriptionchecklistID returns a boolean if a field has been set.

### GetFkiFolderID

`func (o *CustomAttachmentResponse) GetFkiFolderID() int32`

GetFkiFolderID returns the FkiFolderID field if non-nil, zero value otherwise.

### GetFkiFolderIDOk

`func (o *CustomAttachmentResponse) GetFkiFolderIDOk() (*int32, bool)`

GetFkiFolderIDOk returns a tuple with the FkiFolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFolderID

`func (o *CustomAttachmentResponse) SetFkiFolderID(v int32)`

SetFkiFolderID sets FkiFolderID field to given value.

### HasFkiFolderID

`func (o *CustomAttachmentResponse) HasFkiFolderID() bool`

HasFkiFolderID returns a boolean if a field has been set.

### GetFkiRejectedoffertopurchaseID

`func (o *CustomAttachmentResponse) GetFkiRejectedoffertopurchaseID() int32`

GetFkiRejectedoffertopurchaseID returns the FkiRejectedoffertopurchaseID field if non-nil, zero value otherwise.

### GetFkiRejectedoffertopurchaseIDOk

`func (o *CustomAttachmentResponse) GetFkiRejectedoffertopurchaseIDOk() (*int32, bool)`

GetFkiRejectedoffertopurchaseIDOk returns a tuple with the FkiRejectedoffertopurchaseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiRejectedoffertopurchaseID

`func (o *CustomAttachmentResponse) SetFkiRejectedoffertopurchaseID(v int32)`

SetFkiRejectedoffertopurchaseID sets FkiRejectedoffertopurchaseID field to given value.

### HasFkiRejectedoffertopurchaseID

`func (o *CustomAttachmentResponse) HasFkiRejectedoffertopurchaseID() bool`

HasFkiRejectedoffertopurchaseID returns a boolean if a field has been set.

### GetFkiDisclosureID

`func (o *CustomAttachmentResponse) GetFkiDisclosureID() int32`

GetFkiDisclosureID returns the FkiDisclosureID field if non-nil, zero value otherwise.

### GetFkiDisclosureIDOk

`func (o *CustomAttachmentResponse) GetFkiDisclosureIDOk() (*int32, bool)`

GetFkiDisclosureIDOk returns a tuple with the FkiDisclosureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiDisclosureID

`func (o *CustomAttachmentResponse) SetFkiDisclosureID(v int32)`

SetFkiDisclosureID sets FkiDisclosureID field to given value.

### HasFkiDisclosureID

`func (o *CustomAttachmentResponse) HasFkiDisclosureID() bool`

HasFkiDisclosureID returns a boolean if a field has been set.

### GetFkiReconciliationID

`func (o *CustomAttachmentResponse) GetFkiReconciliationID() int32`

GetFkiReconciliationID returns the FkiReconciliationID field if non-nil, zero value otherwise.

### GetFkiReconciliationIDOk

`func (o *CustomAttachmentResponse) GetFkiReconciliationIDOk() (*int32, bool)`

GetFkiReconciliationIDOk returns a tuple with the FkiReconciliationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiReconciliationID

`func (o *CustomAttachmentResponse) SetFkiReconciliationID(v int32)`

SetFkiReconciliationID sets FkiReconciliationID field to given value.

### HasFkiReconciliationID

`func (o *CustomAttachmentResponse) HasFkiReconciliationID() bool`

HasFkiReconciliationID returns a boolean if a field has been set.

### GetFkiEzsigndocumentIDReference

`func (o *CustomAttachmentResponse) GetFkiEzsigndocumentIDReference() int32`

GetFkiEzsigndocumentIDReference returns the FkiEzsigndocumentIDReference field if non-nil, zero value otherwise.

### GetFkiEzsigndocumentIDReferenceOk

`func (o *CustomAttachmentResponse) GetFkiEzsigndocumentIDReferenceOk() (*int32, bool)`

GetFkiEzsigndocumentIDReferenceOk returns a tuple with the FkiEzsigndocumentIDReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzsigndocumentIDReference

`func (o *CustomAttachmentResponse) SetFkiEzsigndocumentIDReference(v int32)`

SetFkiEzsigndocumentIDReference sets FkiEzsigndocumentIDReference field to given value.

### HasFkiEzsigndocumentIDReference

`func (o *CustomAttachmentResponse) HasFkiEzsigndocumentIDReference() bool`

HasFkiEzsigndocumentIDReference returns a boolean if a field has been set.

### GetEAttachmentDocumenttype

`func (o *CustomAttachmentResponse) GetEAttachmentDocumenttype() FieldEAttachmentDocumenttype`

GetEAttachmentDocumenttype returns the EAttachmentDocumenttype field if non-nil, zero value otherwise.

### GetEAttachmentDocumenttypeOk

`func (o *CustomAttachmentResponse) GetEAttachmentDocumenttypeOk() (*FieldEAttachmentDocumenttype, bool)`

GetEAttachmentDocumenttypeOk returns a tuple with the EAttachmentDocumenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentDocumenttype

`func (o *CustomAttachmentResponse) SetEAttachmentDocumenttype(v FieldEAttachmentDocumenttype)`

SetEAttachmentDocumenttype sets EAttachmentDocumenttype field to given value.


### GetSAttachmentName

`func (o *CustomAttachmentResponse) GetSAttachmentName() string`

GetSAttachmentName returns the SAttachmentName field if non-nil, zero value otherwise.

### GetSAttachmentNameOk

`func (o *CustomAttachmentResponse) GetSAttachmentNameOk() (*string, bool)`

GetSAttachmentNameOk returns a tuple with the SAttachmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentName

`func (o *CustomAttachmentResponse) SetSAttachmentName(v string)`

SetSAttachmentName sets SAttachmentName field to given value.


### GetEAttachmentPrivacy

`func (o *CustomAttachmentResponse) GetEAttachmentPrivacy() FieldEAttachmentPrivacy`

GetEAttachmentPrivacy returns the EAttachmentPrivacy field if non-nil, zero value otherwise.

### GetEAttachmentPrivacyOk

`func (o *CustomAttachmentResponse) GetEAttachmentPrivacyOk() (*FieldEAttachmentPrivacy, bool)`

GetEAttachmentPrivacyOk returns a tuple with the EAttachmentPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentPrivacy

`func (o *CustomAttachmentResponse) SetEAttachmentPrivacy(v FieldEAttachmentPrivacy)`

SetEAttachmentPrivacy sets EAttachmentPrivacy field to given value.


### GetFkiUserIDSpecific

`func (o *CustomAttachmentResponse) GetFkiUserIDSpecific() int32`

GetFkiUserIDSpecific returns the FkiUserIDSpecific field if non-nil, zero value otherwise.

### GetFkiUserIDSpecificOk

`func (o *CustomAttachmentResponse) GetFkiUserIDSpecificOk() (*int32, bool)`

GetFkiUserIDSpecificOk returns a tuple with the FkiUserIDSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDSpecific

`func (o *CustomAttachmentResponse) SetFkiUserIDSpecific(v int32)`

SetFkiUserIDSpecific sets FkiUserIDSpecific field to given value.

### HasFkiUserIDSpecific

`func (o *CustomAttachmentResponse) HasFkiUserIDSpecific() bool`

HasFkiUserIDSpecific returns a boolean if a field has been set.

### GetEAttachmentType

`func (o *CustomAttachmentResponse) GetEAttachmentType() FieldEAttachmentType`

GetEAttachmentType returns the EAttachmentType field if non-nil, zero value otherwise.

### GetEAttachmentTypeOk

`func (o *CustomAttachmentResponse) GetEAttachmentTypeOk() (*FieldEAttachmentType, bool)`

GetEAttachmentTypeOk returns a tuple with the EAttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentType

`func (o *CustomAttachmentResponse) SetEAttachmentType(v FieldEAttachmentType)`

SetEAttachmentType sets EAttachmentType field to given value.


### GetIAttachmentSize

`func (o *CustomAttachmentResponse) GetIAttachmentSize() int32`

GetIAttachmentSize returns the IAttachmentSize field if non-nil, zero value otherwise.

### GetIAttachmentSizeOk

`func (o *CustomAttachmentResponse) GetIAttachmentSizeOk() (*int32, bool)`

GetIAttachmentSizeOk returns a tuple with the IAttachmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAttachmentSize

`func (o *CustomAttachmentResponse) SetIAttachmentSize(v int32)`

SetIAttachmentSize sets IAttachmentSize field to given value.


### GetIAttachmentEDMmoduleflag

`func (o *CustomAttachmentResponse) GetIAttachmentEDMmoduleflag() int32`

GetIAttachmentEDMmoduleflag returns the IAttachmentEDMmoduleflag field if non-nil, zero value otherwise.

### GetIAttachmentEDMmoduleflagOk

`func (o *CustomAttachmentResponse) GetIAttachmentEDMmoduleflagOk() (*int32, bool)`

GetIAttachmentEDMmoduleflagOk returns a tuple with the IAttachmentEDMmoduleflag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAttachmentEDMmoduleflag

`func (o *CustomAttachmentResponse) SetIAttachmentEDMmoduleflag(v int32)`

SetIAttachmentEDMmoduleflag sets IAttachmentEDMmoduleflag field to given value.

### HasIAttachmentEDMmoduleflag

`func (o *CustomAttachmentResponse) HasIAttachmentEDMmoduleflag() bool`

HasIAttachmentEDMmoduleflag returns a boolean if a field has been set.

### GetSAttachmentMD5

`func (o *CustomAttachmentResponse) GetSAttachmentMD5() string`

GetSAttachmentMD5 returns the SAttachmentMD5 field if non-nil, zero value otherwise.

### GetSAttachmentMD5Ok

`func (o *CustomAttachmentResponse) GetSAttachmentMD5Ok() (*string, bool)`

GetSAttachmentMD5Ok returns a tuple with the SAttachmentMD5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAttachmentMD5

`func (o *CustomAttachmentResponse) SetSAttachmentMD5(v string)`

SetSAttachmentMD5 sets SAttachmentMD5 field to given value.


### GetBAttachmentDeleted

`func (o *CustomAttachmentResponse) GetBAttachmentDeleted() bool`

GetBAttachmentDeleted returns the BAttachmentDeleted field if non-nil, zero value otherwise.

### GetBAttachmentDeletedOk

`func (o *CustomAttachmentResponse) GetBAttachmentDeletedOk() (*bool, bool)`

GetBAttachmentDeletedOk returns a tuple with the BAttachmentDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAttachmentDeleted

`func (o *CustomAttachmentResponse) SetBAttachmentDeleted(v bool)`

SetBAttachmentDeleted sets BAttachmentDeleted field to given value.


### GetBAttachmentValid

`func (o *CustomAttachmentResponse) GetBAttachmentValid() bool`

GetBAttachmentValid returns the BAttachmentValid field if non-nil, zero value otherwise.

### GetBAttachmentValidOk

`func (o *CustomAttachmentResponse) GetBAttachmentValidOk() (*bool, bool)`

GetBAttachmentValidOk returns a tuple with the BAttachmentValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBAttachmentValid

`func (o *CustomAttachmentResponse) SetBAttachmentValid(v bool)`

SetBAttachmentValid sets BAttachmentValid field to given value.


### GetEAttachmentVerified

`func (o *CustomAttachmentResponse) GetEAttachmentVerified() FieldEAttachmentVerified`

GetEAttachmentVerified returns the EAttachmentVerified field if non-nil, zero value otherwise.

### GetEAttachmentVerifiedOk

`func (o *CustomAttachmentResponse) GetEAttachmentVerifiedOk() (*FieldEAttachmentVerified, bool)`

GetEAttachmentVerifiedOk returns a tuple with the EAttachmentVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEAttachmentVerified

`func (o *CustomAttachmentResponse) SetEAttachmentVerified(v FieldEAttachmentVerified)`

SetEAttachmentVerified sets EAttachmentVerified field to given value.


### GetTAttachmentRejectioncomment

`func (o *CustomAttachmentResponse) GetTAttachmentRejectioncomment() string`

GetTAttachmentRejectioncomment returns the TAttachmentRejectioncomment field if non-nil, zero value otherwise.

### GetTAttachmentRejectioncommentOk

`func (o *CustomAttachmentResponse) GetTAttachmentRejectioncommentOk() (*string, bool)`

GetTAttachmentRejectioncommentOk returns a tuple with the TAttachmentRejectioncomment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTAttachmentRejectioncomment

`func (o *CustomAttachmentResponse) SetTAttachmentRejectioncomment(v string)`

SetTAttachmentRejectioncomment sets TAttachmentRejectioncomment field to given value.

### HasTAttachmentRejectioncomment

`func (o *CustomAttachmentResponse) HasTAttachmentRejectioncomment() bool`

HasTAttachmentRejectioncomment returns a boolean if a field has been set.

### GetFkiUserIDOwner

`func (o *CustomAttachmentResponse) GetFkiUserIDOwner() int32`

GetFkiUserIDOwner returns the FkiUserIDOwner field if non-nil, zero value otherwise.

### GetFkiUserIDOwnerOk

`func (o *CustomAttachmentResponse) GetFkiUserIDOwnerOk() (*int32, bool)`

GetFkiUserIDOwnerOk returns a tuple with the FkiUserIDOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiUserIDOwner

`func (o *CustomAttachmentResponse) SetFkiUserIDOwner(v int32)`

SetFkiUserIDOwner sets FkiUserIDOwner field to given value.

### HasFkiUserIDOwner

`func (o *CustomAttachmentResponse) HasFkiUserIDOwner() bool`

HasFkiUserIDOwner returns a boolean if a field has been set.

### GetObjAudit

`func (o *CustomAttachmentResponse) GetObjAudit() CommonAudit`

GetObjAudit returns the ObjAudit field if non-nil, zero value otherwise.

### GetObjAuditOk

`func (o *CustomAttachmentResponse) GetObjAuditOk() (*CommonAudit, bool)`

GetObjAuditOk returns a tuple with the ObjAudit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAudit

`func (o *CustomAttachmentResponse) SetObjAudit(v CommonAudit)`

SetObjAudit sets ObjAudit field to given value.

### HasObjAudit

`func (o *CustomAttachmentResponse) HasObjAudit() bool`

HasObjAudit returns a boolean if a field has been set.

### GetObjAttachmentProof

`func (o *CustomAttachmentResponse) GetObjAttachmentProof() AttachmentResponseCompound`

GetObjAttachmentProof returns the ObjAttachmentProof field if non-nil, zero value otherwise.

### GetObjAttachmentProofOk

`func (o *CustomAttachmentResponse) GetObjAttachmentProofOk() (*AttachmentResponseCompound, bool)`

GetObjAttachmentProofOk returns a tuple with the ObjAttachmentProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAttachmentProof

`func (o *CustomAttachmentResponse) SetObjAttachmentProof(v AttachmentResponseCompound)`

SetObjAttachmentProof sets ObjAttachmentProof field to given value.

### HasObjAttachmentProof

`func (o *CustomAttachmentResponse) HasObjAttachmentProof() bool`

HasObjAttachmentProof returns a boolean if a field has been set.

### GetObjAttachmentProofdocument

`func (o *CustomAttachmentResponse) GetObjAttachmentProofdocument() AttachmentResponseCompound`

GetObjAttachmentProofdocument returns the ObjAttachmentProofdocument field if non-nil, zero value otherwise.

### GetObjAttachmentProofdocumentOk

`func (o *CustomAttachmentResponse) GetObjAttachmentProofdocumentOk() (*AttachmentResponseCompound, bool)`

GetObjAttachmentProofdocumentOk returns a tuple with the ObjAttachmentProofdocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAttachmentProofdocument

`func (o *CustomAttachmentResponse) SetObjAttachmentProofdocument(v AttachmentResponseCompound)`

SetObjAttachmentProofdocument sets ObjAttachmentProofdocument field to given value.

### HasObjAttachmentProofdocument

`func (o *CustomAttachmentResponse) HasObjAttachmentProofdocument() bool`

HasObjAttachmentProofdocument returns a boolean if a field has been set.

### GetAObjAttachmentAttachment

`func (o *CustomAttachmentResponse) GetAObjAttachmentAttachment() []AttachmentResponseCompound`

GetAObjAttachmentAttachment returns the AObjAttachmentAttachment field if non-nil, zero value otherwise.

### GetAObjAttachmentAttachmentOk

`func (o *CustomAttachmentResponse) GetAObjAttachmentAttachmentOk() (*[]AttachmentResponseCompound, bool)`

GetAObjAttachmentAttachmentOk returns a tuple with the AObjAttachmentAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttachmentAttachment

`func (o *CustomAttachmentResponse) SetAObjAttachmentAttachment(v []AttachmentResponseCompound)`

SetAObjAttachmentAttachment sets AObjAttachmentAttachment field to given value.

### HasAObjAttachmentAttachment

`func (o *CustomAttachmentResponse) HasAObjAttachmentAttachment() bool`

HasAObjAttachmentAttachment returns a boolean if a field has been set.

### GetAObjAttachmentVersion

`func (o *CustomAttachmentResponse) GetAObjAttachmentVersion() []AttachmentResponseCompound`

GetAObjAttachmentVersion returns the AObjAttachmentVersion field if non-nil, zero value otherwise.

### GetAObjAttachmentVersionOk

`func (o *CustomAttachmentResponse) GetAObjAttachmentVersionOk() (*[]AttachmentResponseCompound, bool)`

GetAObjAttachmentVersionOk returns a tuple with the AObjAttachmentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjAttachmentVersion

`func (o *CustomAttachmentResponse) SetAObjAttachmentVersion(v []AttachmentResponseCompound)`

SetAObjAttachmentVersion sets AObjAttachmentVersion field to given value.

### HasAObjAttachmentVersion

`func (o *CustomAttachmentResponse) HasAObjAttachmentVersion() bool`

HasAObjAttachmentVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


