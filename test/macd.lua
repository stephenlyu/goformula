--
-- Created by IntelliJ IDEA.
-- User: admin
-- Date: 2017/6/19
-- Time: 下午3:59
-- To change self template use File | Settings | File Templates.
--


MACDClass = {}


function MACDClass:new(data, short, long, mid)
    o = {}
    setmetatable(o, self)
    self.__index = self

    o.data = data
    o.short = short
    o.long = long
    o.mid = mid
    o.close = CLOSE(data)
    o.const2 = Scalar(2)
    o.ema_close_short = EMA(o.close, o.short)
    o.ema_close_long = EMA(o.close, o.long)
    o.dif = SUB(o.ema_close_short, o.ema_close_long)
    o.dea = EMA(o.dif, o.mid)

    o.dif_sub_dea = SUB(o.dif, o.dea)
    o.macd = MUL(o.dif_sub_dea, o.const2)
    o.enter_long = CROSS(o.dif, o.dea)
    o.enter_short = CROSS(o.dea, o.dif)
    return o
end

function MACDClass:updateLastValue()
    self.close.UpdateLastValue()
    self.ema_close_short.UpdateLastValue()
    self.ema_close_long.UpdateLastValue()
    self.dif.UpdateLastValue()
    self.dea.UpdateLastValue()
    self.dif_sub_dea.UpdateLastValue()
    self.macd.UpdateLastValue()
    self.enter_long.UpdateLastValue()
    self.enter_short.UpdateLastValue()
end

function MACDClass:Len()
    return self.data.Len()
end


function MACDClass:Get(index)
    return {
        self.dif.Get(index),
        self.dea.Get(index),
        self.macd.Get(index),
        self.enter_long.Get(index),
        self.enter_short.Get(index),
    }
end

FormulaClass = MACDClass
