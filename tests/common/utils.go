package common

import "encoding/json"

func ToPrettyJsonString(obj interface{}) (string, error) {
    out, err := json.MarshalIndent(obj, "", "    ")
    return string(out), err
}