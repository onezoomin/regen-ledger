 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/divvy/v1/types.proto](#regen/divvy/v1/types.proto)
    - [Allocator](#regen.ecocredit.v1alpha2.Allocator)
    - [Recipient](#regen.ecocredit.v1alpha2.Recipient)
    - [SlowReleaseStream](#regen.ecocredit.v1alpha2.SlowReleaseStream)
  
- [regen/divvy/v1/query.proto](#regen/divvy/v1/query.proto)
    - [Query](#regen.ecocredit.v1alpha2.Query)
  
- [regen/divvy/v1/tx.proto](#regen/divvy/v1/tx.proto)
    - [MsgCreateAllocator](#regen.ecocredit.v1alpha2.MsgCreateAllocator)
    - [MsgCreateAllocatorResp](#regen.ecocredit.v1alpha2.MsgCreateAllocatorResp)
    - [MsgCreateSlowReleaseStream](#regen.ecocredit.v1alpha2.MsgCreateSlowReleaseStream)
    - [MsgEditSlowReleaseStream](#regen.ecocredit.v1alpha2.MsgEditSlowReleaseStream)
    - [MsgEmptyResp](#regen.ecocredit.v1alpha2.MsgEmptyResp)
    - [MsgPauseSlowReleaseStream](#regen.ecocredit.v1alpha2.MsgPauseSlowReleaseStream)
    - [MsgRemoveAllocator](#regen.ecocredit.v1alpha2.MsgRemoveAllocator)
    - [MsgSetAllocationMap](#regen.ecocredit.v1alpha2.MsgSetAllocationMap)
    - [MsgUpdateAllocatorSetting](#regen.ecocredit.v1alpha2.MsgUpdateAllocatorSetting)
  
    - [Msg](#regen.ecocredit.v1alpha2.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/divvy/v1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/divvy/v1/types.proto



<a name="regen.ecocredit.v1alpha2.Allocator"></a>

### Allocator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account that creates the allocator and signs the message |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution, min = 1s |
| name | [string](#string) |  | name of the allocator |
| url | [string](#string) |  | url with metadata |
| paused | [bool](#bool) |  |  |
| entries | [Recipient](#regen.ecocredit.v1alpha2.Recipient) | repeated | Invariant: * sum of shares in entires must equal to 100% (1mln) list of allocation entries |






<a name="regen.ecocredit.v1alpha2.Recipient"></a>

### Recipient



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address wallet address |
| share | [uint32](#uint32) |  | allocation share. 100% = 1e6. |






<a name="regen.ecocredit.v1alpha2.SlowReleaseStream"></a>

### SlowReleaseStream



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | signer and creator of the stream |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | when the stream starts |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution |
| destination | [string](#string) |  | Allocator address |
| name | [string](#string) |  | name of the allocator |
| paused | [bool](#bool) |  |  |
| fixed_amount | [string](#string) |  | fixed amount of tokens streamed in each round. If there is a zero balance available then then nothing will be streamed. If only fraction is available then the it will be fully streamed. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/divvy/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/divvy/v1/query.proto


 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.v1alpha2.Query"></a>

### Query
Msg is the regen.ecocredit.v1alpha1 Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 <!-- end services -->



<a name="regen/divvy/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/divvy/v1/tx.proto



<a name="regen.ecocredit.v1alpha2.MsgCreateAllocator"></a>

### MsgCreateAllocator
MsgCreateClass is the Msg/CreateClass request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account that creates the allocator and signs the message |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution |
| name | [string](#string) |  | name of the allocator |
| url | [string](#string) |  | url with metadata |
| recipients | [Recipient](#regen.ecocredit.v1alpha2.Recipient) | repeated | Initial allocator mapping. Invariants: * sum of shares in recipients must equal to 100% (1mln) |






<a name="regen.ecocredit.v1alpha2.MsgCreateAllocatorResp"></a>

### MsgCreateAllocatorResp
MsgCreateClassResponse is the Msg/CreateClass response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | id is the unique ID of the newly created allocator. |






<a name="regen.ecocredit.v1alpha2.MsgCreateSlowReleaseStream"></a>

### MsgCreateSlowReleaseStream



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | signer and creator of the stream |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | when the stream starts |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution |
| destination | [string](#string) |  | Allocator address |
| paused | [bool](#bool) |  | when paused, stream won't send funds |
| fixed_amount | [string](#string) |  | fixed amount of tokens streamed in each round. If there is a zero balance available then then nothing will be streamed. If only fraction is available then the it will be fully streamed. |






<a name="regen.ecocredit.v1alpha2.MsgEditSlowReleaseStream"></a>

### MsgEditSlowReleaseStream



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender must the the Stream admin |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | when the stream starts |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution |
| destination | [string](#string) |  | Allocator address |
| fixed_amount | [string](#string) |  | fixed amount of tokens streamed in each round. If there is a zero balance available then then nothing will be streamed. If only fraction is available then the it will be fully streamed. |






<a name="regen.ecocredit.v1alpha2.MsgEmptyResp"></a>

### MsgEmptyResp







<a name="regen.ecocredit.v1alpha2.MsgPauseSlowReleaseStream"></a>

### MsgPauseSlowReleaseStream



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender must the the Stream admin |
| paused | [bool](#bool) |  | the pause value to set |






<a name="regen.ecocredit.v1alpha2.MsgRemoveAllocator"></a>

### MsgRemoveAllocator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender must the the Allocator admin |






<a name="regen.ecocredit.v1alpha2.MsgSetAllocationMap"></a>

### MsgSetAllocationMap



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender must the the Allocator admin |
| recipients | [Recipient](#regen.ecocredit.v1alpha2.Recipient) | repeated | New allocator mapping. Invariants: * sum of shares in recipients must equal to 100% (1mln) |






<a name="regen.ecocredit.v1alpha2.MsgUpdateAllocatorSetting"></a>

### MsgUpdateAllocatorSetting



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender must the the Allocator admin |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution, min 1s |
| name | [string](#string) |  | name of the allocator |
| url | [string](#string) |  | url with metadata |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.ecocredit.v1alpha2.Msg"></a>

### Msg
Msg is the regen.ecocredit.v1alpha1 Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateAllocator | [MsgCreateAllocator](#regen.ecocredit.v1alpha2.MsgCreateAllocator) | [MsgCreateAllocatorResp](#regen.ecocredit.v1alpha2.MsgCreateAllocatorResp) | Allocator is a distribution engine, which distributes everything which is comming in configurable interval periods to registered recipients. Each allocator has only one owner. Ideally this can be managed by a group module. |
| UpdateAllocatorSetting | [MsgUpdateAllocatorSetting](#regen.ecocredit.v1alpha2.MsgUpdateAllocatorSetting) | [MsgEmptyResp](#regen.ecocredit.v1alpha2.MsgEmptyResp) | Updates all allocator settings except admin and recipient map. |
| SetAllocationMap | [MsgSetAllocationMap](#regen.ecocredit.v1alpha2.MsgSetAllocationMap) | [MsgEmptyResp](#regen.ecocredit.v1alpha2.MsgEmptyResp) | Allocator owner can update the recipient list by setting a new allocation map. |
| RemoveAllocator | [MsgRemoveAllocator](#regen.ecocredit.v1alpha2.MsgRemoveAllocator) | [MsgCreateAllocatorResp](#regen.ecocredit.v1alpha2.MsgCreateAllocatorResp) | Removes allocator and disables all streamers! |
| CreateSlowReleaseStream | [MsgCreateSlowReleaseStream](#regen.ecocredit.v1alpha2.MsgCreateSlowReleaseStream) | [MsgEmptyResp](#regen.ecocredit.v1alpha2.MsgEmptyResp) | Creates a new stream to feed an address User creates a stream. Parameters: * % of total amount to be streamed to allocator every second. * destination address where the stream will go (ideally allocator) |
| PauseSlowReleaseStream | [MsgPauseSlowReleaseStream](#regen.ecocredit.v1alpha2.MsgPauseSlowReleaseStream) | [MsgEmptyResp](#regen.ecocredit.v1alpha2.MsgEmptyResp) |  |
| EditSlowReleaseStream | [MsgEditSlowReleaseStream](#regen.ecocredit.v1alpha2.MsgEditSlowReleaseStream) | [MsgEmptyResp](#regen.ecocredit.v1alpha2.MsgEmptyResp) |  |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

