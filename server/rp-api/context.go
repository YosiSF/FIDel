//Copyright 2020 WHTCORPS INC ALL RIGHTS RESERVED.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Go originally introduced the context package to the standard library to unify the context propagation
inside the same process. So the entire library- and framework-space can work against the standard context
and we can avoid fragmentation.

*/

package rp_api

import ("context"
		"github.com/YosiSF/FIDel/server/cluster"
)

type contextKey int

const (
	clusterKey contextKey = iota + 1
)

func withSolitonClusterCtx(ctx context.Context, cluster *cluster.RaftSolitonCluster) context.Context {
	return context.WithValue(ctx, clusterKey, cluster)
}

func getSolitonCluster(ctx context.Context) *cluster.RaftSolitonCluster {
	return ctx.Value(clusterKey).(*cluster.RaftSolitonCluster)
}
