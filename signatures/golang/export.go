package main

import "github.com/aquasecurity/tracee/types/detect"

// ExportedSignatures fulfills the goplugins contract required by the rule-engine
// this is a list of signatures that this plugin exports
var ExportedSignatures = []detect.Signature{
	&StdioOverSocket{},
	&K8sApiConnection{},
	&AslrInspection{},
	&DroppedExecutable{},
	&ProcMemCodeInjection{},
	&DockerAbuse{},
	&ScheduledTaskModification{},
	&LdPreload{},
	&CgroupNotifyOnReleaseModification{},
	&DefaultLoaderModification{},
	&SudoersModification{},
	&SchedDebugRecon{},
	&SystemRequestKeyConfigModification{},
	&CgroupReleaseAgentModification{},
	&RcdModification{},
	&CorePatternModification{},
	&ProcKcoreRead{},
	&ProcMemAccess{},
	&HiddenFileCreated{},
	&AntiDebuggingPtraceme{},
	&PtraceCodeInjection{},
	&ProcessVmWriteCodeInjection{},
	&DiskMount{},
	&DynamicCodeLoading{},
	&FilelessExecution{},
	&IllegitimateShell{},
	&KernelModuleLoading{},
	&KubernetesCertificateTheftAttempt{},
	&ProcFopsHooking{},
	&SyscallTableHooking{},
	&K8SServiceAccountToken{},
}
