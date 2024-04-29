/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package srvhelper

import (
	"reflect"

	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/models"
)

type ConnectionModelInfo interface {
	ModelInfo
	GetConnectionId(any) uint64
}

// ConnectionSrvHelper
type AnyConnectionSrvHelper struct {
	*AnyModelSrvHelper
	connModelInfo        ConnectionModelInfo
	scopeModelInfo       ScopeModelInfo
	scopeConfigModelInfo ScopeConfigModelInfo
	pluginName           string
}

// NewConnectionSrvHelper creates a ConnectionDalHelper for connection management
func NewAnyConnectionSrvHelper(
	basicRes context.BasicRes,
	connModelInfo ConnectionModelInfo,
	scopeModelInfo ScopeModelInfo,
	scopeConfigModelInfo ScopeConfigModelInfo,
	pluginName string,
) *AnyConnectionSrvHelper {
	return &AnyConnectionSrvHelper{
		AnyModelSrvHelper:    NewAnyModelSrvHelper(basicRes, connModelInfo, nil),
		connModelInfo:        connModelInfo,
		scopeModelInfo:       scopeModelInfo,
		scopeConfigModelInfo: scopeConfigModelInfo,
		pluginName:           pluginName,
	}
}

func (connSrv *AnyConnectionSrvHelper) DeleteConnectionAny(connection any) (refs *DsRefs, err errors.Error) {
	err = connSrv.NoRunningPipeline(func(tx dal.Transaction) errors.Error {
		// make sure no blueprint is using the connection
		connectionId := connSrv.connModelInfo.GetConnectionId(connection)
		refs = toDsRefs(connSrv.getAllBlueprinsByConnection(connectionId))
		if refs != nil {
			return errors.Conflict.New("Cannot delete the connection because it is referenced by blueprints")
		}
		scopeCount := errors.Must1(connSrv.db.Count(dal.From(connSrv.scopeModelInfo.TableName()), dal.Where("connection_id = ?", connectionId)))
		if scopeCount > 0 {
			return errors.Conflict.New("Please delete all data scope(s) before you delete this Data Connection.")
		}
		errors.Must(tx.Delete(connection))
		sc := connSrv.scopeConfigModelInfo.New()
		if reflect.TypeOf(sc) != reflect.TypeOf(new(NoScopeConfig)) {
			errors.Must(connSrv.db.Delete(sc, dal.Where("connection_id = ?", connectionId)))
		}
		return nil
	})
	return
}

func (connSrv *AnyConnectionSrvHelper) getAllBlueprinsByConnection(connectionId uint64) []*models.Blueprint {
	blueprints := make([]*models.Blueprint, 0)
	errors.Must(connSrv.db.All(
		&blueprints,
		dal.From("_devlake_blueprints bp"),
		dal.Join("JOIN _devlake_blueprint_connections cn ON cn.blueprint_id = bp.id"),
		dal.Where(
			"mode = ? AND cn.connection_id = ? AND cn.plugin_name = ?",
			"NORMAL",
			connectionId,
			connSrv.pluginName,
		),
	))
	return blueprints
}
