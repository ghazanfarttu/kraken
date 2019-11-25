/* HostDiscovery.go: This module performs CPU Thermal discovery of HPC nodes through in-band (through OS) mechanism.
 *
 * Author: Ghazanfar Ali, ghazanfar.ali@ttu.edu; Kevin Pelzel <kevinpelzel22@gmail.com>; J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2018, Triad National Security, LLC
 * See LICENSE file for details.
 */

//go:generate protoc -I ../../core/proto/include -I proto --go_out=plugins=grpc:proto proto/pxe.proto

package hostdiscovery

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/hpc/kraken/core"
	cpb "github.com/hpc/kraken/core/proto"
	thpb "github.com/hpc/kraken/extensions/HostThermal/proto"
	"github.com/hpc/kraken/lib"
	pb "github.com/hpc/kraken/modules/hostdiscovery/proto"
)

//CPUTempObj is strututure for node CPU temperature
type CPUTempObj struct {
	TimeStamp   time.Time
	HostAddress string
	CPUTemp     int
}

const (
	// HostThermalStateURL points to Thermal extension
	HostThermalStateURL = "type.googleapis.com/proto.HostThermal/State"
	// ModuleStateURL refers to module state
	ModuleStateURL = "/Services/hostdiscovery/State"
)

var _ lib.Module = (*HostDisc)(nil)
var _ lib.ModuleWithConfig = (*HostDisc)(nil)
var _ lib.ModuleWithDiscovery = (*HostDisc)(nil)
var _ lib.ModuleSelfService = (*HostDisc)(nil)

// HostDisc provides hostdiscovery module capabilities
type HostDisc struct {
	prevTemp   int
	api        lib.APIClient
	cfg        *pb.HostDiscoveryConfig
	dchan      chan<- lib.Event
	pollTicker *time.Ticker
}

// Name returns the FQDN of the module
func (*HostDisc) Name() string { return "github.com/hpc/kraken/modules/hostdiscovery" }

// NewConfig returns a fully initialized default config
func (*HostDisc) NewConfig() proto.Message {
	r := &pb.HostDiscoveryConfig{
		PollingInterval: "10s",
	}
	return r
}

// UpdateConfig updates the running config
func (hostDisc *HostDisc) UpdateConfig(cfg proto.Message) (e error) {
	if rcfg, ok := cfg.(*pb.HostDiscoveryConfig); ok {
		hostDisc.cfg = rcfg
		return
	}
	return fmt.Errorf("invalid config type")
}

// ConfigURL gives the any resolver URL for the config
func (*HostDisc) ConfigURL() string {
	cfg := &pb.HostDiscoveryConfig{}
	any, _ := ptypes.MarshalAny(cfg)
	return any.GetTypeUrl()
}

// SetDiscoveryChan sets the current discovery channel
// this is generally done by the API
func (hostDisc *HostDisc) SetDiscoveryChan(c chan<- lib.Event) { hostDisc.dchan = c }

func init() {
	module := &HostDisc{}
	discovers := make(map[string]map[string]reflect.Value)
	hostThermDisc := make(map[string]reflect.Value)

	hostThermDisc[thpb.HostThermal_CPU_TEMP_NONE.String()] = reflect.ValueOf(thpb.HostThermal_CPU_TEMP_NONE)
	hostThermDisc[thpb.HostThermal_CPU_TEMP_NORMAL.String()] = reflect.ValueOf(thpb.HostThermal_CPU_TEMP_NORMAL)
	hostThermDisc[thpb.HostThermal_CPU_TEMP_HIGH.String()] = reflect.ValueOf(thpb.HostThermal_CPU_TEMP_HIGH)
	hostThermDisc[thpb.HostThermal_CPU_TEMP_CRITICAL.String()] = reflect.ValueOf(thpb.HostThermal_CPU_TEMP_CRITICAL)

	discovers[HostThermalStateURL] = hostThermDisc

	discovers[ModuleStateURL] = map[string]reflect.Value{
		"RUN": reflect.ValueOf(cpb.ServiceInstance_RUN)}

	si := core.NewServiceInstance("hostdiscovery", module.Name(), module.Entry, nil)

	// Register it all
	core.Registry.RegisterModule(module)
	core.Registry.RegisterServiceInstance(module, map[string]lib.ServiceInstance{si.ID(): si})
	core.Registry.RegisterDiscoverable(module, discovers)
}

// Init is used to intialize an executable module prior to entrypoint
func (hostDisc *HostDisc) Init(api lib.APIClient) {
	hostDisc.api = api
	hostDisc.cfg = hostDisc.NewConfig().(*pb.HostDiscoveryConfig)
}

// Stop should perform a graceful exit
func (hostDisc *HostDisc) Stop() {
	os.Exit(0)
}

// Entry is the module's executable entrypoint
func (hostDisc *HostDisc) Entry() {
	url := lib.NodeURLJoin(hostDisc.api.Self().String(), ModuleStateURL)
	ev := core.NewEvent(
		lib.Event_DISCOVERY,
		url,
		&core.DiscoveryEvent{
			Module:  hostDisc.Name(),
			URL:     url,
			ValueID: "RUN",
		},
	)
	hostDisc.dchan <- ev

	// setup a ticker for polling discovery
	dur, _ := time.ParseDuration(hostDisc.cfg.GetPollingInterval())
	hostDisc.pollTicker = time.NewTicker(dur)

	// main loop
	for {
		select {
		case <-hostDisc.pollTicker.C:
			go hostDisc.discoverHostCPUTemp()
			break
		}
	}
}

// discoverHostCPUTemp is used to acquire CPU thermal locally.
func (hostDisc *HostDisc) discoverHostCPUTemp() {
	hostCPUTemp := hostDisc.GetCPUTemp()
	if hostCPUTemp.CPUTemp == hostDisc.prevTemp {
		// no change in temperature so no need to generate discovery event
		return
	}
	hostDisc.prevTemp = hostCPUTemp.CPUTemp

	vid, _ := lambdaStateDiscovery(hostCPUTemp)
	url := lib.NodeURLJoin(hostDisc.api.Self().String(), HostThermalStateURL)

	// Generating discovery event for CPU Thermal state
	v := core.NewEvent(
		lib.Event_DISCOVERY,
		url,
		&core.DiscoveryEvent{
			Module:  hostDisc.Name(),
			URL:     url,
			ValueID: vid,
		},
	)
	hostDisc.dchan <- v

}

// GetCPUTemp returns CPU temperature
func (hostDisc *HostDisc) GetCPUTemp() CPUTempObj {

	hostIP := hostDisc.GetNodeIPAddress()

	// Its a mockup CPU temperature
	cpuTempObj := CPUTempObj{}
	cpuTempObj.TimeStamp = time.Now()
	cpuTempObj.HostAddress = hostIP

	tempVal := hostDisc.ReadCPUTemp()
	cpuTempObj.CPUTemp = tempVal

	return cpuTempObj
}

// ReadCPUTemp function reads the CPU thermal sensor
func (hostDisc *HostDisc) ReadCPUTemp() int {

	tempSensorPath := "/sys/devices/virtual/thermal/thermal_zone0/temp"
	cpuTemp, err := ioutil.ReadFile(tempSensorPath)
	if err != nil {
		hostDisc.api.Logf(lib.LLERROR, "Reading CPU thermal sensor failed: %v", err)
		return 0
	}
	cpuTempInt, err := strconv.Atoi(strings.TrimSuffix(string(cpuTemp), "\n"))

	if err != nil {
		hostDisc.api.Logf(lib.LLERROR, "String to Int conversion failed: %v", err)
		return 0
	}

	return cpuTempInt
}

// GetNodeIPAddress returns non local loop IP address
func (hostDisc *HostDisc) GetNodeIPAddress() string {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		hostDisc.api.Logf(lib.LLERROR, "Can not find network interfaces: %v", err)
	}
	ip := ""
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			} else {
				hostDisc.api.Logf(lib.LLERROR, "Can not format IP address")
			}
		}
	}

	return ip
}

// Discovers state of the CPU based on CPU temperature thresholds
func lambdaStateDiscovery(v CPUTempObj) (string, int) {
	cpuTemp := v.CPUTemp
	cpuTempState := thpb.HostThermal_CPU_TEMP_NONE

	if cpuTemp <= 3000 || cpuTemp >= 70000 {
		cpuTempState = thpb.HostThermal_CPU_TEMP_CRITICAL
	} else if cpuTemp >= 60000 && cpuTemp < 70000 {
		cpuTempState = thpb.HostThermal_CPU_TEMP_HIGH
	} else if cpuTemp > 3000 && cpuTemp < 60000 {
		cpuTempState = thpb.HostThermal_CPU_TEMP_NORMAL
	}
	return cpuTempState.String(), cpuTemp

}