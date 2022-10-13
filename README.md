> 参照资料: 
> 1. https://zhuanlan.zhihu.com/p/297029161
> 2. https://www.runoob.com/design-pattern/design-pattern-tutorial.html

# 目录结构介绍

```mermaid
graph LR;

/ --> Behavioral_Patterns
Behavioral_Patterns --> 01_Template
Behavioral_Patterns --> 02_Strategy

/ --> Creational_Patterns
Creational_Patterns --> 01_SimpleFactory
Creational_Patterns --> 02_FactoryMethod
Creational_Patterns --> 03_AbstractFactory
Creational_Patterns --> 04_Builder

/ --> Structural_Patterns
Structural_Patterns --> 01_Adapter
Structural_Patterns --> 02_Flyweight
```

# 1. Behavioral_Patterns(行为型模式)

> 这些设计模式特别关注对象之间的通信.

## 1.1 Template(模板模式)

> 在模板模式（Template Pattern）中，一个抽象类公开定义了执行它的方法的方式/模板。它的子类可以按需要重写方法实现，但调用将以抽象类中定义的方式进行。

## 1.2 Strategy(策略模式)

> 在策略模式（Strategy Pattern）中，一个类的行为或其算法可以在运行时更改。这种类型的设计模式属于行为型模式。

---

# 2. Creational_Patterns(创建型模式)

> 提供了一种在创建对象的同时隐藏创建逻辑的方式, 而不是使用"new"运算符直接实例化对象, 这使得程序在判断针对某个给定实例需要创建哪些对象时更加灵活.

## 2.1 SimpleFactory(简单工厂模式)

## 2.2 FactoryMethod(工厂模式)

> 工厂模式（Factory Pattern）是 Java 中最常用的设计模式之一。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。 在工厂模式中，我们在创建对象时不会对客户端暴露创建逻辑，并且是通过使用一个共同的接口来指向新创建的对象。

## 2.3 AbstractFactory(抽象工厂模式)

> 抽象工厂模式（Abstract Factory Pattern）是围绕一个超级工厂创建其他工厂。该超级工厂又称为其他工厂的工厂。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。 在抽象工厂模式中，接口是负责创建一个相关对象的工厂，不需要显式指定它们的类。每个生成的工厂都能按照工厂模式提供对象。

## 2.4 Builder(建造者模式)

> 建造者模式（Builder Pattern）使用多个简单的对象一步一步构建成一个复杂的对象。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。 一个 Builder 类会一步一步构造最终的对象。该 Builder 类是独立于其他对象的。

---

# 3. Structural_Patterns(结构型模式)

> 这些设计模式关注类和对象的组合. 继承的概念被用来组合接口和定义组合对象获得新功能的方式.

## 3.1 Adapter(适配器模式)

> 适配器模式（Adapter Pattern）是作为两个不兼容的接口之间的桥梁。这种类型的设计模式属于结构型模式，它结合了两个独立接口的功能。 这种模式涉及到一个单一的类，该类负责加入独立的或不兼容的接口功能。举个真实的例子，读卡器是作为内存卡和笔记本之间的适配器。您将内存卡插入读卡器，再将读卡器插入笔记本，这样就可以通过笔记本来读取内存卡。

## 3.2 Flyweight(享元模式)

> 享元模式（Flyweight Pattern）主要用于减少创建对象的数量，以减少内存占用和提高性能。这种类型的设计模式属于结构型模式，它提供了减少对象数量从而改善应用所需的对象结构的方式。 享元模式尝试重用现有的同类对象，如果未找到匹配的对象，则创建新对象。
