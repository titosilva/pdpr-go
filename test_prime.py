from math import sqrt,ceil
import re

# Python3 program Miller-Rabin primality test
import random 
 
# Utility function to do
# modular exponentiation.
# It returns (x^y) % p
def power(x, y, p):
     
    # Initialize result
    res = 1 
     
    # Update x if it is more than or
    # equal to p
    x = x % p 
    while (y > 0):
         
        # If y is odd, multiply
        # x with result
        if (y & 1):
            res = (res * x) % p
 
        # y must be even now
        y = y>>1 # y = y/2
        x = (x * x) % p
     
    return res
 
# This function is called
# for all k trials. It returns
# false if n is composite and 
# returns false if n is
# probably prime. d is an odd 
# number such that d*2<sup>r</sup> = n-1
# for some r >= 1
def miillerTest(d, n):
     
    # Pick a random number in [2..n-2]
    # Corner cases make sure that n > 4
    a = 2 + random.randint(1, n - 4)
 
    # Compute a^d % n
    x = power(a, d, n)
 
    if (x == 1 or x == n - 1):
        return True
 
    # Keep squaring x while one 
    # of the following doesn't 
    # happen
    # (i) d does not reach n-1
    # (ii) (x^2) % n is not 1
    # (iii) (x^2) % n is not n-1
    while (d != n - 1):
        x = (x * x) % n
        d *= 2
 
        if (x == 1):
            return False
        if (x == n - 1):
            return True
 
    # Return composite
    return False
 
# It returns false if n is 
# composite and returns true if n
# is probably prime. k is an 
# input parameter that determines
# accuracy level. Higher value of 
# k indicates more accuracy.
def isPrime( n, k):
     
    # Corner cases
    if (n <= 1 or n == 4):
        return False
    if (n <= 3):
        return True
 
    # Find r such that n = 
    # 2^d * r + 1 for some r >= 1
    d = n - 1
    while (d % 2 == 0):
        d //= 2
 
    # Iterate given number of 'k' times
    for i in range(k):
        if (miillerTest(d, n) == False):
            return False
 
    return True
 
# Driver Code
# Number of iterations
k = 20
 
# This code is contributed by mits

number_hex = """
FFFFFFFF FFFFFFFF C90FDAA2 2168C234 C4C6628B 80DC1CD1
29024E08 8A67CC74 020BBEA6 3B139B22 514A0879 8E3404DD
EF9519B3 CD3A431B 302B0A6D F25F1437 4FE1356D 6D51C245
E485B576 625E7EC6 F44C42E9 A637ED6B 0BFF5CB6 F406B7ED
EE386BFB 5A899FA5 AE9F2411 7C4B1FE6 49286651 ECE65381
FFFFFFFF FFFFFFFF
""" # Oakley Group

if __name__ == "__main__":
    print("Test prime")

    number_hex = re.sub(r"\s+", "", number_hex)
    number = int(number_hex, 16)
    number_be = number.to_bytes((number.bit_length() // 8), 'big', signed=False)
    print(list(number_be))

    # is_prime = True
    # for i in range(2, number // 2 + 1):
    #     if number % i == 0:
    #         print(f"{i} is a factor of {number}")
    #         is_prime = False
    #         break

    number_recovered = int.from_bytes(number_be, "big", signed=False)
    print(number_recovered % 2)

    if isPrime(number_recovered, k):
        print(f"{number_recovered} is a prime number")
    
    print("End test prime")