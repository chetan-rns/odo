package tasks

import (
	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha2"
	corev1 "k8s.io/api/core/v1"
)

// GenerateDeployFromSourceTask will return a github-status-task
func GenerateDeployFromSourceTask() pipelinev1.Task {
	task := pipelinev1.Task{
		TypeMeta:   createTaskTypeMeta(),
		ObjectMeta: createTaskObjectMeta("deploy-from-source-task"),
		Spec: pipelinev1.TaskSpec{
			Inputs: createInputsForDeployFromSourceTask(),
			TaskSpec: v1alpha2.TaskSpec{
				Steps: createStepsForDeployFromSourceTask(),
			},
		},
	}
	return task
}

func createStepsForDeployFromSourceTask() []pipelinev1.Step {
	return []pipelinev1.Step{
		pipelinev1.Step{
			Container: corev1.Container{
				Name:       "run-kubectl",
				Image:      "quay.io/kmcdermo/k8s-kubectl:latest",
				WorkingDir: "/workspace/source",
				Command:    []string{"kubectl"},
				Args:       argsForRunKubectlStep(),
			},
		},
	}
}

func argsForRunKubectlStep() []string {
	return []string{
		"apply",
		"--dry-run=$(inputs.params.DRYRUN)",
		"-n",
		"$(inputs.params.NAMESPACE)",
		"-k",
		"$(inputs.params.PATHTODEPLOYMENT)",
	}
}

func createInputsForDeployFromSourceTask() *pipelinev1.Inputs {
	return &pipelinev1.Inputs{
		Resources: []pipelinev1.TaskResource{
			createTaskResource("source", "git"),
		},
		Params: []pipelinev1.ParamSpec{
			createStringParamSpecWithDefault(
				"PATHTODEPLOYMENT",
				"Path to the manifest to apply",
				"deploy",
			),
			createStringParamSpec(
				"NAMESPACE",
				"Namespace to deploy into",
			),
			createStringParamSpecWithDefault(
				"DRYRUN",
				"If true run a server-side dryrun.",
				"false",
			),
		},
	}
}
