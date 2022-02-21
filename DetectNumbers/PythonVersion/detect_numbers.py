import re
import sys


def string_is_number(st):
    return re.match(r'^[\-]*\d*[.e]?\d+$', st)


if __name__ == "__main__":
    input_st = input("Please enter the string: ")

    sys.stderr.write("Checking your input...\n")
    is_number = string_is_number(input_st)
    sys.stderr.write("You entered a number\n" if is_number else "You did not enter a number\n")
