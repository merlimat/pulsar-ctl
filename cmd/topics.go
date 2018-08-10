// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

// tenantsCmd represents the tenants command
var topicsCmd = &cobra.Command{
	Use:   "topics",
	Short: "Operations about Pulsar's topics",
	Long: `Manage topics

For example, getting the stats for a topic:

    pulsar-ctl topics stats my-topic

Delete a topic

    pulsar-ctl topics delete my-topic
`,
}

func topicsList() {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "Get the list of topics under a namespace",
		// Long: `Manage tenants`,
		Example: "pulsar-ctl topics list my-tenant/my-namespace",
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			RestPrintStringList()
		},
	}

	topicsCmd.AddCommand(listCmd);
}


func init() {
	topicsList()

	rootCmd.AddCommand(topicsCmd)
}
