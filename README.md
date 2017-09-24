# lucky-money

编译后拷贝log.xml到执行目录后直接执行即可<br><br>
接口:<br>
所有接口都需要带query : ?account=该用户id, 所有红包以分为单位<br>
http://localhost/luckmoney/envelops 该用户打开的所有红包记录(GET请求)<br>
http://localhost/luckmoney/envelop/open 打开某个口令的红包(POST { "code" : "12345678"} )<br>
http://localhost/luckmoney/envelop 生成一个红包 返回该红包的口令(POST { "amount":111,"number":4 })<br>
http://localhost/luckmoney/account/balance 该用户的余额(GET 返回{"balance":11111})<br>
