package config

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"

	envconfig "github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v3"
)

// From read from an io (file, byte buffer) and populates the specified struct based on
// the inputs and environment variables.
//
//	c := struct {
//		Name string
//		St   struct {
//			Name string
//			Port int
//			List []string
//		} `yaml:"astruct"`
//	}{}
func From(in io.Reader, i interface{}) error {
	// Unmarshall the input to the given struct
	if err := yaml.NewDecoder(in).Decode(i); err != nil {
		return err
	}

	// In order to apply the env var override, we need to process it at this stage.
	return envconfig.Process(context.TODO(), i)
}

// FromConfigMap read a config from the path `./config.yaml`.
// ol
func FromConfigMap(i interface{}) error {
	// because we want to be able to read file from env and configmap
	// might not be given or mounted, we need to continue
	b, err := os.ReadFile("/etc/app/config.yaml")
	if err != nil {
		return envconfig.Process(context.TODO(), i)
	}
	return From(bytes.NewBuffer(b), i)
}

// Process populates the specified struct based on environment variables.
// Structs declare their environment dependencies using the `env` tag with the key being the name of
// the environment variable, case sensitive.
//
//	type MyStruct struct {
//	    A string `env:"A"` // resolves A to A
//	    B string `env:"B,required"` // resolves B to B, errors if $B is unset
//	    C string `env:"C,default=foo"` // resolves C to C, defaults to "foo"
//	    D string `env:"D,required,default=foo"` // error, cannot be required and default
//	    E string `env:""` // error, must specify key
//	}
func Process(i interface{}) error {
	return envconfig.Process(context.TODO(), i)
}

// LookupEnv retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value is returned.
// Otherwise, the returned value will the defaultValue.
func LookupEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return v
}

// LookupEnvAsSlice retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value  is returned as a slice.
// Otherwise, the returned value will the defaultValue.
//
// The expected value of the env should be coma separated as follows:
//
// export MY_ENV_KEY=val1,val2,val3
func LookupEnvAsSlice(key string, defaultValue []string) []string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return strings.Split(v, ",")
}
