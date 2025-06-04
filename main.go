package main

import (
	"errors"
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Pin struct {
	Position   rl.Vector2
	Color      rl.Color
	Connection *Pin
	IsOut      bool
}

type Node struct {
	Type     string
	Position rl.Vector2
	Size     rl.Vector2
	OutPins  []Pin
	InPins   []Pin
}

type Graph struct {
	OutputNode *Node
	Nodes      []Node
}

// State
var graph Graph

func main() {
	// Start up window
	fmt.Println("Starting App...")
	rl.InitWindow(800, 450, "GLSL Generator")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// State
	var selectedNode *Node
	var selectedPin *Pin

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		// Draw whole graph
		if len(graph.Nodes) == 0 {
			rl.DrawText("Press [SPACE] to add a node!", 190, 200, 20, rl.LightGray)
		} else {
			for _, node := range graph.Nodes {
				drawNode(node)
			}
		}

		// Process Input
		if rl.IsKeyPressed(rl.KeySpace) {
			addNode()
		}

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			var err error
			selectedPin, err = selectPin()
			if err != nil {
				selectedNode, err = selectNode()
			}
		}

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			moveNode(selectedNode)
			drawLineToPin(selectedPin)
		}

		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			if selectedPin != nil {
				tryConnectPin(selectedPin)
			}
			deselectPin(&selectedPin)
			deselectNode(&selectedNode)
		}

		rl.EndDrawing()
	}
}

func addNode() {
	name := "Add"
	nodeSize := rl.Vector2{X: 100, Y: 100}
	mousePos := rl.Vector2Add(rl.GetMousePosition(), rl.Vector2Scale(nodeSize, -0.5))
	pinPos := rl.Vector2{X: 0, Y: 0}
	p1 := Pin{IsOut: true, Color: rl.White, Position: pinPos}
	p2 := Pin{IsOut: false, Color: rl.White, Position: pinPos}
	p3 := Pin{IsOut: false, Color: rl.White, Position: pinPos}
	outPins := make([]Pin, 0)
	inPins := make([]Pin, 0)
	outPins = append(outPins, p1)
	inPins = append(inPins, p2, p3)
	node := Node{Type: name, Position: mousePos, Size: nodeSize, OutPins: outPins, InPins: inPins}
	setPinPositions(&node)
	graph.Nodes = append(graph.Nodes, node)
}

func selectPin() (*Pin, error) {
	for _, node := range graph.Nodes {
		for j, pin := range node.OutPins {
			if rl.CheckCollisionPointCircle(rl.GetMousePosition(), pin.Position, 10) {
				return &node.OutPins[j], nil
			}
		}
		for j, pin := range node.InPins {
			if rl.CheckCollisionPointCircle(rl.GetMousePosition(), pin.Position, 10) {
				return &node.InPins[j], nil
			}
		}
	}
	return nil, errors.New("No Pin Found")
}

func selectNode() (*Node, error) {
	for i, node := range graph.Nodes {
		rec := rl.Rectangle{X: node.Position.X, Y: node.Position.Y, Width: node.Size.X, Height: node.Size.Y}
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), rec) {
			return &graph.Nodes[i], nil
		}
	}
	return nil, fmt.Errorf("No Node Found")
}

func moveNode(n *Node) {
	if n != nil {
		n.Position = rl.Vector2Add(rl.GetMousePosition(), rl.Vector2Scale(n.Size, -0.5))
		setPinPositions(n)
	}
}

func setPinPositions(node *Node) {
	for i := range node.OutPins {
		node.OutPins[i].Position = rl.Vector2{X: node.Position.X + 100, Y: node.Position.Y + float32(i*30.0)}
	}
	for i := range node.InPins {
		node.InPins[i].Position = rl.Vector2{X: node.Position.X, Y: node.Position.Y + float32(i*30.0)}
	}
}

func deselectNode(n **Node) {
	if *n != nil {
		*n = nil
	}
}

func deselectPin(p **Pin) {
	if *p != nil {
		*p = nil
	}
}

func drawNode(n Node) {
	// Draw Node Base
	rl.DrawRectangleV(n.Position, n.Size, rl.Black)

	// Draw Node Type
	width := rl.MeasureText(n.Type, 20)
	rl.DrawText(n.Type, int32(n.Position.X+n.Size.X/2)-width/2, int32(n.Position.Y), 20, rl.White)

	// Draw Pins
	for _, p := range n.InPins {
		rl.DrawCircle(int32(p.Position.X), int32(p.Position.Y), 10, rl.White)
	}
	for _, p := range n.OutPins {
		rl.DrawCircle(int32(p.Position.X), int32(p.Position.Y), 10, rl.White)
	}
	// Draw Connections
	for _, p := range n.OutPins {
		if p.Connection != nil {
			rl.DrawLineV(p.Position, p.Connection.Position, rl.White)
		}
	}
}

func drawLineToPin(p *Pin) {
	if p != nil {
		rl.DrawLineV(rl.GetMousePosition(), p.Position, rl.White)
	}
}

func tryConnectPin(p *Pin) error {
	pin, err := selectPin()
	if err == nil {
		if pin == nil {
			log.Fatal("Pin Not Found")
		}
		if p.IsOut == pin.IsOut {
			return errors.New("Pin IsOut Matches")
		}
		p.Connection = pin
		pin.Connection = p
	}
	return err
}
