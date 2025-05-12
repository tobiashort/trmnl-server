package main

import (
	"testing"
)

func TestParsePayload(t *testing.T) {
	payload := `{"log":{"logs_array":[{"creation_timestamp":1746793033,"device_status_stamp":{"wifi_rssi_level":-54,"wifi_status":"connected","refresh_rate":900,"time_since_last_sleep_start":617,"current_fw_version":"1.5.4","special_function":"none","battery_voltage":3.894,"wakeup_reason":"button","free_heap_size":185936,"max_alloc_size":159732},"log_id":10,"log_message":"Error fetching API display: 7, detail: HTTP Client failed with error: (501)","log_codeline":579,"log_sourcefile":"src/bl.cpp","additional_info":{"filename_current":"","filename_new":"","retry_attempt":5}},{"creation_timestamp":1746792363,"device_status_stamp":{"wifi_rssi_level":-56,"wifi_status":"connected","refresh_rate":5,"time_since_last_sleep_start":6,"current_fw_version":"1.5.4","special_function":"none","battery_voltage":3.9,"wakeup_reason":"timer","free_heap_size":215032,"max_alloc_size":192500},"log_id":6,"log_message":"Error fetching API display: 7, detail: HTTP Client failed with error: connection refused(-1)","log_codeline":579,"log_sourcefile":"src/bl.cpp","additional_info":{"filename_current":"","filename_new":"","retry_attempt":2}},{"creation_timestamp":1746792374,"device_status_stamp":{"wifi_rssi_level":-54,"wifi_status":"connected","refresh_rate":10,"time_since_last_sleep_start":11,"current_fw_version":"1.5.4","special_function":"none","battery_voltage":3.898,"wakeup_reason":"timer","free_heap_size":215132,"max_alloc_size":192500},"log_id":7,"log_message":"Error fetching API display: 7, detail: HTTP Client failed with error: connection refused(-1)","log_codeline":579,"log_sourcefile":"src/bl.cpp","additional_info":{"filename_current":"","filename_new":"","retry_attempt":3}},{"creation_timestamp":1746792405,"device_status_stamp":{"wifi_rssi_level":-57,"wifi_status":"connected","refresh_rate":30,"time_since_last_sleep_start":31,"current_fw_version":"1.5.4","special_function":"none","battery_voltage":3.896,"wakeup_reason":"timer","free_heap_size":215164,"max_alloc_size":192500},"log_id":8,"log_message":"Error fetching API display: 7, detail: HTTP Client failed with error: connection refused(-1)","log_codeline":579,"log_sourcefile":"src/bl.cpp","additional_info":{"filename_current":"","filename_new":"","retry_attempt":4}},{"creation_timestamp":1746793030,"device_status_stamp":{"wifi_rssi_level":-53,"wifi_status":"connected","refresh_rate":900,"time_since_last_sleep_start":617,"current_fw_version":"1.5.4","special_function":"none","battery_voltage":3.894,"wakeup_reason":"button","free_heap_size":186684,"max_alloc_size":159732},"log_id":9,"log_message":"Failed to resolve hostname after 5 attempts, continuing...","log_codeline":566,"log_sourcefile":"src/bl.cpp","additional_info":{"filename_current":"","filename_new":"","retry_attempt":5}}]}}`

	h := LogHandler{}
	err := h.parsePayload([]byte(payload))
	if err != nil {
		t.Fatal(err)
	}
}
