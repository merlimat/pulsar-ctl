package cmd

import (
	"github.com/spf13/cobra"
	"encoding/json"
	"fmt"
)

func clusterFailureDomainsPath(cluster string, domain string) string {
	return fmt.Sprintf("%s/%s/failureDomains/%s", clustersBasePath, cluster, domain)
}

type FailureDomain struct {
	Brokers []string `json:"brokers"`
}

var failureDomainsCmd = &cobra.Command{
	Use:   "failure-domains",
	Short: "Manage failure domains for clusters",
	Example: `pulsar-ctl clusters failure-domains create us-west my-domain
		--brokers pulsar://host-1:6650,pulsar://host-2:6650`,
}

func failuresDomainsCreate() {
	failureDomain := FailureDomain{}

	var createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new failure domain for a cluster",
		Example: `pulsar-ctl clusters failure-domains create us-west my-domain
		--brokers pulsar://host-1:6650,pulsar://host-2:6650`,
		Args: cobra.ExactArgs(2),

		Run: func(cmd *cobra.Command, args []string) {
			clusterName := args[0]
			domainName := args[1]

			RestPost(clusterFailureDomainsPath(clusterName, domainName), failureDomain)
		},
	}

	createCmd.Flags().StringSliceVarP(&failureDomain.Brokers, "brokers", "b", nil,
		"Comma separated list of brokers to be included in the domain")

	failureDomainsCmd.AddCommand(createCmd)
}

func failuresDomainsUpdate() {
	failureDomain := FailureDomain{}

	var updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an existing failure domain for a cluster",
		Example: `pulsar-ctl clusters failure-domains update us-west my-domain
		--brokers pulsar://host-1:6650,pulsar://host-2:6650`,
		Args: cobra.ExactArgs(2),

		Run: func(cmd *cobra.Command, args []string) {
			clusterName := args[0]
			domainName := args[1]

			RestPost(clusterFailureDomainsPath(clusterName, domainName), failureDomain)
		},
	}

	updateCmd.Flags().StringSliceVarP(&failureDomain.Brokers, "brokers", "b", nil,
		"Comma separated list of brokers to be included in the domain")

	failureDomainsCmd.AddCommand(updateCmd)
}

func failuresDomainsList() {
	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List existing failure domains for a cluster",
		Example: `pulsar-ctl clusters failure-domains list us-west`,
		Args:    cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			clusterName := args[0]

			response := RestGet(clusterFailureDomainsPath(clusterName, ""))
			var domains map[string]interface{}
			json.Unmarshal([]byte(response), &domains)
			for name, _ := range domains {
				fmt.Println(name)
			}
		},
	}

	failureDomainsCmd.AddCommand(listCmd)
}

func failuresDomainsGet() {
	var listCmd = &cobra.Command{
		Use:     "get",
		Short:   "Get the information regarding a particular failure domain",
		Example: `pulsar-ctl clusters failure-domains get us-west my-domain`,
		Args:    cobra.ExactArgs(2),

		Run: func(cmd *cobra.Command, args []string) {
			clusterName := args[0]
			domainName := args[1]

			RestPrint(clusterFailureDomainsPath(clusterName, domainName))
		},
	}

	failureDomainsCmd.AddCommand(listCmd)
}

func failuresDomainsDelete() {
	var listCmd = &cobra.Command{
		Use:     "delete",
		Short:   "Deletes an existing failure-domain",
		Example: `pulsar-ctl clusters failure-domains delete us-west my-domain`,
		Args:    cobra.ExactArgs(2),

		Run: func(cmd *cobra.Command, args []string) {
			clusterName := args[0]
			domainName := args[1]

			RestDelete(clusterFailureDomainsPath(clusterName, domainName))
		},
	}

	failureDomainsCmd.AddCommand(listCmd)
}

func init() {
	failuresDomainsCreate()
	failuresDomainsUpdate()
	failuresDomainsList()
	failuresDomainsGet()
	failuresDomainsDelete()

	clustersCmd.AddCommand(failureDomainsCmd)
}
