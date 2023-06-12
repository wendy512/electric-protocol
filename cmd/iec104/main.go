package main

import (
	"fmt"
	"github.com/thinkgos/go-iecp5/asdu"
	"github.com/thinkgos/go-iecp5/cs104"
	"log"
	"time"
)

type IEC104ClientContext struct{}

// 总召唤响应
func (ctx *IEC104ClientContext) InterrogationHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	fmt.Println("InterrogationHandler")
	fmt.Println("InterrogationHandler RX ", rxAsdu.String())

	if rxAsdu.Coa.Cause == asdu.ActivationTerm {
		reply := rxAsdu.Reply(asdu.ActivationTerm, 1)
		conn.Send(reply)
	}
	return nil
}

func (ctx *IEC104ClientContext) CounterInterrogationHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	fmt.Println("CounterInterrogationHandler")
	return nil
}
func (ctx *IEC104ClientContext) ReadHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	fmt.Println("ReadHandler")
	return nil
}

func (ctx *IEC104ClientContext) TestCommandHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	fmt.Println("TestCommandHandler")
	return nil
}

func (ctx *IEC104ClientContext) ClockSyncHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	fmt.Println("ClockSyncHandler")
	return nil
}

func (ctx *IEC104ClientContext) ResetProcessHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	fmt.Println("ResetProcessHandler")
	return nil
}

func (ctx *IEC104ClientContext) DelayAcquisitionHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	fmt.Println("DelayAcquisitionHandler")
	return nil
}

func (ctx *IEC104ClientContext) ASDUHandler(conn asdu.Connect, rxAsdu *asdu.ASDU) error {
	fmt.Println("ASDUHandler")
	fmt.Println("ASDUHandler RX ", rxAsdu.String())

	//fmt.Println("rxAsdu params ", rxAsdu.Identifier.Variable)
	//fmt.Printf("ASDU params %#v\n", rxAsdu.Params)
	params := rxAsdu.Params
	identifier := rxAsdu.Identifier
	fmt.Printf("Params --> CauseSize: %d, OriginAddr: %d ,CommonAddrSize: %d,InfoObjAddrSize: %d, InfoObjTimeZone: %s\n", params.CauseSize, params.OrigAddress, params.CommonAddrSize, params.InfoObjAddrSize, params.InfoObjTimeZone.String())
	fmt.Printf("Identifier --> Type: %d, Variable: {Number: %d, IsSequence: %t}, Coa: {IsTest: %t,IsNegative: %t,Cause: %d}, OriginAddr: %d, CommonAddr: %d\n", identifier.Type, identifier.Variable.Number, identifier.Variable.IsSequence, identifier.Coa.IsTest, identifier.Coa.IsNegative, identifier.Coa.Cause, identifier.OrigAddr, identifier.CommonAddr)

	/*parameterFloat := rxAsdu.GetParameterFloat()
	fmt.Printf("parameterFloat: %#v\n", parameterFloat)

	singlePoints := rxAsdu.GetSinglePoint()
	for i, point := range singlePoints {
		fmt.Printf("singlePoint %d, info: %#v\n", i, point)
	}*/

	rxAsdu.DecodeFloat32()
	fmt.Printf("infoObjAddr: %d\n", rxAsdu.DecodeInfoObjAddr())
	fmt.Printf("infoObjValue: %d\n", rxAsdu.DecodeUint16())
	return nil
}

func main() {
	ctx := &IEC104ClientContext{}
	opts := cs104.NewOption()
	opts.AddRemoteServer("tcp://127.0.0.1:2404")

	client := cs104.NewClient(ctx, opts)
	client.LogMode(true)
	client.SetOnConnectHandler(func(c *cs104.Client) {
		log.Println("Start startdt cmd")
		client.SendStartDt()

		go func() {
			time.Sleep(time.Second * 5)
			log.Println("Send interrogation cmd")
			coa := asdu.CauseOfTransmission{false, false, asdu.Activation}
			if err := client.InterrogationCmd(coa, 1, asdu.QOIStation); err != nil {
				log.Println(err)
			}
		}()
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
