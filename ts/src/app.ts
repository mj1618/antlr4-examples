import antlr4, { TerminalNode } from "antlr4";
import CalcLexer from "./calc/CalcLexer";
import CalcParser, {
  ExpressionContext,
  ProgramContext,
} from "./calc/CalcParser";
import CalcVisitor from "./calc/CalcVisitor";

class Visitor extends CalcVisitor<number> {
  visitProgram: (ctx: ProgramContext) => number = (ctx: ProgramContext) => {
    ctx.expression_list().map((exp) => console.log(this.visitExpression(exp)));
    return 0;
  };

  visitExpression: (ctx: ExpressionContext) => number = (
    ctx: ExpressionContext
  ) => {
    const numbers: number[] = [];
    const children = ctx.children?.slice(1, ctx.children.length - 1);

    children?.map((child) => {
      if (child instanceof ExpressionContext) {
        numbers.push(this.visitExpression(child));
      } else if (
        child instanceof TerminalNode &&
        child.symbol.type === CalcParser.NUMBER
      ) {
        // console.log(child.symbol.type);
        numbers.push(parseInt(child.getText()));
      }
    });

    const operator = children?.[0].getText();
    // console.log(children?.[0]);
    const operand = ctx._operand.text;

    switch (operand) {
      case "+":
        return numbers.reduce((acc, curr) => acc + curr, 0);
      case "-":
        return numbers.reduce((acc, curr) => acc - curr, 0);
      case "*":
        return numbers.reduce((acc, curr) => acc * curr, 1);
      case "/":
        return numbers.slice(1).reduce((acc, curr) => acc / curr, numbers[0]);
    }
    console.log("returning dfault");
    return 0;
  };
}

const main = () => {
  const is = antlr4.CharStreams.fromString("(+ (/ 26 (* 2 (+ 1 2 3))))");
  const lexer = new CalcLexer(is);
  const tokens = new antlr4.CommonTokenStream(lexer);
  const parser = new CalcParser(tokens);
  const tree = parser.program();

  console.log(tree.toStringTree(parser.ruleNames, parser));
  console.log(parser.ruleNames);

  const visitor = new Visitor();
  visitor.visitProgram(tree);
};

main();
