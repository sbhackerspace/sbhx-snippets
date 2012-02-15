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


// {"pk": 4827, "model": "autocall.call", "fields": {"user": 1,
// "recip_number": "7073010751"}}

func jsonParser(jsonStr []byte) map[string]string {
	var m map[string]string
	var call interface{}
	err := json.Unmarshal(jsonStr, &call)
    if err != nil {
		fmt.Printf("%v\n", err)
        return m
    }
	c, ok := call.(map[string]interface{})
	if !ok {
		fmt.Printf("Problem!\n")
		return m
	}
	// c is a map of strings to empty interfaces
	for k, v := range c {
		vStr, ok := v.(string)
		if ok {
			m[k] = string(v)
		}
		vInt, ok := v.(int)
		if ok {
			m[k] = string(v)
		}
		if !ok {
			fmt.Printf("Not a string or int: %v\n", v)
		}
	}
	return m
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
	m := jsonParser( []byte(jsonStr) )
	fmt.Printf("\nResult: %v\n", m)
	// http.HandleFunc("/", jsonHandler)
	// http.ListenAndServe(":9999", nil)
}