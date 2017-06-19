--
-- Created by IntelliJ IDEA.
-- User: admin
-- Date: 2017/6/19
-- Time: 下午4:02
-- To change this template use File | Settings | File Templates.
--

Account = {value = 0}
function Account:new(o)  -- 这里是冒号哦
    o = o or {}  -- 如果用户没有提供table，则创建一个
    setmetatable(o, self)
    self.__index = self
    return o
end

function Account:display()
    self.value = self.value + 100
    self.value1 = 1000
    print(self.value)
    print(self.value1)
end
--
--local a = Account:new{} -- 这里使用原型Account创建了一个对象a
--a:display() --(1)
--a:display() --(2)
--
--local b = Account.new(Account, {value=1000})
--b.display(b)
