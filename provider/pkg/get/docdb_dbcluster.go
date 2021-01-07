package get

import (
	"context"

	"github.com/aws/aws-sdk-go/service/docdb"
)

func (g *Getter) getDocDBDBClusterAttributes(ctx context.Context, physicalResourceID string) (map[string]interface{}, error) {
	resp, err := g.docdb.DescribeDBClustersWithContext(ctx, &docdb.DescribeDBClustersInput{
		// Add params here
	})
	_ = resp
	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ClusterResourceId": nil,
		"Endpoint": nil,
		"Port": nil,
		"ReadEndpoint": nil,
	}
	return attrs, nil
}