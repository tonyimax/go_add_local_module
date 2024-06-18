module com.test/hello //包名

go 1.22.2 //go运行时版本

require (
	//本地模块
	com.test/greetings v0.0.0-00010101000000-000000000000 //使用 go mod edit -replace com.test/greetings=../greetings 生成  go mod tidy 更新
	com.test/room v0.0.0-00010101000000-000000000000 //使用 go mod edit -replace com.test/room=../room   生成  go mod tidy 更新
	rsc.io/quote v1.5.2
)

require (
	//公网模块
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

//替换本地模块路径
replace com.test/greetings => ../greetings

replace com.test/room => ../room
