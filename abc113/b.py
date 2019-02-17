n = int(input())
t, a = map(int,input().split())
li = list(map(int,input().split()))
n_temp = li[0]
answer = 0
for idx, temp in enumerate(li):
    i_temp = t - temp * 0.006
    if (abs(a - n_temp) >= abs(a -i_temp)):
        n_temp = i_temp
        answer = idx

print(answer+1)
