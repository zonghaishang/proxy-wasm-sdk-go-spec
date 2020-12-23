package spec

import "github.com/zonghaishang/proxy-wasm-sdk-go-spec/spec/types"

//export proxy_log
func ProxyLog(logLevel types.LogLevel, buffer *byte, len int) types.Status

//export proxy_get_header_map_value
func ProxyGetHeaderMapValue(mapType types.MapType, keyData *byte, keySize int, returnValueData **byte, returnValueSize *int) types.Status

//export proxy_add_header_map_value
func ProxyAddHeaderMapValue(mapType types.MapType, keyData *byte, keySize int, valueData *byte, valueSize int) types.Status

//export proxy_replace_header_map_value
func ProxyReplaceHeaderMapValue(mapType types.MapType, keyData *byte, keySize int, valueData *byte, valueSize int) types.Status

//export proxy_remove_header_map_value
func ProxyRemoveHeaderMapValue(mapType types.MapType, keyData *byte, keySize int) types.Status

//export proxy_get_header_map_pairs
func ProxyGetHeaderMapPairs(mapType types.MapType, returnValueData **byte, returnValueSize *int) types.Status

//export proxy_set_header_map_pairs
func ProxySetHeaderMapPairs(mapType types.MapType, mapData *byte, mapSize int) types.Status

//export proxy_get_buffer_bytes
func ProxyGetBufferBytes(bt types.BufferType, start int, maxSize int, returnBufferData **byte, returnBufferSize *int) types.Status

//export proxy_set_buffer_bytes
func ProxySetBufferBytes(bt types.BufferType, start int, maxSize int, bufferData *byte, bufferSize int) types.Status

//export proxy_continue_stream
func ProxyContinueStream(streamType types.StreamType) types.Status

//export proxy_close_stream
func ProxyCloseStream(streamType types.StreamType) types.Status

//export proxy_set_tick_period_milliseconds
func ProxySetTickPeriodMilliseconds(period uint32) types.Status

//export proxy_get_current_time_nanoseconds
func ProxyGetCurrentTimeNanoseconds(returnTime *int64) types.Status

//export proxy_set_effective_context
func ProxySetEffectiveContext(contextID uint32) types.Status

//export proxy_done
func ProxyDone() types.Status

//export proxy_get_property
func ProxyGetProperty(pathData *byte, pathSize int, returnValueData **byte, returnValueSize *int) types.Status

//export proxy_set_property
func ProxySetProperty(pathData *byte, pathSize int, valueData *byte, valueSize int) types.Status
