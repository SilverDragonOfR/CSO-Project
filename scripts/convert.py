from sympy.logic.boolalg import to_anf, to_cnf, to_dnf, to_nnf
import sys

check = sys.argv[1]
expression = sys.argv[2]

if check=="1":
    print("anf: ",to_anf(expression))
elif check=="2":
    print("cnf: ",to_cnf(expression))
elif check=="3":
    print("dnf: ",to_dnf(expression))
elif check=="4":
    print("nnf: ",to_nnf(expression))