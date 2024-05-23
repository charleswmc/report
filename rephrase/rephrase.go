package rephrase

import (
	"encoding/json"
	"fmt"
	"report/dao"
	"strings"
)

func RepharseJobToReportJob(getJobReply dao.GetJobReply) (reportJob []dao.ReportJob, createdAt string, err error) {
	var rephrasedJob []dao.RephrasedJob
	dataByte, _ := json.Marshal(getJobReply.Data.Job[0])
	var tempJob dao.RephrasedJob
	err = json.Unmarshal(dataByte, &tempJob)
	if err != nil {
		return nil, "", fmt.Errorf("json.Unmarshal(dataByte, &temp) err: %v", err)
	}
	createdAt = tempJob.CreatedAt
	rephrasedJob = append(rephrasedJob, tempJob)
	var tempReportJob dao.ReportJob
	tempReportJob.Job = rephrasedJob
	if tempReportJob.Job[0].Type == "installation" {
		// fmt.Println("jobId:", tempReportJob.Job[0].Id)
		if len(tempReportJob.Job[0].Hardware) > 0 {
			fmt.Println("jobId:", tempReportJob.Job[0].Id)
			itemIdSeg := strings.Split(tempReportJob.Job[0].Hardware[0].ItemId, "-")
			tempReportJob.Job[0].Sn = itemIdSeg[2]
			createDate := tempReportJob.Job[0].CreatedAt[:10]
			createDateSeg := strings.Split(createDate, "-")
			tempReportJob.Job[0].DeliveryNoteNo = "DN" + createDateSeg[0] + createDateSeg[1] + createDateSeg[2] + "_" + tempReportJob.Job[0].Info[0].TerminalId
		}
	}

	reportJob = append(reportJob, tempReportJob)
	return reportJob, createdAt, err
}

func RepharseJobToOrderFormInfo(getJobReply dao.GetJobReply) (rephrasedOrderFormInfo []dao.OrderFormInfo, err error) {
	dataByte, _ := json.Marshal(getJobReply.Data.Job[0])
	var tempOrderFormInfo dao.OrderFormInfo
	err = json.Unmarshal(dataByte, &tempOrderFormInfo)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(dataByte, &tempOrderFormInfo) err: %v", err)
	}
	if tempOrderFormInfo.ScheduleDate != "" {
		tempOrderFormInfo.ScheduleDate = tempOrderFormInfo.ScheduleDate[:10]
	}

	rephrasedOrderFormInfo = append(rephrasedOrderFormInfo, tempOrderFormInfo)

	return rephrasedOrderFormInfo, err
}
