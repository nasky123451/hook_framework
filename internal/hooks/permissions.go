package hooks

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	permissionTree      map[string][]string
	resolvedPermissions map[string]map[string]bool
	once                sync.Once
)

// 載入權限繼承圖（從 JSON 檔）
func loadPermissionTreeFromFile(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read permission tree: %w", err)
	}
	err = json.Unmarshal(file, &permissionTree)
	if err != nil {
		return fmt.Errorf("invalid permission tree format: %w", err)
	}
	return nil
}

func buildResolvedPermissions() {
	resolvedPermissions = make(map[string]map[string]bool)
	for perm := range permissionTree {
		visited := make(map[string]bool)
		resolvePermission(perm, visited)
		resolvedPermissions[perm] = visited
	}
}

func resolvePermission(perm string, visited map[string]bool) {
	if visited[perm] {
		return
	}
	visited[perm] = true
	for _, child := range permissionTree[perm] {
		resolvePermission(child, visited)
	}
}

func InitPermissions(path string) error {
	once = sync.Once{} // 允許重載
	err := loadPermissionTreeFromFile(path)
	if err != nil {
		return err
	}
	once.Do(buildResolvedPermissions)
	return nil
}

func GetInheritablePermissions(perm string) map[string]bool {
	once.Do(buildResolvedPermissions)
	return resolvedPermissions[perm]
}

func GetFormattedRoleTree() []string {
	var lines []string

	// 找出所有角色被繼承的情況
	childSet := map[string]bool{}
	for _, children := range permissionTree {
		for _, c := range children {
			childSet[c] = true
		}
	}

	// 找出 root nodes（沒被任何人繼承的角色）
	var roots []string
	for role := range permissionTree {
		if !childSet[role] {
			roots = append(roots, role)
		}
	}

	// 遞迴渲染角色（允許重複展開，但避免無限循環）
	var render func(role string, indent int, path map[string]bool)
	render = func(role string, indent int, path map[string]bool) {
		prefix := strings.Repeat("  ", indent)
		lines = append(lines, fmt.Sprintf("%s- %s", prefix, role))

		// 防止遞迴循環（如 cyclic inheritance）
		if path[role] {
			lines = append(lines, fmt.Sprintf("%s  ⚠️ Circular reference detected!", prefix))
			return
		}

		// 複製 path map（避免共用）
		newPath := make(map[string]bool)
		for k, v := range path {
			newPath[k] = v
		}
		newPath[role] = true

		for _, child := range permissionTree[role] {
			render(child, indent+1, newPath)
		}
	}

	for _, root := range roots {
		render(root, 0, map[string]bool{})
	}
	return lines
}

func ExportPermissionTree(as string) interface{} {
	once.Do(buildResolvedPermissions)
	switch as {
	case "json":
		return permissionTree
	case "text":
		out := ""
		for parent, children := range permissionTree {
			out += fmt.Sprintf("%s -> %v\n", parent, children)
		}
		return out
	default:
		return resolvedPermissions
	}
}

func (hm *HookManager) CheckPermissions(ctx *HookContext, handlers []HookHandler) bool {
	perm, ok := ctx.GetUserData("permissions").(string)
	if !ok {
		return false
	}
	userPermissions := GetInheritablePermissions(perm)

	for _, handler := range handlers {
		hookName := handler.Name()

		var requiredPermission string
		for _, meta := range registeredMetadata {
			if meta.Name == hookName {
				requiredPermission = meta.Permissions
				break
			}
		}

		if requiredPermission == "" {
			continue
		}

		if !userPermissions[requiredPermission] {
			return false
		}
	}
	return true
}
