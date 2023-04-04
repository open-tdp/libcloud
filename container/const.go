package container

type ContainerState string

const (
	ContainerStateRUNNING    ContainerState = "running"
	ContainerStateREBOOTING  ContainerState = "rebooting"
	ContainerStateTERMINATED ContainerState = "terminated"
	ContainerStatePENDING    ContainerState = "pending"
	ContainerStateUNKNOWN    ContainerState = "unknown"
	ContainerStateSTOPPED    ContainerState = "stopped"
	ContainerStateSUSPENDED  ContainerState = "suspended"
	ContainerStateERROR      ContainerState = "error"
	ContainerStatePAUSED     ContainerState = "paused"
)
