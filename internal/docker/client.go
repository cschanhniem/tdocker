package docker

import (
	"context"

	tea "charm.land/bubbletea/v2"
)

type ContainerClient interface {
	FetchContainers(all bool) tea.Cmd
	StopContainer(id string) tea.Cmd
	StartContainer(id string) tea.Cmd
	RestartContainer(id string) tea.Cmd
	DeleteContainer(id string) tea.Cmd
	PauseContainer(id string) tea.Cmd
	UnpauseContainer(id string) tea.Cmd
	RenameContainer(id, newName string) tea.Cmd
}

type ComposeClient interface {
	StopCompose(project string) tea.Cmd
	StartCompose(project string) tea.Cmd
	RestartCompose(project string) tea.Cmd
}

type ExecClient interface {
	CheckShellAvailable(id string) tea.Cmd
	ExecContainer(id, shell string) tea.Cmd
	CheckDebugAvailable(id string) tea.Cmd
	DebugContainer(id string) tea.Cmd
}

type InspectClient interface {
	InspectContainer(id string) tea.Cmd
	InspectContainerExpand(id string) tea.Cmd
}

type LogsClient interface {
	StartLogs(ctx context.Context, id string, tail string, timestamps bool, grep string, gen int) tea.Cmd
	SupportsGrep() tea.Cmd
}

type StatsClient interface {
	FetchStats(id string) tea.Cmd
}

type EventsClient interface {
	StartEvents(ctx context.Context, gen int) tea.Cmd
}

type ContextClient interface {
	FetchContexts() tea.Cmd
	SwitchContext(name string) tea.Cmd
}

type Client interface {
	ContainerClient
	ComposeClient
	ExecClient
	InspectClient
	LogsClient
	StatsClient
	EventsClient
	ContextClient
}
