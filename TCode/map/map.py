l = [1,2,3]

dic = {}
#dic[l] = "windows" # TypeError: unhashable type: 'list'

l2 = (1,2,3)
dic[l2] = "windows"
print(dic)
print(dic[(1,2,3)])