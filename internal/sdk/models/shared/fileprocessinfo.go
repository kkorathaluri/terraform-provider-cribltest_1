// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type FileProcessInfo struct {
	Pid     float64 `json:"pid"`
	Process string  `json:"process"`
}

func (o *FileProcessInfo) GetPid() float64 {
	if o == nil {
		return 0.0
	}
	return o.Pid
}

func (o *FileProcessInfo) GetProcess() string {
	if o == nil {
		return ""
	}
	return o.Process
}
