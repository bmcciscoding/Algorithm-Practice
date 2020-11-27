# Gin 框架简单使用手册


问题

[gin] listen tcp :3000: bind: address already in use

解决：
lsof -wni tcp:3000
kill pid