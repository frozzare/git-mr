package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunners(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/index.json")
	runners, err := gitlab.Runners(0, 2)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(runners), 2)
	defer ts.Close()
}

func TestAllRunners(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/index.json")
	runners, err := gitlab.AllRunners(0, 0)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(runners), 2)
	defer ts.Close()
}

func TestRunner(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/show.json")
	runner, err := gitlab.Runner(6)

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Runner), runner)
	assert.Equal(t, runner.Id, 6)
	assert.Equal(t, runner.IsShared, false)
	assert.Equal(t, runner.Description, "test-1-20150125")
	assert.Equal(t, runner.Token, "205086a8e3b9a2b818ffac9b89d102")
	assert.Equal(t, len(runner.TagList), 2)
	assert.Equal(t, runner.ContactedAt, "2016-01-25T16:39:48.066Z")
	defer ts.Close()
}

func TestUpdateRunner(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/update.json")

	runner := Runner{
		Description: "New Runner Description",
	}

	runner_resp, err := gitlab.UpdateRunner(6, &runner)
	assert.Equal(t, err, nil)
	assert.IsType(t, new(Runner), runner_resp)
	assert.Equal(t, runner_resp.Description, "New Runner Description")
	defer ts.Close()
}

func TestDeleteRunner(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/delete.json")
	resp, err := gitlab.DeleteRunner(6)

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Runner), resp)
	assert.IsType(t, resp.Id, 6)
	defer ts.Close()
}

func TestProjectRunners(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/projects/index.json")
	runners, err := gitlab.ProjectRunners("1", 0, 2)

	assert.Equal(t, err, nil)
	assert.Equal(t, len(runners), 2)
	assert.Equal(t, runners[0].IsShared, false)
	defer ts.Close()
}

func TestEnableProjectRunner(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/projects/enable.json")
	runner, err := gitlab.EnableProjectRunner("1", 9)

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Runner), runner)
	assert.Equal(t, runner.Id, 9)
	defer ts.Close()
}

func TestDisableProjectRunner(t *testing.T) {
	ts, gitlab := Stub("stubs/runners/projects/disable.json")
	runner, err := gitlab.DisableProjectRunner("1", 9)

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Runner), runner)
	assert.Equal(t, runner.Id, 9)
	defer ts.Close()
}
