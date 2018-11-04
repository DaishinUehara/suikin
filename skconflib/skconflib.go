package skconflib

import "os"

const (
	envPrefix          = "SUIKIN_"
	envPasswordPrefix  = envPrefix + "PASSWD_"
	envLogConfigPrefix = envPrefix + "LOG_"
	envConfigPrefix    = envPrefix + "CONFIG_"
)

// SkConf is Struct for get and set Suikin Configure.
type SkConf struct {
	// Password is KeyValue Password
	password  map[string]string
	logConfig map[string]string
	config    map[string]string
}

// GetPassword is get password from System Environment.
func (cf *SkConf) GetPassword(key string) string {
	var ok bool
	var value string
	// If does not initialize the map, then initialize the map.
	if cf.password == nil {
		cf.password = map[string]string{}
	}
	if value, ok = cf.password[key]; ok == false {
		// If does not set key-value in the map,
		// then get value from System Envirionment and set key value in the map.
		envkey := envPasswordPrefix + key
		value = os.Getenv(envkey)
		if value != "" {
			// If set System Environment, then set key-value in the map.
			cf.password[key] = value
		}
	}
	return value
}

// GetConfig is get config from System Environment.
func (cf *SkConf) GetConfig(key string) string {
	var ok bool
	var value string
	// If does not initialize the map, then initialize the map.
	if cf.config == nil {
		cf.config = map[string]string{}
	}
	if value, ok = cf.config[key]; ok == false {
		// If does not set key-value in the map,
		// then get value from System Envirionment and set key value in the map.
		envkey := envConfigPrefix + key
		value = os.Getenv(envkey)
		if value != "" {
			// If set System Environment, then set key-value in the map.
			cf.config[key] = value
		}
	}
	return value
}

// GetLogConfig is get logconfig from System Environment.
func (cf *SkConf) GetLogConfig(key string) string {
	var ok bool
	var value string
	// If does not initialize the map, then initialize the map.
	if cf.logConfig == nil {
		cf.logConfig = map[string]string{}
	}
	if value, ok = cf.logConfig[key]; ok == false {
		// If does not set key-value in the map,
		// then get value from System Envirionment and set key value in the map.
		envkey := envLogConfigPrefix + key
		value = os.Getenv(envkey)
		if value != "" {
			// If set System Environment, then set key-value in the map.
			cf.logConfig[key] = value
		}
	}
	return value
}

// SigHup is initialize the map.
func (cf *SkConf) SigHup() {
	cf.password = map[string]string{}
	cf.config = map[string]string{}
	cf.logConfig = map[string]string{}
}
