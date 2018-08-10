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
	ServiceUrl          string `json:"serviceUrl"`
	ServiceUrlTls       string `json:"serviceUrlTls"`
	BrokerServiceUrl    string `json:"brokerServiceUrl"`
	BrokerServiceUrlTls string `json:"brokerServiceUrlTls"`

	// For given Cluster1(us-west1, us-east1) and Cluster2(us-west2, us-east2)
	// Peer: [us-west1 -> us-west2] and [us-east1 -> us-east2]
	PeerClusterNames []string `json:"peerClusterNames"`
}

const(
	clustersBasePath = "/admin/v2/clusters"
)

func clusterPath(name string) string {
	return clustersBasePath + "/" + name
}

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
			RestPrint(clusterPath(args[0]))
		},
	}

	clustersCmd.AddCommand(getCmd);
}

func clustersCreate() {
	cluster := ClusterData{}

	var createCmd = &cobra.Command{
		Use:     "create",
		Short:   "Configure a new cluster",
		Example: "pulsar-ctl clusters create us-west --service-url pulsar://host:6650",
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			RestPut(clusterPath(args[0]), cluster)
		},
	}

	createCmd.Flags().StringVar(&cluster.BrokerServiceUrl, "service-url", "",
		"Service endpoint for the cluster. eg: pulsar://example.com:6650")
	createCmd.Flags().StringVar(&cluster.BrokerServiceUrlTls, "service-url-tls", "",
		"TLS enabled service endpoint for the cluster. eg: pulsar+ssl://example.com:6651")

	createCmd.Flags().StringVar(&cluster.ServiceUrl, "admin-service-url", "",
		"Admin service endpoint for the cluster. eg: http://example.com:8080")
	createCmd.Flags().StringVar(&cluster.ServiceUrlTls, "admin-service-url-tls", "",
		"TLS enabled admin service endpoint for the cluster. eg: https://example.com:8443")

	createCmd.Flags().StringSliceVar(&cluster.PeerClusterNames, "peer-clusters", nil,
		"Comma separated peer-cluster names")

	createCmd.MarkFlagRequired("service-url")
	clustersCmd.AddCommand(createCmd);
}

func clustersUpdate() {
	cluster := ClusterData{}

	var updateCmd = &cobra.Command{
		Use:     "update",
		Short:   "Update configuration for an existing cluster",
		Example: "pulsar-ctl clusters update us-west --service-url pulsar://host:6650",
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			RestPost(clusterPath(args[0]), cluster)
		},
	}

	updateCmd.Flags().StringVar(&cluster.BrokerServiceUrl, "service-url", "",
		"Service endpoint for the cluster. eg: pulsar://example.com:6650")
	updateCmd.Flags().StringVar(&cluster.BrokerServiceUrlTls, "service-url-tls", "",
		"TLS enabled service endpoint for the cluster. eg: pulsar+ssl://example.com:6651")

	updateCmd.Flags().StringVar(&cluster.ServiceUrl, "admin-service-url", "",
		"Admin service endpoint for the cluster. eg: http://example.com:8080")
	updateCmd.Flags().StringVar(&cluster.ServiceUrlTls, "admin-service-url-tls", "",
		"TLS enabled admin service endpoint for the cluster. eg: https://example.com:8443")

	updateCmd.Flags().StringSliceVar(&cluster.PeerClusterNames, "peer-clusters", nil,
		"Comma separated peer-cluster names")

	updateCmd.MarkFlagRequired("service-url")
	clustersCmd.AddCommand(updateCmd);
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
	clustersCreate()
	clustersUpdate()
	clustersList()
	clustersGet()
	clustersDelete()

	rootCmd.AddCommand(clustersCmd)
}
