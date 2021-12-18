package froniusv1

type Head struct {
	Status struct {
		Code        int    `json:"Code"`
		Reason      string `json:"Reason"`
		UserMessage string `json:"UserMessage"`
	} `json:"Status"`
	TimeStamp string `json:"TimeStamp"`
}

type GetInverterRealtimeData struct {
	Head Head `json:"Head"`
	Body struct {
		Data struct {
			IacL1 struct {
				Unit  string `json:"Unit"`
				Value int    `json:"Value"`
			} `json:"IAC_L1"`
			IacL2 struct {
				Unit  string `json:"Unit"`
				Value int    `json:"Value"`
			} `json:"IAC_L2"`
			IacL3 struct {
				Unit  string `json:"Unit"`
				Value int    `json:"Value"`
			} `json:"IAC_L3"`
			TAmbient struct {
				Unit  string `json:"Unit"`
				Value int    `json:"Value"`
			} `json:"T_AMBIENT"`
			UacL1 struct {
				Unit  string `json:"Unit"`
				Value int    `json:"Value"`
			} `json:"UAC_L1"`
			UacL2 struct {
				Unit  string `json:"Unit"`
				Value int    `json:"Value"`
			} `json:"UAC_L2"`
			UacL3 struct {
				Unit  string `json:"Unit"`
				Value int    `json:"Value"`
			} `json:"UAC_L3"`
		} `json:"Data"`
	} `json:"Body"`
}

type GetInverterInfo struct {
	Head Head `json:"Head"`
	Body struct {
		Data struct {
			Num1 struct {
				CustomName    string `json:"CustomName"`
				Dt            int    `json:"DT"`
				ErrorCode     int    `json:"ErrorCode"`
				PVPower       int    `json:"PVPower"`
				Show          int    `json:"Show"`
				StatusCode    int    `json:"StatusCode"`
				InverterState string `json:"InverterState"`
				UniqueID      string `json:"UniqueID"`
			} `json:"1"`
		} `json:"Data"`
	} `json:"Body"`
}

type GetPowerFlowRealtimeData struct {
	Head Head `json:"Head"`
	Body struct {
		Data struct {
			Inverters struct {
				Num1 struct {
					BatteryMode string  `json:"Battery_Mode"`
					Dt          int     `json:"DT"`
					EDay        int     `json:"E_Day"`
					EYear       float64 `json:"E_Year"`
					ETotal      int     `json:"E_Total"`
					P           int     `json:"P"`
					Soc         int     `json:"SOC"`
				} `json:"1"`
				Site struct {
					BackupMode         int    `json:"BackupMode"`
					BatteryStandby     string `json:"BatteryStandby"`
					EDay               int    `json:"E_Day"`
					EYear              int    `json:"E_Year"`
					ETotal             int    `json:"E_Total"`
					MeterLocation      string `json:"Meter_Location"`
					Mode               string `json:"Mode"`
					PAkku              int    `json:"P_Akku"`
					PGrid              int    `json:"P_Grid"`
					PLoad              int    `json:"P_Load"`
					PPv                int    `json:"P_PV"`
					RelAutonomy        int    `json:"rel_Autonomy"`
					RelSelfConsumption int    `json:"rel_SelfConsumption"`
				} `json:"Site"`
				Version string `json:"Version"`
			} `json:"inverters"`
		} `json:"Data"`
	} `json:"Body"`
}

type GetActiveDeviceInfo struct {
	Head Head `json:"Head"`
	Body struct {
		Data struct {
			Inverter struct {
				Num1 struct {
					Dt     int    `json:"DT"`
					Serial string `json:"Serial"`
				} `json:"1"`
			} `json:"Inverter"`
			Meter struct {
				AdditionalProp1 struct {
					Dt     int    `json:"DT"`
					Serial string `json:"Serial"`
				} `json:"additionalProp1"`
				AdditionalProp2 struct {
					Dt     int    `json:"DT"`
					Serial string `json:"Serial"`
				} `json:"additionalProp2"`
				AdditionalProp3 struct {
					Dt     int    `json:"DT"`
					Serial string `json:"Serial"`
				} `json:"additionalProp3"`
			} `json:"Meter"`
			Storage struct {
				AdditionalProp1 struct {
					Dt     int    `json:"DT"`
					Serial string `json:"Serial"`
				} `json:"additionalProp1"`
				AdditionalProp2 struct {
					Dt     int    `json:"DT"`
					Serial string `json:"Serial"`
				} `json:"additionalProp2"`
				AdditionalProp3 struct {
					Dt     int    `json:"DT"`
					Serial string `json:"Serial"`
				} `json:"additionalProp3"`
			} `json:"Storage"`
			Ohmpilot struct {
				AdditionalProp1 struct {
					Dt     int    `json:"DT"`
					Serial string `json:"Serial"`
				} `json:"additionalProp1"`
				AdditionalProp2 struct {
					Dt     int    `json:"DT"`
					Serial string `json:"Serial"`
				} `json:"additionalProp2"`
				AdditionalProp3 struct {
					Dt     int    `json:"DT"`
					Serial string `json:"Serial"`
				} `json:"additionalProp3"`
			} `json:"Ohmpilot"`
		} `json:"Data"`
	} `json:"Body"`
}

type GetMeterRealtimeData struct {
	Head Head `json:"Head"`
	Body struct {
		Data struct {
			AdditionalProp1 struct {
				Details struct {
					Manufacturer string `json:"Manufacturer"`
					Serial       string `json:"Serial"`
					Model        string `json:"Model"`
				} `json:"Details"`
				CurrentACPhase1                   int `json:"Current_AC_Phase_1"`
				CurrentACSum                      int `json:"Current_AC_Sum"`
				Enable                            int `json:"Enable"`
				EnergyReactiveVArACPhase1Consumed int `json:"EnergyReactive_VArAC_Phase_1_Consumed"`
				EnergyReactiveVArACPhase1Produced int `json:"EnergyReactive_VArAC_Phase_1_Produced"`
				EnergyReactiveVArACSumConsumed    int `json:"EnergyReactive_VArAC_Sum_Consumed"`
				EnergyReactiveVArACSumProduced    int `json:"EnergyReactive_VArAC_Sum_Produced"`
				EnergyRealWACMinusAbsolute        int `json:"EnergyReal_WAC_Minus_Absolute"`
				EnergyRealWACPhase1Consumed       int `json:"EnergyReal_WAC_Phase_1_Consumed"`
				EnergyRealWACPhase1Produced       int `json:"EnergyReal_WAC_Phase_1_Produced"`
				EnergyRealWACPlusAbsolute         int `json:"EnergyReal_WAC_Plus_Absolute"`
				EnergyRealWACSumConsumed          int `json:"EnergyReal_WAC_Sum_Consumed"`
				EnergyRealWACSumProduced          int `json:"EnergyReal_WAC_Sum_Produced"`
				FrequencyPhaseAverage             int `json:"Frequency_Phase_Average"`
				MeterLocationCurrent              int `json:"Meter_Location_Current"`
				PowerApparentSPhase1              int `json:"PowerApparent_S_Phase_1"`
				PowerApparentSSum                 int `json:"PowerApparent_S_Sum"`
				PowerFactorPhase1                 int `json:"PowerFactor_Phase_1"`
				PowerFactorSum                    int `json:"PowerFactor_Sum"`
				PowerReactiveQPhase1              int `json:"PowerReactive_Q_Phase_1"`
				PowerReactiveQSum                 int `json:"PowerReactive_Q_Sum"`
				PowerRealPPhase1                  int `json:"PowerReal_P_Phase_1"`
				PowerRealPSum                     int `json:"PowerReal_P_Sum"`
				Timestamp                         int `json:"Timestamp"`
				Visible                           int `json:"Visible"`
				VoltageACPhase1                   int `json:"Voltage_AC_Phase_1"`
			} `json:"additionalProp1"`
			AdditionalProp2 struct {
				Details struct {
					Manufacturer string `json:"Manufacturer"`
					Serial       string `json:"Serial"`
					Model        string `json:"Model"`
				} `json:"Details"`
				CurrentACPhase1                   int `json:"Current_AC_Phase_1"`
				CurrentACSum                      int `json:"Current_AC_Sum"`
				Enable                            int `json:"Enable"`
				EnergyReactiveVArACPhase1Consumed int `json:"EnergyReactive_VArAC_Phase_1_Consumed"`
				EnergyReactiveVArACPhase1Produced int `json:"EnergyReactive_VArAC_Phase_1_Produced"`
				EnergyReactiveVArACSumConsumed    int `json:"EnergyReactive_VArAC_Sum_Consumed"`
				EnergyReactiveVArACSumProduced    int `json:"EnergyReactive_VArAC_Sum_Produced"`
				EnergyRealWACMinusAbsolute        int `json:"EnergyReal_WAC_Minus_Absolute"`
				EnergyRealWACPhase1Consumed       int `json:"EnergyReal_WAC_Phase_1_Consumed"`
				EnergyRealWACPhase1Produced       int `json:"EnergyReal_WAC_Phase_1_Produced"`
				EnergyRealWACPlusAbsolute         int `json:"EnergyReal_WAC_Plus_Absolute"`
				EnergyRealWACSumConsumed          int `json:"EnergyReal_WAC_Sum_Consumed"`
				EnergyRealWACSumProduced          int `json:"EnergyReal_WAC_Sum_Produced"`
				FrequencyPhaseAverage             int `json:"Frequency_Phase_Average"`
				MeterLocationCurrent              int `json:"Meter_Location_Current"`
				PowerApparentSPhase1              int `json:"PowerApparent_S_Phase_1"`
				PowerApparentSSum                 int `json:"PowerApparent_S_Sum"`
				PowerFactorPhase1                 int `json:"PowerFactor_Phase_1"`
				PowerFactorSum                    int `json:"PowerFactor_Sum"`
				PowerReactiveQPhase1              int `json:"PowerReactive_Q_Phase_1"`
				PowerReactiveQSum                 int `json:"PowerReactive_Q_Sum"`
				PowerRealPPhase1                  int `json:"PowerReal_P_Phase_1"`
				PowerRealPSum                     int `json:"PowerReal_P_Sum"`
				Timestamp                         int `json:"Timestamp"`
				Visible                           int `json:"Visible"`
				VoltageACPhase1                   int `json:"Voltage_AC_Phase_1"`
			} `json:"additionalProp2"`
			AdditionalProp3 struct {
				Details struct {
					Manufacturer string `json:"Manufacturer"`
					Serial       string `json:"Serial"`
					Model        string `json:"Model"`
				} `json:"Details"`
				CurrentACPhase1                   int `json:"Current_AC_Phase_1"`
				CurrentACSum                      int `json:"Current_AC_Sum"`
				Enable                            int `json:"Enable"`
				EnergyReactiveVArACPhase1Consumed int `json:"EnergyReactive_VArAC_Phase_1_Consumed"`
				EnergyReactiveVArACPhase1Produced int `json:"EnergyReactive_VArAC_Phase_1_Produced"`
				EnergyReactiveVArACSumConsumed    int `json:"EnergyReactive_VArAC_Sum_Consumed"`
				EnergyReactiveVArACSumProduced    int `json:"EnergyReactive_VArAC_Sum_Produced"`
				EnergyRealWACMinusAbsolute        int `json:"EnergyReal_WAC_Minus_Absolute"`
				EnergyRealWACPhase1Consumed       int `json:"EnergyReal_WAC_Phase_1_Consumed"`
				EnergyRealWACPhase1Produced       int `json:"EnergyReal_WAC_Phase_1_Produced"`
				EnergyRealWACPlusAbsolute         int `json:"EnergyReal_WAC_Plus_Absolute"`
				EnergyRealWACSumConsumed          int `json:"EnergyReal_WAC_Sum_Consumed"`
				EnergyRealWACSumProduced          int `json:"EnergyReal_WAC_Sum_Produced"`
				FrequencyPhaseAverage             int `json:"Frequency_Phase_Average"`
				MeterLocationCurrent              int `json:"Meter_Location_Current"`
				PowerApparentSPhase1              int `json:"PowerApparent_S_Phase_1"`
				PowerApparentSSum                 int `json:"PowerApparent_S_Sum"`
				PowerFactorPhase1                 int `json:"PowerFactor_Phase_1"`
				PowerFactorSum                    int `json:"PowerFactor_Sum"`
				PowerReactiveQPhase1              int `json:"PowerReactive_Q_Phase_1"`
				PowerReactiveQSum                 int `json:"PowerReactive_Q_Sum"`
				PowerRealPPhase1                  int `json:"PowerReal_P_Phase_1"`
				PowerRealPSum                     int `json:"PowerReal_P_Sum"`
				Timestamp                         int `json:"Timestamp"`
				Visible                           int `json:"Visible"`
				VoltageACPhase1                   int `json:"Voltage_AC_Phase_1"`
			} `json:"additionalProp3"`
		} `json:"Data"`
	} `json:"Body"`
}

type GetOhmpilotRealtimeData struct {
	Head Head `json:"Head"`
	Body struct {
		Data struct {
			AdditionalProp1 struct {
				Details struct {
					Manufacturer string `json:"Manufacturer"`
					Serial       string `json:"Serial"`
					Model        string `json:"Model"`
					Hardware     string `json:"Hardware"`
					Software     string `json:"Software"`
				} `json:"Details"`
				CodeOfState                       int `json:"CodeOfState"`
				EnergyRealWACSumConsumed          int `json:"EnergyReal_WAC_Sum_Consumed"`
				PowerRealPACSum                   int `json:"PowerReal_PAC_Sum"`
				TemperatureChannel1               int `json:"Temperature_Channel_1"`
				EnergyReactiveVArACPhase1Produced int `json:"EnergyReactive_VArAC_Phase_1_Produced"`
			} `json:"additionalProp1"`
			AdditionalProp2 struct {
				Details struct {
					Manufacturer string `json:"Manufacturer"`
					Serial       string `json:"Serial"`
					Model        string `json:"Model"`
					Hardware     string `json:"Hardware"`
					Software     string `json:"Software"`
				} `json:"Details"`
				CodeOfState                       int `json:"CodeOfState"`
				EnergyRealWACSumConsumed          int `json:"EnergyReal_WAC_Sum_Consumed"`
				PowerRealPACSum                   int `json:"PowerReal_PAC_Sum"`
				TemperatureChannel1               int `json:"Temperature_Channel_1"`
				EnergyReactiveVArACPhase1Produced int `json:"EnergyReactive_VArAC_Phase_1_Produced"`
			} `json:"additionalProp2"`
			AdditionalProp3 struct {
				Details struct {
					Manufacturer string `json:"Manufacturer"`
					Serial       string `json:"Serial"`
					Model        string `json:"Model"`
					Hardware     string `json:"Hardware"`
					Software     string `json:"Software"`
				} `json:"Details"`
				CodeOfState                       int `json:"CodeOfState"`
				EnergyRealWACSumConsumed          int `json:"EnergyReal_WAC_Sum_Consumed"`
				PowerRealPACSum                   int `json:"PowerReal_PAC_Sum"`
				TemperatureChannel1               int `json:"Temperature_Channel_1"`
				EnergyReactiveVArACPhase1Produced int `json:"EnergyReactive_VArAC_Phase_1_Produced"`
			} `json:"additionalProp3"`
		} `json:"Data"`
	} `json:"Body"`
}

type GetStorageRealtimeData struct {
	Head Head `json:"Head"`
	Body struct {
		Data struct {
			AdditionalProp1 struct {
				Controller struct {
					Details struct {
						Manufacturer string `json:"Manufacturer"`
						Serial       string `json:"Serial"`
						Model        string `json:"Model"`
					} `json:"Details"`
					CapacityMaximum       int `json:"Capacity_Maximum"`
					CurrentDC             int `json:"Current_DC"`
					DesignedCapacity      int `json:"DesignedCapacity"`
					Enable                int `json:"Enable"`
					StateOfChargeRelative int `json:"StateOfCharge_Relative"`
					StatusBatteryCell     int `json:"Status_BatteryCell"`
					TemperatureCell       int `json:"Temperature_Cell"`
					TimeStamp             int `json:"TimeStamp"`
					VoltageDC             int `json:"Voltage_DC"`
				} `json:"Controller"`
				Modules []string `json:"Modules"`
			} `json:"additionalProp1"`
			AdditionalProp2 struct {
				Controller struct {
					Details struct {
						Manufacturer string `json:"Manufacturer"`
						Serial       string `json:"Serial"`
						Model        string `json:"Model"`
					} `json:"Details"`
					CapacityMaximum       int `json:"Capacity_Maximum"`
					CurrentDC             int `json:"Current_DC"`
					DesignedCapacity      int `json:"DesignedCapacity"`
					Enable                int `json:"Enable"`
					StateOfChargeRelative int `json:"StateOfCharge_Relative"`
					StatusBatteryCell     int `json:"Status_BatteryCell"`
					TemperatureCell       int `json:"Temperature_Cell"`
					TimeStamp             int `json:"TimeStamp"`
					VoltageDC             int `json:"Voltage_DC"`
				} `json:"Controller"`
				Modules []string `json:"Modules"`
			} `json:"additionalProp2"`
			AdditionalProp3 struct {
				Controller struct {
					Details struct {
						Manufacturer string `json:"Manufacturer"`
						Serial       string `json:"Serial"`
						Model        string `json:"Model"`
					} `json:"Details"`
					CapacityMaximum       int `json:"Capacity_Maximum"`
					CurrentDC             int `json:"Current_DC"`
					DesignedCapacity      int `json:"DesignedCapacity"`
					Enable                int `json:"Enable"`
					StateOfChargeRelative int `json:"StateOfCharge_Relative"`
					StatusBatteryCell     int `json:"Status_BatteryCell"`
					TemperatureCell       int `json:"Temperature_Cell"`
					TimeStamp             int `json:"TimeStamp"`
					VoltageDC             int `json:"Voltage_DC"`
				} `json:"Controller"`
				Modules []string `json:"Modules"`
			} `json:"additionalProp3"`
		} `json:"Data"`
	} `json:"Body"`
}
