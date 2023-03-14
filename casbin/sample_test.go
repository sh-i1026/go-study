package casbin

import (
	"strconv"
	"testing"

	casbin "github.com/casbin/casbin/v2"
)

func TestSampleRbacProcess(t *testing.T) {
	e, _ := casbin.NewEnforcer("conf/allow_sample_rbac.conf", "policy/sample_rbac.csv")
	testRbac(t, e, [][]interface{}{
		{"alice", "data1", "read"},
		{"alice", "data1", "write"},
		{"alice", "data1", "delete"},
		{"mike", "data1", "write"},
		{"mike", "data1", "delete"},
	}, []bool{
		true, true, true, true, true,
	})
}

func TestDenyListSampleRbacProcess(t *testing.T) {
	e, _ := casbin.NewEnforcer("conf/deny_sample_rbac.conf", "policy/sample_rbac.csv")
	testRbac(t, e, [][]interface{}{
		{"alice", "data1", "read"},
		{"alice", "data1", "write"},
		{"alice", "data1", "delete"},
		{"mike", "data1", "write"},
		{"mike", "data1", "delete"},
		// denyに合致しないため許容させる
		{"john", "data1", "delete"},
	}, []bool{
		true, true, true, true, false, true,
	})
}

func TestAllowDenySampleRbacProcess(t *testing.T) {
	e, _ := casbin.NewEnforcer("conf/allow_deny_sample_rbac.conf", "policy/sample_rbac.csv")
	testRbac(t, e, [][]interface{}{
		{"alice", "data1", "read"},
		{"alice", "data1", "write"},
		{"alice", "data1", "delete"},
		{"mike", "data1", "write"},
		// denyに合致しdenyが優先される
		{"mike", "data1", "delete"},
		{"john", "data1", "delete"},
	}, []bool{
		true, true, true, true, false, false,
	})
}

func TestSplitPlicySampleRbacProcess(t *testing.T) {
	e, _ := casbin.NewEnforcer("conf/split_policy_rbac.conf", "policy/split_policy_sample_rbac.csv")
	enforceContext := casbin.NewEnforceContext("2")
	enforceContext.EType = "e"
	testRbac(t, e, [][]interface{}{
		// flag立っている場合(p2指定)
		{enforceContext, "john", "data1", "write", "1"},
		// flag立っていない場合(p2指定)
		{enforceContext, "john", "data2", "write", "0"},
		// pを指定
		{"alice", "data1", "read"},
	}, []bool{
		true, false, true,
	})
}

func TestCheckByFuncRbacProcess(t *testing.T) {
	e, _ := casbin.NewEnforcer("conf/with_function_rbac.conf", "policy/with_function_sample_rbac.csv")
	enforceContext := casbin.NewEnforceContext("3")
	enforceContext.EType = "e"
	e.AddFunction("CheckByBit", ChckeByBitOperation)
	testRbac(t, e, [][]interface{}{
		{enforceContext, "john", "data1", "delete", "1", "1"},
		{enforceContext, "john", "data2", "delete", "1", "1"},
	}, []bool{
		false, true,
	})
}

func ChckeByBitOperation(args ...interface{}) (interface{}, error) {
	requestValue, err := strconv.Atoi(args[0].(string))
	if err != nil {
		return false, err
	}

	policyValue, err := strconv.Atoi(args[1].(string))
	if err != nil {
		return false, err
	}

	return requestValue&policyValue > 0, nil
}

func testRbac(t *testing.T, e *casbin.Enforcer, requests [][]interface{}, expectResults []bool) {
	t.Helper()
	results, _ := e.BatchEnforce(requests)
	if len(results) != len(expectResults) {
		t.Errorf("結果の値：%v 期待する値： %v", results, expectResults)
	}
	for i, v := range results {
		if v != expectResults[i] {
			t.Errorf("結果の値：%v 期待する値： %v", results, expectResults)
		}
	}
}
