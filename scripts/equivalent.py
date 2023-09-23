import sympy
import sys

ans = True
x = sympy.parsing.sympy_parser.parse_expr
expressionList = sys.argv[1:]
for i in range(1,len(expressionList)):
    ans = ans and (x(expressionList[i]).equals(x(expressionList[i-1])))
    # print("hi",ans)
print(ans)