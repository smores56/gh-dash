package data

import (
	"reflect"
	"testing"
)

// Issue queries used to embed Repository, whose branchProtectionRules field a
// private repository on the Free plan cannot read. Keep the issue-only type slim.
func TestIssueRepositoryOmitsBranchProtectionRules(t *testing.T) {
	field, ok := reflect.TypeOf(IssueData{}).FieldByName("Repository")
	if !ok {
		t.Fatal("IssueData has no Repository field")
	}
	if _, ok := field.Type.FieldByName("BranchProtectionRules"); ok {
		t.Fatal("issue queries must not request branchProtectionRules")
	}
}
