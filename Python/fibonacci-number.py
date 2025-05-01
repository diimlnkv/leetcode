#Time:
#Space:

class Solution:
    def fib(self, n: int) -> int:
        predefined = [0, 1]

        if n < 2:
            return n

        for i in range(2, n+1):
            s = predefined[i-1] + predefined[i-2]
            predefined.append(s)

        return predefined[n]