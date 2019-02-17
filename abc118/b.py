N, M = map(int, input().split())
dic = {}
for i in range(0, N):
    Ak = list(map(int, input().split()))[1:]
   # print(Ak)
    for v in Ak:
        if v in dic.keys():
            dic[v] += 1
        else:
            dic[v] = 1    
result = 0
for vals in dic.values():
    if vals == N:
        result += 1

print(result)