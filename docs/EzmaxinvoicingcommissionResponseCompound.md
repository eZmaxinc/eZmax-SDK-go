# EzmaxinvoicingcommissionResponseCompound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkiEzmaxinvoicingcommissionID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingcommission | [optional] 
**FkiEzmaxinvoicingsummaryglobalID** | Pointer to **int32** | The unique ID of the Ezmaxinvoicingsummaryglobal | [optional] 
**FkiEzmaxpartnerID** | Pointer to **int32** | The unique ID of the Ezmaxpartner | [optional] 
**FkiEzmaxrepresentativeID** | Pointer to **int32** | The unique ID of the Ezmaxrepresentative | [optional] 
**DtEzmaxinvoicingcommissionStart** | **string** | The start date for the Ezmaxinvoicingcommission | 
**DtEzmaxinvoicingcommissionEnd** | **string** | The end date for the Ezmaxinvoicingcommission | 
**IEzmaxinvoicingcommissionDays** | **int32** | This is the number of days during the month on which the Ezmaxinvoigcommission applies | 
**DEzmaxinvoicingcommissionAmount** | **string** | The amount of Ezmaxinvoicingcommission | 
**ObjContactName** | Pointer to [**CustomContactNameResponse**](CustomContactNameResponse.md) |  | [optional] 

## Methods

### NewEzmaxinvoicingcommissionResponseCompound

`func NewEzmaxinvoicingcommissionResponseCompound(dtEzmaxinvoicingcommissionStart string, dtEzmaxinvoicingcommissionEnd string, iEzmaxinvoicingcommissionDays int32, dEzmaxinvoicingcommissionAmount string, ) *EzmaxinvoicingcommissionResponseCompound`

NewEzmaxinvoicingcommissionResponseCompound instantiates a new EzmaxinvoicingcommissionResponseCompound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzmaxinvoicingcommissionResponseCompoundWithDefaults

`func NewEzmaxinvoicingcommissionResponseCompoundWithDefaults() *EzmaxinvoicingcommissionResponseCompound`

NewEzmaxinvoicingcommissionResponseCompoundWithDefaults instantiates a new EzmaxinvoicingcommissionResponseCompound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkiEzmaxinvoicingcommissionID

`func (o *EzmaxinvoicingcommissionResponseCompound) GetPkiEzmaxinvoicingcommissionID() int32`

GetPkiEzmaxinvoicingcommissionID returns the PkiEzmaxinvoicingcommissionID field if non-nil, zero value otherwise.

### GetPkiEzmaxinvoicingcommissionIDOk

`func (o *EzmaxinvoicingcommissionResponseCompound) GetPkiEzmaxinvoicingcommissionIDOk() (*int32, bool)`

GetPkiEzmaxinvoicingcommissionIDOk returns a tuple with the PkiEzmaxinvoicingcommissionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiEzmaxinvoicingcommissionID

`func (o *EzmaxinvoicingcommissionResponseCompound) SetPkiEzmaxinvoicingcommissionID(v int32)`

SetPkiEzmaxinvoicingcommissionID sets PkiEzmaxinvoicingcommissionID field to given value.

### HasPkiEzmaxinvoicingcommissionID

`func (o *EzmaxinvoicingcommissionResponseCompound) HasPkiEzmaxinvoicingcommissionID() bool`

HasPkiEzmaxinvoicingcommissionID returns a boolean if a field has been set.

### GetFkiEzmaxinvoicingsummaryglobalID

`func (o *EzmaxinvoicingcommissionResponseCompound) GetFkiEzmaxinvoicingsummaryglobalID() int32`

GetFkiEzmaxinvoicingsummaryglobalID returns the FkiEzmaxinvoicingsummaryglobalID field if non-nil, zero value otherwise.

### GetFkiEzmaxinvoicingsummaryglobalIDOk

`func (o *EzmaxinvoicingcommissionResponseCompound) GetFkiEzmaxinvoicingsummaryglobalIDOk() (*int32, bool)`

GetFkiEzmaxinvoicingsummaryglobalIDOk returns a tuple with the FkiEzmaxinvoicingsummaryglobalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxinvoicingsummaryglobalID

`func (o *EzmaxinvoicingcommissionResponseCompound) SetFkiEzmaxinvoicingsummaryglobalID(v int32)`

SetFkiEzmaxinvoicingsummaryglobalID sets FkiEzmaxinvoicingsummaryglobalID field to given value.

### HasFkiEzmaxinvoicingsummaryglobalID

`func (o *EzmaxinvoicingcommissionResponseCompound) HasFkiEzmaxinvoicingsummaryglobalID() bool`

HasFkiEzmaxinvoicingsummaryglobalID returns a boolean if a field has been set.

### GetFkiEzmaxpartnerID

`func (o *EzmaxinvoicingcommissionResponseCompound) GetFkiEzmaxpartnerID() int32`

GetFkiEzmaxpartnerID returns the FkiEzmaxpartnerID field if non-nil, zero value otherwise.

### GetFkiEzmaxpartnerIDOk

`func (o *EzmaxinvoicingcommissionResponseCompound) GetFkiEzmaxpartnerIDOk() (*int32, bool)`

GetFkiEzmaxpartnerIDOk returns a tuple with the FkiEzmaxpartnerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxpartnerID

`func (o *EzmaxinvoicingcommissionResponseCompound) SetFkiEzmaxpartnerID(v int32)`

SetFkiEzmaxpartnerID sets FkiEzmaxpartnerID field to given value.

### HasFkiEzmaxpartnerID

`func (o *EzmaxinvoicingcommissionResponseCompound) HasFkiEzmaxpartnerID() bool`

HasFkiEzmaxpartnerID returns a boolean if a field has been set.

### GetFkiEzmaxrepresentativeID

`func (o *EzmaxinvoicingcommissionResponseCompound) GetFkiEzmaxrepresentativeID() int32`

GetFkiEzmaxrepresentativeID returns the FkiEzmaxrepresentativeID field if non-nil, zero value otherwise.

### GetFkiEzmaxrepresentativeIDOk

`func (o *EzmaxinvoicingcommissionResponseCompound) GetFkiEzmaxrepresentativeIDOk() (*int32, bool)`

GetFkiEzmaxrepresentativeIDOk returns a tuple with the FkiEzmaxrepresentativeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFkiEzmaxrepresentativeID

`func (o *EzmaxinvoicingcommissionResponseCompound) SetFkiEzmaxrepresentativeID(v int32)`

SetFkiEzmaxrepresentativeID sets FkiEzmaxrepresentativeID field to given value.

### HasFkiEzmaxrepresentativeID

`func (o *EzmaxinvoicingcommissionResponseCompound) HasFkiEzmaxrepresentativeID() bool`

HasFkiEzmaxrepresentativeID returns a boolean if a field has been set.

### GetDtEzmaxinvoicingcommissionStart

`func (o *EzmaxinvoicingcommissionResponseCompound) GetDtEzmaxinvoicingcommissionStart() string`

GetDtEzmaxinvoicingcommissionStart returns the DtEzmaxinvoicingcommissionStart field if non-nil, zero value otherwise.

### GetDtEzmaxinvoicingcommissionStartOk

`func (o *EzmaxinvoicingcommissionResponseCompound) GetDtEzmaxinvoicingcommissionStartOk() (*string, bool)`

GetDtEzmaxinvoicingcommissionStartOk returns a tuple with the DtEzmaxinvoicingcommissionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzmaxinvoicingcommissionStart

`func (o *EzmaxinvoicingcommissionResponseCompound) SetDtEzmaxinvoicingcommissionStart(v string)`

SetDtEzmaxinvoicingcommissionStart sets DtEzmaxinvoicingcommissionStart field to given value.


### GetDtEzmaxinvoicingcommissionEnd

`func (o *EzmaxinvoicingcommissionResponseCompound) GetDtEzmaxinvoicingcommissionEnd() string`

GetDtEzmaxinvoicingcommissionEnd returns the DtEzmaxinvoicingcommissionEnd field if non-nil, zero value otherwise.

### GetDtEzmaxinvoicingcommissionEndOk

`func (o *EzmaxinvoicingcommissionResponseCompound) GetDtEzmaxinvoicingcommissionEndOk() (*string, bool)`

GetDtEzmaxinvoicingcommissionEndOk returns a tuple with the DtEzmaxinvoicingcommissionEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEzmaxinvoicingcommissionEnd

`func (o *EzmaxinvoicingcommissionResponseCompound) SetDtEzmaxinvoicingcommissionEnd(v string)`

SetDtEzmaxinvoicingcommissionEnd sets DtEzmaxinvoicingcommissionEnd field to given value.


### GetIEzmaxinvoicingcommissionDays

`func (o *EzmaxinvoicingcommissionResponseCompound) GetIEzmaxinvoicingcommissionDays() int32`

GetIEzmaxinvoicingcommissionDays returns the IEzmaxinvoicingcommissionDays field if non-nil, zero value otherwise.

### GetIEzmaxinvoicingcommissionDaysOk

`func (o *EzmaxinvoicingcommissionResponseCompound) GetIEzmaxinvoicingcommissionDaysOk() (*int32, bool)`

GetIEzmaxinvoicingcommissionDaysOk returns a tuple with the IEzmaxinvoicingcommissionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIEzmaxinvoicingcommissionDays

`func (o *EzmaxinvoicingcommissionResponseCompound) SetIEzmaxinvoicingcommissionDays(v int32)`

SetIEzmaxinvoicingcommissionDays sets IEzmaxinvoicingcommissionDays field to given value.


### GetDEzmaxinvoicingcommissionAmount

`func (o *EzmaxinvoicingcommissionResponseCompound) GetDEzmaxinvoicingcommissionAmount() string`

GetDEzmaxinvoicingcommissionAmount returns the DEzmaxinvoicingcommissionAmount field if non-nil, zero value otherwise.

### GetDEzmaxinvoicingcommissionAmountOk

`func (o *EzmaxinvoicingcommissionResponseCompound) GetDEzmaxinvoicingcommissionAmountOk() (*string, bool)`

GetDEzmaxinvoicingcommissionAmountOk returns a tuple with the DEzmaxinvoicingcommissionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEzmaxinvoicingcommissionAmount

`func (o *EzmaxinvoicingcommissionResponseCompound) SetDEzmaxinvoicingcommissionAmount(v string)`

SetDEzmaxinvoicingcommissionAmount sets DEzmaxinvoicingcommissionAmount field to given value.


### GetObjContactName

`func (o *EzmaxinvoicingcommissionResponseCompound) GetObjContactName() CustomContactNameResponse`

GetObjContactName returns the ObjContactName field if non-nil, zero value otherwise.

### GetObjContactNameOk

`func (o *EzmaxinvoicingcommissionResponseCompound) GetObjContactNameOk() (*CustomContactNameResponse, bool)`

GetObjContactNameOk returns a tuple with the ObjContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjContactName

`func (o *EzmaxinvoicingcommissionResponseCompound) SetObjContactName(v CustomContactNameResponse)`

SetObjContactName sets ObjContactName field to given value.

### HasObjContactName

`func (o *EzmaxinvoicingcommissionResponseCompound) HasObjContactName() bool`

HasObjContactName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


