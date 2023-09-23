from sympy.logic.boolalg import is_anf, is_cnf, is_dnf, is_nnf
import sys

check = sys.argv[1]
expression = sys.argv[2]

if check=="1":
    print("anf: ",is_anf(expression))
elif check=="2":
    print("cnf: ",is_cnf(expression))
elif check=="3":
    print("dnf: ",is_dnf(expression))
elif check=="4":
    print("nnf: ",is_nnf(expression))