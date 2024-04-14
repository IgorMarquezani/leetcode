class Solution:
    def evalRPN(self, tokens: list[str]) -> int:
        stack = []
        for v in tokens:
            match v:
                case "+":
                    stack.append(int(stack.pop()) + int(stack.pop()))
                case "-":
                    a, b = int(stack.pop()), int(stack.pop())
                    stack.append(b - a)
                case "*":
                    stack.append(int(stack.pop()) * int(stack.pop()))
                case "/":
                    a, b = int(stack.pop()), int(stack.pop())
                    stack.append(int(b / a))
                case _:
                    stack.append(int(v))

        return int(stack[0])


s = Solution()
print(s.evalRPN(["2", "1", "+", "3", "*"]))
print(s.evalRPN(["4", "13", "5", "/", "+"]))
