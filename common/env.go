package common

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	ENV_TEST = `test`
	ENV_DEV  = `dev`
	ENV_PROD = `prod`

	EnvVersion   = `API_VERSION`
	EnvBuildTime = `API_BUILDTIME`

	REQUEST_MAX_HEADER_BYTES       = 1 << 9
	REQUEST_READ_TIMEOUT_DURATION  = 10 * time.Second
	REQUEST_WRITE_TIMEOUT_DURATION = 10 * time.Second
	MIDDLEWARE_TIMEOUT_DURATION    = 60 * time.Second
)

var (
	Version   string // Set on compile
	BuildTime string // Set on compile

	ServerEnv  = os.Getenv("API_ENV")
	ServerPort = os.Getenv("API_PORT")

	ServerEnvVar  = EnvVar{VarName: `API_ENV`, VarVal: ServerEnv}
	ServerPortVar = EnvVar{VarName: `API_PORT`, VarVal: ServerPort}
)

type EnvVar struct {
	VarName string
	VarVal  string
}

func (v EnvVar) String() string {
	return fmt.Sprintf("[ ENV_VAR ] name=%s \t val=%s", v.VarName, v.VarVal)
}

type ErrEnvVarNotSet struct{ EnvVar string }

func (e *ErrEnvVarNotSet) Error() string {
	return fmt.Sprintf(`Environment variable %s is not defined`, e.EnvVar)
}

// Add environment variables here
func EnvVars() []EnvVar {
	return []EnvVar{
		ServerEnvVar,
		ServerPortVar,
		EnvVar{VarName: `API_DB_NAME`, VarVal: DBName},
		EnvVar{VarName: `API_DB_USER`, VarVal: DBUser},
		EnvVar{VarName: `API_DB_PWD`, VarVal: DBPwd},
		EnvVar{VarName: `API_DB_HOST`, VarVal: DBHost},
		EnvVar{VarName: `API_DB_PORT`, VarVal: DBPort},
	}
}

func ExitOnError(e error) {
	if e != nil {
		log.Panicf(`API failed during startup: %s`, e)
	}
}

// Checks if the required environment variables are set
func ValidateEnv() error {
	for _, v := range EnvVars() {
		if v.VarVal == `` {
			return &ErrEnvVarNotSet{v.VarName}
		}
		log.Println(v)
	}
	return nil
}
