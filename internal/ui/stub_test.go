package ui

import (
	"context"

	tea "charm.land/bubbletea/v2"
)

var noopCmd tea.Cmd = func() tea.Msg { return nil }

type stubClient struct {
	fetchContainers        func(bool) tea.Cmd
	stopContainer          func(string) tea.Cmd
	startContainer         func(string) tea.Cmd
	restartContainer       func(string) tea.Cmd
	deleteContainer        func(string) tea.Cmd
	checkShellAvail        func(string) tea.Cmd
	execContainer          func(string, string) tea.Cmd
	checkDebugAvail        func(string) tea.Cmd
	debugContainer         func(string) tea.Cmd
	inspectContainer       func(string) tea.Cmd
	inspectContainerExpand func(string) tea.Cmd
	fetchStats             func(string) tea.Cmd
	startLogs              func(context.Context, string, string, bool, string, int) tea.Cmd
	supportsGrep           func() tea.Cmd
	startEvents            func(context.Context, int) tea.Cmd
	fetchContexts          func() tea.Cmd
	switchContext          func(string) tea.Cmd
	pauseContainer         func(string) tea.Cmd
	unpauseContainer       func(string) tea.Cmd
	renameContainer        func(string, string) tea.Cmd
	stopCompose            func(string) tea.Cmd
	startCompose           func(string) tea.Cmd
	restartCompose         func(string) tea.Cmd
}

func (c *stubClient) FetchContainers(all bool) tea.Cmd {
	if c.fetchContainers != nil {
		return c.fetchContainers(all)
	}
	return noopCmd
}
func (c *stubClient) StopContainer(id string) tea.Cmd {
	if c.stopContainer != nil {
		return c.stopContainer(id)
	}
	return noopCmd
}
func (c *stubClient) StartContainer(id string) tea.Cmd {
	if c.startContainer != nil {
		return c.startContainer(id)
	}
	return noopCmd
}
func (c *stubClient) RestartContainer(id string) tea.Cmd {
	if c.restartContainer != nil {
		return c.restartContainer(id)
	}
	return noopCmd
}
func (c *stubClient) DeleteContainer(id string) tea.Cmd {
	if c.deleteContainer != nil {
		return c.deleteContainer(id)
	}
	return noopCmd
}
func (c *stubClient) CheckShellAvailable(id string) tea.Cmd {
	if c.checkShellAvail != nil {
		return c.checkShellAvail(id)
	}
	return noopCmd
}
func (c *stubClient) ExecContainer(id, shell string) tea.Cmd {
	if c.execContainer != nil {
		return c.execContainer(id, shell)
	}
	return noopCmd
}
func (c *stubClient) CheckDebugAvailable(id string) tea.Cmd {
	if c.checkDebugAvail != nil {
		return c.checkDebugAvail(id)
	}
	return noopCmd
}
func (c *stubClient) DebugContainer(id string) tea.Cmd {
	if c.debugContainer != nil {
		return c.debugContainer(id)
	}
	return noopCmd
}
func (c *stubClient) InspectContainer(id string) tea.Cmd {
	if c.inspectContainer != nil {
		return c.inspectContainer(id)
	}
	return noopCmd
}
func (c *stubClient) InspectContainerExpand(id string) tea.Cmd {
	if c.inspectContainerExpand != nil {
		return c.inspectContainerExpand(id)
	}
	return noopCmd
}
func (c *stubClient) FetchStats(id string) tea.Cmd {
	if c.fetchStats != nil {
		return c.fetchStats(id)
	}
	return noopCmd
}
func (c *stubClient) StartLogs(ctx context.Context, id string, tail string, timestamps bool, grep string, gen int) tea.Cmd {
	if c.startLogs != nil {
		return c.startLogs(ctx, id, tail, timestamps, grep, gen)
	}
	return noopCmd
}
func (c *stubClient) SupportsGrep() tea.Cmd {
	if c.supportsGrep != nil {
		return c.supportsGrep()
	}
	return noopCmd
}
func (c *stubClient) StartEvents(ctx context.Context, gen int) tea.Cmd {
	if c.startEvents != nil {
		return c.startEvents(ctx, gen)
	}
	return noopCmd
}
func (c *stubClient) FetchContexts() tea.Cmd {
	if c.fetchContexts != nil {
		return c.fetchContexts()
	}
	return noopCmd
}
func (c *stubClient) SwitchContext(name string) tea.Cmd {
	if c.switchContext != nil {
		return c.switchContext(name)
	}
	return noopCmd
}
func (c *stubClient) PauseContainer(id string) tea.Cmd {
	if c.pauseContainer != nil {
		return c.pauseContainer(id)
	}
	return noopCmd
}
func (c *stubClient) UnpauseContainer(id string) tea.Cmd {
	if c.unpauseContainer != nil {
		return c.unpauseContainer(id)
	}
	return noopCmd
}
func (c *stubClient) RenameContainer(id, newName string) tea.Cmd {
	if c.renameContainer != nil {
		return c.renameContainer(id, newName)
	}
	return noopCmd
}
func (c *stubClient) StopCompose(project string) tea.Cmd {
	if c.stopCompose != nil {
		return c.stopCompose(project)
	}
	return noopCmd
}
func (c *stubClient) StartCompose(project string) tea.Cmd {
	if c.startCompose != nil {
		return c.startCompose(project)
	}
	return noopCmd
}
func (c *stubClient) RestartCompose(project string) tea.Cmd {
	if c.restartCompose != nil {
		return c.restartCompose(project)
	}
	return noopCmd
}
