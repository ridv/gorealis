// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"apache/aurora"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Response setQuota(string ownerRole, ResourceAggregate quota)")
	fmt.Fprintln(os.Stderr, "  Response forceTaskState(string taskId, ScheduleStatus status)")
	fmt.Fprintln(os.Stderr, "  Response performBackup()")
	fmt.Fprintln(os.Stderr, "  Response listBackups()")
	fmt.Fprintln(os.Stderr, "  Response stageRecovery(string backupId)")
	fmt.Fprintln(os.Stderr, "  Response queryRecovery(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response deleteRecoveryTasks(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response commitRecovery()")
	fmt.Fprintln(os.Stderr, "  Response unloadRecovery()")
	fmt.Fprintln(os.Stderr, "  Response startMaintenance(Hosts hosts)")
	fmt.Fprintln(os.Stderr, "  Response drainHosts(Hosts hosts)")
	fmt.Fprintln(os.Stderr, "  Response maintenanceStatus(Hosts hosts)")
	fmt.Fprintln(os.Stderr, "  Response endMaintenance(Hosts hosts)")
	fmt.Fprintln(os.Stderr, "  Response snapshot()")
	fmt.Fprintln(os.Stderr, "  Response rewriteConfigs(RewriteConfigsRequest request)")
	fmt.Fprintln(os.Stderr, "  Response triggerExplicitTaskReconciliation(ExplicitReconciliationSettings settings)")
	fmt.Fprintln(os.Stderr, "  Response triggerImplicitTaskReconciliation()")
	fmt.Fprintln(os.Stderr, "  Response createJob(JobConfiguration description)")
	fmt.Fprintln(os.Stderr, "  Response scheduleCronJob(JobConfiguration description)")
	fmt.Fprintln(os.Stderr, "  Response descheduleCronJob(JobKey job)")
	fmt.Fprintln(os.Stderr, "  Response startCronJob(JobKey job)")
	fmt.Fprintln(os.Stderr, "  Response restartShards(JobKey job,  shardIds)")
	fmt.Fprintln(os.Stderr, "  Response killTasks(JobKey job,  instances)")
	fmt.Fprintln(os.Stderr, "  Response addInstances(InstanceKey key, i32 count)")
	fmt.Fprintln(os.Stderr, "  Response replaceCronTemplate(JobConfiguration config)")
	fmt.Fprintln(os.Stderr, "  Response startJobUpdate(JobUpdateRequest request, string message)")
	fmt.Fprintln(os.Stderr, "  Response pauseJobUpdate(JobUpdateKey key, string message)")
	fmt.Fprintln(os.Stderr, "  Response resumeJobUpdate(JobUpdateKey key, string message)")
	fmt.Fprintln(os.Stderr, "  Response abortJobUpdate(JobUpdateKey key, string message)")
	fmt.Fprintln(os.Stderr, "  Response rollbackJobUpdate(JobUpdateKey key, string message)")
	fmt.Fprintln(os.Stderr, "  Response pulseJobUpdate(JobUpdateKey key)")
	fmt.Fprintln(os.Stderr, "  Response getRoleSummary()")
	fmt.Fprintln(os.Stderr, "  Response getJobSummary(string role)")
	fmt.Fprintln(os.Stderr, "  Response getTasksStatus(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response getTasksWithoutConfigs(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response getPendingReason(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response getConfigSummary(JobKey job)")
	fmt.Fprintln(os.Stderr, "  Response getJobs(string ownerRole)")
	fmt.Fprintln(os.Stderr, "  Response getQuota(string ownerRole)")
	fmt.Fprintln(os.Stderr, "  Response populateJobConfig(JobConfiguration description)")
	fmt.Fprintln(os.Stderr, "  Response getJobUpdateSummaries(JobUpdateQuery jobUpdateQuery)")
	fmt.Fprintln(os.Stderr, "  Response getJobUpdateDetails(JobUpdateQuery query)")
	fmt.Fprintln(os.Stderr, "  Response getJobUpdateDiff(JobUpdateRequest request)")
	fmt.Fprintln(os.Stderr, "  Response getTierConfigs()")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := aurora.NewAuroraAdminClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "setQuota":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SetQuota requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg352 := flag.Arg(2)
		mbTrans353 := thrift.NewTMemoryBufferLen(len(arg352))
		defer mbTrans353.Close()
		_, err354 := mbTrans353.WriteString(arg352)
		if err354 != nil {
			Usage()
			return
		}
		factory355 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt356 := factory355.GetProtocol(mbTrans353)
		argvalue1 := aurora.NewResourceAggregate()
		err357 := argvalue1.Read(jsProt356)
		if err357 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.SetQuota(value0, value1))
		fmt.Print("\n")
		break
	case "forceTaskState":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ForceTaskState requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		tmp1, err := (strconv.Atoi(flag.Arg(2)))
		if err != nil {
			Usage()
			return
		}
		argvalue1 := aurora.ScheduleStatus(tmp1)
		value1 := argvalue1
		fmt.Print(client.ForceTaskState(value0, value1))
		fmt.Print("\n")
		break
	case "performBackup":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "PerformBackup requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.PerformBackup())
		fmt.Print("\n")
		break
	case "listBackups":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "ListBackups requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.ListBackups())
		fmt.Print("\n")
		break
	case "stageRecovery":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "StageRecovery requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.StageRecovery(value0))
		fmt.Print("\n")
		break
	case "queryRecovery":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "QueryRecovery requires 1 args")
			flag.Usage()
		}
		arg360 := flag.Arg(1)
		mbTrans361 := thrift.NewTMemoryBufferLen(len(arg360))
		defer mbTrans361.Close()
		_, err362 := mbTrans361.WriteString(arg360)
		if err362 != nil {
			Usage()
			return
		}
		factory363 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt364 := factory363.GetProtocol(mbTrans361)
		argvalue0 := aurora.NewTaskQuery()
		err365 := argvalue0.Read(jsProt364)
		if err365 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.QueryRecovery(value0))
		fmt.Print("\n")
		break
	case "deleteRecoveryTasks":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DeleteRecoveryTasks requires 1 args")
			flag.Usage()
		}
		arg366 := flag.Arg(1)
		mbTrans367 := thrift.NewTMemoryBufferLen(len(arg366))
		defer mbTrans367.Close()
		_, err368 := mbTrans367.WriteString(arg366)
		if err368 != nil {
			Usage()
			return
		}
		factory369 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt370 := factory369.GetProtocol(mbTrans367)
		argvalue0 := aurora.NewTaskQuery()
		err371 := argvalue0.Read(jsProt370)
		if err371 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DeleteRecoveryTasks(value0))
		fmt.Print("\n")
		break
	case "commitRecovery":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "CommitRecovery requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.CommitRecovery())
		fmt.Print("\n")
		break
	case "unloadRecovery":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "UnloadRecovery requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.UnloadRecovery())
		fmt.Print("\n")
		break
	case "startMaintenance":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "StartMaintenance requires 1 args")
			flag.Usage()
		}
		arg372 := flag.Arg(1)
		mbTrans373 := thrift.NewTMemoryBufferLen(len(arg372))
		defer mbTrans373.Close()
		_, err374 := mbTrans373.WriteString(arg372)
		if err374 != nil {
			Usage()
			return
		}
		factory375 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt376 := factory375.GetProtocol(mbTrans373)
		argvalue0 := aurora.NewHosts()
		err377 := argvalue0.Read(jsProt376)
		if err377 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.StartMaintenance(value0))
		fmt.Print("\n")
		break
	case "drainHosts":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DrainHosts requires 1 args")
			flag.Usage()
		}
		arg378 := flag.Arg(1)
		mbTrans379 := thrift.NewTMemoryBufferLen(len(arg378))
		defer mbTrans379.Close()
		_, err380 := mbTrans379.WriteString(arg378)
		if err380 != nil {
			Usage()
			return
		}
		factory381 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt382 := factory381.GetProtocol(mbTrans379)
		argvalue0 := aurora.NewHosts()
		err383 := argvalue0.Read(jsProt382)
		if err383 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DrainHosts(value0))
		fmt.Print("\n")
		break
	case "maintenanceStatus":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "MaintenanceStatus requires 1 args")
			flag.Usage()
		}
		arg384 := flag.Arg(1)
		mbTrans385 := thrift.NewTMemoryBufferLen(len(arg384))
		defer mbTrans385.Close()
		_, err386 := mbTrans385.WriteString(arg384)
		if err386 != nil {
			Usage()
			return
		}
		factory387 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt388 := factory387.GetProtocol(mbTrans385)
		argvalue0 := aurora.NewHosts()
		err389 := argvalue0.Read(jsProt388)
		if err389 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.MaintenanceStatus(value0))
		fmt.Print("\n")
		break
	case "endMaintenance":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "EndMaintenance requires 1 args")
			flag.Usage()
		}
		arg390 := flag.Arg(1)
		mbTrans391 := thrift.NewTMemoryBufferLen(len(arg390))
		defer mbTrans391.Close()
		_, err392 := mbTrans391.WriteString(arg390)
		if err392 != nil {
			Usage()
			return
		}
		factory393 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt394 := factory393.GetProtocol(mbTrans391)
		argvalue0 := aurora.NewHosts()
		err395 := argvalue0.Read(jsProt394)
		if err395 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.EndMaintenance(value0))
		fmt.Print("\n")
		break
	case "snapshot":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Snapshot requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Snapshot())
		fmt.Print("\n")
		break
	case "rewriteConfigs":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RewriteConfigs requires 1 args")
			flag.Usage()
		}
		arg396 := flag.Arg(1)
		mbTrans397 := thrift.NewTMemoryBufferLen(len(arg396))
		defer mbTrans397.Close()
		_, err398 := mbTrans397.WriteString(arg396)
		if err398 != nil {
			Usage()
			return
		}
		factory399 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt400 := factory399.GetProtocol(mbTrans397)
		argvalue0 := aurora.NewRewriteConfigsRequest()
		err401 := argvalue0.Read(jsProt400)
		if err401 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.RewriteConfigs(value0))
		fmt.Print("\n")
		break
	case "triggerExplicitTaskReconciliation":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TriggerExplicitTaskReconciliation requires 1 args")
			flag.Usage()
		}
		arg402 := flag.Arg(1)
		mbTrans403 := thrift.NewTMemoryBufferLen(len(arg402))
		defer mbTrans403.Close()
		_, err404 := mbTrans403.WriteString(arg402)
		if err404 != nil {
			Usage()
			return
		}
		factory405 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt406 := factory405.GetProtocol(mbTrans403)
		argvalue0 := aurora.NewExplicitReconciliationSettings()
		err407 := argvalue0.Read(jsProt406)
		if err407 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TriggerExplicitTaskReconciliation(value0))
		fmt.Print("\n")
		break
	case "triggerImplicitTaskReconciliation":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "TriggerImplicitTaskReconciliation requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.TriggerImplicitTaskReconciliation())
		fmt.Print("\n")
		break
	case "createJob":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CreateJob requires 1 args")
			flag.Usage()
		}
		arg408 := flag.Arg(1)
		mbTrans409 := thrift.NewTMemoryBufferLen(len(arg408))
		defer mbTrans409.Close()
		_, err410 := mbTrans409.WriteString(arg408)
		if err410 != nil {
			Usage()
			return
		}
		factory411 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt412 := factory411.GetProtocol(mbTrans409)
		argvalue0 := aurora.NewJobConfiguration()
		err413 := argvalue0.Read(jsProt412)
		if err413 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.CreateJob(value0))
		fmt.Print("\n")
		break
	case "scheduleCronJob":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ScheduleCronJob requires 1 args")
			flag.Usage()
		}
		arg414 := flag.Arg(1)
		mbTrans415 := thrift.NewTMemoryBufferLen(len(arg414))
		defer mbTrans415.Close()
		_, err416 := mbTrans415.WriteString(arg414)
		if err416 != nil {
			Usage()
			return
		}
		factory417 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt418 := factory417.GetProtocol(mbTrans415)
		argvalue0 := aurora.NewJobConfiguration()
		err419 := argvalue0.Read(jsProt418)
		if err419 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ScheduleCronJob(value0))
		fmt.Print("\n")
		break
	case "descheduleCronJob":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DescheduleCronJob requires 1 args")
			flag.Usage()
		}
		arg420 := flag.Arg(1)
		mbTrans421 := thrift.NewTMemoryBufferLen(len(arg420))
		defer mbTrans421.Close()
		_, err422 := mbTrans421.WriteString(arg420)
		if err422 != nil {
			Usage()
			return
		}
		factory423 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt424 := factory423.GetProtocol(mbTrans421)
		argvalue0 := aurora.NewJobKey()
		err425 := argvalue0.Read(jsProt424)
		if err425 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DescheduleCronJob(value0))
		fmt.Print("\n")
		break
	case "startCronJob":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "StartCronJob requires 1 args")
			flag.Usage()
		}
		arg426 := flag.Arg(1)
		mbTrans427 := thrift.NewTMemoryBufferLen(len(arg426))
		defer mbTrans427.Close()
		_, err428 := mbTrans427.WriteString(arg426)
		if err428 != nil {
			Usage()
			return
		}
		factory429 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt430 := factory429.GetProtocol(mbTrans427)
		argvalue0 := aurora.NewJobKey()
		err431 := argvalue0.Read(jsProt430)
		if err431 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.StartCronJob(value0))
		fmt.Print("\n")
		break
	case "restartShards":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "RestartShards requires 2 args")
			flag.Usage()
		}
		arg432 := flag.Arg(1)
		mbTrans433 := thrift.NewTMemoryBufferLen(len(arg432))
		defer mbTrans433.Close()
		_, err434 := mbTrans433.WriteString(arg432)
		if err434 != nil {
			Usage()
			return
		}
		factory435 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt436 := factory435.GetProtocol(mbTrans433)
		argvalue0 := aurora.NewJobKey()
		err437 := argvalue0.Read(jsProt436)
		if err437 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg438 := flag.Arg(2)
		mbTrans439 := thrift.NewTMemoryBufferLen(len(arg438))
		defer mbTrans439.Close()
		_, err440 := mbTrans439.WriteString(arg438)
		if err440 != nil {
			Usage()
			return
		}
		factory441 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt442 := factory441.GetProtocol(mbTrans439)
		containerStruct1 := aurora.NewAuroraAdminRestartShardsArgs()
		err443 := containerStruct1.ReadField2(jsProt442)
		if err443 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.ShardIds
		value1 := argvalue1
		fmt.Print(client.RestartShards(value0, value1))
		fmt.Print("\n")
		break
	case "killTasks":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "KillTasks requires 2 args")
			flag.Usage()
		}
		arg444 := flag.Arg(1)
		mbTrans445 := thrift.NewTMemoryBufferLen(len(arg444))
		defer mbTrans445.Close()
		_, err446 := mbTrans445.WriteString(arg444)
		if err446 != nil {
			Usage()
			return
		}
		factory447 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt448 := factory447.GetProtocol(mbTrans445)
		argvalue0 := aurora.NewJobKey()
		err449 := argvalue0.Read(jsProt448)
		if err449 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg450 := flag.Arg(2)
		mbTrans451 := thrift.NewTMemoryBufferLen(len(arg450))
		defer mbTrans451.Close()
		_, err452 := mbTrans451.WriteString(arg450)
		if err452 != nil {
			Usage()
			return
		}
		factory453 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt454 := factory453.GetProtocol(mbTrans451)
		containerStruct1 := aurora.NewAuroraAdminKillTasksArgs()
		err455 := containerStruct1.ReadField2(jsProt454)
		if err455 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Instances
		value1 := argvalue1
		fmt.Print(client.KillTasks(value0, value1))
		fmt.Print("\n")
		break
	case "addInstances":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "AddInstances requires 2 args")
			flag.Usage()
		}
		arg456 := flag.Arg(1)
		mbTrans457 := thrift.NewTMemoryBufferLen(len(arg456))
		defer mbTrans457.Close()
		_, err458 := mbTrans457.WriteString(arg456)
		if err458 != nil {
			Usage()
			return
		}
		factory459 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt460 := factory459.GetProtocol(mbTrans457)
		argvalue0 := aurora.NewInstanceKey()
		err461 := argvalue0.Read(jsProt460)
		if err461 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err462 := (strconv.Atoi(flag.Arg(2)))
		if err462 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.AddInstances(value0, value1))
		fmt.Print("\n")
		break
	case "replaceCronTemplate":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ReplaceCronTemplate requires 1 args")
			flag.Usage()
		}
		arg463 := flag.Arg(1)
		mbTrans464 := thrift.NewTMemoryBufferLen(len(arg463))
		defer mbTrans464.Close()
		_, err465 := mbTrans464.WriteString(arg463)
		if err465 != nil {
			Usage()
			return
		}
		factory466 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt467 := factory466.GetProtocol(mbTrans464)
		argvalue0 := aurora.NewJobConfiguration()
		err468 := argvalue0.Read(jsProt467)
		if err468 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ReplaceCronTemplate(value0))
		fmt.Print("\n")
		break
	case "startJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "StartJobUpdate requires 2 args")
			flag.Usage()
		}
		arg469 := flag.Arg(1)
		mbTrans470 := thrift.NewTMemoryBufferLen(len(arg469))
		defer mbTrans470.Close()
		_, err471 := mbTrans470.WriteString(arg469)
		if err471 != nil {
			Usage()
			return
		}
		factory472 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt473 := factory472.GetProtocol(mbTrans470)
		argvalue0 := aurora.NewJobUpdateRequest()
		err474 := argvalue0.Read(jsProt473)
		if err474 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.StartJobUpdate(value0, value1))
		fmt.Print("\n")
		break
	case "pauseJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "PauseJobUpdate requires 2 args")
			flag.Usage()
		}
		arg476 := flag.Arg(1)
		mbTrans477 := thrift.NewTMemoryBufferLen(len(arg476))
		defer mbTrans477.Close()
		_, err478 := mbTrans477.WriteString(arg476)
		if err478 != nil {
			Usage()
			return
		}
		factory479 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt480 := factory479.GetProtocol(mbTrans477)
		argvalue0 := aurora.NewJobUpdateKey()
		err481 := argvalue0.Read(jsProt480)
		if err481 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.PauseJobUpdate(value0, value1))
		fmt.Print("\n")
		break
	case "resumeJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ResumeJobUpdate requires 2 args")
			flag.Usage()
		}
		arg483 := flag.Arg(1)
		mbTrans484 := thrift.NewTMemoryBufferLen(len(arg483))
		defer mbTrans484.Close()
		_, err485 := mbTrans484.WriteString(arg483)
		if err485 != nil {
			Usage()
			return
		}
		factory486 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt487 := factory486.GetProtocol(mbTrans484)
		argvalue0 := aurora.NewJobUpdateKey()
		err488 := argvalue0.Read(jsProt487)
		if err488 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.ResumeJobUpdate(value0, value1))
		fmt.Print("\n")
		break
	case "abortJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "AbortJobUpdate requires 2 args")
			flag.Usage()
		}
		arg490 := flag.Arg(1)
		mbTrans491 := thrift.NewTMemoryBufferLen(len(arg490))
		defer mbTrans491.Close()
		_, err492 := mbTrans491.WriteString(arg490)
		if err492 != nil {
			Usage()
			return
		}
		factory493 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt494 := factory493.GetProtocol(mbTrans491)
		argvalue0 := aurora.NewJobUpdateKey()
		err495 := argvalue0.Read(jsProt494)
		if err495 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.AbortJobUpdate(value0, value1))
		fmt.Print("\n")
		break
	case "rollbackJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "RollbackJobUpdate requires 2 args")
			flag.Usage()
		}
		arg497 := flag.Arg(1)
		mbTrans498 := thrift.NewTMemoryBufferLen(len(arg497))
		defer mbTrans498.Close()
		_, err499 := mbTrans498.WriteString(arg497)
		if err499 != nil {
			Usage()
			return
		}
		factory500 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt501 := factory500.GetProtocol(mbTrans498)
		argvalue0 := aurora.NewJobUpdateKey()
		err502 := argvalue0.Read(jsProt501)
		if err502 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.RollbackJobUpdate(value0, value1))
		fmt.Print("\n")
		break
	case "pulseJobUpdate":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "PulseJobUpdate requires 1 args")
			flag.Usage()
		}
		arg504 := flag.Arg(1)
		mbTrans505 := thrift.NewTMemoryBufferLen(len(arg504))
		defer mbTrans505.Close()
		_, err506 := mbTrans505.WriteString(arg504)
		if err506 != nil {
			Usage()
			return
		}
		factory507 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt508 := factory507.GetProtocol(mbTrans505)
		argvalue0 := aurora.NewJobUpdateKey()
		err509 := argvalue0.Read(jsProt508)
		if err509 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.PulseJobUpdate(value0))
		fmt.Print("\n")
		break
	case "getRoleSummary":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetRoleSummary requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetRoleSummary())
		fmt.Print("\n")
		break
	case "getJobSummary":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobSummary requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetJobSummary(value0))
		fmt.Print("\n")
		break
	case "getTasksStatus":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTasksStatus requires 1 args")
			flag.Usage()
		}
		arg511 := flag.Arg(1)
		mbTrans512 := thrift.NewTMemoryBufferLen(len(arg511))
		defer mbTrans512.Close()
		_, err513 := mbTrans512.WriteString(arg511)
		if err513 != nil {
			Usage()
			return
		}
		factory514 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt515 := factory514.GetProtocol(mbTrans512)
		argvalue0 := aurora.NewTaskQuery()
		err516 := argvalue0.Read(jsProt515)
		if err516 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetTasksStatus(value0))
		fmt.Print("\n")
		break
	case "getTasksWithoutConfigs":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTasksWithoutConfigs requires 1 args")
			flag.Usage()
		}
		arg517 := flag.Arg(1)
		mbTrans518 := thrift.NewTMemoryBufferLen(len(arg517))
		defer mbTrans518.Close()
		_, err519 := mbTrans518.WriteString(arg517)
		if err519 != nil {
			Usage()
			return
		}
		factory520 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt521 := factory520.GetProtocol(mbTrans518)
		argvalue0 := aurora.NewTaskQuery()
		err522 := argvalue0.Read(jsProt521)
		if err522 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetTasksWithoutConfigs(value0))
		fmt.Print("\n")
		break
	case "getPendingReason":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPendingReason requires 1 args")
			flag.Usage()
		}
		arg523 := flag.Arg(1)
		mbTrans524 := thrift.NewTMemoryBufferLen(len(arg523))
		defer mbTrans524.Close()
		_, err525 := mbTrans524.WriteString(arg523)
		if err525 != nil {
			Usage()
			return
		}
		factory526 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt527 := factory526.GetProtocol(mbTrans524)
		argvalue0 := aurora.NewTaskQuery()
		err528 := argvalue0.Read(jsProt527)
		if err528 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetPendingReason(value0))
		fmt.Print("\n")
		break
	case "getConfigSummary":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetConfigSummary requires 1 args")
			flag.Usage()
		}
		arg529 := flag.Arg(1)
		mbTrans530 := thrift.NewTMemoryBufferLen(len(arg529))
		defer mbTrans530.Close()
		_, err531 := mbTrans530.WriteString(arg529)
		if err531 != nil {
			Usage()
			return
		}
		factory532 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt533 := factory532.GetProtocol(mbTrans530)
		argvalue0 := aurora.NewJobKey()
		err534 := argvalue0.Read(jsProt533)
		if err534 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetConfigSummary(value0))
		fmt.Print("\n")
		break
	case "getJobs":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobs requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetJobs(value0))
		fmt.Print("\n")
		break
	case "getQuota":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetQuota requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetQuota(value0))
		fmt.Print("\n")
		break
	case "populateJobConfig":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "PopulateJobConfig requires 1 args")
			flag.Usage()
		}
		arg537 := flag.Arg(1)
		mbTrans538 := thrift.NewTMemoryBufferLen(len(arg537))
		defer mbTrans538.Close()
		_, err539 := mbTrans538.WriteString(arg537)
		if err539 != nil {
			Usage()
			return
		}
		factory540 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt541 := factory540.GetProtocol(mbTrans538)
		argvalue0 := aurora.NewJobConfiguration()
		err542 := argvalue0.Read(jsProt541)
		if err542 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.PopulateJobConfig(value0))
		fmt.Print("\n")
		break
	case "getJobUpdateSummaries":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobUpdateSummaries requires 1 args")
			flag.Usage()
		}
		arg543 := flag.Arg(1)
		mbTrans544 := thrift.NewTMemoryBufferLen(len(arg543))
		defer mbTrans544.Close()
		_, err545 := mbTrans544.WriteString(arg543)
		if err545 != nil {
			Usage()
			return
		}
		factory546 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt547 := factory546.GetProtocol(mbTrans544)
		argvalue0 := aurora.NewJobUpdateQuery()
		err548 := argvalue0.Read(jsProt547)
		if err548 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetJobUpdateSummaries(value0))
		fmt.Print("\n")
		break
	case "getJobUpdateDetails":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobUpdateDetails requires 1 args")
			flag.Usage()
		}
		arg549 := flag.Arg(1)
		mbTrans550 := thrift.NewTMemoryBufferLen(len(arg549))
		defer mbTrans550.Close()
		_, err551 := mbTrans550.WriteString(arg549)
		if err551 != nil {
			Usage()
			return
		}
		factory552 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt553 := factory552.GetProtocol(mbTrans550)
		argvalue0 := aurora.NewJobUpdateQuery()
		err554 := argvalue0.Read(jsProt553)
		if err554 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetJobUpdateDetails(value0))
		fmt.Print("\n")
		break
	case "getJobUpdateDiff":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobUpdateDiff requires 1 args")
			flag.Usage()
		}
		arg555 := flag.Arg(1)
		mbTrans556 := thrift.NewTMemoryBufferLen(len(arg555))
		defer mbTrans556.Close()
		_, err557 := mbTrans556.WriteString(arg555)
		if err557 != nil {
			Usage()
			return
		}
		factory558 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt559 := factory558.GetProtocol(mbTrans556)
		argvalue0 := aurora.NewJobUpdateRequest()
		err560 := argvalue0.Read(jsProt559)
		if err560 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetJobUpdateDiff(value0))
		fmt.Print("\n")
		break
	case "getTierConfigs":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetTierConfigs requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetTierConfigs())
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
