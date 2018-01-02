package main

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var racadmOutput string = `
Sensor Type : POWER
<Sensor Name>                   <Status>                 <Type>
PS1 Status                      Present                  AC
PS2 Status                      Present                  AC

Sensor Type : TEMPERATURE
<Sensor Name>            <Status>    <Reading> <lc> <uc>  <lnc>[R/W]  <unc>[R/W]
[Key = iDRAC.Embedded.1#SystemBoardInletTemp]
System Board Inlet Temp       Ok      21C      -7C  47C      3C [Y]      42C [Y]

[Key = iDRAC.Embedded.1#SystemBoardExhaustTemp]
System Board Exhaust Temp     Ok      35C      0C   75C      0C [N]      70C [N]

[Key = iDRAC.Embedded.1#CPU1Temp]
CPU1 Temp                     Ok      49C      3C   73C      8C [N]      68C [N]

[Key = iDRAC.Embedded.1#CPU2Temp]
CPU2 Temp                     Ok      52C      3C   73C      8C [N]      68C [N]



Sensor Type : FAN
<Sensor Name>                   <Status>    <Reading>   <lc>        <uc>        <PWM %>
System Board Fan1A              Ok          5880RPM     600RPM      NA          23%
System Board Fan2A              Ok          6960RPM     600RPM      NA          28%
System Board Fan3A              Ok          7920RPM     600RPM      NA          34%
System Board Fan4A              Ok          6720RPM     600RPM      NA          28%
System Board Fan5A              Ok          7440RPM     600RPM      NA          31%
System Board Fan6A              Ok          6840RPM     600RPM      NA          28%
System Board Fan7A              Ok          5880RPM     600RPM      NA          23%
System Board Fan1B              Ok          4680RPM     600RPM      NA          23%
System Board Fan2B              Ok          5400RPM     600RPM      NA          28%
System Board Fan3B              Ok          6120RPM     600RPM      NA          34%
System Board Fan4B              Ok          5400RPM     600RPM      NA          28%
System Board Fan5B              Ok          5880RPM     600RPM      NA          31%
System Board Fan6B              Ok          5400RPM     600RPM      NA          28%
System Board Fan7B              Ok          4680RPM     600RPM      NA          23%

Sensor Type : VOLTAGE
<Sensor Name>                   <Status>    <Reading>   <lc>        <uc>
CPU1 VCORE PG                   Ok          Good                NA          NA
CPU2 VCORE PG                   Ok          Good                NA          NA
System Board 3.3V PG            Ok          Good                NA          NA
System Board 5V AUX PG          Ok          Good                NA          NA
CPU2 M23 VPP PG                 Ok          Good                NA          NA
CPU1 M23 VPP PG                 Ok          Good                NA          NA
System Board 2.5V AUX PG        Ok          Good                NA          NA
System Board 1.05V PG           Ok          Good                NA          NA
CPU1 M23 VDDQ PG                Ok          Good                NA          NA
CPU1 M23 VTT PG                 Ok          Good                NA          NA
System Board 5V SWITCH PG       Ok          Good                NA          NA
System Board DIMM PG            Ok          Good                NA          NA
System Board VCCIO PG           Ok          Good                NA          NA
CPU2 M01 VDDQ PG                Ok          Good                NA          NA
CPU1 M01 VDDQ PG                Ok          Good                NA          NA
CPU2 M23 VTT PG                 Ok          Good                NA          NA
CPU2 M01 VTT PG                 Ok          Good                NA          NA
System Board NDC PG             Ok          Good                NA          NA
CPU2 M01 VPP PG                 Ok          Good                NA          NA
CPU1 M01 VPP PG                 Ok          Good                NA          NA
CPU2 M23 VDDQ PG                Ok          Good                NA          NA
System Board 1.5V PG            Ok          Good                NA          NA
System Board 1.5V AUX PG        Ok          Good                NA          NA
CPU1 FIVR PG                    Ok          Good                NA          NA
CPU2 FIVR PG                    Ok          Good                NA          NA
System Board BP1 5V PG          Ok          Good                NA          NA
System Board BP2 5V PG          Ok          Good                NA          NA
CPU1 M01 VTT PG                 Ok          Good                NA          NA
System Board PS1 PG Fail        Ok          Good                NA          NA
PS1 Voltage 1                   Ok          222.00V             NA          NA
System Board PS2 PG Fail        Ok          Good                NA          NA
PS2 Voltage 2                   Ok          230.00V             NA          NA

Sensor Type : CURRENT
<Sensor Name>            <Status>    <Reading> <lc> <uc>  <lnc>[R/W]  <unc>[R/W]
[Key = iDRAC.Embedded.1#SystemBoardPwrConsumption]
System Board Pwr Consumption  Ok      224Watts NA   980Watts 0Watts [N]      896Watts [Y]

[Key = iDRAC.Embedded.1#PS1Current1]
PS1 Current 1                 Ok      1.0Amps  NA   NA       0Amps [N]      0Amps [N]

[Key = iDRAC.Embedded.1#PS2Current2]
PS2 Current 2                 Ok      0.0Amps  NA   NA       0Amps [N]      0Amps [N]


Sensor Type : PROCESSOR
<Sensor Name>                   <Status>    <State>             <lc>        <uc>
CPU1 Status                     Ok          Presence_Detected   NA          NA
CPU2 Status                     Ok          Presence_Detected   NA          NA

Sensor Type : MEMORY
<Sensor Name>                   <Status>    <State>             <lc>        <uc>
DIMM A1                         Ok          Presence_Detected   NA          NA
DIMM A2                         Ok          Presence_Detected   NA          NA
DIMM A3                         Ok          Presence_Detected   NA          NA
DIMM A4                         Ok          Presence_Detected   NA          NA
DIMM A5                         N/A         Absent              NA          NA
DIMM A6                         N/A         Absent              NA          NA
DIMM A7                         N/A         Absent              NA          NA
DIMM A8                         N/A         Absent              NA          NA
DIMM A9                         N/A         Absent              NA          NA
DIMM A10                        N/A         Absent              NA          NA
DIMM A11                        N/A         Absent              NA          NA
DIMM A12                        N/A         Absent              NA          NA
DIMM B1                         Ok          Presence_Detected   NA          NA
DIMM B2                         Ok          Presence_Detected   NA          NA
DIMM B3                         Ok          Presence_Detected   NA          NA
DIMM B4                         Ok          Presence_Detected   NA          NA
DIMM B5                         N/A         Absent              NA          NA
DIMM B6                         N/A         Absent              NA          NA
DIMM B7                         N/A         Absent              NA          NA
DIMM B8                         N/A         Absent              NA          NA
DIMM B9                         N/A         Absent              NA          NA
DIMM B10                        N/A         Absent              NA          NA
DIMM B11                        N/A         Absent              NA          NA
DIMM B12                        N/A         Absent              NA          NA

Sensor Type : BATTERY
<Sensor Name>                   <Status>    <Reading>   <lc>        <uc>
System Board CMOS Battery       Ok          Present             NA          NA
PERC1 ROMB Battery              Ok          Present             NA          NA

Sensor Type : PERFORMANCE
<Sensor Name>                   <Status>    <State>        <lc>      <uc>
System Board Power Optimized    Ok          Not Degraded   NA        NA

Sensor Type : INTRUSION
<Sensor Name>                   <Intrusion>    <Status>
System Board Intrusion          Closed         Power ON

Sensor Type : REDUNDANCY
<Sensor Name>                   <Status>                 <Type>
System Board Fan Redundancy     Full Redundant           Fan
System Board PS Redundancy      Full Redundant           PSU

Sensor Type : SYSTEM PERFORMANCE
<Sensor Name>            <Status>    <Reading> <lc> <uc>  <lnc>[R/W]  <unc>[R/W]
[Key = iDRAC.Embedded.1#SystemBoardCPUUsage]
System Board CPU Usage   Ok             99%     NA  NA    NA  [N]      101% [Y]

[Key = iDRAC.Embedded.1#SystemBoardIOUsage]
System Board IO Usage    Ok             0%      NA  NA    NA  [N]      101% [Y]

[Key = iDRAC.Embedded.1#SystemBoardMEMUsage]
System Board MEM Usage   Ok             5%      NA  NA    NA  [N]      101% [Y]

[Key = iDRAC.Embedded.1#SystemBoardSYSUsage]
System Board SYS Usage   Ok             94%     NA  NA    NA  [N]      101% [Y]
	`
	var metricsSensorExpected = map[string]int{
		"POWER":       1,
		"TEMPERATURE": 1,
		"FAN":         1,
		"VOLTAGE":     1,
		"CURRENT":     1,
		"PROCESSOR":   1,
		"MEMORY":      1,
		"BATTERY":     1,
		"PERFORMANCE": 1,
		"INTRUSION":   1,
		"REDUNDANCY":  1,
		"SYSTEM":      1,
	}
	metricsSensor := parseRacadmOutput(racadmOutput)
	assert.NotEmpty(t, metricsSensor, "Should not be empty")
	assert.EqualValues(t, metricsSensor, metricsSensorExpected, "Metrics should be in the array")
	assert.Equal(t, 123, 123, "they should be equal")
}

func Test_getSensorsTypeHeaders(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		isOk     bool
	}{
		{name: "power_match", input: "Sensor Type : POWER", expected: "POWER", isOk: true},
		{name: "temp_match", input: "Sensor Type : TEMPERATURE", expected: "TEMPERATURE", isOk: true},
		{name: "fan_match", input: "Sensor Type : FAN", expected: "FAN", isOk: true},
		{name: "voltage_match", input: "Sensor Type : VOLTAGE", expected: "VOLTAGE", isOk: true},
		{name: "current_match", input: "Sensor Type : CURRENT", expected: "CURRENT", isOk: true},
		{name: "processor_match", input: "Sensor Type : PROCESSOR", expected: "PROCESSOR", isOk: true},
		{name: "memory_match", input: "Sensor Type : MEMORY", expected: "MEMORY", isOk: true},
		{name: "battery_match", input: "Sensor Type : BATTERY", expected: "BATTERY", isOk: true},
		{name: "performance_match", input: "Sensor Type : PERFORMANCE", expected: "PERFORMANCE", isOk: true},
		{name: "intrusion_match", input: "Sensor Type : INTRUSION", expected: "INTRUSION", isOk: true},
		{name: "redundancy_match", input: "Sensor Type : REDUNDANCY", expected: "REDUNDANCY", isOk: true},
		{name: "system_match", input: "Sensor Type : SYSTEM", expected: "SYSTEM", isOk: true},
		{name: "system_match", input: "Sensor Type : SYSTEM", expected: "SYSTEM", isOk: true},
		{name: "notOK", input: "SYSTEM", expected: "", isOk: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok, match := getSensorsTypeHeaders(test.input)
			assert.EqualValues(t, match, test.expected)
			assert.EqualValues(t, ok, test.isOk)
		})
	}
}

func Test_getSensorHeaders(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
		isOk     bool
	}{
		{name: "POWER", input: `<Sensor Name>                   <Status>                 <Type>`, expected: []string{"sensor_name", "status", "type"}, isOk: true},
		{name: "TEMPERATURE", input: `<Sensor Name>            <Status>    <Reading> <lc> <uc>  <lnc>[R/W]  <unc>[R/W]`, expected: []string{"sensor_name", "status", "reading", "lc", "uc", "lnc", "unc"}, isOk: true},
		{name: "FAN", input: `<Sensor Name>                   <Status>    <Reading>   <lc>        <uc>        <PWM %>`, expected: []string{"sensor_name", "status", "reading", "lc", "uc", "pwm_%"}, isOk: true},
		{name: "VOLTAGE", input: `<Sensor Name>                   <Status>    <Reading>   <lc>        <uc>`, expected: []string{"sensor_name", "status", "reading", "lc", "uc"}, isOk: true},
		{name: "CURRENT", input: `<Sensor Name>            <Status>    <Reading> <lc> <uc>  <lnc>[R/W]  <unc>[R/W]`, expected: []string{"sensor_name", "status", "reading", "lc", "uc", "lnc", "unc"}, isOk: true},
		{name: "PROCESSOR", input: `<Sensor Name>                   <Status>    <State>             <lc>        <uc>`, expected: []string{"sensor_name", "status", "state", "lc", "uc"}, isOk: true},
		{name: "MEMORY", input: `<Sensor Name>                   <Status>    <State>             <lc>        <uc>`, expected: []string{"sensor_name", "status", "state", "lc", "uc"}, isOk: true},
		{name: "BATTERY", input: `<Sensor Name>                   <Status>    <Reading>   <lc>        <uc>`, expected: []string{"sensor_name", "status", "reading", "lc", "uc"}, isOk: true},
		{name: "PERFORMANCE", input: `<Sensor Name>                   <Status>    <State>        <lc>      <uc>`, expected: []string{"sensor_name", "status", "state", "lc", "uc"}, isOk: true},
		{name: "INTRUSION", input: `<Sensor Name>                   <Intrusion>    <Status>`, expected: []string{"sensor_name", "intrusion", "status"}, isOk: true},
		{name: "REDUNDANCY", input: `<Sensor Name>                   <Status>                 <Type>`, expected: []string{"sensor_name", "status", "type"}, isOk: true},
		{name: "SYSTEM", input: `<Sensor Name>            <Status>    <Reading> <lc> <uc>  <lnc>[R/W]  <unc>[R/W]`, expected: []string{"sensor_name", "status", "reading", "lc", "uc", "lnc", "unc"}, isOk: true},
		{name: "notOk", input: `Sensor Type : SYSTEM`, expected: []string(nil), isOk: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok, match := getSensorHeaders(test.input)
			assert.EqualValues(t, match, test.expected)
			assert.EqualValues(t, ok, test.isOk)
		})
	}
}

func Test_getSensorData(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
		isOk     bool
	}{
		{name: "POWER", input: `PS1 Status                      Present                  AC`, expected: []string{"ps1_status", "present", "ac"}, isOk: true},
		{name: "TEMPERATURE", input: `System Board Inlet Temp       Ok      21C      -7C  47C      3C [Y]      42C [Y]`, expected: []string{"system_board_inlet_temp", "ok", "21c", "-7c", "47c", "3c", "42c"}, isOk: true},
		{name: "TEMPERATURE_CPU", input: `CPU1 Temp                     Ok      49C      3C   73C      8C [N]      68C [N]`, expected: []string{"cpu1_temp", "ok", "49c", "3c", "73c", "8c", "68c"}, isOk: true},
		{name: "FAN", input: `System Board Fan1A              Ok          5880RPM     600RPM      NA          23%`, expected: []string{"system_board_fan1a", "ok", "5880rpm", "600rpm", "na", "23%"}, isOk: true},
		{name: "VOLTAGE", input: `CPU1 VCORE PG                   Ok          Good                NA          NA`, expected: []string{"cpu1_vcore_pg", "ok", "good", "na", "na"}, isOk: true},
		{name: "DIMM", input: `DIMM A1                         Ok          Presence_Detected   NA          NA`, expected: []string{"dimm_a1", "ok", "presence_detected", "na", "na"}, isOk: true},
		{name: "notOkHeaders", input: `<Sensor Name>            <Status>    <Reading> <lc> <uc>  <lnc>[R/W]  <unc>[R/W]`, expected: []string(nil), isOk: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ok, match := getSensorData(test.input)
			assert.EqualValues(t, test.expected, match)
			assert.EqualValues(t, test.isOk, ok)
		})
	}
}

func Test_metricsPower(t *testing.T) {
	tests := []struct {
		name                string
		inputMatch          []string
		inputHeaders        []string
		expected            *prometheus.GaugeVec
		expectedStatus      float64
		expectedLabelValues []string
	}{
		{
			name:         "POWER",
			inputMatch:   []string{"ps1_status", "present", "ac"},
			inputHeaders: []string{"sensor_name", "status", "type"},
			expected: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name:      "power_sensor_name",
				Namespace: Namespace,
			}, []string{"status", "PSU"}),
			expectedStatus:      0,
			expectedLabelValues: []string{"present", "ps1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			match := metricsPower(test.inputMatch, test.inputHeaders)
			test.expected.WithLabelValues(test.expectedLabelValues...).Set(test.expectedStatus)
			expectedGauge, _ := test.expected.GetMetricWithLabelValues(test.expectedLabelValues...)
			providedGauge, _ := match.GetMetricWithLabelValues(test.expectedLabelValues...)
			assert.EqualValues(t, expectedGauge, providedGauge)
		})
	}
}
