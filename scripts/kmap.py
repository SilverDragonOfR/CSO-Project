import sys
import sympy
from sympy import symbols
from sympy.logic import POSform, SOPform
x = sympy.parsing.sympy_parser.parse_expr

check = sys.argv[1]
num = int(sys.argv[2])
minterms = [int(x) for x in sys.argv[3].split(",")]
dontcare = [int(x) for x in sys.argv[4].split(",")]
k = ["a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"]

if check=="1":
    print("SOP form : ",SOPform(symbols(" ".join(k[:num])), minterms=minterms, dontcares=dontcare))
elif check=="2":
    print("POS form : ",POSform(symbols(" ".join(k[:num])), minterms=minterms, dontcares=dontcare))
