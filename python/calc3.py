import sys
from antlr4 import *
from CalcPlusLexer import CalcPlusLexer
from CalcPlusParser import CalcPlusParser
from Calculator import Calculator


def main():
    if len(sys.argv) < 2:
        print("Usage: python calc3.py <file_path>")
        return

    input_path = sys.argv[1]

    input_stream = FileStream(input_path, encoding='utf-8')

    lexer = CalcPlusLexer(input_stream)
    stream = CommonTokenStream(lexer)

    parser = CalcPlusParser(stream)
    tree = parser.calc3()

    visitor = Calculator()
    visitor.visit(tree)


if __name__ == '__main__':
    main()