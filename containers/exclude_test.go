package containers

import "testing"

func TestIsSelfContainer(t *testing.T) {
	if !isSelfContainer("beacon", "nginx:latest") {
		t.Fatal("expected beacon name match")
	}
	if !isSelfContainer("/beacon", "") {
		t.Fatal("expected trimmed beacon name match")
	}
	if !isSelfContainer("api", "example/beacon:latest") {
		t.Fatal("expected beacon image match")
	}
	if isSelfContainer("api", "nginx:latest") {
		t.Fatal("expected unrelated container to be false")
	}
}

func TestIsExcludedContainer(t *testing.T) {
	excludedContainerNames = []string{"redis", "proxy"}

	if !IsExcludedContainer("beacon", "nginx:latest") {
		t.Fatal("beacon self must always be excluded")
	}
	if !IsExcludedContainer("redis", "redis:7") {
		t.Fatal("expected redis in exclude list")
	}
	if !IsExcludedContainer("proxy", "nginx:latest") {
		t.Fatal("expected proxy in exclude list")
	}
	if IsExcludedContainer("api", "node:20") {
		t.Fatal("api should not be excluded")
	}
}
