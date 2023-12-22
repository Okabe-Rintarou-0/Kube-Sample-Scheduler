package plugins

import (
	"context"
	"fmt"
	"math/rand"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "MySchedulerPlugin"

type MyScheduler struct {
	handle framework.Handle
}

// Score implements framework.ScorePlugin.
func (*MyScheduler) Score(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) (int64, *framework.Status) {
	scores := int64(rand.Intn(10))
	fmt.Printf("score pod %s with %d\n", p.Name, scores)
	return scores, framework.NewStatus(framework.Success, "")
}

// ScoreExtensions implements framework.ScorePlugin.
func (*MyScheduler) ScoreExtensions() framework.ScoreExtensions {
	return nil
}

// PreFilter implements framework.PreFilterPlugin.
func (*MyScheduler) PreFilter(ctx context.Context, state *framework.CycleState, p *v1.Pod) *framework.Status {
	fmt.Printf("prefilter pod: %s\n", p.Name)
	return framework.NewStatus(framework.Success, "")
}

// PreFilterExtensions implements framework.PreFilterPlugin.
func (*MyScheduler) PreFilterExtensions() framework.PreFilterExtensions {
	panic("unimplemented")
}

// PreBind implements framework.PreBindPlugin.
func (*MyScheduler) PreBind(ctx context.Context, state *framework.CycleState, p *v1.Pod, nodeName string) *framework.Status {
	fmt.Printf("pre bind pod %s to node: %s\n", p.Name, nodeName)
	return framework.NewStatus(framework.Success, "")
}

// Filter implements framework.FilterPlugin.
func (*MyScheduler) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	fmt.Printf("filter pod: %v, node: %v\n", pod.Name, nodeInfo.Node().Name)
	return framework.NewStatus(framework.Success, "")
}

func (s *MyScheduler) Name() string {
	return Name
}

// var _ = framework.QueueSortPlugin(&MyScheduler{})

// type PluginFactory = func(configuration *runtime.Unknown, f FrameworkHandle) (Plugin, error)
func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	fmt.Println("My plugin start")
	return &MyScheduler{
		handle: h,
	}, nil
}
