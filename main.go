package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Node struct {
	Position rl.Vector2
	Size     rl.Vector2
}

type Graph struct {
	Nodes []Node
}

var graph Graph

func main() {
	fmt.Println("Starting App...")
	rl.InitWindow(800, 450, "GLSL Generator")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		if len(graph.Nodes) == 0 {
			rl.DrawText("Press [SPACE] to add a node!", 190, 200, 20, rl.LightGray)
		} else {
			for _, node := range graph.Nodes {
				rl.DrawRectangleV(node.Position, node.Size, rl.Black)
			}
		}

		rl.EndDrawing()

		if rl.IsKeyPressed(rl.KeySpace) {
			addNode()
		}

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			for i, node := range graph.Nodes {
				rec := rl.Rectangle{X: node.Position.X, Y: node.Position.Y, Width: node.Size.X, Height: node.Size.Y}
				if rl.CheckCollisionPointRec(rl.GetMousePosition(), rec) {
					graph.Nodes[i].Position = rl.Vector2Add(rl.GetMousePosition(), rl.Vector2Scale(node.Size, -0.5))
				}
			}
		}

	}
}

func addNode() {
	nodeSize := rl.Vector2{X: 100, Y: 100}
	mousePos := rl.Vector2Add(rl.GetMousePosition(), rl.Vector2Scale(nodeSize, -0.5))
	node := Node{Position: mousePos, Size: nodeSize}
	graph.Nodes = append(graph.Nodes, node)
}
