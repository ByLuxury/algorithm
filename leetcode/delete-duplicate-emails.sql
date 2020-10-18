
# 采用了自连接
# delete p1删除满足where条件的记录
delete p1 from Person p1,Person p2 where p1.Email = p2.Email and p1.Id > p2.Id