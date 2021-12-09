 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/divvy/v1/events.proto](#regen/divvy/v1/events.proto)
    - [EventCreateAllocator](#regen.divvy.v1.EventCreateAllocator)
    - [EventCreateStream](#regen.divvy.v1.EventCreateStream)
  
- [regen/divvy/v1/types.proto](#regen/divvy/v1/types.proto)
    - [Allocator](#regen.divvy.v1.Allocator)
    - [Recipient](#regen.divvy.v1.Recipient)
    - [SlowReleaseStream](#regen.divvy.v1.SlowReleaseStream)
    - [StoreAllocator](#regen.divvy.v1.StoreAllocator)
    - [StoreRecipient](#regen.divvy.v1.StoreRecipient)
    - [StoreSlowReleaseStream](#regen.divvy.v1.StoreSlowReleaseStream)
    - [StreamStrategy](#regen.divvy.v1.StreamStrategy)
  
- [regen/divvy/v1/query.proto](#regen/divvy/v1/query.proto)
    - [QueryAllocator](#regen.divvy.v1.QueryAllocator)
    - [QueryAllocators](#regen.divvy.v1.QueryAllocators)
    - [QueryAllocatorsByOwner](#regen.divvy.v1.QueryAllocatorsByOwner)
    - [QueryAllocatorsResp](#regen.divvy.v1.QueryAllocatorsResp)
    - [QueryStream](#regen.divvy.v1.QueryStream)
    - [QueryStreams](#regen.divvy.v1.QueryStreams)
    - [QueryStreamsResp](#regen.divvy.v1.QueryStreamsResp)
  
    - [Query](#regen.divvy.v1.Query)
  
- [regen/divvy/v1/tx.proto](#regen/divvy/v1/tx.proto)
    - [MsgClaimAllocations](#regen.divvy.v1.MsgClaimAllocations)
    - [MsgClaimAllocationsResp](#regen.divvy.v1.MsgClaimAllocationsResp)
    - [MsgCreateAllocator](#regen.divvy.v1.MsgCreateAllocator)
    - [MsgCreateAllocatorResp](#regen.divvy.v1.MsgCreateAllocatorResp)
    - [MsgCreateSlowReleaseStream](#regen.divvy.v1.MsgCreateSlowReleaseStream)
    - [MsgCreateSlowReleaseStreamResp](#regen.divvy.v1.MsgCreateSlowReleaseStreamResp)
    - [MsgEditSlowReleaseStream](#regen.divvy.v1.MsgEditSlowReleaseStream)
    - [MsgEmptyResp](#regen.divvy.v1.MsgEmptyResp)
    - [MsgPauseSlowReleaseStream](#regen.divvy.v1.MsgPauseSlowReleaseStream)
    - [MsgRemoveAllocator](#regen.divvy.v1.MsgRemoveAllocator)
    - [MsgSetAllocatorRecipients](#regen.divvy.v1.MsgSetAllocatorRecipients)
    - [MsgUpdateAllocatorSettings](#regen.divvy.v1.MsgUpdateAllocatorSettings)
  
    - [Msg](#regen.divvy.v1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/divvy/v1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/divvy/v1/events.proto



<a name="regen.divvy.v1.EventCreateAllocator"></a>

### EventCreateAllocator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="regen.divvy.v1.EventCreateStream"></a>

### EventCreateStream



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/divvy/v1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/divvy/v1/types.proto



<a name="regen.divvy.v1.Allocator"></a>

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
| address | [string](#string) |  | submodule address of the given allocator |
| recipients | [Recipient](#regen.divvy.v1.Recipient) | repeated | Invariant: * sum of shares in entires must equal to 100% (1mln) list of allocation entries |
| next_claim | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp when anyone can call for the next time. |






<a name="regen.divvy.v1.Recipient"></a>

### Recipient



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address wallet address |
| share | [uint32](#uint32) |  | allocation share. 100% = 1e6. |
| name | [string](#string) |  |  |






<a name="regen.divvy.v1.SlowReleaseStream"></a>

### SlowReleaseStream



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | signer and creator of the stream |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | when the stream starts |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution |
| destination | [string](#string) |  | Allocator address |
| name | [string](#string) |  | name of the allocator |
| paused | [bool](#bool) |  |  |
| strategy | [StreamStrategy](#regen.divvy.v1.StreamStrategy) |  |  |
| address | [string](#string) |  | submodule address of the given stream |






<a name="regen.divvy.v1.StoreAllocator"></a>

### StoreAllocator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | admin is the address of the account that creates the allocator and signs the message |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution, min = 1s |
| name | [string](#string) |  | name of the allocator |
| url | [string](#string) |  | url with metadata |
| paused | [bool](#bool) |  |  |
| recipients | [StoreRecipient](#regen.divvy.v1.StoreRecipient) | repeated | Invariant: * sum of shares in entires must equal to 100% (1mln) list of allocation entries |
| next_claim | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp when anyone can call for the next time. |






<a name="regen.divvy.v1.StoreRecipient"></a>

### StoreRecipient



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [bytes](#bytes) |  | address wallet address |
| share | [uint32](#uint32) |  | allocation share. 100% = 1e6. |
| name | [string](#string) |  |  |






<a name="regen.divvy.v1.StoreSlowReleaseStream"></a>

### StoreSlowReleaseStream



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [bytes](#bytes) |  | signer and creator of the stream |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | when the stream starts |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution |
| destination | [bytes](#bytes) |  | Allocator address |
| name | [string](#string) |  | name of the allocator |
| paused | [bool](#bool) |  |  |
| strategy | [StreamStrategy](#regen.divvy.v1.StreamStrategy) |  |  |






<a name="regen.divvy.v1.StreamStrategy"></a>

### StreamStrategy



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fixed_amount | [string](#string) |  | fixed amount of tokens streamed in each round. If there is a zero balance available then then nothing will be streamed. If only fraction is available then the it will be fully streamed. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/divvy/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/divvy/v1/query.proto



<a name="regen.divvy.v1.QueryAllocator"></a>

### QueryAllocator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="regen.divvy.v1.QueryAllocators"></a>

### QueryAllocators



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.divvy.v1.QueryAllocatorsByOwner"></a>

### QueryAllocatorsByOwner



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner | [string](#string) |  |  |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.divvy.v1.QueryAllocatorsResp"></a>

### QueryAllocatorsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allocator | [Allocator](#regen.divvy.v1.Allocator) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="regen.divvy.v1.QueryStream"></a>

### QueryStream



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="regen.divvy.v1.QueryStreams"></a>

### QueryStreams



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="regen.divvy.v1.QueryStreamsResp"></a>

### QueryStreamsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| streams | [SlowReleaseStream](#regen.divvy.v1.SlowReleaseStream) | repeated |  |
| pagination | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.divvy.v1.Query"></a>

### Query
Msg is the divvy Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AllocatorByAddress | [QueryAllocator](#regen.divvy.v1.QueryAllocator) | [Allocator](#regen.divvy.v1.Allocator) |  |
| Allocators | [QueryAllocators](#regen.divvy.v1.QueryAllocators) | [QueryAllocatorsResp](#regen.divvy.v1.QueryAllocatorsResp) |  |
| AllocatorsByOwner | [QueryAllocatorsByOwner](#regen.divvy.v1.QueryAllocatorsByOwner) | [QueryAllocatorsResp](#regen.divvy.v1.QueryAllocatorsResp) |  |
| StreamByAddress | [QueryStream](#regen.divvy.v1.QueryStream) | [SlowReleaseStream](#regen.divvy.v1.SlowReleaseStream) |  |
| Streams | [QueryStreams](#regen.divvy.v1.QueryStreams) | [QueryStreamsResp](#regen.divvy.v1.QueryStreamsResp) |  |

 <!-- end services -->



<a name="regen/divvy/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/divvy/v1/tx.proto



<a name="regen.divvy.v1.MsgClaimAllocations"></a>

### MsgClaimAllocations



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | signer, anyone can claim rewards |
| allocator | [string](#string) |  | allocator address |






<a name="regen.divvy.v1.MsgClaimAllocationsResp"></a>

### MsgClaimAllocationsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| coins | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | distributed allocations |






<a name="regen.divvy.v1.MsgCreateAllocator"></a>

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
| recipients | [Recipient](#regen.divvy.v1.Recipient) | repeated | Initial allocator mapping. Invariants: * sum of shares in recipients must equal to 100% (1mln) |






<a name="regen.divvy.v1.MsgCreateAllocatorResp"></a>

### MsgCreateAllocatorResp
MsgCreateClassResponse is the Msg/CreateAllocator response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | Address is a unique address of newly created Allocator |






<a name="regen.divvy.v1.MsgCreateSlowReleaseStream"></a>

### MsgCreateSlowReleaseStream



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [string](#string) |  | signer and creator of the stream |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | when the stream starts |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution |
| name | [string](#string) |  |  |
| destination | [string](#string) |  | Allocator address |
| paused | [bool](#bool) |  | when paused, stream won't send funds |
| strategy | [StreamStrategy](#regen.divvy.v1.StreamStrategy) |  |  |






<a name="regen.divvy.v1.MsgCreateSlowReleaseStreamResp"></a>

### MsgCreateSlowReleaseStreamResp
MsgCreateSlowReleaseStreamResp is response for
Msg/CreateSlowReleaseStreamResp


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address of the newly created streamer |






<a name="regen.divvy.v1.MsgEditSlowReleaseStream"></a>

### MsgEditSlowReleaseStream



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address of a stream |
| sender | [string](#string) |  | sender must the the Stream admin |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | when the stream starts |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution |
| name | [string](#string) |  |  |
| destination | [string](#string) |  | Allocator address |
| paused | [bool](#bool) |  | when paused, stream won't send funds |
| strategy | [StreamStrategy](#regen.divvy.v1.StreamStrategy) |  |  |






<a name="regen.divvy.v1.MsgEmptyResp"></a>

### MsgEmptyResp







<a name="regen.divvy.v1.MsgPauseSlowReleaseStream"></a>

### MsgPauseSlowReleaseStream



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address of a stream |
| sender | [string](#string) |  | sender must the the Stream admin |
| paused | [bool](#bool) |  | the pause value to set |






<a name="regen.divvy.v1.MsgRemoveAllocator"></a>

### MsgRemoveAllocator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address of the allocator |
| sender | [string](#string) |  | sender must the the Allocator admin |






<a name="regen.divvy.v1.MsgSetAllocatorRecipients"></a>

### MsgSetAllocatorRecipients



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address of the allocator |
| sender | [string](#string) |  | sender must the the Allocator admin |
| recipients | [Recipient](#regen.divvy.v1.Recipient) | repeated | New allocator mapping. Invariants: * sum of shares in recipients must equal to 100% (1mln) |






<a name="regen.divvy.v1.MsgUpdateAllocatorSettings"></a>

### MsgUpdateAllocatorSettings



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | address of the allocator |
| sender | [string](#string) |  | sender must the the Allocator admin |
| start | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| end | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| interval | [google.protobuf.Duration](#google.protobuf.Duration) |  | how often we do a distribution |
| name | [string](#string) |  | name of the allocator |
| url | [string](#string) |  | url with metadata |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.divvy.v1.Msg"></a>

### Msg
Msg is the divvy Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateAllocator | [MsgCreateAllocator](#regen.divvy.v1.MsgCreateAllocator) | [MsgCreateAllocatorResp](#regen.divvy.v1.MsgCreateAllocatorResp) | Allocator is a distribution engine, which "divvys out" all incoming funds, at configurable time intervals to all registered recipients. Each allocator has only one owner. Ideally this can be managed by a group module. |
| UpdateAllocatorSettings | [MsgUpdateAllocatorSettings](#regen.divvy.v1.MsgUpdateAllocatorSettings) | [MsgEmptyResp](#regen.divvy.v1.MsgEmptyResp) | Updates all allocator settings except admin and recipient map. |
| SetAllocatorRecipients | [MsgSetAllocatorRecipients](#regen.divvy.v1.MsgSetAllocatorRecipients) | [MsgEmptyResp](#regen.divvy.v1.MsgEmptyResp) | Allocator owner can update the recipient list by setting a new allocation map. |
| RemoveAllocator | [MsgRemoveAllocator](#regen.divvy.v1.MsgRemoveAllocator) | [MsgEmptyResp](#regen.divvy.v1.MsgEmptyResp) | Removes allocator and disables all streamers! |
| ClaimAllocations | [MsgClaimAllocations](#regen.divvy.v1.MsgClaimAllocations) | [MsgClaimAllocationsResp](#regen.divvy.v1.MsgClaimAllocationsResp) |  |
| CreateSlowReleaseStream | [MsgCreateSlowReleaseStream](#regen.divvy.v1.MsgCreateSlowReleaseStream) | [MsgCreateSlowReleaseStreamResp](#regen.divvy.v1.MsgCreateSlowReleaseStreamResp) | Creates a new stream to feed an address User creates a stream. Parameters: * % of total amount to be streamed to allocator every second. * destination address where the stream will go (ideally allocator) |
| PauseSlowReleaseStream | [MsgPauseSlowReleaseStream](#regen.divvy.v1.MsgPauseSlowReleaseStream) | [MsgEmptyResp](#regen.divvy.v1.MsgEmptyResp) |  |
| EditSlowReleaseStream | [MsgEditSlowReleaseStream](#regen.divvy.v1.MsgEditSlowReleaseStream) | [MsgEmptyResp](#regen.divvy.v1.MsgEmptyResp) |  |

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

