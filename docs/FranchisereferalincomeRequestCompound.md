# FranchisereferalincomeRequestCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjAddress** | [**AddressRequest**](AddressRequest.md) |  | 
**AObjContact** | [**[]ContactRequestCompound**](ContactRequestCompound.md) |  | 
**PkiFranchisereferalincomeID** | Pointer to **int32** | The unique ID of the Franchisereferalincome | [optional] 
**FkiFranchisebrokerID** | **int32** | The unique ID of the Franchisebroker | 
**FkiFranchisereferalincomeprogramID** | **int32** | The unique ID of the Franchisereferalincomeprogram | 
**FkiPeriodID** | **int32** | The unique ID of the Period | 
**DFranchisereferalincomeLoan** | **string** | The loan amount | 
**DFranchisereferalincomeFranchiseamount** | **string** | The amount that will be given to the franchise | 
**DFranchisereferalincomeFranchisoramount** | **string** | The amount that will be kept by the franchisor | 
**DFranchisereferalincomeAgentamount** | **string** | The amount that will be given to the agent | 
**DtFranchisereferalincomeDisbursed** | **string** | The date the amounts were disbursed | 
**TFranchisereferalincomeComment** | **string** | A comment about the transaction | 
**FkiFranchiseofficeID** | **int32** | The unique ID of the Franchisereoffice | 
**SFranchisereferalincomeRemoteid** | **string** |  | 

## Methods

### NewFranchisereferalincomeRequestCompound

`func NewFranchisereferalincomeRequestCompound(objAddress AddressRequest, aObjContact []ContactRequestCompound, fkiFranchisebrokerID int32, fkiFranchisereferalincomeprogramID int32, fkiPeriodID int32, dFranchisereferalincomeLoan string, dFranchisereferalincomeFranchiseamount string, dFranchisereferalincomeFranchisoramount string, dFranchisereferalincomeAgentamount string, dtFranchisereferalincomeDisbursed string, tFranchisereferalincomeComment string, fkiFranchiseofficeID int32, sFranchisereferalincomeRemoteid string, ) *FranchisereferalincomeRequestCompound`

NewFranchisereferalincomeRequestCompound instantiates a new FranchisereferalincomeRequestCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFranchisereferalincomeRequestCompoundWithDefaults

`func NewFranchisereferalincomeRequestCompoundWithDefaults() *FranchisereferalincomeRequestCompound`

NewFranchisereferalincomeRequestCompoundWithDefaults instantiates a new FranchisereferalincomeRequestCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjAddress

`func (o *FranchisereferalincomeRequestCompound) GetObjAddress() AddressRequest`

GetObjAddress returns the ObjAddress field if non-nil, zero value otherwise.

### GetObjAddressOk

`func (o *FranchisereferalincomeRequestCompound) GetObjAddressOk() (*AddressRequest, bool)`

GetObjAddressOk returns a tuple with the ObjAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAddress

`func (o *FranchisereferalincomeRequestCompound) SetObjAddress(v AddressRequest)`

SetObjAddress sets ObjAddress field to given value.


### GetAObjContact

`func (o *FranchisereferalincomeRequestCompound) GetAObjContact() []ContactRequestCompound`

GetAObjContact returns the AObjContact field if non-nil, zero value otherwise.

### GetAObjContactOk

`func (o *FranchisereferalincomeRequestCompound) GetAObjContactOk() (*[]ContactRequestCompound, bool)`

GetAObjContactOk returns a tuple with the AObjContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAObjContact

`func (o *FranchisereferalincomeRequestCompound) SetAObjContact(v []ContactRequestCompound)`

SetAObjContact sets AObjContact field to given value.


### GetPkiFranchisereferalincomeID

`func (o *FranchisereferalincomeRequestCompound) GetPkiFranchisereferalincomeID() int32`

GetPkiFranchisereferalincomeID returns the PkiFranchisereferalincomeID field if non-nil, zero value otherwise.

### GetPkiFranchisereferalincomeIDOk

`func (o *FranchisereferalincomeRequestCompound) GetPkiFranchisereferalincomeIDOk() (*int32, bool)`

GetPkiFranchisereferalincomeIDOk returns a tuple with the PkiFranchisereferalincomeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiFranchisereferalincomeID

`func (o *FranchisereferalincomeRequestCompound) SetPkiFranchisereferalincomeID(v int32)`

SetPkiFranchisereferalincomeID sets PkiFranchisereferalincomeID field to given value.

### HasPkiFranchisereferalincomeID

`func (o *FranchisereferalincomeRequestCompound) HasPkiFranchisereferalincomeID() bool`

HasPkiFranchisereferalincomeID returns a boolean if a field has been set.

### GetFkiFranchisebrokerID

`func (o *FranchisereferalincomeRequestCompound) GetFkiFranchisebrokerID() int32`

GetFkiFranchisebrokerID returns the FkiFranchisebrokerID field if non-nil, zero value otherwise.

### GetFkiFranchisebrokerIDOk

`func (o *FranchisereferalincomeRequestCompound) GetFkiFranchisebrokerIDOk() (*int32, bool)`

GetFkiFranchisebrokerIDOk returns a tuple with the FkiFranchisebrokerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisebrokerID

`func (o *FranchisereferalincomeRequestCompound) SetFkiFranchisebrokerID(v int32)`

SetFkiFranchisebrokerID sets FkiFranchisebrokerID field to given value.


### GetFkiFranchisereferalincomeprogramID

`func (o *FranchisereferalincomeRequestCompound) GetFkiFranchisereferalincomeprogramID() int32`

GetFkiFranchisereferalincomeprogramID returns the FkiFranchisereferalincomeprogramID field if non-nil, zero value otherwise.

### GetFkiFranchisereferalincomeprogramIDOk

`func (o *FranchisereferalincomeRequestCompound) GetFkiFranchisereferalincomeprogramIDOk() (*int32, bool)`

GetFkiFranchisereferalincomeprogramIDOk returns a tuple with the FkiFranchisereferalincomeprogramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchisereferalincomeprogramID

`func (o *FranchisereferalincomeRequestCompound) SetFkiFranchisereferalincomeprogramID(v int32)`

SetFkiFranchisereferalincomeprogramID sets FkiFranchisereferalincomeprogramID field to given value.


### GetFkiPeriodID

`func (o *FranchisereferalincomeRequestCompound) GetFkiPeriodID() int32`

GetFkiPeriodID returns the FkiPeriodID field if non-nil, zero value otherwise.

### GetFkiPeriodIDOk

`func (o *FranchisereferalincomeRequestCompound) GetFkiPeriodIDOk() (*int32, bool)`

GetFkiPeriodIDOk returns a tuple with the FkiPeriodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiPeriodID

`func (o *FranchisereferalincomeRequestCompound) SetFkiPeriodID(v int32)`

SetFkiPeriodID sets FkiPeriodID field to given value.


### GetDFranchisereferalincomeLoan

`func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeLoan() string`

GetDFranchisereferalincomeLoan returns the DFranchisereferalincomeLoan field if non-nil, zero value otherwise.

### GetDFranchisereferalincomeLoanOk

`func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeLoanOk() (*string, bool)`

GetDFranchisereferalincomeLoanOk returns a tuple with the DFranchisereferalincomeLoan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDFranchisereferalincomeLoan

`func (o *FranchisereferalincomeRequestCompound) SetDFranchisereferalincomeLoan(v string)`

SetDFranchisereferalincomeLoan sets DFranchisereferalincomeLoan field to given value.


### GetDFranchisereferalincomeFranchiseamount

`func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeFranchiseamount() string`

GetDFranchisereferalincomeFranchiseamount returns the DFranchisereferalincomeFranchiseamount field if non-nil, zero value otherwise.

### GetDFranchisereferalincomeFranchiseamountOk

`func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeFranchiseamountOk() (*string, bool)`

GetDFranchisereferalincomeFranchiseamountOk returns a tuple with the DFranchisereferalincomeFranchiseamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDFranchisereferalincomeFranchiseamount

`func (o *FranchisereferalincomeRequestCompound) SetDFranchisereferalincomeFranchiseamount(v string)`

SetDFranchisereferalincomeFranchiseamount sets DFranchisereferalincomeFranchiseamount field to given value.


### GetDFranchisereferalincomeFranchisoramount

`func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeFranchisoramount() string`

GetDFranchisereferalincomeFranchisoramount returns the DFranchisereferalincomeFranchisoramount field if non-nil, zero value otherwise.

### GetDFranchisereferalincomeFranchisoramountOk

`func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeFranchisoramountOk() (*string, bool)`

GetDFranchisereferalincomeFranchisoramountOk returns a tuple with the DFranchisereferalincomeFranchisoramount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDFranchisereferalincomeFranchisoramount

`func (o *FranchisereferalincomeRequestCompound) SetDFranchisereferalincomeFranchisoramount(v string)`

SetDFranchisereferalincomeFranchisoramount sets DFranchisereferalincomeFranchisoramount field to given value.


### GetDFranchisereferalincomeAgentamount

`func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeAgentamount() string`

GetDFranchisereferalincomeAgentamount returns the DFranchisereferalincomeAgentamount field if non-nil, zero value otherwise.

### GetDFranchisereferalincomeAgentamountOk

`func (o *FranchisereferalincomeRequestCompound) GetDFranchisereferalincomeAgentamountOk() (*string, bool)`

GetDFranchisereferalincomeAgentamountOk returns a tuple with the DFranchisereferalincomeAgentamount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDFranchisereferalincomeAgentamount

`func (o *FranchisereferalincomeRequestCompound) SetDFranchisereferalincomeAgentamount(v string)`

SetDFranchisereferalincomeAgentamount sets DFranchisereferalincomeAgentamount field to given value.


### GetDtFranchisereferalincomeDisbursed

`func (o *FranchisereferalincomeRequestCompound) GetDtFranchisereferalincomeDisbursed() string`

GetDtFranchisereferalincomeDisbursed returns the DtFranchisereferalincomeDisbursed field if non-nil, zero value otherwise.

### GetDtFranchisereferalincomeDisbursedOk

`func (o *FranchisereferalincomeRequestCompound) GetDtFranchisereferalincomeDisbursedOk() (*string, bool)`

GetDtFranchisereferalincomeDisbursedOk returns a tuple with the DtFranchisereferalincomeDisbursed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtFranchisereferalincomeDisbursed

`func (o *FranchisereferalincomeRequestCompound) SetDtFranchisereferalincomeDisbursed(v string)`

SetDtFranchisereferalincomeDisbursed sets DtFranchisereferalincomeDisbursed field to given value.


### GetTFranchisereferalincomeComment

`func (o *FranchisereferalincomeRequestCompound) GetTFranchisereferalincomeComment() string`

GetTFranchisereferalincomeComment returns the TFranchisereferalincomeComment field if non-nil, zero value otherwise.

### GetTFranchisereferalincomeCommentOk

`func (o *FranchisereferalincomeRequestCompound) GetTFranchisereferalincomeCommentOk() (*string, bool)`

GetTFranchisereferalincomeCommentOk returns a tuple with the TFranchisereferalincomeComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTFranchisereferalincomeComment

`func (o *FranchisereferalincomeRequestCompound) SetTFranchisereferalincomeComment(v string)`

SetTFranchisereferalincomeComment sets TFranchisereferalincomeComment field to given value.


### GetFkiFranchiseofficeID

`func (o *FranchisereferalincomeRequestCompound) GetFkiFranchiseofficeID() int32`

GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field if non-nil, zero value otherwise.

### GetFkiFranchiseofficeIDOk

`func (o *FranchisereferalincomeRequestCompound) GetFkiFranchiseofficeIDOk() (*int32, bool)`

GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiFranchiseofficeID

`func (o *FranchisereferalincomeRequestCompound) SetFkiFranchiseofficeID(v int32)`

SetFkiFranchiseofficeID sets FkiFranchiseofficeID field to given value.


### GetSFranchisereferalincomeRemoteid

`func (o *FranchisereferalincomeRequestCompound) GetSFranchisereferalincomeRemoteid() string`

GetSFranchisereferalincomeRemoteid returns the SFranchisereferalincomeRemoteid field if non-nil, zero value otherwise.

### GetSFranchisereferalincomeRemoteidOk

`func (o *FranchisereferalincomeRequestCompound) GetSFranchisereferalincomeRemoteidOk() (*string, bool)`

GetSFranchisereferalincomeRemoteidOk returns a tuple with the SFranchisereferalincomeRemoteid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSFranchisereferalincomeRemoteid

`func (o *FranchisereferalincomeRequestCompound) SetSFranchisereferalincomeRemoteid(v string)`

SetSFranchisereferalincomeRemoteid sets SFranchisereferalincomeRemoteid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


