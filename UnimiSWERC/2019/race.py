from math import log2
import sys

def main():
    n = int(input())
    runners = []
    for i in range(n):
        head_start, speed = input().split()
        runners.append([i, int(speed), int(head_start)])
        
    for _ in range(int(log2(n))):
        for runner in runners:
            runner[2] += runner[1]
            
        runners.sort(key=lambda x: (x[2], x[0]), reverse=True)
        runners = runners[:len(runners) // 2]
        
    print(runners[0][0])

if __name__ == "__main__":
    main()