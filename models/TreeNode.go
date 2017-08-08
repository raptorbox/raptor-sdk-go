package models

//TreeNode a generic tree structure
type TreeNode struct {
	parent *TreeNode

	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Order    int         `json:"order"`
	ParentID string      `json:"parentId"`
	UserID   string      `json:"userId"`
	Children []*TreeNode `json:"children"`
}

//GetParent return the parent if any
func (t *TreeNode) GetParent() *TreeNode {
	return t.parent
}

//IsRoot check if this node is at the tree root
func (t *TreeNode) IsRoot() bool {
	return t.parent == nil
}

//HasParent check if a parent exists
func (t *TreeNode) HasParent() bool {
	return t.parent != nil
}

//SetParent set the parent
func (t *TreeNode) SetParent(parent *TreeNode) {
	t.parent = parent
	t.ParentID = parent.ID
}

//HasChildren check if children exists
func (t *TreeNode) HasChildren() bool {
	return len(t.Children) > 0
}
