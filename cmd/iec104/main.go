package main

import (
	"fmt"
	"github.com/thinkgos/go-iecp5/asdu"
	"github.com/thinkgos/go-iecp5/cs104"
	"time"
)

type IEC104ClientContext struct{}

func (ctx *IEC104ClientContext) InterrogationHandler(conn asdu.Connect, asdu *asdu.ASDU) error {
	fmt.Println("InterrogationHandler")
	return nil
}

func (ctx *IEC104ClientContext) CounterInterrogationHandler(conn asdu.Connect, asdu *asdu.ASDU) error {
	fmt.Println("CounterInterrogationHandler")
	return nil
}
func (ctx *IEC104ClientContext) ReadHandler(conn asdu.Connect, asdu *asdu.ASDU) error {
	fmt.Println("ReadHandler")
	return nil
}

func (ctx *IEC104ClientContext) TestCommandHandler(conn asdu.Connect, asdu *asdu.ASDU) error {
	fmt.Println("TestCommandHandler")
	return nil
}

func (ctx *IEC104ClientContext) ClockSyncHandler(conn asdu.Connect, asdu *asdu.ASDU) error {
	fmt.Println("ClockSyncHandler")
	return nil
}

func (ctx *IEC104ClientContext) ResetProcessHandler(conn asdu.Connect, asdu *asdu.ASDU) error {
	fmt.Println("ResetProcessHandler")
	return nil
}

func (ctx *IEC104ClientContext) DelayAcquisitionHandler(conn asdu.Connect, asdu *asdu.ASDU) error {
	fmt.Println("DelayAcquisitionHandler")
	return nil
}

func (ctx *IEC104ClientContext) ASDUHandler(conn asdu.Connect, asdu *asdu.ASDU) error {
	fmt.Println("ASDUHandler")
	return nil
}

func main() {
	ctx := &IEC104ClientContext{}
	opts := cs104.NewOption()
	opts.AddRemoteServer("tcp://127.0.0.1:2404")

	client := cs104.NewClient(ctx, opts)
	client.LogMode(true)
	client.SetOnConnectHandler(func(c *cs104.Client) {
		fmt.Println("Start SendStartDt")
		client.SendStartDt()
		fmt.Println("End SendStartDt")

		coa := asdu.CauseOfTransmission{false, false, asdu.Activation}
		//client.InterrogationCmd(coa, 1, asdu.QOIUnused)
		//asdu.ASDU{}
		asdu.parse
		client.Send()
	})

	defer client.Close()
	err := client.Start()
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(2 * time.Second)
	}
}
