package g20220901

import "encoding/json"

func testJson() {
	data := []byte(`{"status":200}`)

	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
	}
}
