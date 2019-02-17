N, M = map(int, input().split())
result_list = []
for i in range(M):
    p, y = map(int, input().split())
    result_list.append([p, y , i])
result_list.sort()
# print(result_list)

p_buf = 0
p_index = 1
answer = []
for row in result_list:
    if p_buf != row[0]:
        p_index = 1
        answer.append([str(row[0]).rjust(6, '0')+str(p_index).rjust(6, '0'), row[2]])
        p_buf = row[0]
    else:
        p_index += 1
        answer.append([str(p_buf).rjust(6, '0')+str(p_index).rjust(6, '0'), row[2]])


for r in sorted(answer, key=lambda x:x[1]):
    print(r[0])
