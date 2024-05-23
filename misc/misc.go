package misc

import (
	"encoding/json"
	"fmt"
	"os"
	"report/dao"
)

func WriteToFile(id string, rephrasedJob []dao.ReportJob, fn string) error {
	// write to file, fn - date_jobId
	jobJson, _ := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0777)
	defer jobJson.Close()
	rephrasedJobByte, err := json.Marshal(rephrasedJob)
	if err != nil {
		return fmt.Errorf("json.Marshal(rephrasedJob) err: %v", err)
	}
	_, err = jobJson.Write(rephrasedJobByte)
	if err != nil {
		return fmt.Errorf("jobJson.Write() to fn %s err: %v", fn, err)
	}
	return err
}

func WriteOrderFormInfoToFile(id string, rephrasedOrderFormInfo []dao.OrderFormInfo, fn string) error {
	// write to file, fn - date_jobId
	jobJson, _ := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE, 0777)
	defer jobJson.Close()
	rephrasedOrderFormInfoByte, err := json.Marshal(rephrasedOrderFormInfo)
	if err != nil {
		return fmt.Errorf("json.Marshal(rephrasedOrderFormInfo) err: %v", err)
	}
	// fmt.Println("rephrasedOrderFormInfoByte:", string(rephrasedOrderFormInfoByte))
	_, err = jobJson.Write(rephrasedOrderFormInfoByte)
	if err != nil {
		return fmt.Errorf("jobJson.Write() to fn %s err: %v", fn, err)
	}
	return err
}
