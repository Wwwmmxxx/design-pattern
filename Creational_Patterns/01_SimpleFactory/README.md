# 简单工厂模式

将对象的创建过程, 用一个公共的函数封装起来, 外部只需要调用该函数即可. 很多go源码中都在使用这种方法.当看到NewXXX开头的函数, 基本都是这种模式的体现.