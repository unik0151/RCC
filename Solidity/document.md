# 安装
npm install -g solc

# 版本
window:
solcjs.cmd --version


# 基础入门

## 创建一个sol文件
pragram 提示 声明License ，version 对应编辑器版本

```sol

contract Hello {
    // 在remix执行部署时，会将普通类型压到栈中，每个声明对应一个槽位
    一个槽位256位，
    // 对于数值类型默认为最大 例如：int256 uint256

    // 

    //声明一个普通变量
    string public name ; 
    
}

```
## 数据结构
值类型 一般对于栈
引用数据类型 对于数组 之类的

栈 stack 每个变量执行时都会调用其地址，每个变量放入栈时都会有个keccak256哈希值与这个字节码对应
memory  
storage 

## 访问符
当前按