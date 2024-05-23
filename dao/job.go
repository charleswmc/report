package dao

type ListJobReply struct {
	Data Data `json:"data"`
}

type GetJobReply struct {
	Data Data `json:"data"`
}

type Data struct {
	Job []Job `json:"job"`
}

type ReportJob struct {
	Job []RephrasedJob `json:"job"`
}

type ReportOrderFormInfo struct {
	OrderFormInfo []OrderFormInfo
}

// type Job struct {
// 	Id                   interface{} `json:"id,order:1"`
// 	Info                 interface{} `json:"info,order:2"`
// 	TerminalInfo         interface{} `json:"terminalInfo,order:3"`
// 	OrderFile            interface{} `json:"orderFile,order:4"`
// 	Misc                 interface{} `json:"misc,order:5"`
// 	Onsite               interface{} `json:"onsite,order:6"`
// 	State                interface{} `json:"state,order:7"`
// 	History              interface{} `json:"history,order:8"`
// 	ScheduleDate         interface{} `json:"scheduleDate,order:9"`
// 	ScheduleRemark       interface{} `json:"scheduleRemark,order:10"`
// 	DeliveryInfo         interface{} `json:"deliveryInfo,order:11"`
// 	UpdateRemarks        interface{} `json:"updateRemarks,order:12"`
// 	AssignedWorker       interface{} `json:"assignedWorker,order:13"`
// 	Type                 interface{} `json:"type,order:14"`
// 	Bank                 interface{} `json:"bank,order:15"`
// 	Hardware             interface{} `json:"hardware,order:16"`
// 	ChecklistId          interface{} `json:"checklistId,order:17"`
// 	VerifyChecklistId    interface{} `json:"verifyChecklistId,order:18"`
// 	Reject               interface{} `json:"reject,order:19"`
// 	CRMId                interface{} `json:"CRMId,order:20"`
// 	Signature            interface{} `json:"signature,order:21"`
// 	ProfileUpdate        interface{} `json:"profileUpdate,order:22"`
// 	CollectedAccessories interface{} `json:"collectedAccessories,order:23"`
// 	TimeRecord           interface{} `json:"timeRecord,order:24"`
// 	CreatedAt            interface{} `json:"createAt,order:25"`
// 	OrderFormInfo        interface{} `json:"orderFormInfo,order:26"`
// }
