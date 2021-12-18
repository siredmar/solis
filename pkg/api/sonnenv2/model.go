package sonnenv2

type Status struct {
	ApparentOutput            int         `json:"Apparent_output"`
	BackupBuffer              string      `json:"BackupBuffer"`
	BatteryCharging           bool        `json:"BatteryCharging"`
	BatteryDischarging        bool        `json:"BatteryDischarging"`
	ConsumptionW              int         `json:"Consumption_W"`
	Fac                       float64     `json:"Fac"`
	FlowConsumptionBattery    bool        `json:"FlowConsumptionBattery"`
	FlowConsumptionGrid       bool        `json:"FlowConsumptionGrid"`
	FlowConsumptionProduction bool        `json:"FlowConsumptionProduction"`
	FlowGridBattery           bool        `json:"FlowGridBattery"`
	FlowProductionBattery     bool        `json:"FlowProductionBattery"`
	FlowProductionGrid        bool        `json:"FlowProductionGrid"`
	GridFeedInW               int         `json:"GridFeedIn_W"`
	IsSystemInstalled         int         `json:"IsSystemInstalled"`
	OperatingMode             string      `json:"OperatingMode"`
	PacTotalW                 int         `json:"Pac_total_W"`
	ProductionW               int         `json:"Production_W"`
	Rsoc                      int         `json:"RSOC"`
	Sac1                      int         `json:"Sac1"`
	Sac2                      interface{} `json:"Sac2"`
	Sac3                      interface{} `json:"Sac3"`
	SystemStatus              string      `json:"SystemStatus"`
	Timestamp                 string      `json:"Timestamp"`
	Usoc                      int         `json:"USOC"`
	Uac                       int         `json:"Uac"`
	Ubat                      int         `json:"Ubat"`
	DischargeNotAllowed       bool        `json:"dischargeNotAllowed"`
	GeneratorAutostart        bool        `json:"generator_autostart"`
}

type LastStatus struct {
	ConsumptionW       int    `json:"Consumption_W"`
	FullChargeCapacity int    `json:"FullChargeCapacity"`
	GridFeedInW        int    `json:"GridFeedIn_W"`
	PacTotalW          int    `json:"Pac_total_W"`
	ProductionW        int    `json:"Production_W"`
	Rsoc               int    `json:"RSOC"`
	SetPointW          int    `json:"SetPoint_W"`
	Timestamp          string `json:"Timestamp"`
	Usoc               int    `json:"USOC"`
	UTCOffet           int    `json:"UTC_Offet"`
	IcStatus           struct {
		DCShutdownReason struct {
			CriticalBMSAlarm                  bool `json:"Critical BMS Alarm"`
			ErrorConditionInBMSInitialization bool `json:"Error condition in BMS initialization"`
			HWShutdown                        bool `json:"HW_Shutdown"`
			HardWireOverVoltage               bool `json:"HardWire Over Voltage"`
			HardWiredDrySignalA               bool `json:"HardWired Dry Signal A"`
			HardWiredUnderVoltage             bool `json:"HardWired Under Voltage"`
			InitializationTimeout             bool `json:"Initialization Timeout"`
			InitializationOfACContactorFailed bool `json:"Initialization of AC contactor failed"`
			InitializationOfBMSHardwareFailed bool `json:"Initialization of BMS hardware failed"`
			InitializationOfDCContactorFailed bool `json:"Initialization of DC contactor failed"`
			InitializationOfInverterFailed    bool `json:"Initialization of Inverter failed"`
			InvalidOrNoSystemTypeWasSet       bool `json:"Invalid or no SystemType was set"`
			InverterOverTemperature           bool `json:"Inverter Over Temperature"`
			InverterUnderVoltage              bool `json:"Inverter Under Voltage"`
			ManualShutdownByUser              bool `json:"Manual shutdown by user"`
			MinimumRSOCOfSystemReached        bool `json:"Minimum rSOC of System reached"`
			NoSetpointReceivedByHC            bool `json:"No Setpoint received by HC"`
			ShutdownTimerStarted              bool `json:"Shutdown Timer started"`
			SystemValidationFailed            bool `json:"System Validation failed"`
			VoltageMonitorChanged             bool `json:"Voltage Monitor Changed"`
		} `json:"DC Shutdown Reason"`
		EclipseLed struct {
			PulsingGreen  bool `json:"Pulsing Green"`
			PulsingOrange bool `json:"Pulsing Orange"`
			PulsingWhite  bool `json:"Pulsing White"`
			SolidRed      bool `json:"Solid Red"`
		} `json:"Eclipse Led"`
		FlatStatus struct {
			AutoMode           bool `json:"Auto Mode"`
			Error              bool `json:"Error"`
			FullChargePower    bool `json:"Full Charge Power"`
			FullDischargePower bool `json:"Full Discharge Power"`
			NotConnected       bool `json:"Not Connected"`
			Spare1             bool `json:"Spare 1"`
			Spare2             bool `json:"Spare 2"`
			Spare3             bool `json:"Spare 3"`
			Timeout            bool `json:"Timeout"`
		} `json:"Flat Status"`
		MISCStatusBits struct {
			DischargeNotAllowed bool `json:"Discharge not allowed"`
			MinSystemSOC        bool `json:"Min System SOC"`
			MinUserSOC          bool `json:"Min User SOC"`
		} `json:"MISC Status Bits"`
		MicrogridStatus struct {
			ContiniousPowerViolation       bool `json:"Continious Power Violation"`
			DischargeCurrentLimitViolation bool `json:"Discharge Current Limit Violation"`
			LowTemperature                 bool `json:"Low Temperature"`
			MaxSystemSOC                   bool `json:"Max System SOC"`
			MaxUserSOC                     bool `json:"Max User SOC"`
			MicrogridEnabled               bool `json:"Microgrid Enabled"`
			MinSystemSOC                   bool `json:"Min System SOC"`
			MinUserSOC                     bool `json:"Min User SOC"`
			OverChargeCurrent              bool `json:"Over Charge Current"`
			OverDischargeCurrent           bool `json:"Over Discharge Current"`
			PeakPowerViolation             bool `json:"Peak Power Violation"`
			ProtectIsActivated             bool `json:"Protect is activated"`
			TransitionToOngridPending      bool `json:"Transition to Ongrid Pending"`
		} `json:"Microgrid Status"`
		SetpointPriority struct {
			Bms               bool `json:"BMS"`
			EnergyManager     bool `json:"Energy Manager"`
			Flat              bool `json:"Flat"`
			FullChargeRequest bool `json:"Full Charge Request"`
			Inverter          bool `json:"Inverter"`
			MinUserSOC        bool `json:"Min User SOC"`
			TrickleCharge     bool `json:"Trickle Charge"`
		} `json:"Setpoint Priority"`
		SystemValidation struct {
			CountryCodeSetStatusFlag1  bool `json:"Country Code Set status flag 1"`
			CountryCodeSetStatusFlag2  bool `json:"Country Code Set status flag 2"`
			SelfTestErrorDCWiring      bool `json:"Self test Error DC Wiring"`
			SelfTestPostponed          bool `json:"Self test Postponed"`
			SelfTestPreconditionNotMet bool `json:"Self test Precondition not met"`
			SelfTestRunning            bool `json:"Self test Running"`
			SelfTestSuccessfulFinished bool `json:"Self test successful finished"`
		} `json:"System Validation"`
		Nrbatterymodules       int    `json:"nrbatterymodules"`
		Secondssincefullcharge int    `json:"secondssincefullcharge"`
		Statebms               string `json:"statebms"`
		Statecorecontrolmodule string `json:"statecorecontrolmodule"`
		Stateinverter          string `json:"stateinverter"`
		Timestamp              string `json:"timestamp"`
	} `json:"ic_status"`
}
