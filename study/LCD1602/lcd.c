#include "lcd.h"

void Lcd1602_Delay1ms(u16 c){//Îó²îÎª0ºÁÃë
	u8 a,b;
	for(;c>0;c--){
		for(b=199;b>0;b--){
			for(a=1;a>0;a--){
			}
		}
	}
}

void LcdWriteCmd(u8 c){
	RE = 0;
	RS = 0;
	RW = 0;
	LCD1602_DATAPINS = c;
	Lcd1602_Delay1ms(1);
	
	RE = 1;
	Lcd1602_Delay1ms(5);
	RE = 0;
	
	LCD1602_DATAPINS = c<<4;
	Lcd1602_Delay1ms(1);
	
	RE = 1;
	Lcd1602_Delay1ms(5);
	RE = 0;
}

void LcdWriteData(u8 c){
	RE = 0;
	RS = 1;
	RW = 0;
	LCD1602_DATAPINS = c;
	Lcd1602_Delay1ms(1);
	
	RE = 1;
	Lcd1602_Delay1ms(5);
	RE = 0;
	
	LCD1602_DATAPINS = c<<4;
	Lcd1602_Delay1ms(1);
	
	RE = 1;
	Lcd1602_Delay1ms(5);
	RE = 0;
}

void LcdInit(){
	LcdWriteCmd(0x32);
	LcdWriteCmd(0x28);
	LcdWriteCmd(0x0c);
	LcdWriteCmd(0x06);
	LcdWriteCmd(0x01);
	LcdWriteCmd(0x80);
}