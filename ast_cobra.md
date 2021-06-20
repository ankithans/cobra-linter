```go
package example

import "github.com/spf13/cobra"

func main() {
        cmd := &cobra.Command{
                Use:            "Example",
                Short:          "This is short",
                Long:           "This is long description of the command. As you can see it is so long",
                Example:        "I have an example",
        }

        cmd.Execute()
}:
- &ast.File{Doc:(*ast.CommentGroup)(nil), Package:1, Name:(*ast.Ident)(0xc00000c060), Decls:[]ast.Decl{(*ast.GenDecl)(0xc00001e100), (*ast.FuncDecl)(0xc000010450)}, Scope:(*ast.Scope)(0xc000056240), Imports:[]*ast.ImportSpec{(*ast.ImportSpec)(0xc0000102a0)}, Unresolved:[]*ast.Ident{(*ast.Ident)(0xc00000c100)}, Comments:[]*ast.CommentGroup(nil)}

example:
- &ast.Ident{NamePos:9, Name:"example", Obj:(*ast.Object)(nil)}

import "github.com/spf13/cobra":
- &ast.GenDecl{Doc:(*ast.CommentGroup)(nil), TokPos:20, Tok:75, Lparen:0, Specs:[]ast.Spec{(*ast.ImportSpec)(0xc0000102a0)}, Rparen:0}

"github.com/spf13/cobra":
- &ast.ImportSpec{Doc:(*ast.CommentGroup)(nil), Name:(*ast.Ident)(nil), Path:(*ast.BasicLit)(0xc00000c080), Comment:(*ast.CommentGroup)(nil), EndPos:0}

"github.com/spf13/cobra":
- &ast.BasicLit{ValuePos:27, Kind:9, Value:"\"github.com/spf13/cobra\""}

func main() {
        cmd := &cobra.Command{
                Use:            "Example",
                Short:          "This is short",
                Long:           "This is long description of the command. As you can see it is so long",
                Example:        "I have an example",
        }

        cmd.Execute()
}:
- &ast.FuncDecl{Doc:(*ast.CommentGroup)(nil), Recv:(*ast.FieldList)(nil), Name:(*ast.Ident)(0xc00000c0a0), Type:(*ast.FuncType)(0xc00000c320), Body:(*ast.BlockStmt)(0xc000010420)}

main:
- &ast.Ident{NamePos:60, Name:"main", Obj:(*ast.Object)(0xc0000980f0)}

func():
- &ast.FuncType{Func:55, Params:(*ast.FieldList)(0xc000010300), Results:(*ast.FieldList)(nil)}

:
- &ast.FieldList{Opening:64, List:[]*ast.Field(nil), Closing:65}

{
        cmd := &cobra.Command{
                Use:            "Example",
                Short:          "This is short",
                Long:           "This is long description of the command. As you can see it is so long",
                Example:        "I have an example",
        }

        cmd.Execute()
}:
- &ast.BlockStmt{Lbrace:67, List:[]ast.Stmt{(*ast.AssignStmt)(0xc00001e1c0), (*ast.ExprStmt)(0xc0000562e0)}, Rbrace:287}

cmd := &cobra.Command{
        Use:            "Example",
        Short:          "This is short",
        Long:           "This is long description of the command. As you can see it is so long",
        Example:        "I have an example",
}:
- &ast.AssignStmt{Lhs:[]ast.Expr{(*ast.Ident)(0xc00000c0e0)}, TokPos:75, Tok:47, Rhs:[]ast.Expr{(*ast.UnaryExpr)(0xc00000c280)}}

cmd:
- &ast.Ident{NamePos:71, Name:"cmd", Obj:(*ast.Object)(0xc0000980a0)}

&cobra.Command{
        Use:            "Example",
        Short:          "This is short",
        Long:           "This is long description of the command. As you can see it is so long",
        Example:        "I have an example",
}:
- &ast.UnaryExpr{OpPos:78, Op:17, X:(*ast.CompositeLit)(0xc00001e180)}

cobra.Command{
        Use:            "Example",
        Short:          "This is short",
        Long:           "This is long description of the command. As you can see it is so long",
        Example:        "I have an example",
}:
- &ast.CompositeLit{Type:(*ast.SelectorExpr)(0xc00000c140), Lbrace:92, Elts:[]ast.Expr{(*ast.KeyValueExpr)(0xc000010360), (*ast.KeyValueExpr)(0xc000010390), (*ast.KeyValueExpr)(0xc0000103c0), (*ast.KeyValueExpr)(0xc0000103f0)}, Rbrace:266, Incomplete:false}

cobra.Command:
- &ast.SelectorExpr{X:(*ast.Ident)(0xc00000c100), Sel:(*ast.Ident)(0xc00000c120)}

cobra:
- &ast.Ident{NamePos:79, Name:"cobra", Obj:(*ast.Object)(nil)}

Command:
- &ast.Ident{NamePos:85, Name:"Command", Obj:(*ast.Object)(nil)}

Use: "Example":
- &ast.KeyValueExpr{Key:(*ast.Ident)(0xc00000c160), Colon:100, Value:(*ast.BasicLit)(0xc00000c180)}

Use:
- &ast.Ident{NamePos:97, Name:"Use", Obj:(*ast.Object)(nil)}

"Example":
- &ast.BasicLit{ValuePos:106, Kind:9, Value:"\"Example\""}

Short: "This is short":
- &ast.KeyValueExpr{Key:(*ast.Ident)(0xc00000c1a0), Colon:125, Value:(*ast.BasicLit)(0xc00000c1c0)}

Short:
- &ast.Ident{NamePos:120, Name:"Short", Obj:(*ast.Object)(nil)}

"This is short":
- &ast.BasicLit{ValuePos:129, Kind:9, Value:"\"This is short\""}

Long: "This is long description of the command. As you can see it is so long":
- &ast.KeyValueExpr{Key:(*ast.Ident)(0xc00000c200), Colon:153, Value:(*ast.BasicLit)(0xc00000c220)}

Long:
- &ast.Ident{NamePos:149, Name:"Long", Obj:(*ast.Object)(nil)}

"This is long description of the command. As you can see it is so long":
- &ast.BasicLit{ValuePos:158, Kind:9, Value:"\"This is long description of the command. As you can see it is so long\""}

Example: "I have an example":
- &ast.KeyValueExpr{Key:(*ast.Ident)(0xc00000c240), Colon:241, Value:(*ast.BasicLit)(0xc00000c260)}

Example:
- &ast.Ident{NamePos:234, Name:"Example", Obj:(*ast.Object)(nil)}

"I have an example":
- &ast.BasicLit{ValuePos:243, Kind:9, Value:"\"I have an example\""}

cmd.Execute():
- &ast.ExprStmt{X:(*ast.CallExpr)(0xc00001e200)}

cmd.Execute():
- &ast.CallExpr{Fun:(*ast.SelectorExpr)(0xc00000c2e0), Lparen:283, Args:[]ast.Expr(nil), Ellipsis:0, Rparen:284}

cmd.Execute:
- &ast.SelectorExpr{X:(*ast.Ident)(0xc00000c2a0), Sel:(*ast.Ident)(0xc00000c2c0)}

cmd:
- &ast.Ident{NamePos:272, Name:"cmd", Obj:(*ast.Object)(0xc0000980a0)}

Execute:
- &ast.Ident{NamePos:276, Name:"Execute", Obj:(*ast.Object)(nil)}
```
