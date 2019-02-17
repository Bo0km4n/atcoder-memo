
def print_dp(dp):
    for array in dp:
        print(array)

N, W = (int(i) for i in input().split())
bagage = [[int(i) for i in input().split()] for _ in range(N)]
weight = []
value = []
for b in bagage:
    value.append(b[0])
    weight.append(b[1])

dp = [[0]*6 for i in range(5)]

for j in range(W):
    dp[N][j] = 0

for i in range(0, N, 1):
    for j in range(0, W+1, 1):
        if j < weight[i-1]:
            dp[i][j] = dp[i][j]
        else:
            dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]] + value[i])
            print_dp(dp)
            print("=" * 30)

print(dp[N-1][W])
    