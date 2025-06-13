package hooks

import (
	"fmt"
)

type HookGraph struct {
	nodes map[string]*HookNode
	hm    *HookManager
}

type HookNode struct {
	Name string
	Next []string
}

func NewHookGraph(hm *HookManager) *HookGraph {
	return &HookGraph{
		nodes: make(map[string]*HookNode),
		hm:    hm,
	}
}

func (hg *HookGraph) AddChain(steps ...string) {
	for i := 0; i < len(steps)-1; i++ {
		hg.AddEdge(steps[i], steps[i+1])
	}
}

func (g *HookGraph) AddEdge(from, to string) {
	if _, ok := g.nodes[from]; !ok {
		g.nodes[from] = &HookNode{Name: from}
	}
	if _, ok := g.nodes[to]; !ok {
		g.nodes[to] = &HookNode{Name: to}
	}
	g.nodes[from].Next = append(g.nodes[from].Next, to)
}

// Execute 從起點hook開始執行整個DAG流程，含中斷控制
func (g *HookGraph) Execute(start string, ctx *HookContext) error {
	visited := make(map[string]bool)
	return g.executeRecursive(start, ctx, visited)
}

func (g *HookGraph) executeRecursive(current string, ctx *HookContext, visited map[string]bool) error {
	if visited[current] {
		return nil // 防止循環
	}
	visited[current] = true

	// 權限檢查（基於 ctx 的 role）
	if !g.hm.CheckPermission(current, ctx) {
		return fmt.Errorf("permission denied for hook %s and role %s", current, ctx.GetUserData("role"))
	}

	result := g.hm.ExecuteHookByName(current, ctx)
	ctx.AddExecutionLog(result) // 紀錄流程與結果

	if result.StopExecution {
		fmt.Printf("[HookGraph] Hook %s 停止後續流程\n", current)
		return nil
	}

	for _, next := range g.nodes[current].Next {
		if err := g.executeRecursive(next, ctx, visited); err != nil {
			return err
		}
	}

	return nil
}
