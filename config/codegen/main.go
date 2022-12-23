//go:build codegen
// +build codegen

package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"text/template"
)

const (
	sharedConfigType  = "&SharedConfig{}"
	envConfigType     = "&EnvConfig{}"
	awsConfigType     = "&awsConfig{}"
	ec2IMDSRegionType = "&UseEC2IMDSRegion{}"
	loadOptionsType   = "&LoadOptions{}"
)

var implAsserts = map[string][]string{
	"sharedConfigProfileProvider":              {envConfigType, loadOptionsType},
	"sharedConfigFilesProvider":                {envConfigType, loadOptionsType},
	"customCABundleProvider":                   {envConfigType, sharedConfigType, loadOptionsType},
	"regionProvider":                           {envConfigType, sharedConfigType, loadOptionsType, ec2IMDSRegionType},
	"credentialsProviderProvider":              {loadOptionsType},
	"defaultRegionProvider":                    {loadOptionsType},
	"credentialsCacheOptionsProvider":          {loadOptionsType},
	"bearerAuthTokenProviderProvider":          {loadOptionsType},
	"bearerAuthTokenCacheOptionsProvider":      {loadOptionsType},
	"ssoTokenProviderOptionsProvider":          {loadOptionsType},
	"processCredentialOptions":                 {loadOptionsType},
	"ec2RoleCredentialOptionsProvider":         {loadOptionsType},
	"endpointCredentialOptionsProvider":        {loadOptionsType},
	"assumeRoleCredentialOptionsProvider":      {loadOptionsType},
	"webIdentityRoleCredentialOptionsProvider": {loadOptionsType},
	"httpClientProvider":                       {loadOptionsType},
	"apiOptionsProvider":                       {loadOptionsType},
	"retryProvider":                            {loadOptionsType},
	"endpointResolverProvider":                 {loadOptionsType},
	"loggerProvider":                           {loadOptionsType},
	"clientLogModeProvider":                    {loadOptionsType},
	"logConfigurationWarningsProvider":         {loadOptionsType},
	"ec2IMDSRegionProvider":                    {loadOptionsType},
	"defaultsModeProvider":                     {envConfigType, sharedConfigType, loadOptionsType},
	"defaultsModeIMDSClientProvider":           {loadOptionsType},
	"retryMaxAttemptsProvider":                 {envConfigType, sharedConfigType, loadOptionsType},
	"retryModeProvider":                        {envConfigType, sharedConfigType, loadOptionsType},
}

var tplProviderTests = template.Must(template.New("tplProviderTests").Funcs(map[string]interface{}{
	"SortKeys": func(m map[string][]string) []string {
		var keys []string
		for k := range m {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		return keys
	},
}).Parse(`
// Code generated by oss-sdk-go/config. DO NOT EDIT.

package config

{{ $sortedKeys := SortKeys . }}
{{- range $_, $provider := $sortedKeys }}
	{{- $implementors := index $ $provider -}}
	// {{ $provider }} implementor assertions
	var (
		{{- range $_, $impl := $implementors }}
		_ {{ $provider }} = {{ $impl }}
		{{- end }}
	)
{{ end -}}
`))

var config = struct {
	OutputPath string
}{}

func init() {
	flag.StringVar(&config.OutputPath, "output", "", "output file path")
	flag.Parse()
}

func main() {
	if len(config.OutputPath) == 0 {
		panic(fmt.Errorf("output path must be specified"))
	}

	file, err := os.OpenFile(config.OutputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0744)
	if err != nil {
		panic(fmt.Errorf("failed to open file: %v", err))
	}
	defer file.Close()

	err = tplProviderTests.Execute(file, implAsserts)
	if err != nil {
		panic(fmt.Errorf("failed to execute template: %v", err))
	}
}