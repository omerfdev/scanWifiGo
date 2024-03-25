package main

import (
    "encoding/json"

    "log"
    "net/http"
    "os/exec"
    "strings"
)

type Device struct {
    MACAddress string `json:"mac_address"`
    IPAddress  string `json:"ip_address"`
}

func main() {
    http.HandleFunc("/wifi/devices", getWifiDevices)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getWifiDevices(w http.ResponseWriter, r *http.Request) {
    output, err := exec.Command("arp", "-a").Output()
    if err != nil {
        http.Error(w, "Error getting device list", http.StatusInternalServerError)
        return
    }

    devices := make(map[string]Device)

    lines := strings.Split(string(output), "\n")
    for _, line := range lines {
        parts := strings.Fields(line)
        if len(parts) >= 2 {
            mac := parts[1]
            ip := parts[0]
            devices[mac] = Device{MACAddress: mac, IPAddress: ip}
        }
    }

    response, err := json.Marshal(devices)
    if err != nil {
        http.Error(w, "Error marshaling JSON response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}
