package cli

/*
Copyright © 2021 strixeye keser@strixeye.com

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

import (
	`bufio`
	`bytes`
	`context`
	`os`
	
	`github.com/sirupsen/logrus`
	"github.com/spf13/cobra"
	`github.com/spf13/viper`
	`github.com/strixeyecom/go-loader/internal/app`
	`github.com/strixeyecom/go-loader/internal/config`
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "start load testing your application",
	Long:  `start load testing your application`,
	Run: func(cmd *cobra.Command, args []string) {
		var config config.App
		err := viper.Unmarshal(&config)
		if err != nil {
			panic(err)
		}
		
		// if no endpoints given via flag, use wordlist file
		if len(config.EndpointWordlist) == 0 {
			// read from wordlist file if given
			if config.EndpointWordlistPath != "" {
				data, err := os.ReadFile(config.EndpointWordlistPath)
				if err != nil {
					panic(err)
				}
				
				scanner := bufio.NewScanner(bytes.NewReader(data))
				scanner.Split(bufio.ScanLines)
				for scanner.Scan() {
					config.EndpointWordlist = append(
						config.EndpointWordlist, scanner.Text(),
					)
				}
			} else {
				// Fallback endpoint list
				config.EndpointWordlist = app.DefaultEndpointList
			}
		}
		for i := 0; i < config.VisitorCount; i++ {
			go func() {
				visitor := app.NewVisitor()
				
				visitor.Endpoints = config.EndpointWordlist
				visitor.TargetScheme = config.TargetScheme
				visitor.TargetHost = config.TargetHost
				
				err := visitor.Run(context.Background())
				if err != nil {
					logrus.Error(err)
				}
			}()
		}
		select {}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	
	// local flags which will only run when this command
	// is called directly, e.g.:
	runCmd.Flags().IntP("visitors", "v", 100, "number of concurrent visitors")
	err := viper.BindPFlag("VISITOR_COUNT", runCmd.Flags().Lookup("visitors"))
	if err != nil {
		panic(err)
	}
	
	runCmd.Flags().IntP("session", "l", 50, "number of requests per visitor")
	err = viper.BindPFlag("SESSION_LENGTH", runCmd.Flags().Lookup("session"))
	if err != nil {
		panic(err)
	}
	runCmd.Flags().StringP(
		"ip-header", "i", app.DefaultIPSourceHeader, "header name to use for IP source",
	)
	err = viper.BindPFlag("IP_SOURCE_HEADER", runCmd.Flags().Lookup("ip-header"))
	if err != nil {
		panic(err)
	}
	runCmd.Flags().StringP(
		"port-header", "p", app.DefaultPortSourceHeader, "header name to use for port source",
	)
	err = viper.BindPFlag("PORT_SOURCE_HEADER", runCmd.Flags().Lookup("port-header"))
	if err != nil {
		panic(err)
	}
	runCmd.Flags().StringP("target-host", "t", "target.omer.beer", "target hostname ")
	err = viper.BindPFlag("TARGET_HOST", runCmd.Flags().Lookup("target-host"))
	if err != nil {
		panic(err)
	}
	runCmd.Flags().StringP("target-scheme", "s", "https", "http/https switch")
	err = viper.BindPFlag("TARGET_SCHEME", runCmd.Flags().Lookup("target-scheme"))
	if err != nil {
		panic(err)
	}
	
	runCmd.Flags().StringP(
		"endpoints-path", "f", "", "given file contains target endpoints. "+
			"flag endpoints overrides this flag",
	)
	err = viper.BindPFlag("ENDPOINT_WORDLIST_PATH", runCmd.Flags().Lookup("endpoints-path"))
	if err != nil {
		panic(err)
	}
	
	runCmd.Flags().StringSliceP("endpoints", "e", nil, "list if directories to visit")
	err = viper.BindPFlag("ENDPOINT_WORDLIST", runCmd.Flags().Lookup("endpoints"))
	if err != nil {
		panic(err)
	}
}
