package OS

import "os/user"

//标准库提供了不依赖平台的操作系统接口


//uid 用户id 操作系统用户 当前电脑
//gid  所属组id
// username 用户名

//group name 组名

func main() {
	user.Current()
}