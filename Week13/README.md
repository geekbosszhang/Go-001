毕业总结

老的项目结构, Beego server的MVC三层架构，之后单体server又增加了很多新的功能，就只能扁平化的添加目录、添加功能。

随着功能的增加，维护性和扩展性越来越差，有些业务代码耦合比较紧，模块间划分没有那么清晰，尤其有时候需要去调用某个函数的时候因为发现出现了循环引用的问题，不得不挪函数，重新建个包来处理。

其次错误处理没有用wrap,随意打log，有beego的Log还有自己项目中的gplog，到底遇到问题的时候log不少，但是有用的消息不多。

第三是没有统一的接口文档，注释也写的比较随意，没有一个有效的能track修改的地方。

还有就是测试，由于历史欠的技术债，DB Connection的Mock混合了几种不同的方式，导致test有时候也比较难写。

```
- webserver
    - controller
    - models
    - routers
- bin
- common
- cli
- conn
```
随着客户有了更多需要开放API接口的需求，我们需要去重构已有的系统。
通过Go训练营的课程，学到了一些新的东西可以用来实践
1. 微服务的拆分

想法是将metrics的采集、聚合逻辑拆分成一个原子服务，能够为上层的服务提供接口
比如可以接Alert模块，接展示模块、接WLM模块，可以接外部的prometheus模块等等

2. 错误的处理

之前代码中很多使用 Sentinel Error 去判断，不灵活。而且有很多地方出现即打印Log，又往上层return error, 重复打印无效的error,以及不能在错误输出中很好的展示stack的调用栈关系
```
if err == ErrSomehing {...}
```
另外如果error != nil, 函数的其他返回值是一个不可用的状态，不应该对返回值有任何期待

3. 使用ErrorGroup处理并发

项目中使用了很多goroutine, 但是通常goroutine的状态需要大量的状态跟踪跟chan操作，同时我们之前没有很好的能处理goroutine级联状态的取消的方法，而且由于项目中几年前用到的go的版本相对比较老，所以新的go版本中的context, errGroup等新特性没有能用到。超时控制也可以考虑refactor下。
以及避免使用野生goroutine panic导致主进程退出

4. 项目中的依赖关系的处理

之前还是传统的手动依赖注入方式，可以考虑使用下wire



