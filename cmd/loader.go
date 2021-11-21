package main

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
