#ifndef _LCD_H_
#define _LCD_H_

#include <REGX52.H>

#ifndef u8
#define u8 unsigned char
#endif

#ifndef u16
#define u16 unsigned int
#endif

//PIN�ڶ���
#define LCD1602_DATAPINS P0
sbit RE = P2^7;
sbit RW = P2^5;
sbit RS = P2^6;

//��Ӧ�ĺ�������
void Lcd1602_Delay1ms(u16 c);
void LcdWriteCmd(u8 c);
void LcdWriteData(u8 c);
void LcdInit();
#endif