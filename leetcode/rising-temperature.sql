
# 查找与之前（昨天的）日期相比温度更高的所有日期的 id 。
select a.id from Weather AS a join Weather AS b on
datediff(a.recordDate,b.recordDate) = 1 and
a.temperature > b.temperature