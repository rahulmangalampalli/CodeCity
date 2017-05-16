/* Copyright 2017 Google Inc.
 * https://github.com/NeilFraser/CodeCity
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ast

import (
	"reflect"

	"CodeCity/server/flatpack"
)

func init() {
	var ifaces = reflect.TypeOf(
		struct {
			i1 node
			i2 Node
			i3 statement
			i4 expression
			i5 forStatementInit
			i6 propertyKey
			i7 lValue
		}{})
	for i := 0; i < ifaces.NumField(); i++ {
		flatpack.RegisterType(ifaces.Field(i).Type)
	}

	var examples = []interface{}{
		// From ast.go:
		Statement{},
		(Statements)(nil),
		Expression{},
		(Expressions)(nil),
		ForStatementInit{},
		LValue{},
		PropertyKey{},
		typeOnly{},

		// From astnodes.go:
		nodeStuff{},
		SourceLocation{},
		Position{},
		Identifier{},
		Literal{},
		Program{},
		functionStuff{},
		statementStuff{},
		ExpressionStatement{},
		BlockStatement{},
		EmptyStatement{},
		DebuggerStatement{},
		WithStatement{},
		ReturnStatement{},
		LabeledStatement{},
		BreakStatement{},
		ContinueStatement{},
		IfStatement{},
		SwitchStatement{},
		SwitchCase{},
		ThrowStatement{},
		TryStatement{},
		CatchClause{},
		WhileStatement{},
		DoWhileStatement{},
		ForStatement{},
		ForInStatement{},
		FunctionDeclaration{},
		VariableDeclaration{},
		VariableDeclarator{},
		expressionStuff{},
		ThisExpression{},
		ArrayExpression{},
		ObjectExpression{},
		Property{},
		FunctionExpression{},
		UnaryExpression{},
		UpdateExpression{},
		BinaryExpression{},
		AssignmentExpression{},
		LogicalExpression{},
		MemberExpression{},
		ConditionalExpression{},
		CallExpression{},
		NewExpression{},
		SequenceExpression{},
	}
	for _, val := range examples {
		flatpack.RegisterTypeOf(val)
	}
}
