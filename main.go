package main

import (
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	p "github.com/jjeffcaii/antlr-mysql-example/parser"
)

type myVisitor struct {
	*p.BaseMySqlParserVisitor
}

func (mv *myVisitor) VisitRoot(ctx *p.RootContext) interface{} {
	return ctx.SqlStatements().Accept(mv)
}

func (mv *myVisitor) VisitSqlStatements(ctx *p.SqlStatementsContext) interface{} {
	for _, it := range ctx.AllSqlStatement() {
		it.Accept(mv)
	}
	return nil
}

func (mv *myVisitor) VisitSqlStatement(ctx *p.SqlStatementContext) interface{} {
	log.Println("bingo: sql statement!")
	// ...
	return nil
}

func main() {
	sql := "SELECT DATE_FORMAT(IF(COUNT(*)>0,'2021-01-02','2020-01-01'),'%Y') AS cc FROM student WHERE uid BETWEEN 1 AND ?"

	is := antlr.NewInputStream(sql)

	lexer := p.NewMySqlLexer(p.NewCaseChangingStream(is, true))

	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	mp := p.NewMySqlParser(ts)
	//mp.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	mp.RemoveErrorListeners()
	mp.SetErrorHandler(antlr.NewBailErrorStrategy())

	// ast root
	root := mp.Root()

	visitor := &myVisitor{
		BaseMySqlParserVisitor: new(p.BaseMySqlParserVisitor),
	}

	root.Accept(visitor)
}
