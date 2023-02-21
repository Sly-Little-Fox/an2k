package main

import (
	"fmt"
	"strings"

	"github.com/mattn/anko/ast"
	"github.com/thoas/go-funk"
)

func ConvertStmt(st ast.Stmt, v *[]string) string {
	// fmt.Printf("%#v\n", st)
	// fmt.Println()
	switch st := st.(type) {
	case *ast.StmtsStmt:
		return strings.Join(funk.Map(st.Stmts, func(s ast.Stmt) string {
			return ConvertStmt(s, v)
		}).([]string), "")
	case *ast.IfStmt:
		thing := ConditionIf(ConvertExpr(st.If)) + "\n" + ConvertStmt(st.Then, v)
		if st.Else != nil {
			thing += "иначе\n" + ConvertStmt(st.Else, v) + "\n"
		}
		return thing + ConditionEnd
	case *ast.LoopStmt:
		if st.Stmt == nil {
			return ""
		}
		stmt := any(st.Stmt).(*ast.StmtsStmt)
		return "\n" + LoopStart + "пока " + ConvertExpr(st.Expr) + "\n" + ConvertStmt(stmt, v) + LoopEnd
	case *ast.ExprStmt:
		return ConvertExpr(st.Expr)
	case *ast.LetsStmt:
		ids := funk.Map(st.LHSS, ConvertExpr).([]string)
		decl := funk.FilterString(ids, func(id string) bool { return !funk.ContainsString(*v, id) })
		*v = append(*v, decl...)
		old := []string{}
		for _, id := range ids {
			if !funk.ContainsString(decl, id) {
				old = append(old, id)
			}
		}
		result := ""
		if len(decl) > 0 {
			for k, id := range decl {
				result += Declaration(id, TypeOf(any(st.RHSS[k]).(*ast.LiteralExpr).Literal)) + "\n"
			}
		}
		if len(old) > 0 {
			for k, id := range old {
				result += Assignment(id, ConvertExpr(st.RHSS[k])) + "\n"
			}
		}
		return result
	}
	return ""
}

func ConvertExpr(ex ast.Expr) string {
	// fmt.Printf("%#v\n\n", ex)
	switch ex := ex.(type) {
	case *ast.LiteralExpr:
		if ex.Literal.Type().Name() == "string" {
			return fmt.Sprintf(`"%v"`, ex.Literal)
		}
		return fmt.Sprint(ex.Literal)
	case *ast.IdentExpr:
		return ex.Lit
	case *ast.OpExpr:
		return ConvertOper(ex.Op)
	case *ast.CallExpr:
		name := ex.Name
		args := []any{name}
		for _, expr := range ex.SubExprs {
			args = append(args, sprintCast(expr))
		}
		if name == "use" || name == "использовать" {
			name = "использовать"
			if module, ok := ModuleNames[args[1].(string)]; ok {
				args[1] = module
			}
		}
		if fname, ok := MoveReplaceMap[name]; ok {
			name = fname
		}
		args[0] = name
		return fmt.Sprintln(args...)
	}
	return ""
}

func ConvertOper(op ast.Operator) string {
	oper := any(op).(*ast.AddOperator)
	return fmt.Sprintln(ConvertExpr(oper.LHS), oper.Operator, ConvertExpr(oper.RHS))
}

func sprintCast(expr ast.Expr) string {
	switch expr := expr.(type) {
	case *ast.LiteralExpr:
		return fmt.Sprint(expr.Literal)
	}
	return ""
}
