package nexus_plugin

type NexusPlugin interface {
	ExecuteTask(taskName string, params map[string]string) (string, error)
}
