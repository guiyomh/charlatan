package model

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	type args struct {
		row *Row
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_HasDependencyOf(t *testing.T) {
	type fields struct {
		Left  *Node
		Right *Node
		Key   string
		Value *Row
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Node{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if got := me.HasDependencyOf(tt.args.name); got != tt.want {
				t.Errorf("Node.HasDependencyOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_LessThan(t *testing.T) {
	type fields struct {
		Left  *Node
		Right *Node
		Key   string
		Value *Row
	}
	type args struct {
		other *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Node{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if got := me.LessThan(tt.args.other); got != tt.want {
				t.Errorf("Node.LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_EqualTo(t *testing.T) {
	type fields struct {
		Left  *Node
		Right *Node
		Key   string
		Value *Row
	}
	type args struct {
		other *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Node{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if got := me.EqualTo(tt.args.other); got != tt.want {
				t.Errorf("Node.EqualTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_GreaterThan(t *testing.T) {
	type fields struct {
		Left  *Node
		Right *Node
		Key   string
		Value *Row
	}
	type args struct {
		other *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Node{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if got := me.GreaterThan(tt.args.other); got != tt.want {
				t.Errorf("Node.GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Add(t *testing.T) {
	type fields struct {
		Left  *Node
		Right *Node
		Key   string
		Value *Row
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Node{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if got := me.Add(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Minimum(t *testing.T) {
	type fields struct {
		Left  *Node
		Right *Node
		Key   string
		Value *Row
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Node{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if got := me.Minimum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Minimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Maximum(t *testing.T) {
	type fields struct {
		Left  *Node
		Right *Node
		Key   string
		Value *Row
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Node{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if got := me.Maximum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.Maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_WalkForward(t *testing.T) {
	type fields struct {
		Left  *Node
		Right *Node
		Key   string
		Value *Row
	}
	type args struct {
		iterator func(me *Node)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Node{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			me.WalkForward(tt.args.iterator)
		})
	}
}

func TestNode_WalkBackward(t *testing.T) {
	type fields struct {
		Left  *Node
		Right *Node
		Key   string
		Value *Row
	}
	type args struct {
		iterator func(me *Node)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Node{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			me.WalkBackward(tt.args.iterator)
		})
	}
}

func TestNewTree(t *testing.T) {
	tests := []struct {
		name string
		want *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Add(t *testing.T) {
	type fields struct {
		Records map[string]*Node
		Root    *Node
	}
	type args struct {
		record *Row
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Tree{
				Records: tt.fields.Records,
				Root:    tt.fields.Root,
			}
			me.Add(tt.args.record)
		})
	}
}

func TestTree_Find(t *testing.T) {
	type fields struct {
		Records map[string]*Node
		Root    *Node
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Tree{
				Records: tt.fields.Records,
				Root:    tt.fields.Root,
			}
			if got := me.Find(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_First(t *testing.T) {
	type fields struct {
		Records map[string]*Node
		Root    *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  *Row
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Tree{
				Records: tt.fields.Records,
				Root:    tt.fields.Root,
			}
			got, got1 := me.First()
			if got != tt.want {
				t.Errorf("Tree.First() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Tree.First() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTree_Last(t *testing.T) {
	type fields struct {
		Records map[string]*Node
		Root    *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  *Row
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Tree{
				Records: tt.fields.Records,
				Root:    tt.fields.Root,
			}
			got, got1 := me.Last()
			if got != tt.want {
				t.Errorf("Tree.Last() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Tree.Last() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTree_Walk(t *testing.T) {
	type fields struct {
		Records map[string]*Node
		Root    *Node
	}
	type args struct {
		iterator Iterator
		forward  bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Tree{
				Records: tt.fields.Records,
				Root:    tt.fields.Root,
			}
			me.Walk(tt.args.iterator, tt.args.forward)
		})
	}
}
