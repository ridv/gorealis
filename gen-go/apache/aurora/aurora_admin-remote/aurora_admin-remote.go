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
	fmt.Fprintln(os.Stderr, "  Response triggerExplicitTaskReconciliation(ExplicitReconciliationSettings settings)")
	fmt.Fprintln(os.Stderr, "  Response triggerImplicitTaskReconciliation()")
	fmt.Fprintln(os.Stderr, "  Response pruneTasks(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response createJob(JobConfiguration description)")
	fmt.Fprintln(os.Stderr, "  Response scheduleCronJob(JobConfiguration description)")
	fmt.Fprintln(os.Stderr, "  Response descheduleCronJob(JobKey job)")
	fmt.Fprintln(os.Stderr, "  Response startCronJob(JobKey job)")
	fmt.Fprintln(os.Stderr, "  Response restartShards(JobKey job,  shardIds)")
	fmt.Fprintln(os.Stderr, "  Response killTasks(JobKey job,  instances, string message)")
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
		arg353 := flag.Arg(2)
		mbTrans354 := thrift.NewTMemoryBufferLen(len(arg353))
		defer mbTrans354.Close()
		_, err355 := mbTrans354.WriteString(arg353)
		if err355 != nil {
			Usage()
			return
		}
		factory356 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt357 := factory356.GetProtocol(mbTrans354)
		argvalue1 := aurora.NewResourceAggregate()
		err358 := argvalue1.Read(jsProt357)
		if err358 != nil {
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
		arg361 := flag.Arg(1)
		mbTrans362 := thrift.NewTMemoryBufferLen(len(arg361))
		defer mbTrans362.Close()
		_, err363 := mbTrans362.WriteString(arg361)
		if err363 != nil {
			Usage()
			return
		}
		factory364 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt365 := factory364.GetProtocol(mbTrans362)
		argvalue0 := aurora.NewTaskQuery()
		err366 := argvalue0.Read(jsProt365)
		if err366 != nil {
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
		arg367 := flag.Arg(1)
		mbTrans368 := thrift.NewTMemoryBufferLen(len(arg367))
		defer mbTrans368.Close()
		_, err369 := mbTrans368.WriteString(arg367)
		if err369 != nil {
			Usage()
			return
		}
		factory370 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt371 := factory370.GetProtocol(mbTrans368)
		argvalue0 := aurora.NewTaskQuery()
		err372 := argvalue0.Read(jsProt371)
		if err372 != nil {
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
		arg373 := flag.Arg(1)
		mbTrans374 := thrift.NewTMemoryBufferLen(len(arg373))
		defer mbTrans374.Close()
		_, err375 := mbTrans374.WriteString(arg373)
		if err375 != nil {
			Usage()
			return
		}
		factory376 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt377 := factory376.GetProtocol(mbTrans374)
		argvalue0 := aurora.NewHosts()
		err378 := argvalue0.Read(jsProt377)
		if err378 != nil {
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
		arg379 := flag.Arg(1)
		mbTrans380 := thrift.NewTMemoryBufferLen(len(arg379))
		defer mbTrans380.Close()
		_, err381 := mbTrans380.WriteString(arg379)
		if err381 != nil {
			Usage()
			return
		}
		factory382 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt383 := factory382.GetProtocol(mbTrans380)
		argvalue0 := aurora.NewHosts()
		err384 := argvalue0.Read(jsProt383)
		if err384 != nil {
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
		arg385 := flag.Arg(1)
		mbTrans386 := thrift.NewTMemoryBufferLen(len(arg385))
		defer mbTrans386.Close()
		_, err387 := mbTrans386.WriteString(arg385)
		if err387 != nil {
			Usage()
			return
		}
		factory388 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt389 := factory388.GetProtocol(mbTrans386)
		argvalue0 := aurora.NewHosts()
		err390 := argvalue0.Read(jsProt389)
		if err390 != nil {
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
		arg391 := flag.Arg(1)
		mbTrans392 := thrift.NewTMemoryBufferLen(len(arg391))
		defer mbTrans392.Close()
		_, err393 := mbTrans392.WriteString(arg391)
		if err393 != nil {
			Usage()
			return
		}
		factory394 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt395 := factory394.GetProtocol(mbTrans392)
		argvalue0 := aurora.NewHosts()
		err396 := argvalue0.Read(jsProt395)
		if err396 != nil {
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
	case "triggerExplicitTaskReconciliation":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TriggerExplicitTaskReconciliation requires 1 args")
			flag.Usage()
		}
		arg397 := flag.Arg(1)
		mbTrans398 := thrift.NewTMemoryBufferLen(len(arg397))
		defer mbTrans398.Close()
		_, err399 := mbTrans398.WriteString(arg397)
		if err399 != nil {
			Usage()
			return
		}
		factory400 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt401 := factory400.GetProtocol(mbTrans398)
		argvalue0 := aurora.NewExplicitReconciliationSettings()
		err402 := argvalue0.Read(jsProt401)
		if err402 != nil {
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
	case "pruneTasks":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "PruneTasks requires 1 args")
			flag.Usage()
		}
		arg403 := flag.Arg(1)
		mbTrans404 := thrift.NewTMemoryBufferLen(len(arg403))
		defer mbTrans404.Close()
		_, err405 := mbTrans404.WriteString(arg403)
		if err405 != nil {
			Usage()
			return
		}
		factory406 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt407 := factory406.GetProtocol(mbTrans404)
		argvalue0 := aurora.NewTaskQuery()
		err408 := argvalue0.Read(jsProt407)
		if err408 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.PruneTasks(value0))
		fmt.Print("\n")
		break
	case "createJob":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CreateJob requires 1 args")
			flag.Usage()
		}
		arg409 := flag.Arg(1)
		mbTrans410 := thrift.NewTMemoryBufferLen(len(arg409))
		defer mbTrans410.Close()
		_, err411 := mbTrans410.WriteString(arg409)
		if err411 != nil {
			Usage()
			return
		}
		factory412 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt413 := factory412.GetProtocol(mbTrans410)
		argvalue0 := aurora.NewJobConfiguration()
		err414 := argvalue0.Read(jsProt413)
		if err414 != nil {
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
		arg415 := flag.Arg(1)
		mbTrans416 := thrift.NewTMemoryBufferLen(len(arg415))
		defer mbTrans416.Close()
		_, err417 := mbTrans416.WriteString(arg415)
		if err417 != nil {
			Usage()
			return
		}
		factory418 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt419 := factory418.GetProtocol(mbTrans416)
		argvalue0 := aurora.NewJobConfiguration()
		err420 := argvalue0.Read(jsProt419)
		if err420 != nil {
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
		arg421 := flag.Arg(1)
		mbTrans422 := thrift.NewTMemoryBufferLen(len(arg421))
		defer mbTrans422.Close()
		_, err423 := mbTrans422.WriteString(arg421)
		if err423 != nil {
			Usage()
			return
		}
		factory424 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt425 := factory424.GetProtocol(mbTrans422)
		argvalue0 := aurora.NewJobKey()
		err426 := argvalue0.Read(jsProt425)
		if err426 != nil {
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
		arg427 := flag.Arg(1)
		mbTrans428 := thrift.NewTMemoryBufferLen(len(arg427))
		defer mbTrans428.Close()
		_, err429 := mbTrans428.WriteString(arg427)
		if err429 != nil {
			Usage()
			return
		}
		factory430 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt431 := factory430.GetProtocol(mbTrans428)
		argvalue0 := aurora.NewJobKey()
		err432 := argvalue0.Read(jsProt431)
		if err432 != nil {
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
		arg433 := flag.Arg(1)
		mbTrans434 := thrift.NewTMemoryBufferLen(len(arg433))
		defer mbTrans434.Close()
		_, err435 := mbTrans434.WriteString(arg433)
		if err435 != nil {
			Usage()
			return
		}
		factory436 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt437 := factory436.GetProtocol(mbTrans434)
		argvalue0 := aurora.NewJobKey()
		err438 := argvalue0.Read(jsProt437)
		if err438 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg439 := flag.Arg(2)
		mbTrans440 := thrift.NewTMemoryBufferLen(len(arg439))
		defer mbTrans440.Close()
		_, err441 := mbTrans440.WriteString(arg439)
		if err441 != nil {
			Usage()
			return
		}
		factory442 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt443 := factory442.GetProtocol(mbTrans440)
		containerStruct1 := aurora.NewAuroraAdminRestartShardsArgs()
		err444 := containerStruct1.ReadField2(jsProt443)
		if err444 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.ShardIds
		value1 := argvalue1
		fmt.Print(client.RestartShards(value0, value1))
		fmt.Print("\n")
		break
	case "killTasks":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "KillTasks requires 3 args")
			flag.Usage()
		}
		arg445 := flag.Arg(1)
		mbTrans446 := thrift.NewTMemoryBufferLen(len(arg445))
		defer mbTrans446.Close()
		_, err447 := mbTrans446.WriteString(arg445)
		if err447 != nil {
			Usage()
			return
		}
		factory448 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt449 := factory448.GetProtocol(mbTrans446)
		argvalue0 := aurora.NewJobKey()
		err450 := argvalue0.Read(jsProt449)
		if err450 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg451 := flag.Arg(2)
		mbTrans452 := thrift.NewTMemoryBufferLen(len(arg451))
		defer mbTrans452.Close()
		_, err453 := mbTrans452.WriteString(arg451)
		if err453 != nil {
			Usage()
			return
		}
		factory454 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt455 := factory454.GetProtocol(mbTrans452)
		containerStruct1 := aurora.NewAuroraAdminKillTasksArgs()
		err456 := containerStruct1.ReadField2(jsProt455)
		if err456 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Instances
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.KillTasks(value0, value1, value2))
		fmt.Print("\n")
		break
	case "addInstances":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "AddInstances requires 2 args")
			flag.Usage()
		}
		arg458 := flag.Arg(1)
		mbTrans459 := thrift.NewTMemoryBufferLen(len(arg458))
		defer mbTrans459.Close()
		_, err460 := mbTrans459.WriteString(arg458)
		if err460 != nil {
			Usage()
			return
		}
		factory461 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt462 := factory461.GetProtocol(mbTrans459)
		argvalue0 := aurora.NewInstanceKey()
		err463 := argvalue0.Read(jsProt462)
		if err463 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err464 := (strconv.Atoi(flag.Arg(2)))
		if err464 != nil {
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
		arg465 := flag.Arg(1)
		mbTrans466 := thrift.NewTMemoryBufferLen(len(arg465))
		defer mbTrans466.Close()
		_, err467 := mbTrans466.WriteString(arg465)
		if err467 != nil {
			Usage()
			return
		}
		factory468 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt469 := factory468.GetProtocol(mbTrans466)
		argvalue0 := aurora.NewJobConfiguration()
		err470 := argvalue0.Read(jsProt469)
		if err470 != nil {
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
		arg471 := flag.Arg(1)
		mbTrans472 := thrift.NewTMemoryBufferLen(len(arg471))
		defer mbTrans472.Close()
		_, err473 := mbTrans472.WriteString(arg471)
		if err473 != nil {
			Usage()
			return
		}
		factory474 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt475 := factory474.GetProtocol(mbTrans472)
		argvalue0 := aurora.NewJobUpdateRequest()
		err476 := argvalue0.Read(jsProt475)
		if err476 != nil {
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
		arg478 := flag.Arg(1)
		mbTrans479 := thrift.NewTMemoryBufferLen(len(arg478))
		defer mbTrans479.Close()
		_, err480 := mbTrans479.WriteString(arg478)
		if err480 != nil {
			Usage()
			return
		}
		factory481 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt482 := factory481.GetProtocol(mbTrans479)
		argvalue0 := aurora.NewJobUpdateKey()
		err483 := argvalue0.Read(jsProt482)
		if err483 != nil {
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
		arg485 := flag.Arg(1)
		mbTrans486 := thrift.NewTMemoryBufferLen(len(arg485))
		defer mbTrans486.Close()
		_, err487 := mbTrans486.WriteString(arg485)
		if err487 != nil {
			Usage()
			return
		}
		factory488 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt489 := factory488.GetProtocol(mbTrans486)
		argvalue0 := aurora.NewJobUpdateKey()
		err490 := argvalue0.Read(jsProt489)
		if err490 != nil {
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
		arg492 := flag.Arg(1)
		mbTrans493 := thrift.NewTMemoryBufferLen(len(arg492))
		defer mbTrans493.Close()
		_, err494 := mbTrans493.WriteString(arg492)
		if err494 != nil {
			Usage()
			return
		}
		factory495 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt496 := factory495.GetProtocol(mbTrans493)
		argvalue0 := aurora.NewJobUpdateKey()
		err497 := argvalue0.Read(jsProt496)
		if err497 != nil {
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
		arg499 := flag.Arg(1)
		mbTrans500 := thrift.NewTMemoryBufferLen(len(arg499))
		defer mbTrans500.Close()
		_, err501 := mbTrans500.WriteString(arg499)
		if err501 != nil {
			Usage()
			return
		}
		factory502 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt503 := factory502.GetProtocol(mbTrans500)
		argvalue0 := aurora.NewJobUpdateKey()
		err504 := argvalue0.Read(jsProt503)
		if err504 != nil {
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
		arg506 := flag.Arg(1)
		mbTrans507 := thrift.NewTMemoryBufferLen(len(arg506))
		defer mbTrans507.Close()
		_, err508 := mbTrans507.WriteString(arg506)
		if err508 != nil {
			Usage()
			return
		}
		factory509 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt510 := factory509.GetProtocol(mbTrans507)
		argvalue0 := aurora.NewJobUpdateKey()
		err511 := argvalue0.Read(jsProt510)
		if err511 != nil {
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
		arg513 := flag.Arg(1)
		mbTrans514 := thrift.NewTMemoryBufferLen(len(arg513))
		defer mbTrans514.Close()
		_, err515 := mbTrans514.WriteString(arg513)
		if err515 != nil {
			Usage()
			return
		}
		factory516 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt517 := factory516.GetProtocol(mbTrans514)
		argvalue0 := aurora.NewTaskQuery()
		err518 := argvalue0.Read(jsProt517)
		if err518 != nil {
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
		arg519 := flag.Arg(1)
		mbTrans520 := thrift.NewTMemoryBufferLen(len(arg519))
		defer mbTrans520.Close()
		_, err521 := mbTrans520.WriteString(arg519)
		if err521 != nil {
			Usage()
			return
		}
		factory522 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt523 := factory522.GetProtocol(mbTrans520)
		argvalue0 := aurora.NewTaskQuery()
		err524 := argvalue0.Read(jsProt523)
		if err524 != nil {
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
		arg525 := flag.Arg(1)
		mbTrans526 := thrift.NewTMemoryBufferLen(len(arg525))
		defer mbTrans526.Close()
		_, err527 := mbTrans526.WriteString(arg525)
		if err527 != nil {
			Usage()
			return
		}
		factory528 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt529 := factory528.GetProtocol(mbTrans526)
		argvalue0 := aurora.NewTaskQuery()
		err530 := argvalue0.Read(jsProt529)
		if err530 != nil {
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
		arg531 := flag.Arg(1)
		mbTrans532 := thrift.NewTMemoryBufferLen(len(arg531))
		defer mbTrans532.Close()
		_, err533 := mbTrans532.WriteString(arg531)
		if err533 != nil {
			Usage()
			return
		}
		factory534 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt535 := factory534.GetProtocol(mbTrans532)
		argvalue0 := aurora.NewJobKey()
		err536 := argvalue0.Read(jsProt535)
		if err536 != nil {
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
		arg539 := flag.Arg(1)
		mbTrans540 := thrift.NewTMemoryBufferLen(len(arg539))
		defer mbTrans540.Close()
		_, err541 := mbTrans540.WriteString(arg539)
		if err541 != nil {
			Usage()
			return
		}
		factory542 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt543 := factory542.GetProtocol(mbTrans540)
		argvalue0 := aurora.NewJobConfiguration()
		err544 := argvalue0.Read(jsProt543)
		if err544 != nil {
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
		arg545 := flag.Arg(1)
		mbTrans546 := thrift.NewTMemoryBufferLen(len(arg545))
		defer mbTrans546.Close()
		_, err547 := mbTrans546.WriteString(arg545)
		if err547 != nil {
			Usage()
			return
		}
		factory548 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt549 := factory548.GetProtocol(mbTrans546)
		argvalue0 := aurora.NewJobUpdateQuery()
		err550 := argvalue0.Read(jsProt549)
		if err550 != nil {
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
		arg551 := flag.Arg(1)
		mbTrans552 := thrift.NewTMemoryBufferLen(len(arg551))
		defer mbTrans552.Close()
		_, err553 := mbTrans552.WriteString(arg551)
		if err553 != nil {
			Usage()
			return
		}
		factory554 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt555 := factory554.GetProtocol(mbTrans552)
		argvalue0 := aurora.NewJobUpdateQuery()
		err556 := argvalue0.Read(jsProt555)
		if err556 != nil {
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
		arg557 := flag.Arg(1)
		mbTrans558 := thrift.NewTMemoryBufferLen(len(arg557))
		defer mbTrans558.Close()
		_, err559 := mbTrans558.WriteString(arg557)
		if err559 != nil {
			Usage()
			return
		}
		factory560 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt561 := factory560.GetProtocol(mbTrans558)
		argvalue0 := aurora.NewJobUpdateRequest()
		err562 := argvalue0.Read(jsProt561)
		if err562 != nil {
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
