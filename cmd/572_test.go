package main

import "testing"

func Test572(t *testing.T) {
	t1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
							Right: &TreeNode{
								Val: 1,
								Right: &TreeNode{
									Val: 1,
									Right: &TreeNode{
										Val: 1,
										Right: &TreeNode{
											Val: 1,
											Right: &TreeNode{
												Val: 1,
											},
											Left: &TreeNode{
												Val: 1,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	t2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
							Right: &TreeNode{
								Val: 1,
								Right: &TreeNode{
									Val: 1,
								},
								Left: &TreeNode{
									Val: 1,
								},
							},
						},
					},
				},
			},
		},
	}

	val := isSubtree(t1, t2)

	t.Log(val)
}
