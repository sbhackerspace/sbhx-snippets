// Steve Phillips / elimisteve
// 2011.09.02

package main

import (
    "fmt"
//	"http"
//	"io/ioutil"
	"json"
)

const dialStr = `Action: Login
Username: confuser
Secret: c0nfc0nn3ct

Priority: 1
RetryTime: 60
ActionID: %s
Context: conf
Action: Originate
Exten: 123
MaxRetries: 2
WaitTime: 30
Channel: SIP/voipms/%s
Variable: FOO=%s
Variable: clid=%s
Variable: accountcode=%s
Variable: mycallid=%s

`


// {"pk": 4807, "model": "autocall.call", "fields":
//      {"recip_ext": "", "scheduled_for": "2011-09-02 01:34:00",
//       "created_at": "2011-09-02 01:34:01",
//       "modified_at": "2011-09-02 01:34:01",
//       "attempted": true, "completed": false, "user": 1,
//       "recip_number": "7073010751", "duration": 0,
//       "recip_name": "", "start_time": "1970-01-01 00:00:00"}
//  }

// {"pk": 4827, "model": "autocall.call", "fields": {"user": 1,
// "recip_number": "7073010751"}}

func jsonParser(jsonStr []byte) {
	var call interface{}
	err := json.Unmarshal(jsonStr, &call)
    if err != nil {
		fmt.Printf("%v\n", err)
        return
    }
	switch c := call.(type) {
	case map[string]interface{}:
		for k, v := range c {
			switch vType := v.(type) {
			case int, string:
				fmt.Printf("%v: %v\n", k, v)
			case map[string]interface{}:
				for key, val := range vType {
					fmt.Printf("%v: %v\n", key, val)
				}
			default:
				fmt.Printf("WTF is v?\n")
			}
		}
	default:
		fmt.Printf("All JSON should be key/value pairs!\n")
	}
    return
}

// jsonHandler accepts JSON on port 9999, parses Call data, then tells
// the PBX to dial a call via TCP
// func jsonHandler(w http.ResponseWriter, r *http.Request) {
// 	// Parse incoming JSON
// 	str, _ := ioutil.ReadAll(r.Body)
// 	fmt.Printf("%v\n", str)

// 	json := jsonParser()

// 	// Connect to PBX
// 	pbx, err := net.Dial("tcp", "127.0.0.1:5038")
// 	if err != nil {
// 		fmt.Printf(" ** Couldn't connect to PBX!\n\n")
// 		fmt.Printf("%s\n", err)
// 		os.Exit(1)
// 	}


// 	for _, call := range callData {

// 	}

//     // Send data to PBX
// 	fmt.Fprintf(pbx, dialStr, call.id, call.recip_number,
// 		call.user, call.recip_number, call.user, call.id)
// }

func main() {
	jsonStr := `{"pk": 4827, "model": "autocall.call", "fields": {"user": 1, "recip_number": "7073010751"}}`
	jsonParser( []byte(jsonStr) )
	// http.HandleFunc("/", jsonHandler)
	// http.ListenAndServe(":9999", nil)
}