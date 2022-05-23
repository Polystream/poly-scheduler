package polybind

import (
	"context"
	"encoding/json"

	discovery "github.com/gkarthiks/k8s-discovery"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type Polybind struct {
}

const Name = "Polybind"

var _ framework.PostBindPlugin = &Polybind{}

//var _ framework.BindPlugin = &Polybind{}
//var _ framework.PreBindPlugin = &Polybind{}
//var _ framework.PermitPlugin = &Polybind{}
//var _ framework.ScorePlugin = &Polybind{}

var k8s *discovery.K8s
var returned_pod *v1.Pod

func (pl *Polybind) Name() string {
	return Name
}

func New(_ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &Polybind{}, nil
}

/*func (pl *Polybind) Score(ctx context.Context, _ *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.InfoS("Poly score start")
	for _, container := range pod.Spec.Containers {
		for _, port := range container.Ports {
			port.HostPort = 100
			port.ContainerPort = 200
		}
	}

	Update(ctx, pod)

	klog.InfoS("Poly score end")
	return 1, framework.NewStatus(framework.Success)
}

func (poly *Polybind) ScoreExtensions() framework.ScoreExtensions {
	return poly
}

func (poly *Polybind) NormalizeScore(ctx context.Context, state *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	return framework.NewStatus(framework.Success)
}*/

/*func (pl *Polybind) Permit(ctx context.Context, _ *framework.CycleState, pod *v1.Pod, nodeName string) (*framework.Status, time.Duration) {
	klog.InfoS("Poly permit start")
	for _, container := range pod.Spec.Containers {
		for _, port := range container.Ports {
			port.HostPort = 100
			port.ContainerPort = 200
		}
	}

	Update(ctx, pod)

	klog.InfoS("Poly permit end")
	return framework.NewStatus(framework.Success), 1
}/*

/*func (pl *Polybind) PreBind(ctx context.Context, _ *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	klog.InfoS("Poly prebind start")
	for _, container := range pod.Spec.Containers {
		for _, port := range container.Ports {
			port.HostPort = 100
			port.ContainerPort = 200
		}
	}

	Update(ctx, pod)

	klog.InfoS("Poly prebind end")
	return framework.NewStatus(framework.Success)
}*/

/*func (pl *Polybind) Bind(ctx context.Context, _ *framework.CycleState, pod *v1.Pod, nodeName string) *framework.Status {
	klog.InfoS("Poly bind start")
	for _, container := range pod.Spec.Containers {
		for _, port := range container.Ports {
			port.HostPort = 100
			port.ContainerPort = 200
		}
	}

	Update(ctx, pod)

	klog.InfoS("Poly bind end")
	return framework.NewStatus(framework.Success)
}*/

func (pl *Polybind) PostBind(ctx context.Context, _ *framework.CycleState, pod *v1.Pod, nodeName string) {
	klog.InfoS("Poly postBind start")
	for _, container := range pod.Spec.Containers {
		for _, port := range container.Ports {
			port.HostPort = 100
			port.ContainerPort = 200
		}
	}

	Update(ctx, pod)

	klog.InfoS("Poly postBind end")
}

func Update(ctx context.Context, pod *v1.Pod) {
	k8s, _ = discovery.NewK8s()
	core := k8s.Clientset.CoreV1()
	var err error

	var test = "[{\"op\": \"replace\", \"path\": \"/spec/containers/0/ports/0/port\", \"value\": \"9999\"}]"
	payload, err := json.Marshal(test)
	if err != nil {
		klog.Fatal(err)
	}

	pod.APIVersion = "v1"
	returned_pod, err = core.Pods("kube-system").Patch(ctx, "pod-test", types.ApplyPatchType, payload, meta_v1.PatchOptions{
		FieldManager: "poly-scheduler",
	})

	if err != nil {
		klog.Fatal(err)
	}

	klog.InfoS("Updated: %s\n", returned_pod.Name)
}
