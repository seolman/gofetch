package main

import (
    "fmt"
    "os"
    "runtime"

    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/host"
    "github.com/shirou/gopsutil/mem"
)

func main() {
    // Hostname
    hostname, _ := os.Hostname()
    fmt.Println("Hostname:", hostname)

    // OS and Architecture
    fmt.Println("OS:", runtime.GOOS)
    fmt.Println("Arch:", runtime.GOARCH)

    // CPU Information
    cpuInfo, _ := cpu.Info()
    fmt.Println("CPU Model:", cpuInfo[0].ModelName)

    // Memory Information
    memInfo, _ := mem.VirtualMemory()
    fmt.Println("Total Memory:", memInfo.Total / 1024 / 1024)
    fmt.Println("Free Memory:", memInfo.Free / 1024 / 1024)

    // Uptime
    uptime, _ := host.Uptime()
    fmt.Println("Uptime:", uptime / 60 / 60)
}
