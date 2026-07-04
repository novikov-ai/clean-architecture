package main

import (
	"fmt"

	"clean-architecture/robot/ast/internal/ast"
	"clean-architecture/robot/ast/internal/controller"
	"clean-architecture/robot/ast/internal/domain"
)

func main() {
	// AST structure:
	//
	// Move(100)
	//   └── Turn(-90)
	//        └── SetState(Soap)
	//             └── Move(50)
	//                  └── Stop

	program := ast.Move(100, func(r ast.MoveResponse) ast.Node {
		fmt.Printf("  ↳ moved to (%d, %d)\n", r.NewX, r.NewY)

		return ast.Turn(-90, func(r ast.TurnResponse) ast.Node {
			fmt.Printf("  ↳ now facing %d°\n", r.NewAngle)

			return ast.SetState(domain.Soap, func(r ast.StateResponse) ast.Node {
				fmt.Printf("  ↳ mode switched to %s\n", r.NewState)

				distance := 50
				if !r.Success {
					distance = 0
				}

				return ast.Move(distance, func(r ast.MoveResponse) ast.Node {
					fmt.Printf("  ↳ moved to (%d, %d)\n", r.NewX, r.NewY)
					return ast.Stop()
				})
			})
		})
	})

	final := controller.Run(program, domain.Robot{})
	fmt.Printf("\nFinal state: pos=(%d, %d) angle=%d mode=%s\n",
		final.Position.X, final.Position.Y, final.Angle, final.State)
}
