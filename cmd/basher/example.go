/*
Copyright (C) 2017 Gerard Braad <me@gbraad.nl>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
    "fmt"
    "strings"
    
	ps "github.com/gorillalabs/go-powershell"
	"github.com/gorillalabs/go-powershell/backend"
)

type PowerShell struct {
	powerShell ps.Shell
}

func New() *PowerShell {
	return &PowerShell{
		powerShell: createPowerShell(),
	}
}
func (p *PowerShell) Execute(command string) (stdOut string, stdErr string) {
	stdOut, stdErr, _ = p.powerShell.Execute(command)
	return
}

func createPowerShell() ps.Shell {
	back := &backend.Local{}
	shell, _ := ps.New(back)
	return shell
}

// Example command I need to test
func ConfigureIPAddress(success chan bool) {
	ipAddress := "10.0.75.128"
	machineName := "minishift"
	
	setKVPForIpAddress := `
$vmMgmt = Get-WmiObject -Namespace root\virtualization\v2 -Class Msvm_VirtualSystemManagementService
$vm = Get-WmiObject -Namespace root\virtualization\v2 -Class Msvm_ComputerSystem -Filter {` + fmt.Sprintf("ElementName = '%s'", machineName) + `}
$kvpDataItem = ([WMIClass][String]::Format("\\{0}\{1}:{2}", $VmMgmt.ClassPath.Server, $VmMgmt.ClassPath.NamespacePath, "Msvm_KvpExchangeDataItem")).CreateInstance()
$kvpDataItem.Name = 'IpAddress'
` + fmt.Sprintf("$kvpDataItem.Data = '%s'", ipAddress) + `
$kvpDataItem.Source = 0
$vmMgmt.RemoveKvpItems($vm, $kvpDataItem.PSBase.GetText(1))
$result = $vmMgmt.AddKvpItems($vm, $kvpDataItem.PSBase.GetText(1))
$result.ReturnValue
	`

	posh := New() // createa a new instance?!

	result, _ := posh.Execute(setKVPForIpAddress)
	
	if (strings.Contains(result, "4096")) {
		success <- true
	}
}