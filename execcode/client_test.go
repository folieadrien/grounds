package execcode

import (
	"io"
	"testing"
)

const (
	validEndpoint = "http://localhost:4243"
	invalidEndpoint = ""
)

// FIXME: Verify that excute remove the container and switch to IsBusy: false (call Interrupt ?)
// FIXME: Verify that interrupt remove the container and switch to IsBusy: false
// FIXME: Test to interrupt a real busy container
func TestNewClient(t *testing.T) {
	registry := "test"
	client, err := NewClient(validEndpoint, registry)
	if err != nil {
		t.Fatal(err)
	}
	if client.registry != registry {
		t.Errorf("Expected registry %s. Got %s.", registry, client.registry)
	}
	if client.container != nil {
		t.Errorf("Expected container to be nil. Got %v", client.container)
	}
	if client.IsBusy {
		t.Errorf("New client is busy but it shouldn't.")
	}
}

func TestNewClientInvalidEndpoint(t *testing.T) {
	_, err := NewClient(invalidEndpoint, "")
	if err == nil {
		t.Errorf("Expected error. Got nothing.")
	}
}

func TestExecute(t *testing.T) {
	client, err := NewClient(validEndpoint, "")
	if err != nil {
		t.Fatal(err)
	}
	client.docker = &FakeDockerClient{}
	executed := false
	_, err = client.Execute("ruby", "42", func(out, err io.Reader) {
		executed = true
	})
	if err != nil {
		t.Fatal(err)
	}
	if executed == false {
		t.Errorf("Expected executed to be true. Got false.")
	}
}

func TestExecuteBusyClient(t *testing.T) {
	client, err := NewClient(validEndpoint, "")
	if err != nil {
		t.Fatal(err)
	}
	client.docker = &FakeDockerClient{}
	client.IsBusy = true
	executed := false
	_, err = client.Execute("ruby", "42", func(out, err io.Reader) {
		executed = true
	})
	if err == nil {
		t.Errorf("Expected error. Got nothing.")
	}
	if executed {
		t.Errorf("Block was executed but it shouldn't")
	}
}

func TestInterruptNotBusyClient(t *testing.T) {
	client, err := NewClient(validEndpoint, "")
	if err != nil {
		t.Fatal(err)
	}
	client.docker = &FakeDockerClient{}
	if err := client.Interrupt(); err == nil {
		t.Errorf("Expected error. Got nothing.")
	}
}