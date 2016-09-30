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
	client := aurora.NewAuroraSchedulerManagerClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "createJob":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CreateJob requires 1 args")
			flag.Usage()
		}
		arg163 := flag.Arg(1)
		mbTrans164 := thrift.NewTMemoryBufferLen(len(arg163))
		defer mbTrans164.Close()
		_, err165 := mbTrans164.WriteString(arg163)
		if err165 != nil {
			Usage()
			return
		}
		factory166 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt167 := factory166.GetProtocol(mbTrans164)
		argvalue0 := aurora.NewJobConfiguration()
		err168 := argvalue0.Read(jsProt167)
		if err168 != nil {
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
		arg169 := flag.Arg(1)
		mbTrans170 := thrift.NewTMemoryBufferLen(len(arg169))
		defer mbTrans170.Close()
		_, err171 := mbTrans170.WriteString(arg169)
		if err171 != nil {
			Usage()
			return
		}
		factory172 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt173 := factory172.GetProtocol(mbTrans170)
		argvalue0 := aurora.NewJobConfiguration()
		err174 := argvalue0.Read(jsProt173)
		if err174 != nil {
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
		arg175 := flag.Arg(1)
		mbTrans176 := thrift.NewTMemoryBufferLen(len(arg175))
		defer mbTrans176.Close()
		_, err177 := mbTrans176.WriteString(arg175)
		if err177 != nil {
			Usage()
			return
		}
		factory178 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt179 := factory178.GetProtocol(mbTrans176)
		argvalue0 := aurora.NewJobKey()
		err180 := argvalue0.Read(jsProt179)
		if err180 != nil {
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
		arg181 := flag.Arg(1)
		mbTrans182 := thrift.NewTMemoryBufferLen(len(arg181))
		defer mbTrans182.Close()
		_, err183 := mbTrans182.WriteString(arg181)
		if err183 != nil {
			Usage()
			return
		}
		factory184 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt185 := factory184.GetProtocol(mbTrans182)
		argvalue0 := aurora.NewJobKey()
		err186 := argvalue0.Read(jsProt185)
		if err186 != nil {
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
		arg187 := flag.Arg(1)
		mbTrans188 := thrift.NewTMemoryBufferLen(len(arg187))
		defer mbTrans188.Close()
		_, err189 := mbTrans188.WriteString(arg187)
		if err189 != nil {
			Usage()
			return
		}
		factory190 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt191 := factory190.GetProtocol(mbTrans188)
		argvalue0 := aurora.NewJobKey()
		err192 := argvalue0.Read(jsProt191)
		if err192 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg193 := flag.Arg(2)
		mbTrans194 := thrift.NewTMemoryBufferLen(len(arg193))
		defer mbTrans194.Close()
		_, err195 := mbTrans194.WriteString(arg193)
		if err195 != nil {
			Usage()
			return
		}
		factory196 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt197 := factory196.GetProtocol(mbTrans194)
		containerStruct1 := aurora.NewAuroraSchedulerManagerRestartShardsArgs()
		err198 := containerStruct1.ReadField2(jsProt197)
		if err198 != nil {
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
		arg199 := flag.Arg(1)
		mbTrans200 := thrift.NewTMemoryBufferLen(len(arg199))
		defer mbTrans200.Close()
		_, err201 := mbTrans200.WriteString(arg199)
		if err201 != nil {
			Usage()
			return
		}
		factory202 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt203 := factory202.GetProtocol(mbTrans200)
		argvalue0 := aurora.NewJobKey()
		err204 := argvalue0.Read(jsProt203)
		if err204 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg205 := flag.Arg(2)
		mbTrans206 := thrift.NewTMemoryBufferLen(len(arg205))
		defer mbTrans206.Close()
		_, err207 := mbTrans206.WriteString(arg205)
		if err207 != nil {
			Usage()
			return
		}
		factory208 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt209 := factory208.GetProtocol(mbTrans206)
		containerStruct1 := aurora.NewAuroraSchedulerManagerKillTasksArgs()
		err210 := containerStruct1.ReadField2(jsProt209)
		if err210 != nil {
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
		arg211 := flag.Arg(1)
		mbTrans212 := thrift.NewTMemoryBufferLen(len(arg211))
		defer mbTrans212.Close()
		_, err213 := mbTrans212.WriteString(arg211)
		if err213 != nil {
			Usage()
			return
		}
		factory214 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt215 := factory214.GetProtocol(mbTrans212)
		argvalue0 := aurora.NewInstanceKey()
		err216 := argvalue0.Read(jsProt215)
		if err216 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err217 := (strconv.Atoi(flag.Arg(2)))
		if err217 != nil {
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
		arg218 := flag.Arg(1)
		mbTrans219 := thrift.NewTMemoryBufferLen(len(arg218))
		defer mbTrans219.Close()
		_, err220 := mbTrans219.WriteString(arg218)
		if err220 != nil {
			Usage()
			return
		}
		factory221 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt222 := factory221.GetProtocol(mbTrans219)
		argvalue0 := aurora.NewJobConfiguration()
		err223 := argvalue0.Read(jsProt222)
		if err223 != nil {
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
		arg224 := flag.Arg(1)
		mbTrans225 := thrift.NewTMemoryBufferLen(len(arg224))
		defer mbTrans225.Close()
		_, err226 := mbTrans225.WriteString(arg224)
		if err226 != nil {
			Usage()
			return
		}
		factory227 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt228 := factory227.GetProtocol(mbTrans225)
		argvalue0 := aurora.NewJobUpdateRequest()
		err229 := argvalue0.Read(jsProt228)
		if err229 != nil {
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
		arg231 := flag.Arg(1)
		mbTrans232 := thrift.NewTMemoryBufferLen(len(arg231))
		defer mbTrans232.Close()
		_, err233 := mbTrans232.WriteString(arg231)
		if err233 != nil {
			Usage()
			return
		}
		factory234 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt235 := factory234.GetProtocol(mbTrans232)
		argvalue0 := aurora.NewJobUpdateKey()
		err236 := argvalue0.Read(jsProt235)
		if err236 != nil {
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
		arg238 := flag.Arg(1)
		mbTrans239 := thrift.NewTMemoryBufferLen(len(arg238))
		defer mbTrans239.Close()
		_, err240 := mbTrans239.WriteString(arg238)
		if err240 != nil {
			Usage()
			return
		}
		factory241 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt242 := factory241.GetProtocol(mbTrans239)
		argvalue0 := aurora.NewJobUpdateKey()
		err243 := argvalue0.Read(jsProt242)
		if err243 != nil {
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
		arg245 := flag.Arg(1)
		mbTrans246 := thrift.NewTMemoryBufferLen(len(arg245))
		defer mbTrans246.Close()
		_, err247 := mbTrans246.WriteString(arg245)
		if err247 != nil {
			Usage()
			return
		}
		factory248 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt249 := factory248.GetProtocol(mbTrans246)
		argvalue0 := aurora.NewJobUpdateKey()
		err250 := argvalue0.Read(jsProt249)
		if err250 != nil {
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
		arg252 := flag.Arg(1)
		mbTrans253 := thrift.NewTMemoryBufferLen(len(arg252))
		defer mbTrans253.Close()
		_, err254 := mbTrans253.WriteString(arg252)
		if err254 != nil {
			Usage()
			return
		}
		factory255 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt256 := factory255.GetProtocol(mbTrans253)
		argvalue0 := aurora.NewJobUpdateKey()
		err257 := argvalue0.Read(jsProt256)
		if err257 != nil {
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
		arg259 := flag.Arg(1)
		mbTrans260 := thrift.NewTMemoryBufferLen(len(arg259))
		defer mbTrans260.Close()
		_, err261 := mbTrans260.WriteString(arg259)
		if err261 != nil {
			Usage()
			return
		}
		factory262 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt263 := factory262.GetProtocol(mbTrans260)
		argvalue0 := aurora.NewJobUpdateKey()
		err264 := argvalue0.Read(jsProt263)
		if err264 != nil {
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
		arg266 := flag.Arg(1)
		mbTrans267 := thrift.NewTMemoryBufferLen(len(arg266))
		defer mbTrans267.Close()
		_, err268 := mbTrans267.WriteString(arg266)
		if err268 != nil {
			Usage()
			return
		}
		factory269 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt270 := factory269.GetProtocol(mbTrans267)
		argvalue0 := aurora.NewTaskQuery()
		err271 := argvalue0.Read(jsProt270)
		if err271 != nil {
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
		arg272 := flag.Arg(1)
		mbTrans273 := thrift.NewTMemoryBufferLen(len(arg272))
		defer mbTrans273.Close()
		_, err274 := mbTrans273.WriteString(arg272)
		if err274 != nil {
			Usage()
			return
		}
		factory275 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt276 := factory275.GetProtocol(mbTrans273)
		argvalue0 := aurora.NewTaskQuery()
		err277 := argvalue0.Read(jsProt276)
		if err277 != nil {
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
		arg278 := flag.Arg(1)
		mbTrans279 := thrift.NewTMemoryBufferLen(len(arg278))
		defer mbTrans279.Close()
		_, err280 := mbTrans279.WriteString(arg278)
		if err280 != nil {
			Usage()
			return
		}
		factory281 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt282 := factory281.GetProtocol(mbTrans279)
		argvalue0 := aurora.NewTaskQuery()
		err283 := argvalue0.Read(jsProt282)
		if err283 != nil {
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
		arg284 := flag.Arg(1)
		mbTrans285 := thrift.NewTMemoryBufferLen(len(arg284))
		defer mbTrans285.Close()
		_, err286 := mbTrans285.WriteString(arg284)
		if err286 != nil {
			Usage()
			return
		}
		factory287 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt288 := factory287.GetProtocol(mbTrans285)
		argvalue0 := aurora.NewJobKey()
		err289 := argvalue0.Read(jsProt288)
		if err289 != nil {
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
		arg292 := flag.Arg(1)
		mbTrans293 := thrift.NewTMemoryBufferLen(len(arg292))
		defer mbTrans293.Close()
		_, err294 := mbTrans293.WriteString(arg292)
		if err294 != nil {
			Usage()
			return
		}
		factory295 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt296 := factory295.GetProtocol(mbTrans293)
		argvalue0 := aurora.NewJobConfiguration()
		err297 := argvalue0.Read(jsProt296)
		if err297 != nil {
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
		arg298 := flag.Arg(1)
		mbTrans299 := thrift.NewTMemoryBufferLen(len(arg298))
		defer mbTrans299.Close()
		_, err300 := mbTrans299.WriteString(arg298)
		if err300 != nil {
			Usage()
			return
		}
		factory301 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt302 := factory301.GetProtocol(mbTrans299)
		argvalue0 := aurora.NewJobUpdateQuery()
		err303 := argvalue0.Read(jsProt302)
		if err303 != nil {
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
		arg304 := flag.Arg(1)
		mbTrans305 := thrift.NewTMemoryBufferLen(len(arg304))
		defer mbTrans305.Close()
		_, err306 := mbTrans305.WriteString(arg304)
		if err306 != nil {
			Usage()
			return
		}
		factory307 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt308 := factory307.GetProtocol(mbTrans305)
		argvalue0 := aurora.NewJobUpdateQuery()
		err309 := argvalue0.Read(jsProt308)
		if err309 != nil {
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
		arg310 := flag.Arg(1)
		mbTrans311 := thrift.NewTMemoryBufferLen(len(arg310))
		defer mbTrans311.Close()
		_, err312 := mbTrans311.WriteString(arg310)
		if err312 != nil {
			Usage()
			return
		}
		factory313 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt314 := factory313.GetProtocol(mbTrans311)
		argvalue0 := aurora.NewJobUpdateRequest()
		err315 := argvalue0.Read(jsProt314)
		if err315 != nil {
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
