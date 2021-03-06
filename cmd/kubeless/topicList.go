/*
Copyright 2016 Skippbox, Ltd.

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

package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var topicListCmd = &cobra.Command{
	Use:     "list FLAG",
	Aliases: []string{"ls"},
	Short:   "list all topics created in Kubeless",
	Long:    `list all topics created in Kubeless`,
	Run: func(cmd *cobra.Command, args []string) {
		ctlNamespace, err := cmd.Flags().GetString("kafka-namespace")
		if err != nil {
			logrus.Fatal(err)
		}
		command := []string{"bash", "/opt/bitnami/kafka/bin/kafka-topics.sh", "--zookeeper", "zookeeper." + ctlNamespace + ":2181", "--list"}
		execCommand(command, ctlNamespace)
	},
}
