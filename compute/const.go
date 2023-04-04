package compute

type NodeState string

const (
	NodeStateRUNNING       NodeState = "running"
	NodeStateSTARTING      NodeState = "starting"
	NodeStateREBOOTING     NodeState = "rebooting"
	NodeStateTERMINATED    NodeState = "terminated"
	NodeStatePENDING       NodeState = "pending"
	NodeStateUNKNOWN       NodeState = "unknown"
	NodeStateSTOPPING      NodeState = "stopping"
	NodeStateSTOPPED       NodeState = "stopped"
	NodeStateSUSPENDED     NodeState = "suspended"
	NodeStateERROR         NodeState = "error"
	NodeStatePAUSED        NodeState = "paused"
	NodeStateRECONFIGURING NodeState = "reconfiguring"
	NodeStateMIGRATING     NodeState = "migrating"
	NodeStateNORMAL        NodeState = "normal"
	NodeStateUPDATING      NodeState = "updating"
)

type NodeImageState string

const (
	NodeImageStateACCEPTED NodeImageState = "accepted"
	NodeImageStatePENDING  NodeImageState = "pending"
	NodeImageStateREJECTED NodeImageState = "rejected"
)

type StorageVolumeState string

const (
	StorageVolumeStateAVAILABLE StorageVolumeState = "available"
	StorageVolumeStateERROR     StorageVolumeState = "error"
	StorageVolumeStateINUSE     StorageVolumeState = "inuse"
	StorageVolumeStateCREATING  StorageVolumeState = "creating"
	StorageVolumeStateDELETING  StorageVolumeState = "deleting"
	StorageVolumeStateDELETED   StorageVolumeState = "deleted"
	StorageVolumeStateBACKUP    StorageVolumeState = "backup"
	StorageVolumeStateATTACHING StorageVolumeState = "attaching"
	StorageVolumeStateUNKNOWN   StorageVolumeState = "unknown"
	StorageVolumeStateMIGRATING StorageVolumeState = "migrating"
	StorageVolumeStateUPDATING  StorageVolumeState = "updating"
)

type VolumeSnapshotState string

const (
	VolumeSnapshotStateAVAILABLE VolumeSnapshotState = "available"
	VolumeSnapshotStateERROR     VolumeSnapshotState = "error"
	VolumeSnapshotStateCREATING  VolumeSnapshotState = "creating"
	VolumeSnapshotStateDELETING  VolumeSnapshotState = "deleting"
	VolumeSnapshotStateRESTORING VolumeSnapshotState = "restoring"
	VolumeSnapshotStateUNKNOWN   VolumeSnapshotState = "unknown"
	VolumeSnapshotStateUPDATING  VolumeSnapshotState = "updating"
)

type OSType string

const (
	Linux   OSType = "linux"
	MacOS   OSType = "macos"
	Windows OSType = "windows"
	Unix    OSType = "unix"
	Unknown OSType = "unknown"
)

type Architecture string

const (
	I386    Architecture = "i386"
	X86_X64 Architecture = "x86_64"
	Arm     Architecture = "arm"
	Arm64   Architecture = "arm64"
)

type ComputeError string

const (
	DeploymentError          ComputeError = "DeploymentError"
	KeyPairError             ComputeError = "KeyPairError"
	KeyPairDoesNotExistError ComputeError = "KeyPairDoesNotExistError"
)
