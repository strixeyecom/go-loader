package main

/*
Copyright © 2021 keser keser@strixeye.com

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
	`github.com/sirupsen/logrus`
	`github.com/strixeyecom/go-loader/cli`
)

/*
	Created by aomerk at 2021-11-20 for project strixeye
*/

/*
	go-loader is a simple command line tool to do a load test (kind of) to a web application.
*/

func main() {
	// setup logging library
	logrus.SetFormatter(
		&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		},
	)
	cli.Execute()
}
