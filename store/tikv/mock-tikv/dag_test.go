// Copyright 2017 PingCAP, Inc.
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

package mocktikv_test

//
//import (
//	"testing"
//	"time"
//
//	. "github.com/pingcap/check"
//	"github.com/pingcap/kvproto/pkg/coprocessor"
//	"github.com/pingcap/tidb/context"
//	"github.com/pingcap/tidb/distsql"
//	"github.com/pingcap/tidb/distsql/xeval"
//	"github.com/pingcap/tidb/kv"
//	"github.com/pingcap/tidb/model"
//	"github.com/pingcap/tidb/sessionctx"
//	"github.com/pingcap/tidb/store/tikv"
//	"github.com/pingcap/tidb/table"
//	"github.com/pingcap/tidb/util/testkit"
//	"github.com/pingcap/tipb/go-tipb"
//)
//
//func TestT(t *testing.T) {
//	TestingT(t)
//}
//
//type testMockDAGSuite struct {
//	store kv.Storage
//	tk    *testkit.TestKit
//	tbl   table.Table
//}
//
//var _ = Suite(&testMockDAGSuite{})
//
//func (s *testMockDAGSuite) SetUpSuite(c *C) {
//	//storeID, regionIDs, peers := BootstrapWithMultiRegions(cluster, []byte("g"), []byte("n"), []byte("t"))
//	var err error
//	s.store, err = tikv.NewMockTikvStore()
//	c.Assert(err, IsNil)
//
//	s.tk = testkit.NewTestKit(c, s.store)
//	s.tk.MustExec("use test")
//	s.tk.MustExec("create table t (a int, b int)")
//	s.tk.MustExec("insert into t values(1,2),(4,3)")
//	ctx := s.tk.Se.(context.Context)
//	is := sessionctx.GetDomain(ctx).InfoSchema()
//	tbl, err := is.TableByName(model.NewCIStr("test"), model.NewCIStr("t"))
//	c.Assert(err, IsNil)
//	s.tbl = tbl
//}
//
//func (s *testMockDAGSuite) buildTableScanReq() *tipb.Executor {
//	tblInfo := s.tbl.Meta()
//	tblScan := &tipb.TableScan{
//		TableId: tblInfo.ID,
//		Columns: distsql.ColumnsToProto(tblInfo.Columns, tblInfo.PKIsHandle),
//	}
//	return &tipb.Executor{
//		Tp:      tipb.ExecType_TypeTableScan,
//		TblScan: tblScan,
//	}
//}
//
//func buildDAGReq(startTS uint64, executors []*tipb.Executor) *tipb.DAGRequest {
//	_, offset := time.Now().Zone()
//	return &tipb.DAGRequest{
//		StartTs:        startTS,
//		TimeZoneOffset: int64(offset),
//		Flags:          xeval.FlagIgnoreTruncate,
//		Executors:      executors,
//	}
//}
//
//func buildReq(c *C, dag *tipb.DAGRequest, keys ...string) *coprocessor.Request {
//	req := &coprocessor.Request{}
//	req.Ranges = tikv.BuildKeyRanges(keys...).ToPBRanges()
//	var err error
//	req.Data, err = dag.Marshal()
//	c.Assert(err, IsNil)
//	return req
//}
//
//func (s *testMockDAGSuite) TestDAGTableScan(c *C) {
//	executors := []*tipb.Executor{s.buildTableScanReq()}
//	dagReq := buildDAGReq(5, executors)
//	req := buildReq(c, dagReq, "a", "b")
//
//	cli := tikv.GetMockTiKVClient(s.store)
//	res, err := cli.SendCopReqNew("store1", req, 1*time.Millisecond)
//	c.Assert(err, IsNil)
//	c.Assert(res, NotNil)
//}
