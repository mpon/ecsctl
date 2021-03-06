package command

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/mpon/ecswalk/internal/pkg/awsapi"
	"github.com/spf13/cobra"
)

// NewCmdGetServices represents the get services command
func NewCmdGetServices() *cobra.Command {
	var clusterFlag string
	cmd := &cobra.Command{
		Use:   "services",
		Short: "get all ECS services specified cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGetServicesCmd(clusterFlag)
		},
	}
	cmd.Flags().StringVarP(&clusterFlag, "cluster", "c", "", "AWS ECS cluster")
	_ = cmd.MarkFlagRequired("cluster")

	return cmd
}

func runGetServicesCmd(clusterName string) error {
	client, err := awsapi.NewClient()
	if err != nil {
		return err
	}

	cluster, err := client.GetEcsCluster(clusterName)
	if err != nil {
		return err
	}

	if err := runGetServices(client, cluster); err != nil {
		return err
	}
	return nil
}

func runGetServices(client *awsapi.Client, cluster *ecs.Cluster) error {
	services, err := client.GetAllEcsServices(cluster)
	if err != nil {
		return nil
	}

	if len(services) == 0 {
		fmt.Printf("%s has no services\n", *cluster.ClusterName)
		return nil
	}

	taskDefinitions, err := client.GetEcsTaskDefinitions(cluster, services)
	if err != nil {
		return err
	}

	ecsServiceInfoList := awsapi.NewEcsServiceInfoList(services, taskDefinitions)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, "Name\tTaskDefinition\tImage\tTag\tDesired\tRunning\t")
	for _, s := range ecsServiceInfoList {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%d\t\n",
			*s.Service.ServiceName,
			s.TaskDefinitionArn(),
			s.DockerImageName(),
			s.DockerImageTag(),
			*s.Service.DesiredCount,
			*s.Service.RunningCount)
	}
	w.Flush()
	return nil
}
