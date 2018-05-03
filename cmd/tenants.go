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
	"encoding/json"
)

var tenantBasePath = "/admin/v2/tenants"

type TenantInfo struct {
	AdminRoles      []string `json:"adminRoles"`
	AllowedClusters []string `json:"allowedClusters"`
}

// tenantsCmd represents the tenants command
var tenantsCmd = &cobra.Command{
	Use:   "tenants",
	Short: "Operations about Pulsar's tenants",
	Long: `Manage tenants

For example creating a tenant:

    pulsar-ctl tenants create my-tenant

Update tenant's configuration:

    pulsar-ctl tenants update my-tenant --allowed-clusters us-west,us-east
`,
}

func tenantsList() {
	var adminRoles []string
	var clusters []string

	var createCmd = &cobra.Command{
		Use:   "list",
		Short: "List all the tenants",
		// Long: `Manage tenants`,
		Example: "pulsar-ctl tenants list",
		Args:    cobra.ExactArgs(0),

		Run: func(cmd *cobra.Command, args []string) {
			RestPrintStringList(tenantBasePath)
		},
	}

	createCmd.Flags().StringSliceVarP(&adminRoles, "admin-roles", "r", nil,
		"Comma separated list of auth principal allowed to administrate the tenant")
	createCmd.Flags().StringSliceVarP(&clusters, "clusters", "c", nil,
		"Comma separated allowed clusters. If empty, the tenant will have access to all clusters")

	tenantsCmd.AddCommand(createCmd);
}

func tenantsGet() {
	var getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get information regarding a tenant",
		// Long: `Manage tenants`,
		Example: "pulsar-ctl tenants get my-tenant",
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(RestGet(tenantPath(args[0])))
		},
	}

	tenantsCmd.AddCommand(getCmd);
}

func tenantPath(name string) string {
	return tenantBasePath + "/" + name
}

func tenantsCreate() {
	var adminRoles []string
	var clusters []string

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new tenant",
		// Long: `Manage tenants`,
		Example: "pulsar-ctl tenants create my-tenant",
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			// By default, if no clusters are provided, allow the tenant to use all clusters
			if len(clusters) == 0 {
				clusters = GetClustersList()
			}

			fmt.Println("clusters: ", clusters)
			var tenant = TenantInfo{AdminRoles: adminRoles, AllowedClusters: clusters}

			fmt.Println("tenant: ", tenant)
			RestPut(tenantBasePath + "/" + args[0], tenant)
		},
	}

	createCmd.Flags().StringSliceVarP(&adminRoles, "admin-roles", "r", nil,
		"Comma separated list of auth principal allowed to administrate the tenant")
	createCmd.Flags().StringSliceVarP(&clusters, "clusters", "c", nil,
		"Comma separated allowed clusters. If empty, the tenant will have access to all clusters")

	tenantsCmd.AddCommand(createCmd);
}

func tenantsUpdate() {
	var adminRoles []string
	var clusters []string

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an existing tenant",
		// Long: `Manage tenants`,
		Example: "pulsar-ctl tenants update my-tenant --allowed-clusters us-west,us-east",
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			// By default, if no clusters are provided, allow the tenant to use all clusters

			var tenant = TenantInfo{}
			json.Unmarshal([]byte(RestGet(tenantBasePath + "/" + args[0])), &tenant)

			if len(clusters) != 0 {
				tenant.AllowedClusters = clusters
			}

			if len(adminRoles) != 0 {
				tenant.AdminRoles = adminRoles
			}

			RestPost(tenantPath(args[0]), tenant)
		},
	}

	updateCmd.Flags().StringSliceVarP(&adminRoles, "admin-roles", "r", nil,
		"Comma separated list of auth principal allowed to administrate the tenant. If empty the current set of roles won't be modified")
	updateCmd.Flags().StringSliceVarP(&clusters, "clusters", "c", nil,
		"Comma separated allowed clusters. If omitted, the current set of clusters will be preserved")

	tenantsCmd.AddCommand(updateCmd);
}

func tenantsDelete() {
	var deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a tenant",
		// Long: `Manage tenants`,
		Example: "pulsar-ctl tenants delete my-tenant",
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			RestDelete(tenantPath(args[0]))
		},
	}

	tenantsCmd.AddCommand(deleteCmd);
}

func init() {
	tenantsList()
	tenantsGet()
	tenantsCreate()
	tenantsUpdate()
	tenantsDelete()

	rootCmd.AddCommand(tenantsCmd)
}
