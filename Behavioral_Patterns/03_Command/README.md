# 命令模式

将一个请求封装成一个对象，从而使您可以用不同的请求对客户进行参数化。

## 介绍

**主要解决:** 在软件系统中, 行为请求者与行为实现者通常是一种紧耦合的关系, 但某些场合, 比如需要对行为进行记录、撤销或重做、事务等处理时, 这种无法抵御变化的紧耦合的设计就不太合适.

**应用场景:** struts 1 中的 action 核心控制器 ActionServlet 只有一个, 相当于 Invoker, 而模型层的类会随着不同的应用有不同的模型类, 相当于具体的 Command.

## 优点

1. 降低了系统耦合度
2. 新的命令可以很容易添加到系统中去

## 缺点

1. 使用命令模式可能会导致某些系统有过多的具体命令类