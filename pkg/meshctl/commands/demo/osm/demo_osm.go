package osm

import (
	"context"

	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/demo/common/cleanup"
	"github.com/solo-io/gloo-mesh/pkg/meshctl/commands/demo/common/initialize"
	"github.com/spf13/cobra"
)

const (
	mgmtCluster = "mgmt-cluster"
)

func Command(ctx context.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "osm",
		Short: "Demo Gloo Mesh functionality one OSM control plane deployed",
	}

	cmd.AddCommand(
		initialize.OsmCommand(ctx, mgmtCluster),
		cleanup.Command(ctx, mgmtCluster),
	)

	return cmd
}
