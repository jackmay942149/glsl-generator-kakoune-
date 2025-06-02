package main

import (
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
	rl.InitWindow(800, 450, "GLSL Generator")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		if len(graph.Nodes) == 0 {
			rl.DrawText("Press [SPACE] to add a node!", 190, 200, 20, rl.LightGray)
		} else {
			for _, rect := range graph.Nodes {
				rl.DrawRectangleV(rect.Position, rect.Size, rl.Black)
			}
		}

		rl.EndDrawing()

		if rl.IsKeyPressed(rl.KeySpace) {
			addNode()
		}

	}
}

func addNode() {
	mousePos := rl.GetMousePosition()
	nodeSize := rl.Vector2{X: 100, Y: 100}
	node := Node{Position: mousePos, Size: nodeSize}
	graph.Nodes = append(graph.Nodes, node)
}
