# fakeeyes_agent

代码包说明
* app 程序入口， 加载配置文件
* config 配置文件解析
* controller 控制器 主要控制逻辑， 保持server链接，解析指令、消费指令
* drivers 驱动包实现， 
    * macos Mac设备作为agent实现
    * x86linux x86服务器Linux的agent实现
    * raspberry 树莓派实现
* motion 移动操作控制
* video 视频控制
* audio 控制

## drivers
对接具体设备平台的实现, 不同平台设备下有不同的控制实现

## controller 
控制器， 对接服务器链接，定时心跳保持，接收服务器命令，转发到不同的控制器去具体实现。

## motion/video/audio
控制器调用的实现， 有共用控制逻辑。 接收`controller`的指令，调用`driver`操作。 

和具体`driver`的实现区别: 

比如:

一般车子行进会有了动能回收，也就是如果不按`前进键`, 车子会停下来。 如果直接对接`driver`接口，车子会立刻停下来， 而我们希望能车子缓步停下来， 会有更好的体验。
* `drivers`实现具体的动力控制操作
* `motion` 实现缓步控制操作

