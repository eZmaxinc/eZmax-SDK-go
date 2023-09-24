# EzsignfolderBatchDownloadV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APkiEzsigndocumentID** | **[]int32** |  | 
**AEDocumentType** | **[]string** | The type of document to retrieve.  1. **Signed** Is the final document once all signatures were applied. 2. **Proofdocument** Is the evidence report. 3. **Proof** Is the complete evidence archive including all of the above and more. | 

## Methods

### NewEzsignfolderBatchDownloadV1Request

`func NewEzsignfolderBatchDownloadV1Request(aPkiEzsigndocumentID []int32, aEDocumentType []string, ) *EzsignfolderBatchDownloadV1Request`

NewEzsignfolderBatchDownloadV1Request instantiates a new EzsignfolderBatchDownloadV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEzsignfolderBatchDownloadV1RequestWithDefaults

`func NewEzsignfolderBatchDownloadV1RequestWithDefaults() *EzsignfolderBatchDownloadV1Request`

NewEzsignfolderBatchDownloadV1RequestWithDefaults instantiates a new EzsignfolderBatchDownloadV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPkiEzsigndocumentID

`func (o *EzsignfolderBatchDownloadV1Request) GetAPkiEzsigndocumentID() []int32`

GetAPkiEzsigndocumentID returns the APkiEzsigndocumentID field if non-nil, zero value otherwise.

### GetAPkiEzsigndocumentIDOk

`func (o *EzsignfolderBatchDownloadV1Request) GetAPkiEzsigndocumentIDOk() (*[]int32, bool)`

GetAPkiEzsigndocumentIDOk returns a tuple with the APkiEzsigndocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPkiEzsigndocumentID

`func (o *EzsignfolderBatchDownloadV1Request) SetAPkiEzsigndocumentID(v []int32)`

SetAPkiEzsigndocumentID sets APkiEzsigndocumentID field to given value.


### GetAEDocumentType

`func (o *EzsignfolderBatchDownloadV1Request) GetAEDocumentType() []string`

GetAEDocumentType returns the AEDocumentType field if non-nil, zero value otherwise.

### GetAEDocumentTypeOk

`func (o *EzsignfolderBatchDownloadV1Request) GetAEDocumentTypeOk() (*[]string, bool)`

GetAEDocumentTypeOk returns a tuple with the AEDocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAEDocumentType

`func (o *EzsignfolderBatchDownloadV1Request) SetAEDocumentType(v []string)`

SetAEDocumentType sets AEDocumentType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


