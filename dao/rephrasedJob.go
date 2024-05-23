package dao

type Job struct {
	Id                   interface{}   `json:"id"`
	Info                 []Info        `json:"info"`
	TerminalInfo         interface{}   `json:"terminalInfo"`
	OrderFile            []OrderFile   `json:"orderFile"`
	Misc                 interface{}   `json:"misc"`
	Onsite               interface{}   `json:"onsite"`
	State                interface{}   `json:"state"`
	History              interface{}   `json:"history"`
	ScheduleDate         string        `json:"scheduleDate"`
	ScheduleRemark       interface{}   `json:"scheduleRemark"`
	DeliveryInfo         interface{}   `json:"deliveryInfo"`
	UpdateRemarks        interface{}   `json:"updateRemarks"`
	AssignedWorker       interface{}   `json:"assignedWorker"`
	Type                 interface{}   `json:"type"`
	Bank                 string        `json:"bank"`
	Hardware             []Hardware    `json:"hardware"`
	ChecklistId          interface{}   `json:"checklistId"`
	VerifyChecklistId    interface{}   `json:"verifyChecklistId"`
	Reject               interface{}   `json:"reject"`
	CRMId                interface{}   `json:"CRMId"`
	Signature            interface{}   `json:"signature"`
	ProfileUpdate        interface{}   `json:"profileUpdate"`
	CollectedAccessories interface{}   `json:"collectedAccessories"`
	TimeRecord           interface{}   `json:"timeRecord"`
	CreatedAt            string        `json:"createAt"`
	OrderFormInfo        []interface{} `json:"orderFormInfo"`
}

type RephrasedJob struct {
	Id                   interface{} `json:"id"`
	Info                 []Info      `json:"info"`
	TerminalInfo         interface{} `json:"terminalInfo"`
	OrderFile            []OrderFile `json:"orderFile"`
	Misc                 interface{} `json:"misc"`
	Onsite               interface{} `json:"onsite"`
	State                interface{} `json:"state"`
	History              interface{} `json:"history"`
	ScheduleDate         string      `json:"scheduleDate"`
	ScheduleRemark       interface{} `json:"scheduleRemark"`
	DeliveryInfo         interface{} `json:"deliveryInfo"`
	UpdateRemarks        interface{} `json:"updateRemarks"`
	AssignedWorker       interface{} `json:"assignedWorker"`
	Type                 interface{} `json:"type"`
	Bank                 string      `json:"bank"`
	Hardware             []Hardware  `json:"hardware"`
	ChecklistId          interface{} `json:"checklistId"`
	VerifyChecklistId    interface{} `json:"verifyChecklistId"`
	Reject               interface{} `json:"reject"`
	CRMId                interface{} `json:"CRMId"`
	Signature            interface{} `json:"signature"`
	ProfileUpdate        interface{} `json:"profileUpdate"`
	CollectedAccessories interface{} `json:"collectedAccessories"`
	TimeRecord           interface{} `json:"timeRecord"`
	CreatedAt            string      `json:"createAt"`
	Sn                   string      `json:"sn,omitempty"`
	DeliveryNoteNo       string      `json:"deliveryNoteNo,omitempty"`
	// OrderFormInfo        interface{} `json:"orderFormInfo,omitempty"`
}

type Info struct {
	PrimaryKey      string `json:"primaryKey,omitempty"`
	TerminalId      string `json:"terminalId,omitempty"`
	MerchantId      string `json:"merchantId,omitempty"`
	MerchantName    string `json:"merchantName,omitempty"`
	MerchantAddress string `json:"merchantAddress,omitempty"`
	OutletName      string `json:"outletName,omitempty"`
	ContactPerson   string `json:"contactPerson,omitempty"`
	ContactPhone    string `json:"contactPhone,omitempty"`
	ExpectedDate    string `json:"expectedDate,omitempty"`
}

type OrderFile struct {
	OriOrderFileName string        `json:"oriOrderFilename,omitempty"`
	Name             string        `json:"name,omitempty"`
	Uploaded         bool          `json:"uploaded,omitempty"`
	Info             []interface{} `json:"info,omitempty"`
}

type Hardware struct {
	ItemId    string `json:"itemId"`
	Battery   bool   `json:"battery"`
	Cable     bool   `json:"cable"`
	Charger   bool   `json:"charger"`
	Cover     bool   `json:"cover"`
	Docking   bool   `json:"docking"`
	Sim       bool   `json:"sim"`
	OldItemId string `json:"oldItemId,omitempty"`
}

type OrderFormInfo struct {
	Bank          string        `json:"bank"`
	OrderFormInfo []interface{} `json:"orderFormInfo"`
	ScheduleDate  string        `json:"scheduleDate"`
}
