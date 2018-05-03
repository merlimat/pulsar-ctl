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
	"fmt"
)

type ClusterData struct {
	ServiceUrl          string `json:serviceUrl`
	ServiceUrlTls       string `json:serviceUrlTls`
	BrokerServiceUrl    string `json:brokerServiceUrl`
	BrokerServiceUrlTls string `json:brokerServiceUrlTls`

	// For given Cluster1(us-west1, us-east1) and Cluster2(us-west2, us-east2)
	// Peer: [us-west1 -> us-west2] and [us-east1 -> us-east2]
	PeerClusterNames []string `json:PeerClusterNames`
}

var clustersBasePath = "/admin/v2/clusters"

var clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Operations about Pulsar's clusters",
	Long:  `Manage Clusters`,
}

func GetClustersList() []string {
	filtered := []string{}
	for _, cluster := range RestGetStringList(clustersBasePath) {
		if cluster != "global" {
			filtered = append(filtered, cluster)
		}
	}
	return filtered
}

func clustersList() {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all the clusters",
		// Long: `Manage tenants`,
		Example: "pulsar-ctl clusters list",
		Args:    cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			for _, cluster := range GetClustersList() {
				fmt.Println(cluster)
			}
		},
	}

	clustersCmd.AddCommand(listCmd);
}

func clustersGet() {
	var getCmd = &cobra.Command{
		Use:     "get",
		Short:   "Get information regarding a cluster",
		Example: "pulsar-ctl clusters get us-west",
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(RestGet(clustersBasePath + "/" + args[0]))
		},
	}

	clustersCmd.AddCommand(getCmd);
}

func clustersDelete() {
	var deleteCmd = &cobra.Command{
		Use:     "delete",
		Short:   "Delete an existing cluster",
		Example: "pulsar-ctl clusters delete us-west",
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			RestDelete(clustersBasePath + "/" + args[0])
		},
	}

	clustersCmd.AddCommand(deleteCmd);
}

func init() {
	clustersList()
	clustersGet()
	clustersDelete()

	rootCmd.AddCommand(clustersCmd)
}
