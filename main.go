package main

import (
	"log"
	"report/httpReq"
	"report/misc"
	"report/rephrase"
)

func main() {
	// login & get login cookie
	loginCookie, err := httpReq.LoginItms()
	if err != nil {
		log.Fatalf("httpReq.LoginItms() err: %v", err)
	}
	// get a list of job which has any action in today
	listJobReply, err := httpReq.ListTodayJob(loginCookie)
	if err != nil {
		log.Fatalf("httpReq.ListTodayJob() err: %v", err)
	}
	// listJobReply, err := httpReq.ListRephrasedJob(loginCookie, "2024-05-17", "2024-05-20")
	// if err != nil {
	// 	log.Fatalf("httpReq.ListTodayJob() err: %v", err)
	// }
	//
	for _, job := range listJobReply.Data.Job {
		id := job.Id.(string)
		getJobReply, err := httpReq.GetJob(loginCookie, id)
		if err != nil {
			log.Fatalf("httpReq.GetJob(%s) err: %v", id, err)
		}
		rephrasedJob, createdAt, err := rephrase.RepharseJobToReportJob(getJobReply)
		if err != nil {
			log.Fatalf("rephrase.RepharseJobStruct(%s) err: %v", id, err)
		}
		createDate := createdAt[:4] + createdAt[5:7] + createdAt[8:10]
		fn := "json/local/job/" + createDate + "_" + id + ".json"
		err = misc.WriteToFile(id, rephrasedJob, fn)
		if err != nil {
			log.Fatalf("misc.WriteToFile(%s, rephrasedJob) err: %v", id, err)
		}
		if len(getJobReply.Data.Job[0].OrderFile) > 0 {
			if len(getJobReply.Data.Job[0].OrderFormInfo) > 0 {
				rephrasedOrderFormInfo, err := rephrase.RepharseJobToOrderFormInfo(getJobReply)
				if err != nil {
					log.Fatalf("rephrase.RepharseJobToOrderFormInfo(%s) err: %v", id, err)
				}
				orderFormInfoFn := "json/local/orderFormInfo/" + getJobReply.Data.Job[0].Info[0].TerminalId + ".json"
				err = misc.WriteOrderFormInfoToFile(id, rephrasedOrderFormInfo, orderFormInfoFn)
				if err != nil {
					log.Fatalf("misc.WriteOrderFormInfoToFile(%s, rephrasedOrderFormInfo) err: %v", id, err)
				}
			}
		}
	}
}
