{参数}
SHORT=>12; LONG=>26; MID=>9;
{DRAWTEXT(1,3,'大阳线');}
DIF:EMA(CLOSE,SHORT)-EMA(CLOSE,LONG);
DEA:EMA(DIF,MID);
MACD:(DIF-DEA)*2,COLORSTICK;

做多:CROSS(DIF,DEA);
做空:CROSS(DEA,DIF);