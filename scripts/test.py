import sys
import ttg

expression = sys.argv[1]

var_str = expression.replace("not","").replace("nor","").replace("xor","").replace("or","").replace("nand","").replace("and","").replace("implies","")
variables = sorted(list(set([x for x in var_str if x in "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"])))
print(ttg.Truths(variables, [expression], ascending=True))