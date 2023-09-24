# EzmaxinvoicingagentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicingagentID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingagent | [optional] 
**FkiEzmaxinvoicingID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicing | [optional] 
**FkiBillingentityinternalID** | **int32** | The unique ID of the Billingentityinternal. | 
**SBillingentityinternalDescriptionX** | **string** | The description of the Billingentityinternal in the language of the requester | 
**FkiAgentID** | Pointer to **int32** | The unique ID of the Agent. | [optional] 
**FkiBrokerID** | Pointer to **int32** | The unique ID of the Broker. | [optional] 
**IEzmaxinvoicingagentSession** | **int32** | The number of sessions | 
**IEzmaxinvoicingagentCloned** | **int32** | The number of times this user was cloned | 
**IEzmaxinvoicingagentInvoice** | **int32** | The number of invoices | 
**IEzmaxinvoicingagentInscription** | **int32** | The number of inscriptions | 
**IEzmaxinvoicingagentInscriptionactive** | **int32** | The number of active inscriptions | 
**IEzmaxinvoicingagentSale** | **int32** | The number of sales | 
**IEzmaxinvoicingagentOtherincome** | **int32** | The number of otherincomes | 
**IEzmaxinvoicingagentCommissioncalculation** | **int32** | The number of commission calculations | 
**IEzmaxinvoicingagentEzsigndocument** | **int32** | The number of ezsign documents | 
**BEzmaxinvoicingagentEzsignaccount** | **bool** | Whether the agent has an eZsign account | 
**BEzmaxinvoicingagentBillableezmax** | **bool** | Whether it is billable for eZmax | 
**EEzmaxinvoicingagentVariationezmax** | [**FieldEEzmaxinvoicingagentVariationezmax**](FieldEEzmaxinvoicingagentVariationezmax.md) |  | 
**BEzmaxinvoicingagentBillableezsign** | **bool** | Whether it is billable for eZsign | 
**EEzmaxinvoicingagentVariationezsign** | [**FieldEEzmaxinvoicingagentVariationezsign**](FieldEEzmaxinvoicingagentVariationezsign.md) |  | 

## Methods

### NewEzmaxinvoicingagentResponse

`func NewEzmaxinvoicingagentResponse(fkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string, iEzmaxinvoicingagentSession int32, iEzmaxinvoicingagentCloned int32, iEzmaxinvoicingagentInvoice int32, iEzmaxinvoicingagentInscription int32, iEzmaxinvoicingagentInscriptionactive int32, iEzmaxinvoicingagentSale int32, iEzmaxinvoicingagentOtherincome int32, iEzmaxinvoicingagentCommissioncalculation int32, iEzmaxinvoicingagentEzsigndocument int32, bEzmaxinvoicingagentEzsignaccount bool, bEzmaxinvoicingagentBillableezmax bool, eEzmaxinvoicingagentVariationezmax FieldEEzmaxinvoicingagentVariationezmax, bEzmaxinvoicingagentBillableezsign bool, eEzmaxinvoicingagentVariationezsign FieldEEzmaxinvoicingagentVariationezsign, ) *EzmaxinvoicingagentResponse`

NewEzmaxinvoicingagentResponse instantiates a new EzmaxinvoicingagentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingagentResponseWithDefaults

`func NewEzmaxinvoicingagentResponseWithDefaults() *EzmaxinvoicingagentResponse`

NewEzmaxinvoicingagentResponseWithDefaults instantiates a new EzmaxinvoicingagentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingagentID

`func (o *EzmaxinvoicingagentResponse) GetPkiEzmaxinvoicingagentID() int32`

GetPkiEzmaxinvoicingagentID returns the PkiEzmaxinvoicingagentID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingagentIDOk

`func (o *EzmaxinvoicingagentResponse) GetPkiEzmaxinvoicingagentIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingagentIDOk returns a tuple with the PkiEzmaxinvoicingagentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingagentID

`func (o *EzmaxinvoicingagentResponse) SetPkiEzmaxinvoicingagentID(v int32)`

SetPkiEzmaxinvoicingagentID sets PkiEzmaxinvoicingagentID field to given value.

### HasPkiEzmaxinvoicingagentID

`func (o *EzmaxinvoicingagentResponse) HasPkiEzmaxinvoicingagentID() bool`

HasPkiEzmaxinvoicingagentID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingagentResponse) GetFkiEzmaxinvoicingID() int32`

GetFkiEzmaxinvoicingID returns the FkiEzmaxinvoicingID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingIDOk

`func (o *EzmaxinvoicingagentResponse) GetFkiEzmaxinvoicingIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingIDOk returns a tuple with the FkiEzmaxinvoicingID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingagentResponse) SetFkiEzmaxinvoicingID(v int32)`

SetFkiEzmaxinvoicingID sets FkiEzmaxinvoicingID field to given value.

### HasFkiEzmaxinvoicingID

`func (o *EzmaxinvoicingagentResponse) HasFkiEzmaxinvoicingID() bool`

HasFkiEzmaxinvoicingID returns a boolean if a field has been set.

### GetFkiBillingentityinternalID

`func (o *EzmaxinvoicingagentResponse) GetFkiBillingentityinternalID() int32`

GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field if non-nil, zero value otherwise.

### GetFkiBillingentityinternalIDOk

`func (o *EzmaxinvoicingagentResponse) GetFkiBillingentityinternalIDOk() (*int32, bool)`

GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBillingentityinternalID

`func (o *EzmaxinvoicingagentResponse) SetFkiBillingentityinternalID(v int32)`

SetFkiBillingentityinternalID sets FkiBillingentityinternalID field to given value.


### GetSBillingentityinternalDescriptionX

`func (o *EzmaxinvoicingagentResponse) GetSBillingentityinternalDescriptionX() string`

GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field if non-nil, zero value otherwise.

### GetSBillingentityinternalDescriptionXOk

`func (o *EzmaxinvoicingagentResponse) GetSBillingentityinternalDescriptionXOk() (*string, bool)`

GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSBillingentityinternalDescriptionX

`func (o *EzmaxinvoicingagentResponse) SetSBillingentityinternalDescriptionX(v string)`

SetSBillingentityinternalDescriptionX sets SBillingentityinternalDescriptionX field to given value.


### GetFkiAgentID

`func (o *EzmaxinvoicingagentResponse) GetFkiAgentID() int32`

GetFkiAgentID returns the FkiAgentID field if non-nil, zero value otherwise.

### GetFkiAgentIDOk

`func (o *EzmaxinvoicingagentResponse) GetFkiAgentIDOk() (*int32, bool)`

GetFkiAgentIDOk returns a tuple with the FkiAgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiAgentID

`func (o *EzmaxinvoicingagentResponse) SetFkiAgentID(v int32)`

SetFkiAgentID sets FkiAgentID field to given value.

### HasFkiAgentID

`func (o *EzmaxinvoicingagentResponse) HasFkiAgentID() bool`

HasFkiAgentID returns a boolean if a field has been set.

### GetFkiBrokerID

`func (o *EzmaxinvoicingagentResponse) GetFkiBrokerID() int32`

GetFkiBrokerID returns the FkiBrokerID field if non-nil, zero value otherwise.

### GetFkiBrokerIDOk

`func (o *EzmaxinvoicingagentResponse) GetFkiBrokerIDOk() (*int32, bool)`

GetFkiBrokerIDOk returns a tuple with the FkiBrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiBrokerID

`func (o *EzmaxinvoicingagentResponse) SetFkiBrokerID(v int32)`

SetFkiBrokerID sets FkiBrokerID field to given value.

### HasFkiBrokerID

`func (o *EzmaxinvoicingagentResponse) HasFkiBrokerID() bool`

HasFkiBrokerID returns a boolean if a field has been set.

### GetIEzmaxinvoicingagentSession

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentSession() int32`

GetIEzmaxinvoicingagentSession returns the IEzmaxinvoicingagentSession field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentSessionOk

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentSessionOk() (*int32, bool)`

GetIEzmaxinvoicingagentSessionOk returns a tuple with the IEzmaxinvoicingagentSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentSession

`func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentSession(v int32)`

SetIEzmaxinvoicingagentSession sets IEzmaxinvoicingagentSession field to given value.


### GetIEzmaxinvoicingagentCloned

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentCloned() int32`

GetIEzmaxinvoicingagentCloned returns the IEzmaxinvoicingagentCloned field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentClonedOk

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentClonedOk() (*int32, bool)`

GetIEzmaxinvoicingagentClonedOk returns a tuple with the IEzmaxinvoicingagentCloned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentCloned

`func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentCloned(v int32)`

SetIEzmaxinvoicingagentCloned sets IEzmaxinvoicingagentCloned field to given value.


### GetIEzmaxinvoicingagentInvoice

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInvoice() int32`

GetIEzmaxinvoicingagentInvoice returns the IEzmaxinvoicingagentInvoice field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentInvoiceOk

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInvoiceOk() (*int32, bool)`

GetIEzmaxinvoicingagentInvoiceOk returns a tuple with the IEzmaxinvoicingagentInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentInvoice

`func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentInvoice(v int32)`

SetIEzmaxinvoicingagentInvoice sets IEzmaxinvoicingagentInvoice field to given value.


### GetIEzmaxinvoicingagentInscription

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInscription() int32`

GetIEzmaxinvoicingagentInscription returns the IEzmaxinvoicingagentInscription field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentInscriptionOk

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInscriptionOk() (*int32, bool)`

GetIEzmaxinvoicingagentInscriptionOk returns a tuple with the IEzmaxinvoicingagentInscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentInscription

`func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentInscription(v int32)`

SetIEzmaxinvoicingagentInscription sets IEzmaxinvoicingagentInscription field to given value.


### GetIEzmaxinvoicingagentInscriptionactive

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInscriptionactive() int32`

GetIEzmaxinvoicingagentInscriptionactive returns the IEzmaxinvoicingagentInscriptionactive field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentInscriptionactiveOk

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentInscriptionactiveOk() (*int32, bool)`

GetIEzmaxinvoicingagentInscriptionactiveOk returns a tuple with the IEzmaxinvoicingagentInscriptionactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentInscriptionactive

`func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentInscriptionactive(v int32)`

SetIEzmaxinvoicingagentInscriptionactive sets IEzmaxinvoicingagentInscriptionactive field to given value.


### GetIEzmaxinvoicingagentSale

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentSale() int32`

GetIEzmaxinvoicingagentSale returns the IEzmaxinvoicingagentSale field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentSaleOk

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentSaleOk() (*int32, bool)`

GetIEzmaxinvoicingagentSaleOk returns a tuple with the IEzmaxinvoicingagentSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentSale

`func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentSale(v int32)`

SetIEzmaxinvoicingagentSale sets IEzmaxinvoicingagentSale field to given value.


### GetIEzmaxinvoicingagentOtherincome

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentOtherincome() int32`

GetIEzmaxinvoicingagentOtherincome returns the IEzmaxinvoicingagentOtherincome field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentOtherincomeOk

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentOtherincomeOk() (*int32, bool)`

GetIEzmaxinvoicingagentOtherincomeOk returns a tuple with the IEzmaxinvoicingagentOtherincome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentOtherincome

`func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentOtherincome(v int32)`

SetIEzmaxinvoicingagentOtherincome sets IEzmaxinvoicingagentOtherincome field to given value.


### GetIEzmaxinvoicingagentCommissioncalculation

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentCommissioncalculation() int32`

GetIEzmaxinvoicingagentCommissioncalculation returns the IEzmaxinvoicingagentCommissioncalculation field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentCommissioncalculationOk

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentCommissioncalculationOk() (*int32, bool)`

GetIEzmaxinvoicingagentCommissioncalculationOk returns a tuple with the IEzmaxinvoicingagentCommissioncalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentCommissioncalculation

`func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentCommissioncalculation(v int32)`

SetIEzmaxinvoicingagentCommissioncalculation sets IEzmaxinvoicingagentCommissioncalculation field to given value.


### GetIEzmaxinvoicingagentEzsigndocument

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentEzsigndocument() int32`

GetIEzmaxinvoicingagentEzsigndocument returns the IEzmaxinvoicingagentEzsigndocument field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingagentEzsigndocumentOk

`func (o *EzmaxinvoicingagentResponse) GetIEzmaxinvoicingagentEzsigndocumentOk() (*int32, bool)`

GetIEzmaxinvoicingagentEzsigndocumentOk returns a tuple with the IEzmaxinvoicingagentEzsigndocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingagentEzsigndocument

`func (o *EzmaxinvoicingagentResponse) SetIEzmaxinvoicingagentEzsigndocument(v int32)`

SetIEzmaxinvoicingagentEzsigndocument sets IEzmaxinvoicingagentEzsigndocument field to given value.


### GetBEzmaxinvoicingagentEzsignaccount

`func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentEzsignaccount() bool`

GetBEzmaxinvoicingagentEzsignaccount returns the BEzmaxinvoicingagentEzsignaccount field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingagentEzsignaccountOk

`func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentEzsignaccountOk() (*bool, bool)`

GetBEzmaxinvoicingagentEzsignaccountOk returns a tuple with the BEzmaxinvoicingagentEzsignaccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingagentEzsignaccount

`func (o *EzmaxinvoicingagentResponse) SetBEzmaxinvoicingagentEzsignaccount(v bool)`

SetBEzmaxinvoicingagentEzsignaccount sets BEzmaxinvoicingagentEzsignaccount field to given value.


### GetBEzmaxinvoicingagentBillableezmax

`func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentBillableezmax() bool`

GetBEzmaxinvoicingagentBillableezmax returns the BEzmaxinvoicingagentBillableezmax field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingagentBillableezmaxOk

`func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentBillableezmaxOk() (*bool, bool)`

GetBEzmaxinvoicingagentBillableezmaxOk returns a tuple with the BEzmaxinvoicingagentBillableezmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingagentBillableezmax

`func (o *EzmaxinvoicingagentResponse) SetBEzmaxinvoicingagentBillableezmax(v bool)`

SetBEzmaxinvoicingagentBillableezmax sets BEzmaxinvoicingagentBillableezmax field to given value.


### GetEEzmaxinvoicingagentVariationezmax

`func (o *EzmaxinvoicingagentResponse) GetEEzmaxinvoicingagentVariationezmax() FieldEEzmaxinvoicingagentVariationezmax`

GetEEzmaxinvoicingagentVariationezmax returns the EEzmaxinvoicingagentVariationezmax field if non-nil, zero value otherwise.

### GetEEzmaxinvoicingagentVariationezmaxOk

`func (o *EzmaxinvoicingagentResponse) GetEEzmaxinvoicingagentVariationezmaxOk() (*FieldEEzmaxinvoicingagentVariationezmax, bool)`

GetEEzmaxinvoicingagentVariationezmaxOk returns a tuple with the EEzmaxinvoicingagentVariationezmax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxinvoicingagentVariationezmax

`func (o *EzmaxinvoicingagentResponse) SetEEzmaxinvoicingagentVariationezmax(v FieldEEzmaxinvoicingagentVariationezmax)`

SetEEzmaxinvoicingagentVariationezmax sets EEzmaxinvoicingagentVariationezmax field to given value.


### GetBEzmaxinvoicingagentBillableezsign

`func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentBillableezsign() bool`

GetBEzmaxinvoicingagentBillableezsign returns the BEzmaxinvoicingagentBillableezsign field if non-nil, zero value otherwise.

### GetBEzmaxinvoicingagentBillableezsignOk

`func (o *EzmaxinvoicingagentResponse) GetBEzmaxinvoicingagentBillableezsignOk() (*bool, bool)`

GetBEzmaxinvoicingagentBillableezsignOk returns a tuple with the BEzmaxinvoicingagentBillableezsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBEzmaxinvoicingagentBillableezsign

`func (o *EzmaxinvoicingagentResponse) SetBEzmaxinvoicingagentBillableezsign(v bool)`

SetBEzmaxinvoicingagentBillableezsign sets BEzmaxinvoicingagentBillableezsign field to given value.


### GetEEzmaxinvoicingagentVariationezsign

`func (o *EzmaxinvoicingagentResponse) GetEEzmaxinvoicingagentVariationezsign() FieldEEzmaxinvoicingagentVariationezsign`

GetEEzmaxinvoicingagentVariationezsign returns the EEzmaxinvoicingagentVariationezsign field if non-nil, zero value otherwise.

### GetEEzmaxinvoicingagentVariationezsignOk

`func (o *EzmaxinvoicingagentResponse) GetEEzmaxinvoicingagentVariationezsignOk() (*FieldEEzmaxinvoicingagentVariationezsign, bool)`

GetEEzmaxinvoicingagentVariationezsignOk returns a tuple with the EEzmaxinvoicingagentVariationezsign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEEzmaxinvoicingagentVariationezsign

`func (o *EzmaxinvoicingagentResponse) SetEEzmaxinvoicingagentVariationezsign(v FieldEEzmaxinvoicingagentVariationezsign)`

SetEEzmaxinvoicingagentVariationezsign sets EEzmaxinvoicingagentVariationezsign field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


